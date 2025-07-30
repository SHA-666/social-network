package repositories

import (
	"context"

	entities "social-network/core/entities"
)

type SessionRepo interface {
	Save(ctx context.Context, s *entities.Session) error
	Delete(ctx context.Context, id string) error
	GetID(ctx context.Context, token string) (string, error)
}
