package application

import (
	"context"

	entities "social-network/core/entities"
	repo "social-network/core/repositories"
)

type PostUsecase struct {
	repo repo.PostRepo
}

func NewPostUsecase(repo repo.PostRepo) PostUsecase {
	return PostUsecase{
		repo: repo,
	}
}

func (u *PostUsecase) CreatePost(ctx context.Context, p *entities.Post) error {
	err := u.repo.Save(ctx, p)
	if err != nil {
		return err
	}
	return nil
}

func (u *PostUsecase) LoadPost(ctx context.Context, id string) ([]*entities.Post, string, error) {
	res, userPic, err := u.repo.Load(ctx, id)
	if err != nil {
		return nil, "", err
	}
	return res, userPic, nil
}
