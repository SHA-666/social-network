package repositories

import (
	"context"

	entities "social-network/core/entities"
)

type FriendsRepo interface {
	Add(ctx context.Context, userId, receiverId string) error
	Load(ctx context.Context, userId string) ([]entities.Friends, error)
	Get(ctx context.Context, userId string) ([]entities.Friends, error)
	Delete(ctx context.Context, userId, receiver string) error
	Accepte(ctx context.Context, userId, receiver string) error
}
