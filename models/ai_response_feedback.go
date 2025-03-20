package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AIResponseFeedback struct {
	gorm.Model
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	OrganizationID string
	AIResponseID   uuid.UUID
	Feedback       string
}

func (aiResponseFeedback *AIResponseFeedback) BeforeCreate(tx *gorm.DB) (err error) {
	aiResponseFeedback.ID = uuid.New()
	return
}
