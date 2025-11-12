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

func ConvertToDomainComplaint(complaint pg.Complaint) domain.Complaint {
	return domain.Complaint{
		ComplaintID:		complaint.ID,
		UserID:    		complaint.UserID,
		BookID:    		complaint.BookID,
		Text:      		complaint.Text,
		CreatedAt: 		ConvertFromPgTimestamptz(complaint.CreatedAt),
	}
}

func ConvertToDomainComplaints(pgComplaints []pg.Complaint) []domain.Complaint {
	domainComplaints := make([]domain.Complaint, 0, len(pgComplaints))
	for _, pgComplaint := range pgComplaints {
		domainComplaints = append(domainComplaints, ConvertToDomainComplaint(pgComplaint))
	}

	return domainComplaints
}
