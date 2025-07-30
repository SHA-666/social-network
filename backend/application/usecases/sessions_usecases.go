package application

import (
	"context"

	entities "social-network/core/entities"
	repo "social-network/core/repositories"
)

type SessionUsecase struct {
	repo repo.SessionRepo
}

func NewSessionUsecase(repo repo.SessionRepo) SessionUsecase {
	return SessionUsecase{
		repo: repo,
	}
}

func (s *SessionUsecase) GetIdOnly(ctx context.Context, token string) string {
	id, err := s.repo.GetID(ctx, token)
	if err != nil {
		return ""
	}
	return id
}

func (s *SessionUsecase) CreateSession(ctx context.Context, session *entities.Session) error {
	return s.repo.Save(ctx, session)
}

func (s *SessionUsecase) DeleteSession(ctx context.Context, token string) error {
	return s.repo.Delete(ctx, token)
}
