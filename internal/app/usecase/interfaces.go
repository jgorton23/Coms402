// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"
	"net/http"

	"github.com/google/uuid"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// ConfigRepo -
	ConfigRepo interface {
		Load(string) error
		Get() *domain.Config
	}

	// UserRepo -
	UserRepo interface {
		Get(context.Context) ([]domain.User, error)
		GetByUUID(context.Context, uuid.UUID) (domain.User, error)
		GetByEmail(context.Context, string) (domain.User, error)
		Exists(context.Context, string) (bool, error)
		Create(context.Context, domain.User) (domain.User, error)
		Update(context.Context, domain.User) error
	}

	// LoggerRepo -
	LoggerRepo interface {
		Trace(msg string, fields ...map[string]interface{})
		Debug(msg string, fields ...map[string]interface{})
		Info(msg string, fields ...map[string]interface{})
		Warn(msg string, fields ...map[string]interface{})
		Error(msg string, fields ...map[string]interface{})
	}

	// IAuthorizationEnforcerRepo -
	IAuthorizationEnforcerRepo interface {
		AddPolicyRolePathMethod(role, path, method string) (bool, error)
		RemovePolicyRolePathMethod(role, path, method string) (bool, error)
		EnforceRolePathMethod(role, path, method string) (bool, error)

		// AddPolicyRBAC is responsible for adding policy actions.
		// For example either adding a policy for a group like
		// "role:company1:member, data_group:company1:certs, read"
		// Or for an individual like
		// "matt, data_group:company1:certs, read"
		AddPolicyRBAC(role, data, action string) (bool, error)
		// RemovePolicyRBAC is responsible for  removing policy actions
		RemovePolicyRBAC(role, data, action string) (bool, error)

		// AddGroupRBAC is responsible for adding a group mapping for
		// role to role mappings and user to role mappings like
		// "role:company1:primary_owner, role:company1:owner"
		// which means the primary_owner has all the policies that a
		// owner has. With all policies being additive this works.
		// Adding a user to a  group is the very similar to before
		// "bob, role:company1:primary_owner"
		AddGroupRBAC(role1, role2 string) (bool, error)
		// RemoveGroupRBAC is responsible for  removing group mappings
		RemoveGroupRBAC(role1, role2 string) (bool, error)

		// AddGroupDataRBAC is responsible for adding group data mappings
		// like "data:cert1, data_group:company1:certs"
		AddGroupDataRBAC(data, role string) (bool, error)
		// RemoveGroupDataRBAC is responsible for removing a group data mapping
		RemoveGroupDataRBAC(data, role string) (bool, error)

		// EnforceRBAC determines if a user can perform an action on a data
		EnforceRBAC(user, data, action string) (bool, error)
	}

	// SessionStateRepo -
	SessionStateRepo interface {
		Get(r *http.Request) (map[interface{}]interface{}, error)
		Save(r *http.Request, w http.ResponseWriter, values map[interface{}]interface{}) error
	}

	// CompanyRepo -
	CompanyRepo interface {
		ExistsNamed(context.Context, string) (bool, error)
		ExistsUUID(context.Context, uuid.UUID) (bool, error)
		Create(context.Context, domain.Company) (domain.Company, error)
		Update(context.Context, domain.Company) error
		GetByUUID(context.Context, uuid.UUID) (domain.Company, error)
	}

	// UserToCompanyRepo -
	UserToCompanyRepo interface {
		Exists(ctx context.Context, userUUID uuid.UUID, companyUUID uuid.UUID) (bool, error)
		Create(context.Context, domain.UserToCompany) (domain.UserToCompany, error)
		Update(context.Context, domain.UserToCompany) error
		GetByUUIDS(ctx context.Context, userUUID uuid.UUID, companyUUID uuid.UUID) (domain.UserToCompany, error)
		GetByCompanyUUID(ctx context.Context, companyUUID uuid.UUID) ([]domain.UserToCompany, error)
		GetByUserUUID(ctx context.Context, userUUID uuid.UUID) ([]domain.UserToCompany, error)
	}

	ItemBatchRepo interface {
		Exists(ctx context.Context, itemNumber string) (bool, error)
		Create(ctx context.Context, itemBatch domain.ItemBatch) (domain.ItemBatch, error)
		Update(ctx context.Context, itemBatch domain.ItemBatch) error
		GetByCompanyUUID(ctx context.Context, companyUUID uuid.UUID) ([]domain.ItemBatch, error)
		GetByUUID(ctx context.Context, uuid uuid.UUID) (domain.ItemBatch, error)
	}

	CertificationRepo interface {
		Exists(ctx context.Context, PrimaryAttribute string) (bool, error)
		Create(ctx context.Context, certification domain.Certification) (domain.Certification, error)
		Update(ctx context.Context, certification domain.Certification) error
		GetByCompanyUUID(ctx context.Context, companyUUID uuid.UUID) ([]domain.Certification, error)
		GetByUUID(ctx context.Context, uuid uuid.UUID) (domain.Certification, error)
	}

	ItemToItemRepo interface {
		Create(ctx context.Context, parentUUID uuid.UUID, childUUID uuid.UUID) error
	}
)
