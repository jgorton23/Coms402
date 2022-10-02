package usecase

import (
	"context"
	"fmt"
	"iseage/bank/internal/entity"
	"net/mail"
	"strings"

	"github.com/volatiletech/authboss/v3"
)

var (
	assertUserUseCase                               = &UserUseCase{}
	_                 authboss.CreatingServerStorer = assertUserUseCase
	// _                 authboss.ConfirmingServerStorer = assertUserUseCase

// 	_                 authboss.RecoveringServerStorer  = assertUserUseCase
// 	_                 authboss.RememberingServerStorer = assertUserUseCase
)

type UserUseCase struct {
	repo UserRepo
	log  LoggerAdapter
}

func NewUserUseCase(r UserRepo, l LoggerAdapter) *UserUseCase {
	return &UserUseCase{
		repo: r,
		log:  l,
	}
}

func (uuc *UserUseCase) Save(ctx context.Context, user authboss.User) error {
	u := user.(*entity.User)
	_, err := uuc.repo.CreateFromEmail(ctx, u.Email)

	if err != nil {
		return err
	}

	return nil
}

func (uuc *UserUseCase) Load(ctx context.Context, key string) (user authboss.User, err error) {

	exists, err := uuc.repo.Exists(ctx, key)

	if err != nil {
		return entity.User{}, err
	}

	// The user already exists
	if !exists {
		uuc.log.Info(fmt.Sprintf("User %v not found in database", key))
		return entity.User{}, authboss.ErrUserNotFound
	}

	user, err = uuc.repo.GetByEmail(ctx, key)

	if err != nil {
		return
	}

	return user, nil
}

func (uuc *UserUseCase) New(_ context.Context) authboss.User {
	return &entity.User{}
}

func (uuc *UserUseCase) Create(ctx context.Context, user authboss.User) error {
	u := user.(*entity.User)

	u.Email = strings.ToLower(u.Email)

	// Verify email is valid
	_, err := mail.ParseAddress(u.Email)

	if err != nil {
		return err
	}

	exists, err := uuc.repo.Exists(ctx, u.Email)

	if err != nil {
		return err
	}

	// The user already exists
	if exists {
		uuc.log.Info(fmt.Sprintf("User %v already exists in database", u.Email))
		return authboss.ErrUserFound
	}

	// Insert into database
	_, err = uuc.repo.CreateFromEmail(ctx, u.Email)

	uuc.log.Info(fmt.Sprintf("User %v created in database", u.Email))

	if err != nil {
		return err
	}

	return nil
}
