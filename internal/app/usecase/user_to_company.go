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

func (c UserToCompany) Create(ctx context.Context, dutc domain.UserToCompany) (err error) {

	exists, err := c.userToCompany.Exists(ctx, dutc.UserUUID, dutc.CompanyUUID)

	if err != nil {
		return err
	}

	// role mapping already exists
	if exists {
		c.logger.Info(fmt.Sprintf("userToCompany %v-%v role mapping already exists in database", dutc.UserUUID, dutc.CompanyUUID))
		return ErrUserToCompanyFound
	}

	// Save role mapping to the database
	_, err = c.userToCompany.Create(ctx, dutc)

	if err != nil {
		return err
	}

	c.logger.Info(fmt.Sprintf("userToCompany %v-%v role mapping created in database", dutc.UserUUID, dutc.CompanyUUID))

	return nil
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
