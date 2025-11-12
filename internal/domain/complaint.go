package domain

import "time"

type Complaint struct {
	ComplaintID int64
	UserID      int64
	BookID      int64
	Text        string
	CreatedAt   *time.Time
}
