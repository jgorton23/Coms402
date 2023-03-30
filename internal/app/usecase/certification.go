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
	ErrCertificationFound    = errors.New("Certification found")
	ErrCertificationNotFound = errors.New("Certification not found")
)

func NewCertification(i *do.Injector) (*Certification, error) {
	c := &Certification{
		certificationRepo: do.MustInvoke[CertificationRepo](i),
		roles:             do.MustInvoke[*UserToCompany](i),
		logger:            do.MustInvoke[*Logger](i),
	}

	return c, nil
}

type Certification struct {
	certificationRepo CertificationRepo
	roles             *UserToCompany
	logger            *Logger
}

// Create
// creates a new cert
func (u Certification) Create(ctx context.Context, c domain.Certification, userUUID uuid.UUID) (domain.Certification, error) {
	// verify a user is allowed to create a certification user a company
	ok, err := u.roles.AllowedToEditData(ctx, userUUID, c.CompanyUUID)

	if err != nil {
		return domain.Certification{}, err
	}

	if !ok {
		return domain.Certification{}, ErrUnauthorized
	}

	exists, err := u.certificationRepo.Exists(ctx, c.PrimaryAttribute)

	if err != nil {
		return domain.Certification{}, err
	}

	if exists {
		u.logger.Info("Certification already exists")
		return domain.Certification{}, ErrCertificationFound
	}

	repoCertification, err := u.certificationRepo.Create(ctx, c)

	if err != nil {
		return domain.Certification{}, err
	}

	u.logger.Info(fmt.Sprintf("Certification %v created in database", repoCertification.UUID))

	return repoCertification, nil
}

// Update
// updates the given cert
func (u Certification) Update(ctx context.Context, c domain.Certification, userUUID uuid.UUID) error {
	// verify a user is allowed to create a certification user a company
	ok, err := u.roles.AllowedToEditData(ctx, userUUID, c.CompanyUUID)

	if err != nil {
		return err
	}

	if !ok {
		return ErrUnauthorized
	}

	exists, err := u.certificationRepo.Exists(ctx, c.PrimaryAttribute)

	if err != nil {
		return err
	}

	if !exists {
		u.logger.Info("Certification item number already exists")
		return ErrCertificationNotFound
	}

	err = u.certificationRepo.Update(ctx, c)

	if err != nil {
		return err
	}

	u.logger.Info(fmt.Sprintf("Certification %v updated in database", c.PrimaryAttribute))

	return nil
}

// GetByCompanyUUID
// returns the certs belonging to a given company
func (u Certification) GetByCompanyUUID(ctx context.Context, companyUUID uuid.UUID, userUUID uuid.UUID) ([]domain.Certification, error) {
	//TODO switch to using the RBAC system / change how roles work
	ok, err := u.roles.AllowedToEditData(ctx, userUUID, companyUUID)

	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, ErrUnauthorized
	}

	itemBatches, err := u.certificationRepo.GetByCompanyUUID(ctx, companyUUID)

	if err != nil {
		return nil, err
	}

	return itemBatches, nil
}

// GetByUUID
// returns the cert with the given UUID
func (u Certification) GetByUUID(ctx context.Context, companyUUID uuid.UUID, userUUID uuid.UUID, uuid uuid.UUID) (domain.Certification, error) {
	//TODO switch to using the RBAC system / change how roles work
	ok, err := u.roles.AllowedToEditData(ctx, userUUID, companyUUID)

	if err != nil {
		return domain.Certification{}, err
	}

	if !ok {
		return domain.Certification{}, ErrUnauthorized
	}

	itemBatches, err := u.certificationRepo.GetByUUID(ctx, uuid)

	if err != nil {
		return domain.Certification{}, err
	}

	return itemBatches, nil
}
