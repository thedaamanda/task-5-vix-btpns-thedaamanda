package model

import (
	"context"
	"project/request"
	"project/response"
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

type PhotoRepository interface {
	Create(ctx context.Context, photo *Photo) (*Photo, error)
	Fetch(ctx context.Context) ([]Photo, error)
	FindByID(ctx context.Context, id int) (*Photo, error)
	Update(ctx context.Context, id int, photo *Photo) (*Photo, error)
	Delete(ctx context.Context, id int) error
}

type PhotoUsecase interface {
	CreatePhoto(ctx context.Context, id int, request request.PhotoRequest) (*response.PhotoResponse, error)
	GetPhotoList(ctx context.Context) ([]response.PhotosResponse, error)
	UpdatePhoto(ctx context.Context, id int, request request.PhotoRequest) (*response.UpdatePhotoResponse, error)
	DeletePhoto(ctx context.Context, id int) error
}
