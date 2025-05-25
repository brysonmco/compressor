package vms

import (
	"context"
)

type FirecrackerService struct {
}

func NewFirecrackerService() *FirecrackerService {
	return &FirecrackerService{}
}

// CreateVM creates a new virtual machine using Firecracker. Returns the port and an error if any.
func (f *FirecrackerService) CreateVM(
	ctx context.Context,
) (int, error) {
	return -1, nil
}

// DestroyVM destroys the virtual machine running on the specified port. Returns an error if any.
func (f *FirecrackerService) DestroyVM(ctx context.Context, port int) error {
	return nil
}
