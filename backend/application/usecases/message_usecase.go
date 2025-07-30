package application

import (
	"context"
	"fmt"

	entities "social-network/core/entities"
	repo "social-network/core/repositories"
)

type MessageUsecase struct {
	repo repo.MessageRepo
}

func NewMessageUsecase(repo repo.MessageRepo) MessageUsecase {
	return MessageUsecase{
		repo: repo,
	}
}

func (m *MessageUsecase) SaveMessage(from,to,content string) error {
	fmt.Println("from : ",from,"to",to,"content",content)
	err := m.repo.Save(from,to,content)
	if err != nil {
		fmt.Println("errrrrrrrrrrrrrr",err)
		return err
	}
	return nil
}

func (m *MessageUsecase) LoadConv(ctx context.Context, userId, to string) ([]entities.Message, error) {
	fmt.Println("okokokokokok    user id = ", userId, "  //   friend id = ", to)
	res, err := m.repo.Load(userId, to)
	if err != nil {
		fmt.Println("pkokpcdfoinjdvonozbfvon",err)
		return nil, err
	}
	return res, nil
}

