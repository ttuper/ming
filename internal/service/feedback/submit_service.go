package feedback

import "ming/internal/models"

func (s *FeedbackService) SubmitFeedback(feedback *models.Feedback) error {
	err := s.FeedbackRepo.CreateFeedback(feedback)
	if err != nil {
		return err
	}
	return nil
}
