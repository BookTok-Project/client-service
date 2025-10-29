package repo

import (
	"context"

	"client-service/internal/db/pg"
)

type FavoritesRepo struct {
	db pg.Conn
}

func NewFavoritesRepo(db pg.Conn) *FavoritesRepo {
	return &FavoritesRepo{db: db}
}

func (r *FavoritesRepo) Add(ctx context.Context, userID, bookID int64) error {
	return r.db.Queries(ctx).AddFavoriteBook(ctx, pg.AddFavoriteBookParams{
		UserID: userID,
		BookID: bookID,
	})
}

func (r *FavoritesRepo) Remove(ctx context.Context, userID, bookID int64) error {
	return r.db.Queries(ctx).RemoveFavoriteBook(ctx, pg.RemoveFavoriteBookParams{
		UserID: userID,
		BookID: bookID,
	})
}

func (r *FavoritesRepo) List(ctx context.Context, userID int64) ([]int64, error) {
	return r.db.Queries(ctx).ListFavoriteBooks(ctx, userID)
}
