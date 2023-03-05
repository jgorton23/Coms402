package repo

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/samber/do"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
)

// Pattern to verify _assertAuthorizationEnforcerCasbinImplem conforms to the required interfaces
var (
	_assertAuthorizationEnforcerCasbinImplem                                   = &authorizationEnforcerCasbinImplem{}
	_                                        usecase.AuthorizationEnforcerRepo = _assertAuthorizationEnforcerCasbinImplem
)

// NewAuthorizationEnforcer -.
func NewAuthorizationEnforcer(i *do.Injector) (usecase.AuthorizationEnforcerRepo, error) {
	authorizationEnforcer := &authorizationEnforcerCasbinImplem{
		authorizationPolicy: do.MustInvoke[usecase.AuthorizationPolicyRepo](i),
	}

	text := `
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
	enforcer, err := casbin.NewEnforcer(m, authorizationEnforcer.authorizationPolicy)
	if err != nil {
		return nil, err
	}

	authorizationEnforcer.enforcer = enforcer

	// TODO implement ths in a smarter way
	_, _ = authorizationEnforcer.enforcer.AddPolicy("*", "/metrics", "*")
	_, _ = authorizationEnforcer.enforcer.AddPolicy("*", "/healthz", "*")
	_, _ = authorizationEnforcer.enforcer.AddPolicy("*", "/docs", "*")
	_, _ = authorizationEnforcer.enforcer.AddPolicy("*", "/docs/*", "*")
	_, _ = authorizationEnforcer.enforcer.AddPolicy("*", "/api/auth/*", "*")
	_, _ = authorizationEnforcer.enforcer.AddPolicy("user", "/api/v1/*", "*")

	return authorizationEnforcer, nil
}

type authorizationEnforcerCasbinImplem struct {
	enforcer            *casbin.Enforcer
	authorizationPolicy usecase.AuthorizationPolicyRepo
}

func (a authorizationEnforcerCasbinImplem) ReloadPolicy() error {
	err := a.enforcer.LoadPolicy()
	if err != nil {
		return err
	}

	return nil
}

func (a authorizationEnforcerCasbinImplem) EnforceRolePathMethod(role, path, method string) (bool, error) {
	enforce, err := a.enforcer.Enforce(role, path, method)
	if err != nil {
		return false, err
	}

	return enforce, nil
}
