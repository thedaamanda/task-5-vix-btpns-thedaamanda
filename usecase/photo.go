package usecase

import (
	"context"
	"errors"
	model "project/models"
	"project/request"
	"project/response"
	"time"
)

type photoUsecase struct {
	photoRepo model.PhotoRepository
}

func NewPhotoUsecase(photo model.PhotoRepository) model.PhotoUsecase {
	return &photoUsecase{
		photoRepo: photo,
	}
}

func (p *photoUsecase) CreatePhoto(ctx context.Context, id int, request request.PhotoRequest) (*response.PhotoResponse, error) {
	photo := &model.Photo{
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoURL: request.PhotoURL,
		UserID:   id,
	}

	photo, err := p.photoRepo.Create(ctx, photo)
	if err != nil {
		return nil, err
	}

	resp := new(response.PhotoResponse)
	resp.ID = photo.ID
	resp.Title = photo.Title
	resp.Caption = photo.Caption
	resp.PhotoURL = photo.PhotoURL
	resp.UserID = photo.UserID
	resp.CreatedAt = photo.CreatedAt.Format("2006-01-02 15:04:05")

	return resp, nil
}

func (p *photoUsecase) GetPhotoList(ctx context.Context) ([]response.PhotosResponse, error) {
	photos, err := p.photoRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	var resp []response.PhotosResponse
	for _, photo := range photos {
		response := response.PhotosResponse{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoURL:  photo.PhotoURL,
			UserID:    photo.UserID,
			CreatedAt: photo.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: photo.UpdatedAt.Format("2006-01-02 15:04:05"),
			User: response.PhotoUser{
				Email:    photo.User.Email,
				Username: photo.User.Username,
			},
		}
		resp = append(resp, response)
	}

	return resp, nil
}

func (p *photoUsecase) UpdatePhoto(ctx context.Context, id int, request request.PhotoRequest) (*response.UpdatePhotoResponse, error) {
	photo, err := p.photoRepo.FindByID(ctx, id)

	if err != nil {
		return nil, errors.New("Sorry, we couldn’t find your photo in our records.")
	}

	photo.Title = request.Title
	photo.Caption = request.Caption
	photo.PhotoURL = request.PhotoURL
	photo.UpdatedAt = time.Now()

	photo, err = p.photoRepo.Update(ctx, id, photo)

	if err != nil {
		return nil, err
	}

	resp := new(response.UpdatePhotoResponse)
	resp.ID = photo.ID
	resp.Title = photo.Title
	resp.Caption = photo.Caption
	resp.PhotoURL = photo.PhotoURL
	resp.UserID = photo.UserID
	resp.UpdatedAt = photo.UpdatedAt.Format("2006-01-02 15:04:05")

	return resp, nil
}

func (p *photoUsecase) DeletePhoto(ctx context.Context, id int) error {
	if _, err := p.photoRepo.FindByID(ctx, id); err != nil {
		return err
	}

	err := p.photoRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
