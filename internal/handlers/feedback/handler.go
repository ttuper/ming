package feedback

import (
	feedback_service "ming/internal/service/feedback"
)

type FeedbackHandler struct {
	FeedbackService *feedback_service.FeedbackService
}

func NewFeedbackHandler(fs *feedback_service.FeedbackService) *FeedbackHandler {
	return &FeedbackHandler{FeedbackService: fs}
}
