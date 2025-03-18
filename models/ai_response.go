package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AIResponse struct {
	gorm.Model
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	OrganizationID string
	Query          string
	Response       string
}

func (aiResponse *AIResponse) BeforeCreate(tx *gorm.DB) (err error) {
	aiResponse.ID = uuid.New()
	return
}
