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
	ErrCompanyFound = errors.New("company found")
	// ErrCompanyNotFound should be returned from Get when the record is not found.
	ErrCompanyNotFound = errors.New("company not found")
)

func NewCompany(i *do.Injector) (*Company, error) {
	c := &Company{
		companyRepo:   do.MustInvoke[CompanyRepo](i),
		userToCompany: do.MustInvoke[*UserToCompany](i),
		logger:        do.MustInvoke[*Logger](i),
	}

	return c, nil
}

type Company struct {
	// Repos
	companyRepo CompanyRepo

	// UseCases
	userToCompany *UserToCompany
	logger        *Logger
}

func (c Company) Create(ctx context.Context, dc domain.Company, userID int) (domain.Company, error) {

	exists, err := c.companyRepo.Exists(ctx, dc.Name)

	if err != nil {
		return domain.Company{}, err
	}
	// The user already exists
	if exists {
		c.logger.Info(fmt.Sprintf("Company %v already exists in database", dc.Name))

		return domain.Company{}, ErrCompanyFound
	}

	// Save company into the database

	dc, err = c.companyRepo.Create(ctx, dc)

	if err != nil {
		return domain.Company{}, err
	}

	c.logger.Info(fmt.Sprintf("Company %v created in database", dc.UUID))

	// Add initial Company Owner

	c.userToCompany.Create(ctx, domain.UserToCompany{
		CompanyUUID: dc.UUID,
		UserID:      userID,
		RoleType:    domain.RolePrimaryOwner,
		Approved:    true,
	})

	return dc, nil
}
