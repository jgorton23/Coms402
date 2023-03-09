package repo

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/samber/do"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
)

// Pattern to verify _assertAuthorizationEnforcerCasbinImplem conforms to the required interfaces
var (
	_assertAuthorizationEnforcerCasbinImplem                                    = &authorizationEnforcerCasbinImplem{}
	_                                        usecase.IAuthorizationEnforcerRepo = _assertAuthorizationEnforcerCasbinImplem
)

// NewAuthorizationEnforcer -.
func NewAuthorizationEnforcer(i *do.Injector) (usecase.IAuthorizationEnforcerRepo, error) {
	authorizationEnforcer := &authorizationEnforcerCasbinImplem{
		authorizationPolicy: do.MustInvoke[persist.Adapter](i),
	}

	// A very helpful note... For some reason p1, r1 dose not work. Do not ask me why but I spent way too long
	// trying to track down why it was not working. Also I think there can only be one policy_effect but don't quote me
	text := `
	[request_definition]
	r = sub, obj, act
	r2 = sub, obj, act

	[policy_definition]
	p = sub, obj, act
	p2 = sub, obj, act
	
	[role_definition]
	g  = _, _
	g2 = _, _

	[policy_effect]
	e = some(where (p.eft == allow))

	[matchers]
	m = (r.sub == p.sub || p.sub == "*") && keyMatch(r.obj, p.obj) && (r.act == p.act || p.act == "*")
	m2 = g(r2.sub, p2.sub) && g2(r2.obj, p2.obj) && r2.act == p2.act
	`

	m, _ := model.NewModelFromString(text)

	// Create the enforcer.
	enforcer, err := casbin.NewEnforcer(m, authorizationEnforcer.authorizationPolicy)
	if err != nil {
		return nil, err
	}

	authorizationEnforcer.enforcer = enforcer
	return authorizationEnforcer, nil
}

type authorizationEnforcerCasbinImplem struct {
	enforcer            *casbin.Enforcer
	authorizationPolicy persist.Adapter
}

func (a authorizationEnforcerCasbinImplem) enforce(enforcer casbin.EnforceContext, rvals ...interface{}) (bool, error) {
	err := a.enforcer.LoadPolicy()
	if err != nil {
		return false, err
	}

	enforce, err := a.enforcer.Enforce(enforcer, rvals[0], rvals[1], rvals[2])
	if err != nil {
		return false, err
	}

	return enforce, nil
}

func (a authorizationEnforcerCasbinImplem) EnforceRolePathMethod(role, path, method string) (bool, error) {
	enforceContext := casbin.NewEnforceContext("")
	return a.enforce(enforceContext, role, path, method)
}

func (a authorizationEnforcerCasbinImplem) AddPolicyRolePathMethod(role, path, method string) (bool, error) {
	return a.enforcer.AddNamedPolicy("p", role, path, method)
}

func (a authorizationEnforcerCasbinImplem) RemovePolicyRolePathMethod(role, path, method string) (bool, error) {
	return a.enforcer.RemoveNamedPolicy("p", role, path, method)
}

func (a authorizationEnforcerCasbinImplem) AddPolicyRBAC(role, data, action string) (bool, error) {
	return a.enforcer.AddNamedPolicy("p2", role, data, action)
}

func (a authorizationEnforcerCasbinImplem) RemovePolicyRBAC(role, data, action string) (bool, error) {
	return a.enforcer.RemoveNamedPolicy("p2", role, data, action)
}

func (a authorizationEnforcerCasbinImplem) AddGroupRBAC(role1, role2 string) (bool, error) {
	return a.enforcer.AddNamedGroupingPolicy("g", role1, role2)
}

func (a authorizationEnforcerCasbinImplem) RemoveGroupRBAC(role1, role2 string) (bool, error) {
	return a.enforcer.RemoveNamedGroupingPolicy("g", role1, role2)
}

func (a authorizationEnforcerCasbinImplem) AddGroupDataRBAC(data, role string) (bool, error) {
	return a.enforcer.AddNamedGroupingPolicy("g2", data, role)
}

func (a authorizationEnforcerCasbinImplem) RemoveGroupDataRBAC(data, role string) (bool, error) {
	return a.enforcer.RemoveNamedGroupingPolicy("g2", data, role)
}

func (a authorizationEnforcerCasbinImplem) EnforceRBAC(user, data, action string) (bool, error) {
	enforceContext := casbin.NewEnforceContext("2")
	enforceContext.EType = "e"
	return a.enforce(enforceContext, user, data, action)
}
