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
