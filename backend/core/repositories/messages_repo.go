package repositories

import (
	entities "social-network/core/entities"
)

type MessageRepo interface {
	Save(from, to, content string) error
	Load(user1, user2 string) ([]entities.Message, error)
}
