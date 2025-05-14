package db

import (
	"context"
	"github.com/awesomebfm/compressor/internal/models"
)

func (d *Database) FindSessionByID(
	ctx context.Context,
	id int64,
) (*models.Session, error) {
	query := `SELECT id, token_hash, user_id, expires_at, revoked, created_at
			FROM sessions
			WHERE id = $1`

	row := d.Pool.QueryRow(ctx, query, id)

	var session models.Session
	if err := row.Scan(
		&session.Id,
		&session.TokenHash,
		&session.UserId,
		&session.ExpiresAt,
		&session.Revoked,
		&session.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &session, nil
}

func (d *Database) FindSessionByTokenHash(
	ctx context.Context,
	tokenHash string,
) (*models.Session, error) {
	query := `SELECT id, token_hash, user_id, expires_at, revoked, created_at
			FROM sessions
			WHERE token_hash = $1`

	row := d.Pool.QueryRow(ctx, query, tokenHash)

	var session models.Session
	if err := row.Scan(
		&session.Id,
		&session.TokenHash,
		&session.UserId,
		&session.ExpiresAt,
		&session.Revoked,
		&session.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &session, nil
}

func (d *Database) CreateSession(
	ctx context.Context,
	sessionReq *models.CreateSession,
) (*models.Session, error) {
	query := `INSERT INTO sessions (token_hash, user_id, expires_at, created_at)
			VALUES ($1, $2, $3, $4)
			RETURNING id, token_hash, user_id, expires_at, revoked, created_at`

	var session models.Session
	err := d.Pool.QueryRow(ctx, query,
		sessionReq.TokenHash,
		sessionReq.UserId,
		sessionReq.ExpiresAt,
		sessionReq.CreatedAt,
	).Scan(
		&session.Id,
		&session.TokenHash,
		&session.UserId,
		&session.ExpiresAt,
		&session.Revoked,
		&session.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (d *Database) RevokeSession(
	ctx context.Context,
	id int64,
) error {
	query := `UPDATE sessions SET revoked = true WHERE id = $1`
	_, err := d.Pool.Exec(ctx, query, id)
	return err
}
