package repo

import (
	"context"

	"client-service/internal/db/pg"
)

type LikesRepo struct {
	db pg.Conn
}

func NewLikesRepo(db pg.Conn) *LikesRepo {
	return &LikesRepo{db: db}
}

func (r *LikesRepo) Add(ctx context.Context, userID, bookID int64) error {
	return r.db.Queries(ctx).AddLikeBook(ctx, pg.AddLikeBookParams{
		UserID: userID,
		BookID: bookID,
	})
}

func (r *LikesRepo) Remove(ctx context.Context, userID, bookID int64) error {
	return r.db.Queries(ctx).RemoveLikeBook(ctx, pg.RemoveLikeBookParams{
		UserID: userID,
		BookID: bookID,
	})
}

func (r *LikesRepo) List(ctx context.Context, userID int64) ([]int64, error) {
	return r.db.Queries(ctx).ListLikeBooks(ctx, userID)
}
