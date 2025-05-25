package models

import "time"

type Plan struct {
	Id              int64         `json:"id"`
	Name            string        `json:"name"`
	Tokens          int           `json:"tokens"`
	Priority        string        `json:"priority"`
	StripeProductId string        `json:"stripeProductId"`
	ConcurrentJobs  int           `json:"concurrentJobs"`
	MaxResolution   int64         `json:"maxResolution"`
	MaxFileSize     int64         `json:"maxFileSize"`
	FileRetention   time.Duration `json:"fileRetention"`
	Watermark       bool          `json:"watermark"`
}
