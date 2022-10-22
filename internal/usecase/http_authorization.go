package usecase

import (
	"context"

	"github.com/samber/do"
)

func NewHttpAuthorization(i *do.Injector) (*HttpAuthorization, error) {
	a := &HttpAuthorization{
		userRepo:                do.MustInvoke[UserRepo](i),
		authorizationPolicyRepo: do.MustInvoke[AuthorizationPolicyRepo](i),
		authorizationEnforcer:   do.MustInvoke[AuthorizationEnforcerRepo](i),
	}

	do.MustInvoke[*Logger](i).WithSubsystem("http_authorization").Info("http authorization service started")

	return a, nil
}

type HttpAuthorization struct {
	userRepo                UserRepo
	authorizationPolicyRepo AuthorizationPolicyRepo
	authorizationEnforcer   AuthorizationEnforcerRepo
}

func (a HttpAuthorization) EnforceUser(user, path, method string) (bool, error) {
	role := "anonymous"

	if user != "anonymous" {
		ok, err := a.userRepo.Exists(context.Background(), user)

		if err != nil {
			return false, err
		}

		// User not found
		if !ok {
			return false, nil
		}

		u, err := a.userRepo.GetByEmail(context.Background(), user)

		if err != nil {
			return false, err
		}

		role = u.Role
	}

	err := a.authorizationEnforcer.ReloadPolicy()
	if err != nil {
		return false, err
	}

	enf, err := a.authorizationEnforcer.EnforceRolePathMethod(role, path, method)

	if err != nil {
		return false, err
	}

	return enf, nil

}
