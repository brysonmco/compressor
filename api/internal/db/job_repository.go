package db

import (
	"context"
	"github.com/awesomebfm/compressor/internal/models"
)

func (d *Database) FindJobById(
	ctx context.Context,
	id int64,
) (*models.Job, error) {
	return nil, nil
}

func (d *Database) FindJobsByUserId(
	ctx context.Context,
	userId int64,
) ([]*models.Job, error) {
	return nil, nil
}

func (d *Database) CreateJob(
	ctx context.Context,
	jobReq *models.CreateJob,
) (*models.Job, error) {
	return nil, nil
}

func (d *Database) UpdateJob(
	ctx context.Context,
	jobReq *models.Job,
) (*models.Job, error) {
	return nil, nil
}
