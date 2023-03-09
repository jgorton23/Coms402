package repo_test

import (
	"testing"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase/repo"
	"github.com/samber/do"
)

// Interface Testing for the usecase.IAuthorizationEnforcerRepo casbin implementations

func TestAuthorizationEnforcerCasbinRolePathMethod(t *testing.T) {
	injector := do.New()

	// The AuthorizationEnforcer needs a policy repo
	do.Provide(injector, repo.NewAuthorizationPolicyInMemoryRepo)
	// the repo we are testing
	do.Provide(injector, repo.NewAuthorizationEnforcer)

	auth := do.MustInvoke[usecase.IAuthorizationEnforcerRepo](injector)

	fn := func(ok bool, err error) {
		if err != nil {
			t.Error(err)
		}
		if !ok {
			t.Error("Expected to be true")
		}
	}

	// Policy rule set to be tested against
	fn(auth.AddPolicyRolePathMethod("*", "/metrics", "*"))
	fn(auth.AddPolicyRolePathMethod("user", "/api/v1/*", "*"))

	policyTests := []struct {
		role     string
		path     string
		method   string
		expected bool
	}{
		{"*", "/metrics", "*", true},
		{"user", "/metrics", "*", true},
		{"bob", "/metrics", "*", true},
		{"bob", "/api/v1/test", "*", false},
		{"user", "/api/v1/test", "*", true},
	}

	for _, tt := range policyTests {
		actual, err := auth.EnforceRolePathMethod(tt.role, tt.path, tt.method)

		if err != nil {
			t.Error(err)
		}

		if actual != tt.expected {
			t.Errorf("EnforceRolePathMethod(%v, %v, %v): expected %v, actual %v", tt.role, tt.path, tt.method, tt.expected, actual)
		}
	}
}

func TestAuthorizationEnforcerCasbinRBAC(t *testing.T) {
	injector := do.New()

	// The AuthorizationEnforcer needs a policy repo
	do.Provide(injector, repo.NewAuthorizationPolicyInMemoryRepo)
	// the repo we are testing
	do.Provide(injector, repo.NewAuthorizationEnforcer)

	auth := do.MustInvoke[usecase.IAuthorizationEnforcerRepo](injector)

	fn := func(ok bool, err error) {
		if err != nil {
			t.Error(err)
		}
		if !ok {
			t.Error("Expected to be true")
		}
	}

	// Policy rule set to be tested against
	fn(auth.AddPolicyRBAC("alice", "data1", "read"))
	fn(auth.AddPolicyRBAC("bob", "data2", "write"))
	fn(auth.AddPolicyRBAC("data_group_admin", "data_group", "write"))
	fn(auth.AddGroupRBAC("alice", "data_group_admin"))
	fn(auth.AddGroupDataRBAC("data1", "data_group"))
	fn(auth.AddGroupDataRBAC("data2", "data_group"))

	policyTests := []struct {
		role     string
		data     string
		action   string
		expected bool
	}{
		{"alice", "data1", "read", true},
		{"alice", "data1", "write", true},
		{"alice", "data2", "read", false},
		{"alice", "data2", "write", true},
		{"bob", "data2", "write", true},
		{"bob", "data2", "read", false},
	}

	for _, tt := range policyTests {
		actual, err := auth.EnforceRBAC(tt.role, tt.data, tt.action)

		if err != nil {
			t.Error(err)
		}

		if actual != tt.expected {
			t.Errorf("EnforceRBAC(%v, %v, %v): expected %v, actual %v", tt.role, tt.data, tt.action, tt.expected, actual)
		}
	}
}
