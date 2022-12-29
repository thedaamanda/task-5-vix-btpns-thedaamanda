package model

import (
	"time"
)

type Photo struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"type:varchar(100);not null"`
	Caption   string    `json:"caption" gorm:"type:varchar(200);not null"`
	PhotoURL  string    `json:"photo_url" gorm:"type:varchar(200);not null"`
	UserID    int       `json:"user_id" gorm:"foreignKey:UserID"`
	User      User      `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
