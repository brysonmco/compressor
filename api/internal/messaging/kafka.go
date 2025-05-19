package messaging

import (
	"encoding/json"
	"fmt"
)

type KafkaMessage struct {
	Event   string          `json:"event"`
	JobId   int64           `json:"job_id"`
	Payload json.RawMessage `json:"payload"`
}

type KafkaService struct {
}

func NewKafkaService() *KafkaService {
	return &KafkaService{}
}

func (k *KafkaService) SendNewJobMessage(jobId int64, downloadUrl string) error {
	payload := map[string]string{
		"download_url": downloadUrl,
	}
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	message, err := json.Marshal(KafkaMessage{
		Event:   "new_job",
		JobId:   jobId,
		Payload: payloadJson,
	})
	if err != nil {
		return err
	}

	fmt.Println(string(message))
	// TODO: Send this message
	return nil
}
