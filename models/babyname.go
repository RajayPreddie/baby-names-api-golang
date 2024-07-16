package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BabyName struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	BirthYear      int       `json:"birth_year"`
	Gender         string    `json:"gender"`
	Ethnicity      string    `json:"ethnicity"`
	ChildFirstName string    `json:"child_first_name"`
	Count          int       `json:"count"`
	Rank           int       `json:"rank"`
}

func (babyName *BabyName) BeforeCreate(tx *gorm.DB) (err error) {
	babyName.ID = uuid.New()
	return
}
