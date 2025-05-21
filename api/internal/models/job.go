package models

import "time"

type Job struct {
	Id                         int64     `json:"id"`
	UserId                     int64     `json:"userId"`
	CreatedAt                  time.Time `json:"createdAt"`
	UpdatedAt                  time.Time `json:"updatedAt"`
	FileUploaded               bool      `json:"fileUploaded"`
	FileName                   string    `json:"fileName"`
	Status                     string    `json:"status"`
	InputCodec                 string    `json:"inputCodec"`
	InputContainer             string    `json:"inputContainer"`
	InputResolutionHorizontal  int       `json:"inputResolutionHorizontal"`
	InputResolutionVertical    int       `json:"inputResolutionVertical"`
	InputSize                  int64     `json:"inputSize"`
	OutputCodec                string    `json:"outputCodec"`
	OutputContainer            string    `json:"output_container"`
	OutputResolutionHorizontal int       `json:"outputResolutionHorizontal"`
	OutputResolutionVertical   int       `json:"outputResolutionVertical"`
	OutputSize                 int64     `json:"output_size"`
}

type CreateJob struct {
	UserId         int64  `json:"userId"`
	FileName       string `json:"fileName"`
	InputContainer string `json:"inputContainer"`
	InputSize      int64  `json:"inputSize"`
}
