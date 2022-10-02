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
	assertUserUseCase                               = &AuthBossUseCase{}
	_                 authboss.CreatingServerStorer = assertUserUseCase
	// _                 authboss.ConfirmingServerStorer = assertUserUseCase

// 	_                 authboss.RecoveringServerStorer  = assertUserUseCase
// 	_                 authboss.RememberingServerStorer = assertUserUseCase
)

type AuthBossUseCase struct {
	repo UserRepo
	log  LoggerAdapter
}

func NewAuthBossUseCase(r UserRepo, l LoggerAdapter) *AuthBossUseCase {
	return &AuthBossUseCase{
		repo: r,
		log:  l,
	}
}

func (uuc AuthBossUseCase) Save(ctx context.Context, user authboss.User) error {
	u := user.(*entity.User)

	exists, err := uuc.repo.Exists(ctx, u.Email)

	if err != nil {
		return err
	}

	// The user already exists
	if !exists {
		uuc.log.Info(fmt.Sprintf("User %v does not exists in database", u.Email))
		return authboss.ErrUserNotFound
	}

	err = uuc.repo.Update(ctx, *u)

	if err != nil {
		return err
	}

	return nil
}

func (uuc AuthBossUseCase) Load(ctx context.Context, key string) (user authboss.User, err error) {

	exists, err := uuc.repo.Exists(ctx, key)

	if err != nil {
		return &entity.User{}, err
	}

	// The user already exists
	if !exists {
		uuc.log.Info(fmt.Sprintf("User %v not found in database", key))
		return &entity.User{}, authboss.ErrUserNotFound
	}

	u, err := uuc.repo.GetByEmail(ctx, key)

	user = &u

	if err != nil {
		return
	}

	return user, nil
}

func (uuc AuthBossUseCase) New(_ context.Context) authboss.User {
	return &entity.User{}
}

func (uuc AuthBossUseCase) Create(ctx context.Context, user authboss.User) error {
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
	_, err = uuc.repo.Create(ctx, *u)

	uuc.log.Info(fmt.Sprintf("User %v created in database", u.Email))

	if err != nil {
		return err
	}

	return nil
}
