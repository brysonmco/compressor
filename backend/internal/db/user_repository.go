package db

import (
	"context"
	"github.com/awesomebfm/compressor/internal/models"
)

func (d *Database) FindUserByEmail(
	ctx context.Context,
	email string,
) (*models.User, error) {
	query := `SELECT id, email, first_name, last_name, password_hash
		FROM users
		WHERE email = $1`

	row := d.Pool.QueryRow(ctx, query, email)

	var user models.User
	if err := row.Scan(
		&user.Id,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.PasswordHash,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

func (d *Database) FindUserByID(
	ctx context.Context,
	id int64,
) (*models.User, error) {
	query := `SELECT id, email, first_name, last_name, password_hash
		FROM users
		WHERE id = $1`

	row := d.Pool.QueryRow(ctx, query, id)

	var user models.User
	if err := row.Scan(
		&user.Id,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.PasswordHash,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

func (d *Database) CreateUser(
	ctx context.Context,
	userReq *models.CreateUser,
) (*models.User, error) {
	query := `INSERT INTO users (email, first_name, last_name, password_hash)
		VALUES ($1, $2, $3, $4)
		RETURNING id, email, first_name, last_name, password_hash`

	var user models.User
	err := d.Pool.QueryRow(ctx, query,
		userReq.Email,
		userReq.FirstName,
		userReq.LastName,
		userReq.PasswordHash,
	).Scan(
		&user.Id,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.PasswordHash,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
