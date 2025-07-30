package application

import (
	"context"
	"fmt"

	entities "social-network/core/entities"
	repo "social-network/core/repositories"
)

type FriendUsecase struct {
	repo repo.FriendsRepo
}

func NewFriendUsecase(repo repo.FriendsRepo) FriendUsecase {
	return FriendUsecase{
		repo: repo,
	}
}

func (f *FriendUsecase) Addfriend(ctx context.Context, userId, receiverId string) error {
	fmt.Printf("%s sender, //%s receiver \n", userId, receiverId)
	err := f.repo.Add(ctx, userId, receiverId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (f *FriendUsecase) Acceptefriend(ctx context.Context, userId, receiverId string) error {
	fmt.Printf("%s sender, //%s receiver \n", userId, receiverId)

	err := f.repo.Accepte(ctx,userId,receiverId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (f *FriendUsecase) LoadFriend(ctx context.Context, userId string) ([]entities.Friends, error) {
	frs, err := f.repo.Load(ctx, userId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return frs, nil
}

func (f *FriendUsecase) DeleteFriend(ctx context.Context, userId, receiver string) error {
	fmt.Printf("%s sender, //%s receiver \n", userId, receiver)
	err := f.repo.Delete(ctx, userId, receiver)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (f *FriendUsecase) GetFriend(ctx context.Context, userId string) ([]entities.Friends, error) {
	frs, err := f.repo.Get(ctx, userId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return frs, nil
}