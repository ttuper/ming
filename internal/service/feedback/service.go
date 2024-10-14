package feedback

import "ming/internal/repository"

type FeedbackService struct {
	FeedbackRepo *repository.FeedbackRepository
}

func NewFeedbackService() *FeedbackService {
	return &FeedbackService{
		FeedbackRepo: repository.NewFeedbackRepository(),
	}
}
