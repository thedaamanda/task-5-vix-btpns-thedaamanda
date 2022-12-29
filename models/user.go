package model

import (
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"type:varchar(50);column:username;uniqueIndex"`
	Email     string    `json:"email" gorm:"type:varchar(70);column:email;uniqueIndex"`
	Password  string    `json:"password" gorm:"type:varchar(100)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Photos    []Photo   `json:"photos" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
