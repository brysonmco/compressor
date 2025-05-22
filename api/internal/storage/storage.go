package storage

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"time"
)

type Storage struct {
	Client        *minio.Client
	UploadsBucket string
}

func NewStorage(
	uploadsBucket string,
	endpoint string,
	accessKey string,
	secretKey string,
	secure bool,
) (*Storage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: secure,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create S3 client: %v", err)
	}

	// Check if the bucket exists
	bucketExists, err := client.BucketExists(ctx, uploadsBucket)
	if err != nil {
		return nil, fmt.Errorf("failed to check if bucket exists: %v", err)
	} else if !bucketExists {
		return nil, fmt.Errorf("bucket does not exist")
	}

	return &Storage{
		Client:        client,
		UploadsBucket: uploadsBucket,
	}, nil
}

// GenerateUploadURLForUploads generates a pre-signed URL for the client to upload an uncompressed file.
func (s *Storage) GenerateUploadURLForUploads(
	ctx context.Context,
	id int64,
	fileType string,
	expires time.Time,
	maxFileSize int64,
) (string, map[string]string, error) {
	policy := minio.NewPostPolicy()

	err := policy.SetBucket(s.UploadsBucket)
	if err != nil {
		return "", nil, err
	}
	err = policy.SetKey(fmt.Sprintf("%d.%s", id, fileType))
	if err != nil {
		return "", nil, err
	}
	err = policy.SetContentLengthRange(0, maxFileSize)
	if err != nil {
		return "", nil, err
	}
	err = policy.SetExpires(expires)
	if err != nil {
		return "", nil, err
	}

	url, formData, err := s.Client.PresignedPostPolicy(ctx, policy)
	if err != nil {
		return "", nil, err
	}
	return url.String(), formData, nil
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
	extension string,
) (bool, error) {
	_, err := s.Client.GetObject(ctx, s.UploadsBucket, fmt.Sprintf("%d.%v", id, extension), minio.GetObjectOptions{})
	if err != nil && minio.ToErrorResponse(err).Code == "NoSuchKey" {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func (s *Storage) FileInDownloads(
	ctx context.Context,
	id int64,
) (bool, error) {
	return false, fmt.Errorf("not implemented")
}
