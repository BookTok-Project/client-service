package repo

import (
	"context"

	"client-service/internal/db/pg"
)

type FavoritesRepo interface {
	Add(ctx context.Context, userID, bookID int64) error
	Remove(ctx context.Context, userID, bookID int64) error
	List(ctx context.Context, userID int64) ([]int64, error)
}

type favoritesRepo struct {
	db pg.Conn
}

func NewFavoritesRepo(db pg.Conn) FavoritesRepo {
	return &favoritesRepo{db: db}
}

func (r *favoritesRepo) Add(ctx context.Context, userID, bookID int64) error {
	return r.db.Queries(ctx).AddFavoriteBook(ctx, pg.AddFavoriteBookParams{
		UserID: userID,
		BookID: bookID,
	})
}

func (r *favoritesRepo) Remove(ctx context.Context, userID, bookID int64) error {
	return r.db.Queries(ctx).RemoveFavoriteBook(ctx, pg.RemoveFavoriteBookParams{
		UserID: userID,
		BookID: bookID,
	})
}

func (r *favoritesRepo) List(ctx context.Context, userID int64) ([]int64, error) {
	return r.db.Queries(ctx).ListFavoriteBooks(ctx, userID)
}
