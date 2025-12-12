package repo

import (
	"context"

	"client-service/internal/db/pg"
)

type Repo struct {
	conn pg.Conn
}

func New(conn pg.Conn) *Repo {
	return &Repo{conn: conn}
}

func (r *Repo) GetSubscriberIDs(ctx context.Context, subscribeeID int64) ([]int64, error) {
	return r.conn.Queries(ctx).GetSubscriberIDs(ctx, subscribeeID)
}

func (r *Repo) GetSubscribeeIDs(ctx context.Context, subscriberID int64) ([]int64, error) {
	return r.conn.Queries(ctx).GetSubscribeeIDs(ctx, subscriberID)
}

func (r *Repo) Subscribe(ctx context.Context, subscriberID, subscribeeID int64) error {
	return r.conn.Queries(ctx).Subscribe(ctx, pg.SubscribeParams{
		SubscriberID: subscriberID,
		SubscribeeID: subscribeeID,
	})
}

func (r *Repo) AddFavorite(ctx context.Context, userID, cardID int64) error {
	return r.conn.Queries(ctx).AddFavoriteCard(ctx, pg.AddFavoriteCardParams{
		UserID: userID,
		CardID: cardID,
	})
}

func (r *Repo) RemoveFavorite(ctx context.Context, userID, cardID int64) error {
	return r.conn.Queries(ctx).RemoveFavoriteCard(ctx, pg.RemoveFavoriteCardParams{
		UserID: userID,
		CardID: cardID,
	})
}

func (r *Repo) ListFavorites(ctx context.Context, userID int64) ([]int64, error) {
	return r.conn.Queries(ctx).ListFavoriteCards(ctx, userID)
}

func (r *Repo) AddComment(ctx context.Context, userID, bookID int64, text string) error {
	return r.conn.Queries(ctx).InsertComment(ctx, pg.InsertCommentParams{
		UserID: userID,
		BookID: bookID,
		Text:   text,
	})
}

func (r *Repo) GetCommentsByBookID(ctx context.Context, page, limit, bookID int64) ([]pg.CommentsBook, error) {
	return r.conn.Queries(ctx).GetCommentsByBookId(ctx, pg.GetCommentsByBookIdParams{
		BookID: bookID,
		Limit:  int32(limit),
		Offset: int32(limit * page),
	})
}

func (r *Repo) GetCommentsByUserID(ctx context.Context, page, limit, userID int64) ([]pg.CommentsBook, error) {
	return r.conn.Queries(ctx).GetCommentsByUserId(ctx, pg.GetCommentsByUserIdParams{
		UserID: userID,
		Limit:  int32(limit),
		Offset: int32(limit * page),
	})
}

func (r *Repo) GetCountCommentsByBookID(ctx context.Context, bookID int64) (int64, error) {
	return r.conn.Queries(ctx).GetCountCommentsByBookId(ctx, bookID)
}

func (r *Repo) GetCountCommentsByUserID(ctx context.Context, userID int64) (int64, error) {
	return r.conn.Queries(ctx).GetCountCommentsByUserId(ctx, userID)
}

func (r *Repo) AddLike(ctx context.Context, userID, bookID int64) error {
	return r.conn.Queries(ctx).AddLikeBook(ctx, pg.AddLikeBookParams{
		UserID: userID,
		BookID: bookID,
	})
}

func (r *Repo) RemoveLike(ctx context.Context, userID, bookID int64) error {
	return r.conn.Queries(ctx).RemoveLikeBook(ctx, pg.RemoveLikeBookParams{
		UserID: userID,
		BookID: bookID,
	})
}

func (r *Repo) ListLike(ctx context.Context, userID int64) ([]int64, error) {
	return r.conn.Queries(ctx).ListLikeBooks(ctx, userID)
}
