package repository

import (
	"ming/internal/db"
	"ming/internal/models"
	"ming/pkg/logger"
)

type FeedbackRepository struct{}

func NewFeedbackRepository() *FeedbackRepository {
	return &FeedbackRepository{}
}

func (r *FeedbackRepository) CreateFeedback(feedback *models.Feedback) error {
	err := db.DB.Create(feedback).Error
	if err != nil {
		logger.Error("Failed to create feedback")
		return err
	}
	return nil
}
