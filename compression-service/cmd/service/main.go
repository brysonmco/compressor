package service

import (
	"github.com/brysonmco/compressor/compression-service/internal/containers"
	"log"
)

func main() {
	// Container service
	containerService := containers.NewService()
	err := containerService.InitializeClient()
	if err != nil {
		log.Fatal(err)
	}

}
