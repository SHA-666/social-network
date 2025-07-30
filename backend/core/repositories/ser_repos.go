package repositories

import (
	"context"

	entities "social-network/core/entities"
)

type UserRepo interface {
	// register
	Save(ctx context.Context, user *entities.User) error
	// login
	Connect(ctx context.Context, input, password string) (int, error)
	Get(ctx context.Context, id string) entities.User
	GetName(ctx context.Context, id string) (string,error)

	Hash(password string) string
	Compare(hashed string, plain string) error
}
