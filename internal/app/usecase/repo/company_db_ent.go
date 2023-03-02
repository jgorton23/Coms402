package repo

import (
	"context"

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

// func (ur *companyDBEntImplem) GetByEmail(ctx context.Context, email string) (domain.Company, error) {
// 	u, err := ur.Client.Company.
// 		Query().
// 		Where(Company.Email(email)).
// 		First(ctx)
// 	if err != nil {
// 		return domain.Company{}, err
// 	}

// 	return ur.databaseCompanyToEntityCompany(u), nil
// }

// Exists -.
func (ur *companyDBEntImplem) Exists(ctx context.Context, u string) (bool, error) {
	exists, err := ur.Client.Company.
		Query().
		Where(company.Name(u)).
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

// // Update -.
// func (ur *companyDBEntImplem) Update(ctx context.Context, u domain.Company) error {
// 	_, err := ur.Client.Company.
// 		Update().
// 		SetEmail(u.Email).
// 		SetPasswordHash(u.PasswordHash).
// 		// SetConfirmSelector(u.ConfirmSelector).
// 		// SetConfirmVerifier(u.ConfirmVerifier).
// 		// SetConfirmed(u.Confirmed).
// 		SetAttemptCount(u.AttemptCount).
// 		SetLastAttempt(u.LastAttempt).
// 		SetLocked(u.Locked).
// 		Save(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (ur *companyDBEntImplem) databaseCompanyToEntityCompany(u *ent.Company) domain.Company {
	return domain.Company{
		UUID: u.ID,
		Name: u.Name,
	}
}
