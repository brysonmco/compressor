package db

import (
	"context"
	"fmt"
	"github.com/awesomebfm/compressor/internal/models"
)

func (d *Database) FindUserByEmail(
	ctx context.Context,
	email string,
) (*models.User, error) {
	query := `SELECT id, email, first_name, last_name, password_hash, stripe_customer_id, email_verified, created_at, 
       updated_at, last_login
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
		&user.StripeCustomerId,
		&user.EmailVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.LastLogin,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

func (d *Database) FindUserByID(
	ctx context.Context,
	id int64,
) (*models.User, error) {
	query := `SELECT id, email, first_name, last_name, password_hash, stripe_customer_id, email_verified, created_at, 
       updated_at, last_login
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
		&user.StripeCustomerId,
		&user.EmailVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.LastLogin,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

func (d *Database) FindUserByStripeCustomerID(
	ctx context.Context,
	stripeCustomerId string,
) (*models.User, error) {
	query := `SELECT id, email, first_name, last_name, password_hash, stripe_customer_id, email_verified, created_at, 
       updated_at, last_login
		FROM users
		WHERE stripe_customer_id = $1`

	row := d.Pool.QueryRow(ctx, query, stripeCustomerId)

	var user models.User
	if err := row.Scan(
		&user.Id,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.PasswordHash,
		&user.StripeCustomerId,
		&user.EmailVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.LastLogin,
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
		RETURNING id, email, first_name, last_name, password_hash, email_verified, created_at, 
		    updated_at, last_login`

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
		&user.EmailVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.LastLogin,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (d *Database) UpdateUser(
	ctx context.Context,
	user *models.User,
) error {
	query := `UPDATE users 
		SET email = $1, first_name = $2, last_name = $3, password_hash = $4, stripe_customer_id = $5, created_at = $6,
		    updated_at = $7, last_login = $8
		WHERE id = $9`

	cmdTag, err := d.Pool.Exec(ctx, query,
		user.Email,
		user.FirstName,
		user.LastName,
		user.PasswordHash,
		user.StripeCustomerId,
		user.CreatedAt,
		user.UpdatedAt,
		user.LastLogin,
		user.Id,
	)
	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("could not update user")
	}
	return err
}
