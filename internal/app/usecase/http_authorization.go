package usecase

import (
	"context"

	"github.com/samber/do"
)

func NewHTTPAuthorization(i *do.Injector) (*HTTPAuthorization, error) {
	a := &HTTPAuthorization{
		userRepo:                do.MustInvoke[UserRepo](i),
		authorizationPolicyRepo: do.MustInvoke[AuthorizationPolicyRepo](i),
		authorizationEnforcer:   do.MustInvoke[AuthorizationEnforcerRepo](i),
	}

	do.MustInvoke[*Logger](i).WithSubsystem("http_authorization").Info("http authorization service started")

	return a, nil
}

type HTTPAuthorization struct {
	userRepo                UserRepo
	authorizationPolicyRepo AuthorizationPolicyRepo
	authorizationEnforcer   AuthorizationEnforcerRepo
}

func (a HTTPAuthorization) EnforceUser(user, path, method string) (bool, error) {
	role := "anonymous"

	if user != role {
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
