package model

import (
	"context"
	"project/request"
	"project/response"
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

type UserRepository interface {
	Create(ctx context.Context, user *User) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Update(ctx context.Context, id int, user *User) (*User, error)
	FindByID(ctx context.Context, id int) (*User, error)
	Delete(ctx context.Context, id int) error
}

type UserUsecase interface {
	Register(ctx context.Context, request request.CreateUserRequest) (*response.UserResponse, error)
	Login(ctx context.Context, request request.LoginRequest) (*response.TokenResponse, error)
	UpdateUser(ctx context.Context, id int, request request.UpdateUserRequest) (*response.UpdateUserResponse, error)
	DeleteUser(ctx context.Context, id int) error
}
