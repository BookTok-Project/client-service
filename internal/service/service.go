package service

import (
	"client-service/internal/domain"
	"client-service/internal/service/dto"
	"context"
	"fmt"
)

type Service struct {
	repo repo
}

func New(repo repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Subscribe(ctx context.Context, subscriberID, subscribeeID int64) error {
	return s.repo.Subscribe(ctx, subscriberID, subscribeeID)
}

func (s *Service) GetSubscriberIDs(ctx context.Context, subscribeeID int64) ([]int64, error) {
	return s.repo.GetSubscriberIDs(ctx, subscribeeID)
}

func (s *Service) GetSubscribeeIDs(ctx context.Context, subscriberID int64) ([]int64, error) {
	return s.repo.GetSubscribeeIDs(ctx, subscriberID)
}

func (s *Service) AddFavorite(ctx context.Context, userID, cardID int64) error {
	return s.repo.AddFavorite(ctx, userID, cardID)
}

func (s *Service) RemoveFavorite(ctx context.Context, userID, cardID int64) error {
	return s.repo.RemoveFavorite(ctx, userID, cardID)
}

func (s *Service) ListFavorites(ctx context.Context, userID int64) ([]int64, error) {
	return s.repo.ListFavorites(ctx, userID)
}

func (s *Service) AddComment(ctx context.Context, userID, bookID int64, text string) error {
	return s.repo.AddComment(ctx, userID, bookID, text)
}

func (s *Service) GetCommentsByBookID(ctx context.Context, page, limit, bookID int64) ([]domain.Comment, int64, error) {
	comments, err := s.repo.GetCommentsByBookID(ctx, page, limit, bookID)
	if err != nil {
		return nil, 0, fmt.Errorf("GetCommentsByBookId(%d, %d, %d): %w", page, limit, bookID, err)
	}

	totalComments, err := s.repo.GetCountCommentsByBookID(ctx, bookID)
	if err != nil {
		return nil, 0, fmt.Errorf("GetCountCommentsByBookId(%d): %w", bookID, err)
	}

	return dto.ConvertToDomainComments(comments), totalComments, nil
}

func (s *Service) GetCommentsByUserID(ctx context.Context, page, limit, userID int64) ([]domain.Comment, int64, error) {
	comments, err := s.repo.GetCommentsByUserID(ctx, page, limit, userID)
	if err != nil {
		return nil, 0, fmt.Errorf("GetCommentsByUserId(%d, %d, %d): %w", page, limit, userID, err)
	}

	totalComments, err := s.repo.GetCountCommentsByUserID(ctx, userID)
	if err != nil {
		return nil, 0, fmt.Errorf("GetCountCommentsByUserId(%d): %w", userID, err)
	}

	return dto.ConvertToDomainComments(comments), totalComments, nil
}

func (s *Service) GetComplaints(ctx context.Context) ([]domain.Complaint, error) {
	complaints, err := s.repo.GetComplaints(ctx)
	return dto.ConvertToDomainComplaints(complaints), err
}

func (s *Service) AddComplaint(ctx context.Context, userID, bookID int64, text string) error {
	return s.repo.AddComplaint(ctx, userID, bookID, text)
}

func (s *Service) AddLike(ctx context.Context, userID, bookID int64) error {
	return s.repo.AddLike(ctx, userID, bookID)
}

func (s *Service) RemoveLike(ctx context.Context, userID, bookID int64) error {
	return s.repo.RemoveLike(ctx, userID, bookID)
}

func (s *Service) ListLike(ctx context.Context, userID int64) ([]int64, error) {
	return s.repo.ListLike(ctx, userID)
}
