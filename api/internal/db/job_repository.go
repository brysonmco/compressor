package db

import (
	"context"
	"github.com/awesomebfm/compressor/internal/models"
)

func (d *Database) FindJobById(
	ctx context.Context,
	id int64,
) (*models.Job, error) {
	query := `SELECT id, user_id, created_at, updated_at, input_codec, input_container, input_size, output_codec, 
       output_container, output_size 
		FROM jobs
		WHERE id = $1`

	row := d.Pool.QueryRow(ctx, query, id)

	var job models.Job
	if err := row.Scan(
		&job.Id,
		&job.UserId,
		&job.CreatedAt,
		&job.UpdatedAt,
		&job.InputCodec,
		&job.InputContainer,
		&job.InputSize,
		&job.OutputCodec,
		&job.OutputContainer,
		&job.OutputSize,
	); err != nil {
		return nil, err
	}

	return &job, nil
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
	query := `INSERT INTO jobs (user_id, input_codec, input_container, input_size, output_codec, output_container, 
                  output_size)
    		VALUES ($1, $2, $3, $4, $5, $6, $7)
    		RETURNING id, user_id, created_at, updated_at, input_codec, input_container, input_size, output_codec,
    		    output_container, output_size`

	var job models.Job
	if err := d.Pool.QueryRow(ctx, query,
		jobReq.UserId,
		jobReq.InputCodec,
		jobReq.InputContainer,
		jobReq.InputSize,
		jobReq.OutputCodec,
		jobReq.OutputContainer,
		jobReq.OutputSize,
	).Scan(
		&job.Id,
		&job.UserId,
		&job.CreatedAt,
		&job.UpdatedAt,
		&job.InputCodec,
		&job.InputContainer,
		&job.InputSize,
		&job.OutputCodec,
		&job.OutputContainer,
		&job.OutputSize,
	); err != nil {
		return nil, err
	}

	return &job, nil
}

func (d *Database) UpdateJob(
	ctx context.Context,
	jobReq *models.Job,
) (*models.Job, error) {
	return nil, nil
}
