package repo

import (
	"context"

	"github.com/google/uuid"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/certification"
)

// Pattern to verify databaseEntImplemCertification  conforms to the required interfaces
var (
	_assertDatabaseEntImplemCertification                           = &databaseEntImplemCertification{}
	_                                     usecase.CertificationRepo = _assertDatabaseEntImplemCertification
)

func (implem *DatabaseEnt) RepoCertification() usecase.CertificationRepo {
	return &databaseEntImplemCertification{
		client: implem.client,
	}
}

type databaseEntImplemCertification struct {
	client *ent.Client
}

// Exists
// returns true if the cert exists
func (ur *databaseEntImplemCertification) Exists(ctx context.Context, primaryAttribute string) (bool, error) {

	exists, err := ur.client.Certification.
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
func (ur *databaseEntImplemCertification) Create(ctx context.Context, c domain.Certification) (domain.Certification, error) {
	utc, err := ur.client.Certification.
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
func (ur *databaseEntImplemCertification) Update(ctx context.Context, c domain.Certification) error {
	_, err := ur.client.Certification.
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
func (ur *databaseEntImplemCertification) GetByCompanyUUID(ctx context.Context, companyUUID uuid.UUID) (certificates []domain.Certification, err error) {
	u, err := ur.client.Certification.
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
func (ur *databaseEntImplemCertification) GetByUUID(ctx context.Context, uuid uuid.UUID) (domain.Certification, error) {
	u, err := ur.client.Certification.
		Query().
		Where(certification.ID(uuid)).
		First(ctx)
	if err != nil {
		return domain.Certification{}, err
	}

	return ur.databaseToEntity(u), nil
}

// databaseToEntity
func (ur *databaseEntImplemCertification) databaseToEntity(ib *ent.Certification) domain.Certification {
	return domain.Certification{
		UUID:             ib.ID,
		PrimaryAttribute: ib.PrimaryAttribute,
		ImageUUID:        ib.ImageUUID,
		ItemBatchUUID:    ib.ItemBatchUUID,
		// TemplateUUID:     ib.TemplateUUID,
		CompanyUUID: ib.CompanyUUID,
	}
}
