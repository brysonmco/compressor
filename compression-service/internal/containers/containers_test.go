package containers

import (
	"context"
	"testing"
)

func TestCreateContainer(t *testing.T) {
	t.Setenv("DEPLOYMENT_TARGET", "development")

	containerService := NewService("worker:latest")
	err := containerService.InitializeClient()
	if err != nil {
		t.Errorf("Failed to initialize client: %v", err)
		return
	}
	defer containerService.CloseClient()

	container, err := containerService.NewContainer(32)
	if err != nil {
		t.Errorf("Failed to create container: %v", err)
		return
	}

	if container.Id == "" {
		t.Error("Container ID should not be empty")
		return
	}
	if container.JobId != 32 {
		t.Errorf("Expected JobId 32, got %d", container.JobId)
		return
	}

	err = containerService.RemoveContainer(container.Id)
	if err != nil {
		t.Errorf("Failed to remove container: %v", err)
		return
	}
}

func TestMonitorContainer(t *testing.T) {
	t.Setenv("DEPLOYMENT_TARGET", "development")

	containerService := NewService("worker:latest")
	err := containerService.InitializeClient()
	if err != nil {
		t.Errorf("Failed to initialize client: %v", err)
		return
	}
	defer containerService.CloseClient()

	container, err := containerService.NewContainer(32)
	if err != nil {
		t.Errorf("Failed to create container: %v", err)
		return
	}

	events := make(chan ContainerEvent)

	go func() {
		if err = containerService.MonitorOutput(context.TODO(), container.Id, events); err != nil {
			t.Errorf("Failed to monitor container output: %v", err)
			return
		}
	}()

	gotStarted := false

	for event := range events {
		if event.Type == "APPLICATION_STARTED" {
			gotStarted = true
			close(events)
			err = containerService.RemoveContainer(container.Id)
			if err != nil {
				t.Errorf("Failed to remove container after monitoring: %v", err)
				return
			}
			break
		}
	}

	if !gotStarted {
		t.Error("Expected APPLICATION_STARTED event, but did not receive it")
		return
	}
}
