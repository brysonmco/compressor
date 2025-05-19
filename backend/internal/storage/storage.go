package storage

import (
	"context"
	"fmt"
	"time"
)

type Storage struct {
}

func NewStorage() (*Storage, error) {
	return &Storage{}, nil
}

func (s *Storage) GenerateUploadURL(
	ctx context.Context,
	id int64,
	expires time.Time,
) (string, error) {
	return "", fmt.Errorf("not implemented")
}

func (s *Storage) GenerateDownloadURL(
	ctx context.Context,
	id int64,
	expires time.Time,
) (string, error) {
	return "", fmt.Errorf("not implemented")
}

func (s *Storage) FileInUploads(
	ctx context.Context,
	id int64,
) (bool, error) {
	return false, fmt.Errorf("not implemented")
}

func (s *Storage) FileInDownloads(
	ctx context.Context,
	id int64,
) (bool, error) {
	return false, fmt.Errorf("not implemented")
}
