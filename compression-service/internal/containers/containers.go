package containers

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
	"io"
	"os"
	"strings"
	"time"
)

type Service struct {
	Client      *client.Client
	WorkerImage string
	Containers  []*Container
}

type Container struct {
	Id    string `json:"id"`
	JobId int64  `json:"jobId"`
	Port  int    `json:"port"`
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) InitializeClient() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Initialize client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("error initializing docker client: %v", err)
	}
	s.Client = cli

	// Authenticate against container registry
	username := os.Getenv("DOCKER_USERNAME")
	password := os.Getenv("DOCKER_PASSWORD")
	if len(username) == 0 || len(password) == 0 {
		return fmt.Errorf("missing GHCR credentials")
	}

	authConfig := registry.AuthConfig{
		Username: username,
		Password: password,
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		// Do not pass the error here as it likely contains credentials
		return errors.New("error marshalling auth config")
	}
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)

	// Pull worker image
	s.WorkerImage = os.Getenv("WORKER_IMAGE_URL")
	_, err = cli.ImagePull(ctx, s.WorkerImage, image.PullOptions{
		RegistryAuth: authStr,
	})
	if err != nil {
		return fmt.Errorf("error pulling docker image: %v", err)
	}

	return nil
}

func (s *Service) CloseClient() error {
	return s.Client.Close()
}

func (s *Service) NewContainer(
	jobId int64,
) (*Container, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	containerName := fmt.Sprintf("worker-%d", jobId)

	resp, err := s.Client.ContainerCreate(ctx, &container.Config{
		Image: s.WorkerImage,
	}, nil, nil, nil, containerName)
	if err != nil {
		return nil, fmt.Errorf("error creating container: %v", err)
	}

	cont := &Container{
		Id:    resp.ID,
		JobId: jobId,
	}

	s.Containers = append(s.Containers, cont)

	err = s.Client.ContainerStart(ctx, resp.ID, container.StartOptions{})
	if err != nil {
		return cont, fmt.Errorf("error starting container: %v", err)
	}

	return cont, nil
}

func (s *Service) RemoveContainer(
	containerId string,
) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := s.Client.ContainerRemove(ctx, containerId, container.RemoveOptions{
		Force: true,
	})
	if err != nil {
		return fmt.Errorf("error removing container: %v", err)
	}

	for i, cont := range s.Containers {
		if cont.Id == containerId {
			s.Containers = append(s.Containers[:i], s.Containers[i+1:]...)
			break
		}
	}

	return nil
}

type ContainerEvent struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func (s *Service) MonitorOutput(
	ctx context.Context,
	containerId string,
	events chan<- ContainerEvent,
) error {
	resp, err := s.Client.ContainerAttach(ctx, containerId, container.AttachOptions{
		Stream: true,
		Stderr: true,
		Stdout: true,
		Logs:   true,
	})
	if err != nil {
		return err
	}
	defer resp.Close()

	reader := bufio.NewReader(resp.Reader)

	var collectingJSON bool
	var jsonBuffer strings.Builder

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					events <- ContainerEvent{Type: "EOF", Data: nil}
					return nil
				}
				return fmt.Errorf("error reading container output: %v", err)
			}

			line = strings.TrimSpace(line)

			switch line {
			case "APPLICATION_STARTED":
				events <- ContainerEvent{Type: "APPLICATION_STARTED", Data: nil}

			case "SERVER_FAILED":
				events <- ContainerEvent{Type: "SERVER_FAILED", Data: nil}

			case "DOWNLOAD_FAILED":
				events <- ContainerEvent{Type: "DOWNLOAD_FAILED", Data: nil}

			case "DOWNLOAD_COMPLETED":
				events <- ContainerEvent{Type: "DOWNLOAD_COMPLETED", Data: nil}

			case "PROBE_FAILED":
				events <- ContainerEvent{Type: "PROBE_FAILED", Data: nil}

			case "START_PROBE_DATA":
				collectingJSON = true
				jsonBuffer.Reset()

			case "END_PROBE_DATA":
				collectingJSON = false
				var probeData map[string]interface{}
				if err := json.Unmarshal([]byte(jsonBuffer.String()), &probeData); err != nil {
					events <- ContainerEvent{Type: "ERROR", Data: fmt.Sprintf("error parsing probe data: %v", err)}
				} else {
					events <- ContainerEvent{Type: "PROBE_DATA", Data: probeData}
				}

			case "COMPRESSION_FAILED":
				events <- ContainerEvent{Type: "COMPRESSION_FAILED", Data: nil}

			case "COMPRESSION_STARTED":
				events <- ContainerEvent{Type: "COMPRESSION_STARTED", Data: nil}

			case "COMPRESSION_COMPLETED":
				events <- ContainerEvent{Type: "COMPRESSION_COMPLETED", Data: nil}
			default:
				if collectingJSON {
					jsonBuffer.WriteString(line)
					jsonBuffer.WriteString("\n")
				} else {
					events <- ContainerEvent{Type: "UNRECOGNIZED", Data: line}
				}
			}
		}
	}
}
