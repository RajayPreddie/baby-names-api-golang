package models

import (
	"github.com/google/uuid"
)

type BabyName struct {
	ID             uuid.UUID `gorm:"type:uuid"`
	BirthYear      int       `json:"birth_year"`
	Gender         string    `json:"gender"`
	Ethnicity      string    `json:"ethnicity"`
	ChildFirstName string    `json:"child_first_name"`
	Count          int       `json:"count"`
	Rank           int       `json:"rank"`
}
