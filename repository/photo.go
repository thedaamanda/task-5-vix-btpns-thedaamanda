package repository

import (
	"context"
	"project/config"
	model "project/models"
)

type photoRepository struct {
	Cfg config.Config
}

func NewPhotoRepository(cfg config.Config) model.PhotoRepository {
	return &photoRepository{
		Cfg: cfg,
	}
}

func (u *photoRepository) Create(ctx context.Context, photo *model.Photo) (*model.Photo, error) {
	if err := u.Cfg.Database().
		WithContext(ctx).
		Create(photo).
		Error; err != nil {
		return nil, err
	}

	return photo, nil
}

func (u *photoRepository) Fetch(ctx context.Context) ([]model.Photo, error) {
	var photos []model.Photo

	if err := u.Cfg.Database().
		WithContext(ctx).
		Preload("User").
		Find(&photos).
		Error; err != nil {
		return nil, err
	}

	return photos, nil
}

func (u *photoRepository) FindByID(ctx context.Context, id int) (*model.Photo, error) {
	var photo model.Photo

	if err := u.Cfg.Database().
		WithContext(ctx).
		Preload("User").
		Where("id = ?", id).
		First(&photo).
		Error; err != nil {
		return nil, err
	}

	return &photo, nil
}

func (u *photoRepository) Update(ctx context.Context, id int, photo *model.Photo) (*model.Photo, error) {
	if err := u.Cfg.Database().
		WithContext(ctx).
		Model(&model.Photo{}).
		Where("id = ?", id).
		Updates(photo).
		Error; err != nil {
		return nil, err
	}

	return photo, nil
}

func (u *photoRepository) Delete(ctx context.Context, id int) error {
	if err := u.Cfg.Database().
		WithContext(ctx).
		Where("id = ?", id).
		Delete(&model.Photo{}).
		Error; err != nil {
		return err
	}

	return nil
}
