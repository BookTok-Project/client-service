package service

import "context"

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
