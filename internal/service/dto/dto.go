package dto

import (
	"client-service/internal/db/pg"
	"client-service/internal/domain"
	"time"

	"github.com/jackc/pgx/v5/pgtype"

	repodto "client-service/internal/repo/dto"
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
		ComplaintID: complaint.ID,
		UserID:      complaint.UserID,
		BookID:      complaint.BookID,
		Text:        complaint.Text,
		CreatedAt:   ConvertFromPgTimestamptz(complaint.CreatedAt),
	}
}

func ConvertToDomainComplaints(pgComplaints []pg.Complaint) []domain.Complaint {
	domainComplaints := make([]domain.Complaint, 0, len(pgComplaints))
	for _, pgComplaint := range pgComplaints {
		domainComplaints = append(domainComplaints, ConvertToDomainComplaint(pgComplaint))
	}

	return domainComplaints
}

type GetBooksLikeCountsArgs struct {
	BookIDs []int64
}

func (a GetBooksLikeCountsArgs) ConvertToRepo() repodto.GetBooksLikeCountsArgs {
	args := repodto.GetBooksLikeCountsArgs{
		BookIDs: a.BookIDs,
	}

	return args
}

type GetBooksLikeCountsResult struct {
	Items []GetBooksLikeCountsResultItem
}

func ConvertToGetBooksLikeCountsResult(repoRes repodto.GetBooksLikeCountsResult) GetBooksLikeCountsResult {
	res := GetBooksLikeCountsResult{
		Items: make([]GetBooksLikeCountsResultItem, 0, len(repoRes.Items)),
	}

	for _, item := range repoRes.Items {
		res.Items = append(res.Items, ConvertToGetBooksLikeCountsResultItem(item))
	}

	return res
}

type GetBooksLikeCountsResultItem struct {
	BookID int64
	Count  int64
}

func ConvertToGetBooksLikeCountsResultItem(repoItem repodto.GetBooksLikeCountsResultItem) GetBooksLikeCountsResultItem {
	item := GetBooksLikeCountsResultItem{
		BookID: repoItem.BookID,
		Count:  repoItem.Count,
	}

	return item
}
