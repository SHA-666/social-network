package entities

import "time"

type Session struct {
	Token     string //  UUID
	UserID    int
	CreatedAt time.Time
	ExpiresAt time.Time
}
