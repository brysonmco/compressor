package main

import (
	"github.com/brysonmco/compressor/compression-service/internal/containers"
	"github.com/brysonmco/compressor/compression-service/internal/messaging"
	"log"
	"os"
)

func main() {
	// Container service
	containerService := containers.NewService()
	err := containerService.InitializeClient()
	if err != nil {
		log.Fatal(err)
	}
	defer containerService.CloseClient()

	// Messaging Service
	messagingService := messaging.NewService()
	err = messagingService.Connect(
		os.Getenv("RABBIT_USERNAME"),
		os.Getenv("RABBIT_PASSWORD"),
		os.Getenv("RABBIT_HOST"),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer messagingService.Close()

}
