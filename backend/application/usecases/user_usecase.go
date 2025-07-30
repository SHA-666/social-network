package application

import (
	"context"
	"strings"

	entities "social-network/core/entities"
	core "social-network/core/repositories"
)

type UserUsecase struct {
	repo core.UserRepo
}

func NewUserUsecase(repo core.UserRepo) UserUsecase {
	return UserUsecase{
		repo: repo,
	}
}
func (u *UserUsecase) Register(ctx context.Context, user *entities.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	return u.repo.Save(ctx, user)
}

func (u *UserUsecase) Login(ctx context.Context, userData *entities.User) (int, error) {
	email := strings.TrimSpace(strings.ToLower(userData.Email))
	id, err := u.repo.Connect(ctx, email, userData.Password)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *UserUsecase) UserName(ctx context.Context, user string) (string) {
	name, _ := u.repo.GetName(ctx,user)
	
	return name
}

func (u *UserUsecase) ReceiverName(ctx context.Context, receiverId string) (string, error) {
	name, err := u.repo.GetName(ctx,receiverId)
	if err != nil {
		return "", err
	}
	return name, nil
}

func (u *UserUsecase) GetInfo(ctx context.Context, id string) (entities.User, bool) {
	res := u.repo.Get(ctx, id)
	if res.ID == 0 {
		return entities.User{}, false
	}
	return res, true
}
