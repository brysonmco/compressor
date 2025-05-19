package models

import "time"

type Job struct {
	Id              int64     `json:"id"`
	UserId          int64     `json:"userId"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	InputCodec      string    `json:"inputCodec"`
	InputContainer  string    `json:"inputContainer"`
	InputSize       string    `json:"inputSize"`
	OutputCodec     string    `json:"outputCodec"`
	OutputContainer string    `json:"output_container"`
	OutputSize      string    `json:"output_size"`
}

type CreateJob struct {
	UserId          int64  `json:"userId"`
	InputCodec      string `json:"inputCodec"`
	InputContainer  string `json:"inputContainer"`
	InputSize       string `json:"inputSize"`
	OutputCodec     string `json:"outputCodec"`
	OutputContainer string `json:"outputContainer"`
	OutputSize      string `json:"outputSize"`
}
