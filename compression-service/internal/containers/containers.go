package containers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
	"os"
	"time"
)

type Service struct {
	Client      *client.Client
	WorkerImage string
	Containers  []Container
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
) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	containerName := fmt.Sprintf("worker-%d", jobId)

	resp, err := s.Client.ContainerCreate(ctx, &container.Config{
		Image: s.WorkerImage,
	}, nil, nil, nil, containerName)
	if err != nil {
		return fmt.Errorf("error creating container: %v", err)
	}

	s.Containers = append(s.Containers, Container{
		Id:    resp.ID,
		JobId: jobId,
	})

	err = s.Client.ContainerStart(ctx, resp.ID, container.StartOptions{})
	if err != nil {
		return fmt.Errorf("error starting container: %v", err)
	}

	return nil
}
