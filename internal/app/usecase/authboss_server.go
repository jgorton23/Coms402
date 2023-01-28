package usecase

import (
	"context"
	"fmt"
	"net/mail"
	"strings"

	"github.com/samber/do"
	"github.com/volatiletech/authboss/v3"

	"github.com/MatthewBehnke/apis/internal/app/domain"
)

var (
	_assertAuthBossImplem                               = &AuthBossServer{}
	_                     authboss.CreatingServerStorer = _assertAuthBossImplem
	// If there is ever a need to implement confirming, recovering or remembering
	// uncommenting the below lines will make it a lot easier
	//  _                 authboss.ConfirmingServerStorer = assertUserUseCase
	// 	_                 authboss.RecoveringServerStorer  = assertUserUseCase
	// 	_                 authboss.RememberingServerStorer = assertUserUseCase
)

func NewAuthBossServer(i *do.Injector) (*AuthBossServer, error) {
	abuc := &AuthBossServer{
		authBossLogger: do.MustInvoke[*AuthBossLogger](i),
		userRepo:       do.MustInvoke[UserRepo](i),
	}

	return abuc, nil
}

type AuthBossServer struct {
	userRepo       UserRepo
	authBossLogger *AuthBossLogger
}

func (uuc AuthBossServer) Save(ctx context.Context, user authboss.User) error {
	u := user.(*domain.User)

	exists, err := uuc.userRepo.Exists(ctx, u.Email)

	if err != nil {
		return err
	}

	// The user already exists
	if !exists {
		uuc.authBossLogger.Info(fmt.Sprintf("User %v does not exists in database", u.Email))

		return authboss.ErrUserNotFound
	}

	err = uuc.userRepo.Update(ctx, *u)

	if err != nil {
		return err
	}

	return nil
}

func (uuc AuthBossServer) Load(ctx context.Context, key string) (user authboss.User, err error) {
	exists, err := uuc.userRepo.Exists(ctx, key)

	if err != nil {
		return &domain.User{}, err
	}

	// The user already exists
	if !exists {
		uuc.authBossLogger.Info(fmt.Sprintf("User %v not found in database", key))

		return &domain.User{}, authboss.ErrUserNotFound
	}

	u, err := uuc.userRepo.GetByEmail(ctx, key)

	user = &u

	if err != nil {
		return
	}

	return user, nil
}

func (uuc AuthBossServer) New(_ context.Context) authboss.User {
	return &domain.User{}
}

func (uuc AuthBossServer) Create(ctx context.Context, user authboss.User) error {
	u := user.(*domain.User)

	u.Email = strings.ToLower(u.Email)

	// Verify email is valid
	_, err := mail.ParseAddress(u.Email)

	if err != nil {
		return err
	}

	exists, err := uuc.userRepo.Exists(ctx, u.Email)

	if err != nil {
		return err
	}

	// The user already exists
	if exists {
		uuc.authBossLogger.Info(fmt.Sprintf("User %v already exists in database", u.Email))

		return authboss.ErrUserFound
	}

	// Insert into database
	_, err = uuc.userRepo.Create(ctx, *u)

	uuc.authBossLogger.Info(fmt.Sprintf("User %v created in database", u.Email))

	if err != nil {
		return err
	}

	return nil
}
