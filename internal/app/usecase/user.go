package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/samber/do"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
)

var (
	ErrUserFound    = errors.New("User found")
	ErrUserNotFound = errors.New("User not found")
)

func NewUser(i *do.Injector) (*User, error) {
	c := &User{
		user:   do.MustInvoke[UserRepo](i),
		logger: do.MustInvoke[*Logger](i),
	}

	return c, nil
}

type User struct {
	user   UserRepo
	logger *Logger
}

func (u User) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	return u.user.GetByEmail(ctx, email)
}

func (u User) FindByUUID(ctx context.Context, uuid uuid.UUID) (domain.User, error) {
	return u.user.GetByUUID(ctx, uuid)
}
