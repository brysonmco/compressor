package compression

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/brysonmco/compressor/compression-service/internal/containers"
	"log"
	"net/http"
)

type Service struct {
	ContainerService *containers.Service
}

func NewService() *Service {
	return &Service{}
}

type fFProbeStream struct {
	CodecName  string `json:"codec_name"`
	CodecType  string `json:"codec_type"`
	Width      int    `json:"width,omitempty"`
	Height     int    `json:"height,omitempty"`
	SampleRate string `json:"sample_rate,omitempty"`
}

type fFProbeFormat struct {
	Filename   string `json:"filename"`
	NbStreams  int    `json:"nb_streams"`
	FormatName string `json:"format_name"`
}

type fFProbeOutput struct {
	Streams []fFProbeStream `json:"streams"`
	Format  fFProbeFormat   `json:"format"`
}

func (s *Service) HandleNewJob(
	jobId int64,
	downloadUrl string,
) {
	var err error

	// Create a new container, retrying up to 3 times if it fails
	var container *containers.Container
	for i := 0; i < 3; i++ {
		container, err = s.ContainerService.NewContainer(jobId)
		if err == nil {
			break
		}
		if container != nil {
			// If we have a container, remove it before retrying
			if err := s.ContainerService.RemoveContainer(container.Id); err != nil {
				log.Printf("error removing container %s: %v", container.Id, err)
			}
		}

		if i == 2 {
			log.Printf("failed to create container for job %d after 3 attempts: %v", jobId, err)
			return
		}
	}

	if container == nil {
		log.Printf("failed to create container for job %d after 3 attempts: %v", jobId, err)
		return
	}

	events := make(chan containers.ContainerEvent)

	go func() {
		if err := s.ContainerService.MonitorOutput(context.TODO(), container.Id, events); err != nil {
			log.Printf("error monitoring output for container %s: %v", container.Id, err)
			return
		}
	}()

	for event := range events {
		switch event.Type {
		case "APPLICATION_STARTED":
			// Send download URL to the container
			// TODO: This needs error handling and retries
			body := map[string]string{
				"url":       downloadUrl,
				"container": "mp4",
			}
			bodyBytes, err := json.Marshal(body)
			if err != nil {
				// IRDK
			}

			r, err := http.NewRequest("POST", fmt.Sprintf("http://localhost:%d/download", container.Port), bytes.NewBuffer(bodyBytes))
			if err != nil {
				// IRDK
			}
			r.Header.Add("Content-Type", "application/json")
			client := &http.Client{}
			resp, err := client.Do(r)
			if err != nil {
				// IRDK
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusCreated {
				// IRDK
			}

		case "SERVER_FAILED":
			// Try again??
			// TODO: Implement

		case "DOWNLOAD_FAILED":
			// Try again??
			// TODO: Implement

		case "DOWNLOAD_COMPLETED":
			// Probe downloaded file
			// TODO: This needs error handling and retries
			r, err := http.NewRequest("POST", fmt.Sprintf("http://localhost:%d/probe", container.Port), nil)
			if err != nil {
				// IRDK
			}

			client := &http.Client{}
			resp, err := client.Do(r)
			if err != nil {
				// IRDK
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				// IRDK
			}

		case "PROBE_FAILED":
			// Try again??
			// TODO: Implement

		case "PROBE_DATA":
			var probeData fFProbeOutput
			data, ok := event.Data.(map[string]interface{})
			if !ok {
				// IRDK
			}
			dataBytes, err := json.Marshal(data)
			if err != nil {
				// IRDK
			}
			if err := json.Unmarshal(dataBytes, &probeData); err != nil {
				// IRDK
			}

			// TODO: Pass this data back to the api

		default:
			log.Printf("Unknown event type for container %s: %s", container.Id, event.Type)
		}
	}

}
