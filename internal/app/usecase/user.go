package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/samber/do"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
)

var (
	ErrUserFound    = errors.New("user found")
	ErrUserNotFound = errors.New("user not found")
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

// FindByEmail
// returns the user with the given email
func (u User) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	return u.user.GetByEmail(ctx, email)
}

// FindByUUID
// returns the user with the given UUID
func (u User) FindByUUID(ctx context.Context, uuid uuid.UUID) (domain.User, error) {
	return u.user.GetByUUID(ctx, uuid)
}
