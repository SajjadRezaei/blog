package models

import (
	"blog/internal/modules/user/models"
	"gorm.io/gorm"
	"time"
)

type Article struct {
	ID        uint `gorm:"primarykey"`
	UserID    uint
	Title     string `gorm:"varchar:191"`
	Content   string `gorm:"text"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	User      models.User
}
