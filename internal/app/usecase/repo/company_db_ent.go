package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/samber/do"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/company"
)

// Pattern to verify companyDBEntImplem conforms to the required interfaces
var (
	_assertCompanyImplem                     = &companyDBEntImplem{}
	_                    usecase.CompanyRepo = _assertCompanyImplem
)

// NewCompanyRepo -.
func NewCompanyRepo(i *do.Injector) (usecase.CompanyRepo, error) {
	implem := &companyDBEntImplem{
		do.MustInvoke[*DatabaseConnection](i).Client(),
	}

	return implem, nil
}

type companyDBEntImplem struct {
	*ent.Client
}

// func (ur *companyDBEntImplem) Get(ctx context.Context) ([]domain.Company, error) {
// 	us, err := ur.Client.Company.
// 		Query().
// 		All(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	Companys := make([]domain.Company, len(us))

// 	for i := 0; i < len(us); i++ {
// 		Companys[i] = ur.databaseCompanyToEntityCompany(us[i])
// 	}

// 	return Companys, nil
// }

// func (ur *companyDBEntImplem) GetById(ctx context.Context, id int) (domain.Company, error) {
// 	u, err := ur.Client.Company.
// 		Query().
// 		Where(Company.ID(id)).
// 		First(ctx)
// 	if err != nil {
// 		return domain.Company{}, err
// 	}

// 	return ur.databaseCompanyToEntityCompany(u), nil
// }

func (ur *companyDBEntImplem) GetByUUID(ctx context.Context, companyUuid uuid.UUID) (domain.Company, error) {
	u, err := ur.Client.Company.
		Query().
		Where(company.ID(companyUuid)).
		First(ctx)
	if err != nil {
		return domain.Company{}, err
	}

	return ur.databaseCompanyToEntityCompany(u), nil
}

// ExistsNamed -.
func (ur *companyDBEntImplem) ExistsNamed(ctx context.Context, u string) (bool, error) {
	exists, err := ur.Client.Company.
		Query().
		Where(company.Name(u)).
		Exist(ctx)
	if err != nil {
		return false, err
	}

	return exists, nil
}

// ExistsUUID -.
func (ur *companyDBEntImplem) ExistsUUID(ctx context.Context, u uuid.UUID) (bool, error) {
	exists, err := ur.Client.Company.
		Query().
		Where(company.ID(u)).
		Exist(ctx)
	if err != nil {
		return false, err
	}

	return exists, nil
}

// Create -.
func (ur *companyDBEntImplem) Create(ctx context.Context, usr domain.Company) (domain.Company, error) {
	u, err := ur.Client.Company.
		Create().
		SetName(usr.Name).
		Save(ctx)
	if err != nil {
		return domain.Company{}, err
	}

	return ur.databaseCompanyToEntityCompany(u), nil
}

// Update -.
func (ur *companyDBEntImplem) Update(ctx context.Context, u domain.Company) error {
	_, err := ur.Client.Company. //todo this might be broken
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

func (ur *companyDBEntImplem) databaseCompanyToEntityCompany(u *ent.Company) domain.Company {
	return domain.Company{
		UUID: u.ID,
		Name: u.Name,
	}
}
