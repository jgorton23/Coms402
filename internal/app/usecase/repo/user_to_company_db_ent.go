package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/samber/do"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/userstocompany"
)

// Pattern to verify userToCompanyDBEntImplem conforms to the required interfaces
var (
	_assertUserToCompanyImplem                           = &userToCompanyDBEntImplem{}
	_                          usecase.UserToCompanyRepo = _assertUserToCompanyImplem
)

// NewCompanyRepo -.
func NewUserToCompanyRepo(i *do.Injector) (usecase.UserToCompanyRepo, error) {
	implem := &userToCompanyDBEntImplem{
		do.MustInvoke[*DatabaseConnection](i).Client(),
	}

	return implem, nil
}

type userToCompanyDBEntImplem struct {
	*ent.Client
}

// func (ur *userToCompanyDBEntImplem) Get(ctx context.Context) ([]domain.Company, error) {
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

// func (ur *userToCompanyDBEntImplem) GetById(ctx context.Context, id int) (domain.Company, error) {
// 	u, err := ur.Client.Company.
// 		Query().
// 		Where(Company.ID(id)).
// 		First(ctx)
// 	if err != nil {
// 		return domain.Company{}, err
// 	}

// 	return ur.databaseCompanyToEntityCompany(u), nil
// }

// func (ur *userToCompanyDBEntImplem) GetByEmail(ctx context.Context, email string) (domain.Company, error) {
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
func (ur *userToCompanyDBEntImplem) Exists(ctx context.Context, userUUID uuid.UUID, companyUUID uuid.UUID) (bool, error) {

	exists, err := ur.Client.UsersToCompany.
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

// GetByUUIDS -.
func (ur *userToCompanyDBEntImplem) GetByUUIDS(ctx context.Context, userUUID uuid.UUID, companyUUID uuid.UUID) (domain.UserToCompany, error) {

	utc, err := ur.Client.UsersToCompany.
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

// Create -.
func (ur *userToCompanyDBEntImplem) Create(ctx context.Context, usr domain.UserToCompany) (domain.UserToCompany, error) {
	utc, err := ur.Client.UsersToCompany.
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

func (ur *userToCompanyDBEntImplem) Update(ctx context.Context, role domain.UserToCompany) error {
	_, err := ur.Client.UsersToCompany.
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

func (ur *userToCompanyDBEntImplem) GetByCompanyUUID(ctx context.Context, companyUUID uuid.UUID) ([]domain.UserToCompany, error) {
	utcs, err := ur.Client.UsersToCompany.
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

func (ur *userToCompanyDBEntImplem) GetByUserUUID(ctx context.Context, userUUID uuid.UUID) ([]domain.UserToCompany, error) {
	utcs, err := ur.Client.UsersToCompany.
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

// // Update -.
// func (ur *userToCompanyDBEntImplem) Update(ctx context.Context, u domain.Company) error {
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

func (ur *userToCompanyDBEntImplem) databaseToEntity(u *ent.UsersToCompany) domain.UserToCompany {
	return domain.UserToCompany{
		UUID:        u.ID,
		CompanyUUID: u.CompanyUUID,
		UserUUID:    u.UserUUID,
		RoleType:    domain.Role(u.RoleType),
		Approved:    u.Approved,
	}
}
