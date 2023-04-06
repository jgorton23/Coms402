package repo

import (
	"context"

	"github.com/google/uuid"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/company"
)

// Pattern to verify companyDBEntImplem conforms to the required interfaces
var (
	_assertDatabaseEntImplemCompany                     = &databaseEntImplemCompany{}
	_                               usecase.CompanyRepo = _assertDatabaseEntImplemCompany
)

func (implem *DatabaseEnt) RepoCompany() usecase.CompanyRepo {
	return &databaseEntImplemCompany{
		client: implem.client,
	}
}

type databaseEntImplemCompany struct {
	client *ent.Client
}

func (ur *databaseEntImplemCompany) GetByUUID(ctx context.Context, companyUuid uuid.UUID) (domain.Company, error) {
	u, err := ur.client.Company.
		Query().
		Where(company.ID(companyUuid)).
		First(ctx)
	if err != nil {
		return domain.Company{}, err
	}

	return ur.databaseToEntity(u), nil
}

// ExistsNamed
// returns if a company exists with the given name
func (ur *databaseEntImplemCompany) ExistsNamed(ctx context.Context, u string) (bool, error) {
	exists, err := ur.client.Company.
		Query().
		Where(company.Name(u)).
		Exist(ctx)
	if err != nil {
		return false, err
	}

	return exists, nil
}

// ExistsUUID
// returns if a company exists with the given UUID
func (ur *databaseEntImplemCompany) ExistsUUID(ctx context.Context, u uuid.UUID) (bool, error) {
	exists, err := ur.client.Company.
		Query().
		Where(company.ID(u)).
		Exist(ctx)
	if err != nil {
		return false, err
	}

	return exists, nil
}

// Create
// creates a new company
func (ur *databaseEntImplemCompany) Create(ctx context.Context, usr domain.Company) (domain.Company, error) {
	u, err := ur.client.Company.
		Create().
		SetName(usr.Name).
		Save(ctx)
	if err != nil {
		return domain.Company{}, err
	}

	return ur.databaseToEntity(u), nil
}

// Update
// updates the given company
func (ur *databaseEntImplemCompany) Update(ctx context.Context, u domain.Company) error {
	_, err := ur.client.Company.
		Update().
		Where(
			company.ID(u.UUID),
		).
		SetName(u.Name).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (ur *databaseEntImplemCompany) databaseToEntity(u *ent.Company) domain.Company {
	return domain.Company{
		UUID: u.ID,
		Name: u.Name,
	}
}
