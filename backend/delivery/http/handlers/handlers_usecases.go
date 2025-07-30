package handlers

import (
	application "social-network/application/usecases"
	"social-network/delivery/ws"
)

type Handler struct {
	UserUsecase         application.UserUsecase
	SessionUsecase      application.SessionUsecase
	PostUsecase         application.PostUsecase
	MessageUsecase      application.MessageUsecase
	FriendUsecase       application.FriendUsecase
	NotificationUsecase application.NotificationUsecase
	Hub                 *ws.Hub
}
