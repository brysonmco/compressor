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

// GenerateUploadURLForUploads generates a pre-signed URL for the client to upload an uncompressed file.
func (s *Storage) GenerateUploadURLForUploads(
	ctx context.Context,
	id int64,
	expires time.Time,
) (string, error) {
	return "", fmt.Errorf("not implemented")
}

// GenerateUploadURLForDownloads generates a pre-signed URL for the VM to upload a compressed file.
func (s *Storage) GenerateUploadURLForDownloads(
	ctx context.Context,
	id int64,
	expires time.Time,
) (string, error) {
	return "", fmt.Errorf("not implemented")
}

// GenerateDownloadURLForUploads generates a pre-signed URL for the VM to download an uncompressed file.
func (s *Storage) GenerateDownloadURLForUploads(
	ctx context.Context,
	id int64,
	expires time.Time,
) (string, error) {
	return "", fmt.Errorf("not implemented")
}

// GenerateDownloadURLForDownloads generates a pre-signed URL for the client to download a compressed file.
func (s *Storage) GenerateDownloadURLForDownloads(
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
