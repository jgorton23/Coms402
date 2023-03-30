package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/samber/do"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/certification"
)

// Pattern to verify certificationDBEntImplem conforms to the required interfaces
var (
	_assertCertificationImplem                           = &certificationDBEntImplem{}
	_                          usecase.CertificationRepo = _assertCertificationImplem
)

// NewCompanyRepo -.
func NewCertificationRepo(i *do.Injector) (usecase.CertificationRepo, error) {
	implem := &certificationDBEntImplem{
		do.MustInvoke[*DatabaseConnection](i).Client(),
	}

	return implem, nil
}

type certificationDBEntImplem struct {
	*ent.Client
}

// Exists
// returns true if the cert exists
func (ur *certificationDBEntImplem) Exists(ctx context.Context, primaryAttribute string) (bool, error) {

	exists, err := ur.Client.Certification.
		Query().
		Where(
			certification.PrimaryAttribute(primaryAttribute),
		).
		Exist(ctx)
	if err != nil {
		return false, err
	}

	return exists, nil
}

// Create
// Creates a new cert
func (ur *certificationDBEntImplem) Create(ctx context.Context, c domain.Certification) (domain.Certification, error) {
	utc, err := ur.Client.Certification.
		Create().
		SetPrimaryAttribute(c.PrimaryAttribute).
		SetImageUUID(c.ImageUUID).
		SetCompanyUUID(c.CompanyUUID).
		SetItemBatchUUID(c.ItemBatchUUID).
		// SetTemplateUUID(uuid.Nil).
		Save(ctx)
	if err != nil {
		return domain.Certification{}, err
	}

	return ur.databaseToEntity(utc), nil
}

// Update
// Update a given cert
func (ur *certificationDBEntImplem) Update(ctx context.Context, c domain.Certification) error {
	_, err := ur.Client.Certification.
		Update().
		Where(
			certification.And(
				certification.ID(c.UUID),
				certification.PrimaryAttribute(c.PrimaryAttribute),
			),
		).
		SetImageUUID(c.ImageUUID).
		SetCompanyUUID(c.CompanyUUID).
		SetItemBatchUUID(c.ItemBatchUUID).
		// SetTemplateUUID(c.TemplateUUID).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

// GetByCompanyUUID
// returns certs with the given companyUUID
func (ur *certificationDBEntImplem) GetByCompanyUUID(ctx context.Context, companyUUID uuid.UUID) (certificates []domain.Certification, err error) {
	u, err := ur.Client.Certification.
		Query().
		Where(certification.CompanyUUID(companyUUID)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	certificates = make([]domain.Certification, 0)

	for _, v := range u {
		certificates = append(certificates, ur.databaseToEntity(v))
	}

	return certificates, nil
}

// GetByUUID
// returns the cert with the given UUID
func (ur *certificationDBEntImplem) GetByUUID(ctx context.Context, uuid uuid.UUID) (domain.Certification, error) {
	u, err := ur.Client.Certification.
		Query().
		Where(certification.ID(uuid)).
		First(ctx)
	if err != nil {
		return domain.Certification{}, err
	}

	return ur.databaseToEntity(u), nil
}

// databaseToEntity
func (ur *certificationDBEntImplem) databaseToEntity(ib *ent.Certification) domain.Certification {
	return domain.Certification{
		UUID:             ib.ID,
		PrimaryAttribute: ib.PrimaryAttribute,
		ImageUUID:        ib.ImageUUID,
		ItemBatchUUID:    ib.ItemBatchUUID,
		// TemplateUUID:     ib.TemplateUUID,
		CompanyUUID: ib.CompanyUUID,
	}
}
