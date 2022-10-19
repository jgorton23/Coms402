package usecase

import (
	"context"
	"fmt"
	"net/mail"
	"strings"

	"github.com/samber/do"
	"github.com/volatiletech/authboss/v3"

	"github.com/MatthewBehnke/exampleGoApi/internal/entity"
	"github.com/MatthewBehnke/exampleGoApi/internal/usecase/repo"
)

var (
	assertAuthBossImplem                               = &authBossImplem{}
	_                    authboss.CreatingServerStorer = assertAuthBossImplem
	// If there is ever a need to implement confirming, recovering or remembering
	// uncommenting the below lines will make it a lot easier
	//  _                 authboss.ConfirmingServerStorer = assertUserUseCase
	// 	_                 authboss.RecoveringServerStorer  = assertUserUseCase
	// 	_                 authboss.RememberingServerStorer = assertUserUseCase
)

func NewAuthBoss(i *do.Injector) (AuthBoss, error) {
	abuc := &authBossImplem{}
	abuc.log = do.MustInvoke[Logger](i)
	abuc.userService = do.MustInvoke[repo.DataBaseServiceUser](i)

	return abuc, nil
}

type authBossImplem struct {
	userService repo.DataBaseServiceUser
	log         Logger
}

func (uuc authBossImplem) Save(ctx context.Context, user authboss.User) error {
	u := user.(*entity.User)

	exists, err := uuc.userService.Exists(ctx, u.Email)

	if err != nil {
		return err
	}

	// The user already exists
	if !exists {
		uuc.log.Info(fmt.Sprintf("User %v does not exists in database", u.Email))
		return authboss.ErrUserNotFound
	}

	err = uuc.userService.Update(ctx, *u)

	if err != nil {
		return err
	}

	return nil
}

func (uuc authBossImplem) Load(ctx context.Context, key string) (user authboss.User, err error) {

	exists, err := uuc.userService.Exists(ctx, key)

	if err != nil {
		return &entity.User{}, err
	}

	// The user already exists
	if !exists {
		uuc.log.Info(fmt.Sprintf("User %v not found in database", key))
		return &entity.User{}, authboss.ErrUserNotFound
	}

	u, err := uuc.userService.GetByEmail(ctx, key)

	user = &u

	if err != nil {
		return
	}

	return user, nil
}

func (uuc authBossImplem) New(_ context.Context) authboss.User {
	return &entity.User{}
}

func (uuc authBossImplem) Create(ctx context.Context, user authboss.User) error {
	u := user.(*entity.User)

	u.Email = strings.ToLower(u.Email)

	// Verify email is valid
	_, err := mail.ParseAddress(u.Email)

	if err != nil {
		return err
	}

	exists, err := uuc.userService.Exists(ctx, u.Email)

	if err != nil {
		return err
	}

	// The user already exists
	if exists {
		uuc.log.Info(fmt.Sprintf("User %v already exists in database", u.Email))
		return authboss.ErrUserFound
	}

	// Insert into database
	_, err = uuc.userService.Create(ctx, *u)

	uuc.log.Info(fmt.Sprintf("User %v created in database", u.Email))

	if err != nil {
		return err
	}

	return nil
}
