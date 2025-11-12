package repo

import (
	"client-service/internal/db/pg"
	"context"
)

func (r *Repo) AddComplaint(ctx context.Context, userID, bookID int64, text string) error {
	return r.conn.Queries(ctx).AddComplaint(ctx, pg.AddComplaintParams{
		UserID: userID,
		BookID: bookID,
		Text:   text,
	})
}

func (r *Repo) GetComplaints(ctx context.Context) ([]pg.Complaint, error) {
	return r.conn.Queries(ctx).GetComplaints(ctx)
}
