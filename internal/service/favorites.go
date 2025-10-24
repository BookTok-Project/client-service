package service

import (
	"context"

	repoPkg "client-service/internal/repo"
)

type FavoritesService interface {
	AddFavorite(ctx context.Context, userID, bookID int64) error
	RemoveFavorite(ctx context.Context, userID, bookID int64) error
	ListFavorites(ctx context.Context, userID int64) ([]int64, error)
}

type favoritesService struct {
	repo repoPkg.FavoritesRepo
}

func NewFavoritesService(r repoPkg.FavoritesRepo) FavoritesService {
	return &favoritesService{repo: r}
}

func (s *favoritesService) AddFavorite(ctx context.Context, userID, bookID int64) error {
	return s.repo.Add(ctx, userID, bookID)
}

func (s *favoritesService) RemoveFavorite(ctx context.Context, userID, bookID int64) error {
	return s.repo.Remove(ctx, userID, bookID)
}

func (s *favoritesService) ListFavorites(ctx context.Context, userID int64) ([]int64, error) {
	return s.repo.List(ctx, userID)
}
