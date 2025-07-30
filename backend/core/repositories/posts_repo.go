package repositories

import (
	"context"

	entities "social-network/core/entities"
)

type PostRepo interface {
	Save(ctx context.Context, input *entities.Post) error
	Load(ctx context.Context, id string) ([]*entities.Post,string, error)
}
