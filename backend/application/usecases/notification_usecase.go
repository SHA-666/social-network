package application

import (
	"context"
	"fmt"

	"social-network/core/entities"
	repo "social-network/core/repositories"
)

type NotificationUsecase struct {
	repo repo.NotificationRepo
}

func NewNotificatoinUsecase(repo repo.NotificationRepo) NotificationUsecase {
	return NotificationUsecase{
		repo: repo,
	}
}

func (n *NotificationUsecase) GetAllNotif(ctx context.Context, id string) ([]entities.Notification, error) {
	res, err := n.repo.GetNotifications(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (n *NotificationUsecase) FriendRequest(ctx context.Context, from string, to string, toName string) error {
	message := fmt.Sprintf("%s vous a envoyé une demande d'ami", toName)
	fmt.Println("message", message)
	err := n.repo.NotifUserFriends(ctx, from, to, "friend request", message)
	if err != nil {
		return err
	}
	return nil
}

func (n *NotificationUsecase) RemoveNotif(ctx context.Context, id string) error {
	err := n.repo.DeleteNotif(ctx, id)
	if err != nil {
		fmt.Println("errrrlkfbgd", err)

		return err
	}
	return nil
}
