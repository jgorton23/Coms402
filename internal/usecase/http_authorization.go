package usecase

import (
	"context"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/samber/do"

	"github.com/MatthewBehnke/exampleGoApi/internal/usecase/repo"
)

func NewHttpAuthorization(i *do.Injector) (HttpAuthorization, error) {
	a := &httpAuthorizationImplem{
		log:           do.MustInvoke[Logger](i).WithSubsystem("http_authorization"),
		userService:   do.MustInvoke[repo.DataBaseServiceUser](i),
		casbinService: do.MustInvoke[*repo.DataBaseServiceAuthorizationPolicy](i),
	}

	text :=
		`
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = (r.sub == p.sub || p.sub == "*") && keyMatch(r.obj, p.obj) && (r.act == p.act || p.act == "*")
		`
	m, _ := model.NewModelFromString(text)

	// Create the enforcer.
	enforcer, err := casbin.NewEnforcer(m, a.casbinService)

	if err != nil {
		return nil, err
	}

	a.enforcer = enforcer

	// TODO implement ths in a smarter way
	_, _ = a.enforcer.AddPolicy("*", "/auth/*", "*")
	_, _ = a.enforcer.AddPolicy("user", "/v1/*", "*")

	return a, nil
}

type httpAuthorizationImplem struct {
	userService   repo.DataBaseServiceUser
	casbinService *repo.DataBaseServiceAuthorizationPolicy
	log           Logger
	enforcer      *casbin.Enforcer
}

func (a httpAuthorizationImplem) EnforceUser(user, path, method string) (bool, error) {
	err := a.enforcer.LoadPolicy()

	if err != nil {
		return false, err
	}

	role := "anonymous"

	if user != "anonymous" {
		ok, err := a.userService.Exists(context.Background(), user)

		if err != nil {
			return false, err
		}

		// User not found
		if !ok {
			return false, nil
		}

		u, err := a.userService.GetByEmail(context.Background(), user)

		if err != nil {
			return false, err
		}

		role = u.Role
	}

	enf, err := a.enforcer.Enforce(role, path, method)

	if err != nil {
		return false, err
	}

	return enf, nil

}
