package repo

import (
	"context"

	"github.com/google/uuid"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/userstocompany"
)

// Pattern to verify databaseEntImplemRoles conforms to the required interfaces
var (
	_assertDatabaseEntImplemRoles                           = &databaseEntImplemRoles{}
	_                             usecase.UserToCompanyRepo = _assertDatabaseEntImplemRoles
)

func (implem *DatabaseEnt) RepoRoles() usecase.UserToCompanyRepo {
	return &databaseEntImplemRoles{
		client: implem.client,
	}
}

type databaseEntImplemRoles struct {
	client *ent.Client
}

// Exists
// returns if the given role exists
func (ur *databaseEntImplemRoles) Exists(ctx context.Context, userUUID uuid.UUID, companyUUID uuid.UUID) (bool, error) {

	exists, err := ur.client.UsersToCompany.
		Query().
		Where(
			userstocompany.And(
				userstocompany.CompanyUUID(companyUUID),
				userstocompany.UserUUID(userUUID),
			),
		).
		Exist(ctx)
	if err != nil {
		return false, err
	}

	return exists, nil
}

// GetByUUIDS
// returns the role of a given user in a given company
func (ur *databaseEntImplemRoles) GetByUUIDS(ctx context.Context, userUUID uuid.UUID, companyUUID uuid.UUID) (domain.UserToCompany, error) {

	utc, err := ur.client.UsersToCompany.
		Query().
		Where(
			userstocompany.And(
				userstocompany.CompanyUUID(companyUUID),
				userstocompany.UserUUID(userUUID),
			),
		).
		Only(ctx)

	if err != nil {
		return domain.UserToCompany{}, err
	}

	return ur.databaseToEntity(utc), nil
}

// Create
// creates a new role
func (ur *databaseEntImplemRoles) Create(ctx context.Context, usr domain.UserToCompany) (domain.UserToCompany, error) {
	utc, err := ur.client.UsersToCompany.
		Create().
		SetCompanyUUID(usr.CompanyUUID).
		SetUserID(usr.UserUUID).
		SetRoleType(string(usr.RoleType)).
		SetApproved(usr.Approved).
		Save(ctx)
	if err != nil {
		return domain.UserToCompany{}, err
	}

	return ur.databaseToEntity(utc), nil
}

// Update
// updates the given role
func (ur *databaseEntImplemRoles) Update(ctx context.Context, role domain.UserToCompany) error {
	_, err := ur.client.UsersToCompany.
		Update().
		Where(userstocompany.And(
			userstocompany.CompanyUUID(role.CompanyUUID),
			userstocompany.UserUUID(role.UserUUID),
		),
		).
		SetApproved(role.Approved).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

// GetByCompanyUUID
// returns all roles for the given company
func (ur *databaseEntImplemRoles) GetByCompanyUUID(ctx context.Context, companyUUID uuid.UUID) ([]domain.UserToCompany, error) {
	utcs, err := ur.client.UsersToCompany.
		Query().
		Where(
			userstocompany.CompanyUUID(companyUUID),
		).
		All(ctx)

	if err != nil {
		return nil, err
	}

	roles := make([]domain.UserToCompany, 0)

	for _, v := range utcs {
		roles = append(roles, ur.databaseToEntity(v))
	}
	return roles, nil
}

// GetByUserUUID
// returns all of a users roles
func (ur *databaseEntImplemRoles) GetByUserUUID(ctx context.Context, userUUID uuid.UUID) ([]domain.UserToCompany, error) {
	utcs, err := ur.client.UsersToCompany.
		Query().
		Where(
			userstocompany.UserUUID(userUUID),
		).
		All(ctx)

	if err != nil {
		return nil, err
	}

	roles := make([]domain.UserToCompany, 0)

	for _, v := range utcs {
		roles = append(roles, ur.databaseToEntity(v))
	}
	return roles, nil
}

func (ur *databaseEntImplemRoles) databaseToEntity(u *ent.UsersToCompany) domain.UserToCompany {
	return domain.UserToCompany{
		UUID:        u.ID,
		CompanyUUID: u.CompanyUUID,
		UserUUID:    u.UserUUID,
		RoleType:    domain.Role(u.RoleType),
		Approved:    u.Approved,
	}
}
