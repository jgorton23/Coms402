package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/samber/do"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
)

var (
	// ErrCompanyFound should be returned from Create
	// when the name of the record is found.
	ErrUserToCompanyFound = errors.New("userToCompany found")
	// ErrCompanyNotFound should be returned from Get when the record is not found.
	ErrUserToCompanyNotFound = errors.New("userToCompany not found")
)

func NewUserToCompany(i *do.Injector) (*UserToCompany, error) {
	c := &UserToCompany{
		userToCompany: do.MustInvoke[UserToCompanyRepo](i),
		logger:        do.MustInvoke[*Logger](i),
	}

	return c, nil
}

type UserToCompany struct {
	userToCompany UserToCompanyRepo
	logger        *Logger
}

func (c UserToCompany) Add(ctx context.Context, dutc domain.UserToCompany, assignerUUID uuid.UUID) (role domain.UserToCompany, err error) {

	if dutc.RoleType == domain.RolePrimaryOwner {
		return domain.UserToCompany{}, errors.New("cannot assign primary owner role")
	}

	if assignerUUID == dutc.UserUUID {
		if dutc.Approved {
			return domain.UserToCompany{}, errors.New("cannot approve self-assigned role")
		}
		return c.Create(ctx, dutc)
	}

	allowed, err := c.AllowedToEdit(ctx, assignerUUID, dutc.CompanyUUID)

	if err != nil {
		return domain.UserToCompany{}, err
	}

	if !allowed {
		return domain.UserToCompany{}, ErrUnauthorized
	}

	role, err = c.Create(ctx, dutc)

	if err != nil {
		c.logger.Info(fmt.Sprintf("invalid UserToCompany %v-%v role mapping creation failed", dutc.UserUUID, dutc.CompanyUUID))
		return domain.UserToCompany{}, err
	}

	return role, nil
}

func (c UserToCompany) Create(ctx context.Context, dutc domain.UserToCompany) (role domain.UserToCompany, err error) {

	exists, err := c.userToCompany.Exists(ctx, dutc.UserUUID, dutc.CompanyUUID)

	if err != nil {
		return domain.UserToCompany{}, err
	}

	// role mapping already exists
	if exists {
		c.logger.Info(fmt.Sprintf("userToCompany %v-%v role mapping already exists in database", dutc.UserUUID, dutc.CompanyUUID))
		return domain.UserToCompany{}, ErrUserToCompanyFound
	}

	// Save role mapping to the database
	role, err = c.userToCompany.Create(ctx, dutc)

	if err != nil {
		return domain.UserToCompany{}, err
	}

	c.logger.Info(fmt.Sprintf("userToCompany %v-%v role mapping created in database", dutc.UserUUID, dutc.CompanyUUID))

	return role, nil
}

func (c UserToCompany) AllowedToEdit(ctx context.Context, userUUID uuid.UUID, companyUUID uuid.UUID) (bool, error) {

	utc, err := c.userToCompany.GetByUUIDS(ctx, userUUID, companyUUID)

	if err != nil {
		return false, err
	}

	if !utc.Approved {
		return false, nil
	}

	switch utc.RoleType {
	case domain.RolePrimaryOwner:
		return true, nil
	case domain.RoleOwner:
		return true, nil
	default:
		return false, nil
	}
}
