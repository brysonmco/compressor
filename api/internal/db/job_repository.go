package db

import (
	"context"
	"fmt"
	"github.com/brysonmco/compressor/internal/models"
)

func (d *Database) FindJobById(
	ctx context.Context,
	id int64,
) (*models.Job, error) {
	query := `SELECT id, user_id, created_at, updated_at, file_uploaded, file_name, status, input_codec, input_container, 
       input_resolution_horizontal, input_resolution_vertical, input_size, output_codec, output_container, 
       output_resolution_horizontal, output_resolution_vertical, output_size
		FROM jobs
		WHERE id = $1`

	row := d.Pool.QueryRow(ctx, query, id)

	var job models.Job
	if err := row.Scan(
		&job.Id,
		&job.UserId,
		&job.CreatedAt,
		&job.UpdatedAt,
		&job.FileUploaded,
		&job.FileName,
		&job.Status,
		&job.InputCodec,
		&job.InputContainer,
		&job.InputResolutionHorizontal,
		&job.InputResolutionVertical,
		&job.InputSize,
		&job.OutputCodec,
		&job.OutputContainer,
		&job.OutputResolutionHorizontal,
		&job.OutputResolutionVertical,
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
	query := `INSERT INTO jobs (user_id, file_name, input_container, input_size)
    		VALUES ($1, $2, $3, $4)
    		RETURNING id, user_id, created_at, updated_at, file_uploaded, file_name, status, input_container, input_size`

	var job models.Job
	if err := d.Pool.QueryRow(ctx, query,
		jobReq.UserId,
		jobReq.FileName,
		jobReq.InputContainer,
		jobReq.InputSize,
	).Scan(
		&job.Id,
		&job.UserId,
		&job.CreatedAt,
		&job.UpdatedAt,
		&job.FileUploaded,
		&job.FileName,
		&job.Status,
		&job.InputContainer,
		&job.InputSize,
	); err != nil {
		return nil, err
	}

	return &job, nil
}

func (d *Database) UpdateJob(
	ctx context.Context,
	job *models.Job,
) error {
	query := `UPDATE jobs 
		SET user_id = $1, created_at = $2, updated_at = $3, file_uploaded = $4, file_name = $5, status = $6, 
		    input_codec = $7, input_container = $8, input_resolution_horizontal = $9, input_resolution_vertical = $10,
		    input_size = $11, output_codec = $12, output_container = $13, output_resolution_horizontal = $14, 
		    output_resolution_vertical = $15, output_size = $16
		WHERE id = $17`

	cmdTag, err := d.Pool.Exec(ctx, query,
		job.UserId,
		job.CreatedAt,
		job.UpdatedAt,
		job.FileUploaded,
		job.FileName,
		job.Status,
		job.InputCodec,
		job.InputContainer,
		job.InputResolutionHorizontal,
		job.InputResolutionVertical,
		job.InputSize,
		job.OutputCodec,
		job.OutputContainer,
		job.OutputResolutionHorizontal,
		job.OutputResolutionVertical,
		job.OutputSize,
		job.Id,
	)
	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("could not update job")
	}
	return nil
}
