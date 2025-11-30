package transport

import (
	"client-service/internal/domain"
	"time"
)

type subscribeRequest struct {
	SubscriberID int64 `json:"subscriber_id"`
	SubscribeeID int64 `json:"subscribee_id"`
}

type errorResponse struct {
	Error string `json:"error"`
}

type favoriteRequest struct {
	UserID int64 `json:"user_id" validate:"required"`
	BookID int64 `json:"book_id" validate:"required"`
}

type likeRequest struct {
	UserID int64 `json:"user_id" validate:"required"`
	BookID int64 `json:"book_id" validate:"required"`
}

type addCommentRequest struct {
	UserID int64  `json:"user_id" validate:"required"`
	BookID int64  `json:"book_id" validate:"required"`
	Text   string `json:"text" validate:"required"`
}

type CommentResponse struct {
	CommentID int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	BookID    int64     `json:"book_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

type addComplaintRequest struct {
	UserID int64  `json:"user_id" validate:"required"`
	BookID int64  `json:"book_id" validate:"required"`
	Text   string `json:"text" validate:"required"`
}

type ComplaintResponse struct {
	ComplaintID int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	BookID      int64     `json:"book_id"`
	Text        string    `json:"text"`
	CreatedAt   time.Time `json:"created_at"`
}

func ConvertToCommentResponse(comment domain.Comment) CommentResponse {
	return CommentResponse{
		CommentID: comment.ID,
		UserID:    comment.UserID,
		BookID:    comment.BookID,
		Text:      comment.Text,
		CreatedAt: *comment.CreatedAt,
	}
}

func convertToCommentsResponse(comments []domain.Comment) []CommentResponse {
	domainComments := make([]CommentResponse, 0, len(comments))
	for _, pgComment := range comments {
		domainComments = append(domainComments, ConvertToCommentResponse(pgComment))
	}

	return domainComments
}

type GetCommentsResponse struct {
	Comments []CommentResponse `json:"comments"`
	Total    int64             `json:"total"`
}

func ConvertToComplaintResponse(complaint domain.Complaint) ComplaintResponse {
	return ComplaintResponse{
		ComplaintID: complaint.ComplaintID,
		UserID:      complaint.UserID,
		BookID:      complaint.BookID,
		Text:        complaint.Text,
		CreatedAt:   *complaint.CreatedAt,
	}
}

func ConvertToGetCommentsResponse(comments []domain.Comment, total int64) GetCommentsResponse {
	return GetCommentsResponse{
		Comments: convertToCommentsResponse(comments),
		Total:    total,
	}
}

/*
func ConvertToAddComplaintsResponse(complaint domain.Complaint) addComplaintResponse {
	return addComplaintResponse {
		ComplaintID:	complaint.ComplaintID,
		UserID:		complaint.UserID,
		BookID:		complaint.BookID,
		Text:		complaint.Text,
		CreatedAt:	*complaint.CreatedAt,
	}
}
*/

func ConvertToGetComplaintsResponse(complaints []domain.Complaint) []ComplaintResponse {
	domainComplaints := make([]ComplaintResponse, 0, len(complaints))
	for _, pgComment := range complaints {
		domainComplaints = append(domainComplaints, ConvertToComplaintResponse(pgComment))
	}

	return domainComplaints
}
