package usecase

import (
	"context"
	"errors"

	"github.com/samber/do"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
)

var (
	ErrUserFound    = errors.New("User found")
	ErrUserNotFound = errors.New("User not found")
)

func NewUser(i *do.Injector) (*User, error) {
	c := &User{
		User:   do.MustInvoke[UserRepo](i),
		logger: do.MustInvoke[*Logger](i),
	}

	return c, nil
}

type User struct {
	User   UserRepo
	logger *Logger
}

func (u User) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	return u.User.GetByEmail(ctx, email)
}
