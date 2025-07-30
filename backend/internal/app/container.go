package internal

import (
	"database/sql"

	usecases "social-network/application/usecases"
	infrastructure "social-network/infrastructure/repositories"
)

type App struct {
	UserUsecase         usecases.UserUsecase
	PostUsescase        usecases.PostUsecase
	MessageUsecase      usecases.MessageUsecase
	FriendUsecase       usecases.FriendUsecase
	NotificationUsecase usecases.NotificationUsecase
	SessionUsecase      usecases.SessionUsecase
}

func NewApp(db *sql.DB) *App {
	return &App{
		UserUsecase:         usecases.NewUserUsecase(infrastructure.NewUserRepo(db)),
		PostUsescase:        usecases.NewPostUsecase(infrastructure.NewPostRepo(db)),
		MessageUsecase:      usecases.NewMessageUsecase(infrastructure.NewMessageRepo(db)),
		FriendUsecase:       usecases.NewFriendUsecase(infrastructure.NewFriendRepo(db)),
		NotificationUsecase: usecases.NewNotificatoinUsecase(infrastructure.NewNotificationRepo(db)),
		SessionUsecase:      usecases.NewSessionUsecase(infrastructure.NewSessionRepo(db)),
	}
}
