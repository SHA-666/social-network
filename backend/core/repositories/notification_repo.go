package repositories

import (
	"context"
	"social-network/core/entities"
)

type NotificationRepo interface {
	NotifUserFriends(ctx context.Context, from, to, types, message string) error
	GetNotifications(ctx context.Context, id string) ([]entities.Notification, error)
	DeleteNotif(ctx context.Context, id string) error
}
