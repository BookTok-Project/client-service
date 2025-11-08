package dto

import (
	"client-service/internal/db/pg"
	"client-service/internal/domain"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func ConvertFromPgTimestamptz(ts pgtype.Timestamptz) *time.Time {
	if !ts.Valid {
		return nil
	}

	return &ts.Time
}

func ConvertToDomainComment(comment pg.CommentsBook) domain.Comment {
	return domain.Comment{
		ID:        comment.ID,
		UserID:    comment.UserID,
		BookID:    comment.BookID,
		Text:      comment.Text,
		CreatedAt: ConvertFromPgTimestamptz(comment.CreatedAt),
	}
}

func ConvertToDomainComments(pgComments []pg.CommentsBook) []domain.Comment {
	domainComments := make([]domain.Comment, 0, len(pgComments))
	for _, pgComment := range pgComments {
		domainComments = append(domainComments, ConvertToDomainComment(pgComment))
	}

	return domainComments
}
