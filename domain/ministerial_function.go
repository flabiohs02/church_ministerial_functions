package domain

import (
	"gorm.io/gorm"
)

type MinisterialFunction struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(100);unique;not null"`
	Description string `gorm:"type:varchar(255)"`
}
