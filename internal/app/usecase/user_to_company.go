package usecase

import (
	"context"
	"errors"
	"fmt"

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

	exists, err := c.userToCompany.Exists(ctx, dutc)

	if err != nil {
		return err
	}

	// role mapping already exists
	if exists {
		c.logger.Info(fmt.Sprintf("userToCompany %v-%v role mapping already exists in database", dutc.UserID, dutc.CompanyUUID))
		return ErrUserToCompanyFound
	}

	// Save role mapping to the database
	_, err = c.userToCompany.Create(ctx, dutc)

	if err != nil {
		return err
	}

	c.logger.Info(fmt.Sprintf("userToCompany %v-%v role mapping created in database", dutc.UserID, dutc.CompanyUUID))

	return nil
}
