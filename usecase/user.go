package usecase

import (
	"context"
	"errors"
	helper "project/helpers"
	"project/helpers/token"
	model "project/models"
	"project/request"
	"project/response"
	"time"

	"github.com/jackc/pgconn"
)

type userUsecase struct {
	userRepo model.UserRepository
}

func NewUserUsecase(user model.UserRepository) model.UserUsecase {
	return &userUsecase{
		userRepo: user,
	}
}

func (u *userUsecase) Register(ctx context.Context, request request.CreateUserRequest) (*response.UserResponse, error) {
	hashPassword, _ := helper.HashPassword(request.Password)

	user, err := u.userRepo.Create(ctx, &model.User{
		Username: request.Username,
		Email:    request.Email,
		Password: hashPassword,
	})

	if err != nil {
		duplicateEntryError := &pgconn.PgError{Code: "23505"}
		if errors.As(err, &duplicateEntryError) {
			return nil, errors.New("Email or Username already taken")
		}
		return nil, err
	}

	resp := new(response.UserResponse)
	resp.Email = user.Email
	resp.ID = user.ID
	resp.Username = user.Username

	return resp, nil
}

func (u *userUsecase) Login(ctx context.Context, request request.LoginRequest) (*response.TokenResponse, error) {
	user, err := u.userRepo.FindByEmail(ctx, request.Email)

	if err != nil {
		return nil, errors.New("Sorry, we couldn’t find your email in our records.")
	}

	if err := helper.ComparePassword(user.Password, request.Password); err != nil {
		return nil, errors.New("Sorry, the password you entered do not match. Please try again.")
	}

	token, err := token.NewCustomToken(token.NewTokenRequest{
		UserID:    user.ID,
		UserEmail: user.Email,
	}, token.DurationLong)

	if err != nil {
		return nil, err
	}

	resp := new(response.TokenResponse)
	resp.Token = token.AccessToken

	return resp, nil
}

func (u *userUsecase) UpdateUser(ctx context.Context, id int, request request.UpdateUserRequest) (*response.UpdateUserResponse, error) {
	user, err := u.userRepo.FindByID(ctx, id)

	if err != nil {
		return nil, errors.New("Sorry, we couldn’t find your account in our records.")
	}

	user.Username = request.Username
	user.Email = request.Email
	user.UpdatedAt = time.Now()

	user, err = u.userRepo.Update(ctx, id, user)

	if err != nil {
		return nil, err
	}

	resp := new(response.UpdateUserResponse)
	resp.ID = user.ID
	resp.Email = user.Email
	resp.Username = user.Username
	resp.UpdatedAt = user.UpdatedAt.Format("2006-01-02 15:04:05")

	return resp, nil
}

func (u *userUsecase) DeleteUser(ctx context.Context, id int) error {
	if _, err := u.userRepo.FindByID(ctx, id); err != nil {
		return errors.New("Sorry, we couldn’t find your account in our records.")
	}

	if err := u.userRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
