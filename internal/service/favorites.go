package service

import (
	"context"

	repoPkg "client-service/internal/repo"
)

type FavoritesService struct {
	repo repoPkg.FavoritesRepo
}

func NewFavoritesService(r repoPkg.FavoritesRepo) *FavoritesService {
	return &FavoritesService{repo: r}
}

func (s *FavoritesService) AddFavorite(ctx context.Context, userID, bookID int64) error {
	return s.repo.Add(ctx, userID, bookID)
}

func (s *FavoritesService) RemoveFavorite(ctx context.Context, userID, bookID int64) error {
	return s.repo.Remove(ctx, userID, bookID)
}

func (s *FavoritesService) ListFavorites(ctx context.Context, userID int64) ([]int64, error) {
	return s.repo.List(ctx, userID)
}
