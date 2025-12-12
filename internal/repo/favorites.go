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

func (r *FavoritesRepo) Add(ctx context.Context, userID, cardID int64) error {
	return r.db.Queries(ctx).AddFavoriteCard(ctx, pg.AddFavoriteCardParams{
		UserID: userID,
		CardID: cardID,
	})
}

func (r *FavoritesRepo) Remove(ctx context.Context, userID, cardID int64) error {
	return r.db.Queries(ctx).RemoveFavoriteCard(ctx, pg.RemoveFavoriteCardParams{
		UserID: userID,
		CardID: cardID,
	})
}

func (r *FavoritesRepo) List(ctx context.Context, userID int64) ([]int64, error) {
	return r.db.Queries(ctx).ListFavoriteCards(ctx, userID)
}
