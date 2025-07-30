package repositories

import (
	"context"
	"database/sql"

	entities "social-network/core/entities"
	repositories "social-network/core/repositories"
)

type SessionRepo struct {
	db *sql.DB
}

func NewSessionRepo(db *sql.DB) repositories.SessionRepo {
	return &SessionRepo{db: db}
}

func (r *SessionRepo) Save(ctx context.Context, s *entities.Session) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO sessions (token, user_id, created_at, expires_at)
		VALUES (?, ?, ?, ?)`,
		s.Token, s.UserID, s.CreatedAt, s.ExpiresAt)
	return err
}

func (r *SessionRepo) Delete(ctx context.Context, token string) error {
	_, err := r.db.Exec(`DELETE FROM sessions WHERE token = ?`, token)
	return err
}

func (r *SessionRepo) GetID(ctx context.Context, token string) (string, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT user_id FROM sessions WHERE token = ?`, token)
	var UserID string
	err := row.Scan(&UserID)
	if err != nil {
		return "", err
	}
	return UserID, nil
}

