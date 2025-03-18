package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrganizationSetting struct {
	gorm.Model
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	OrganizationID string
	Key            string
	Value          string
}

func (setting *OrganizationSetting) BeforeCreate(tx *gorm.DB) (err error) {
	setting.ID = uuid.New()
	return
}
