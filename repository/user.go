package repository

import (
	"context"
	"project/config"
	model "project/models"
)

type userRepository struct {
	Cfg config.Config
}

func NewUserRepository(cfg config.Config) model.UserRepository {
	return &userRepository{
		Cfg: cfg,
	}
}

func (u *userRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	if err := u.Cfg.Database().
		WithContext(ctx).
		Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User

	if err := u.Cfg.Database().
		WithContext(ctx).
		Where("email = ?", email).
		First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) Update(ctx context.Context, id int, user *model.User) (*model.User, error) {
	if err := u.Cfg.Database().
		WithContext(ctx).
		Model(&model.User{}).
		Where("id = ?", id).
		Updates(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) FindByID(ctx context.Context, id int) (*model.User, error) {
	var user model.User

	if err := u.Cfg.Database().
		WithContext(ctx).
		Where("id = ?", id).
		First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) Delete(ctx context.Context, id int) error {
	if err := u.Cfg.Database().
		WithContext(ctx).
		Where("id = ?", id).
		Delete(&model.User{}).Error; err != nil {
		return err
	}

	return nil
}
