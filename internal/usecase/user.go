package usecase

import (
	"context"
	"errors"
	"fmt"
	"iseage/bank/internal/entity"
	"net/mail"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	repo UserRepo
}

func NewUserUseCase(r UserRepo) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}

func (uuc *UserUseCase) Create(ctx context.Context, u entity.User) (entity.User, error) {

	u.Email = strings.ToLower(u.Email)

	// Verify email is valid
	_, err := mail.ParseAddress(u.Email)

	if err != nil {
		return entity.User{}, err
	}

	exists, err := uuc.repo.Exists(ctx, u)

	if err != nil {
		return entity.User{}, err
	}

	// The user already exists
	if exists {
		return entity.User{}, errors.New(fmt.Sprintf("failed to create user %v because email already exists", u.Email))
	}

	// Verify passwords match

	if u.Password != u.PasswordConfirmation {
		return entity.User{}, errors.New(fmt.Sprintf("failed to create user %v because passwords do not match", u.Email))
	}

	// Hash password

	ph, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	u.PasswordHash = string(ph)

	if err != nil {
		return entity.User{}, err
	}

	// Insert into database

	user, err := uuc.repo.Create(ctx, u)

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
