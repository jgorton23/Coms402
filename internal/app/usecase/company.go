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
	ErrCompanyFound = errors.New("company found")
	// ErrCompanyNotFound should be returned from Get when the record is not found.
	ErrCompanyNotFound = errors.New("company not found")
	ErrUnauthorized    = errors.New("unauthorized")
)

func NewCompany(i *do.Injector) (*Company, error) {
	c := &Company{
		companyRepo:   do.MustInvoke[CompanyRepo](i),
		userToCompany: do.MustInvoke[*UserToCompany](i),
		user:          do.MustInvoke[*User](i),
		logger:        do.MustInvoke[*Logger](i),
	}

	return c, nil
}

type Company struct {
	// Repos
	companyRepo CompanyRepo

	// UseCases
	userToCompany *UserToCompany
	user          *User
	logger        *Logger
}

func (c Company) Create(ctx context.Context, dc domain.Company, userUUID uuid.UUID) (domain.Company, error) {

	exists, err := c.companyRepo.ExistsNamed(ctx, dc.Name)

	if err != nil {
		return domain.Company{}, err
	}
	// The company already exists
	if exists {
		c.logger.Info(fmt.Sprintf("Company %v already exists in database", dc.Name))

		return domain.Company{}, ErrCompanyFound
	}

	// Save company into the database

	dc, err = c.companyRepo.Create(ctx, dc)

	if err != nil {
		return domain.Company{}, err
	}

	c.logger.Info(fmt.Sprintf("Company %v created in database", dc.Name))

	// Add initial Company Owner

	// TODO check for err
	c.userToCompany.Create(ctx, domain.UserToCompany{
		CompanyUUID: dc.UUID,
		UserUUID:    userUUID,
		RoleType:    domain.RolePrimaryOwner,
		Approved:    true,
	})

	return dc, nil
}

func (c Company) Update(ctx context.Context, dc domain.Company, userUUID uuid.UUID) error {
	existsCompany, err := c.companyRepo.ExistsUUID(ctx, dc.UUID)

	if err != nil {
		return err
	}

	if !existsCompany {
		c.logger.Info(fmt.Sprintf("company %v does not exists in database", dc.UUID))
		return ErrCompanyNotFound
	}

	existsCompanyNamed, err := c.companyRepo.ExistsNamed(ctx, dc.Name)

	if err != nil {
		return err
	}

	if existsCompanyNamed {
		c.logger.Info(fmt.Sprintf("company %v exists in database", dc.Name))
		return ErrCompanyFound
	}

	//Check if user is allowed to update the company
	ok, err := c.userToCompany.AllowedToEdit(ctx, userUUID, dc.UUID)

	if err != nil {
		return err
	}

	if !ok {
		c.logger.Info(fmt.Sprintf("user %v is unauthorized to edit company %v", userUUID, dc.UUID))
		return ErrUnauthorized
	}

	err = c.companyRepo.Update(ctx, dc)

	if err != nil {
		return err
	}

	return nil
}

func (c Company) GetByUUID(ctx context.Context, companyUUID uuid.UUID) (domain.Company, error) {
	return c.companyRepo.GetByUUID(ctx, companyUUID)
}
