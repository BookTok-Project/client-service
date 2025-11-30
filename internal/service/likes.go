package service

import (
	"context"

	repoPkg "client-service/internal/repo"
)

type LikesService struct {
	repo repoPkg.LikesRepo
}

func NewLikesService(r repoPkg.LikesRepo) *LikesService {
	return &LikesService{repo: r}
}

func (s *LikesService) AddLike(ctx context.Context, userID, bookID int64) error {
	return s.repo.Add(ctx, userID, bookID)
}

func (s *LikesService) RemoveLike(ctx context.Context, userID, bookID int64) error {
	return s.repo.Remove(ctx, userID, bookID)
}

func (s *LikesService) ListLike(ctx context.Context, userID int64) ([]int64, error) {
	return s.repo.List(ctx, userID)
}
