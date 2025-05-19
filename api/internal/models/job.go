package models

type Job struct {
	Id              int64  `json:"id"`
	UserId          int64  `json:"user_id"`
	CreatedAt       string `json:"created_at"`
	LastModified    string `json:"last_modified"`
	InputCodec      string `json:"input_codec"`
	InputContainer  string `json:"input_container"`
	InputSize       int64  `json:"input_size"`
	OutputCodec     string `json:"output_codec"`
	OutputContainer string `json:"output_container"`
	OutputSize      int64  `json:"output_size"`
}

type CreateJob struct {
	UserId          int64  `json:"user_id"`
	InputCodec      string `json:"input_codec"`
	InputContainer  string `json:"input_container"`
	InputSize       int64  `json:"input_size"`
	OutputCodec     string `json:"output_codec"`
	OutputContainer string `json:"output_container"`
	OutputSize      int64  `json:"output_size"`
}
