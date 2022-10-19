package controller

// import (
// 	"github.com/casbin/casbin/v2"
// 	"github.com/casbin/casbin/v2/model"
// )

// func NewAuthorization() {
// 	// Initialize the model from a string.
// 	text :=
// 		`
// 		[request_definition]
// 		r = sub, obj, act

// 		[policy_definition]
// 		p = sub, obj, act

// 		[role_definition]
// 		g = _, _

// 		[policy_effect]
// 		e = some(where (p.eft == allow))

// 		[matchers]
// 		m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
// 		`
// 	m, _ := model.NewModelFromString(text)

// 	// Load the policy rules from the .CSV file adapter.
// 	// Replace it with your adapter to avoid files.
// 	//	a := fileadapter.NewAdapter("examples/rbac_policy.csv")

// 	// Create the enforcer.
// 	//	e := casbin.NewEnforcer(m, a)

// }
