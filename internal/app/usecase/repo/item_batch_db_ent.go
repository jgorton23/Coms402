package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/samber/do"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/itembatch"
)

// Pattern to verify itemBatchDBEntImplem conforms to the required interfaces
var (
	_assertItemBatchImplem                       = &itemBatchDBEntImplem{}
	_                      usecase.ItemBatchRepo = _assertItemBatchImplem
)

// NewCompanyRepo -.
func NewItemBatchRepo(i *do.Injector) (usecase.ItemBatchRepo, error) {
	implem := &itemBatchDBEntImplem{
		do.MustInvoke[*DatabaseConnection](i).Client(),
	}

	return implem, nil
}

type itemBatchDBEntImplem struct {
	*ent.Client
}

// Exists
// returns if the given itembatch exists
func (ur *itemBatchDBEntImplem) Exists(ctx context.Context, itemNumber string) (bool, error) {

	exists, err := ur.Client.ItemBatch.
		Query().
		Where(
			itembatch.ItemNumber(itemNumber),
		).
		Exist(ctx)
	if err != nil {
		return false, err
	}

	return exists, nil
}

// Create
// creates a new itembatch
func (ur *itemBatchDBEntImplem) Create(ctx context.Context, itemBatch domain.ItemBatch) (domain.ItemBatch, error) {
	utc, err := ur.Client.ItemBatch.
		Create().
		SetItemNumber(itemBatch.ItemNumber).
		SetDescription(itemBatch.Description).
		SetCompanyID(itemBatch.CompanyUUID).
		Save(ctx)
	if err != nil {
		return domain.ItemBatch{}, err
	}

	return ur.databaseToEntity(utc), nil
}

// Update
// updates the given itembatch
func (ur *itemBatchDBEntImplem) Update(ctx context.Context, ib domain.ItemBatch) error {
	_, err := ur.Client.ItemBatch.
		Update().
		Where(
			itembatch.And(
				itembatch.ItemNumber(ib.ItemNumber),
				itembatch.ID(ib.UUID),
			),
		).
		SetDescription(ib.Description).
		SetCompanyID(ib.CompanyUUID).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

// GetByCompanyUUID
// returns the itembatches belonging to a given company
func (ur *itemBatchDBEntImplem) GetByCompanyUUID(ctx context.Context, companyUUID uuid.UUID) (itembatches []domain.ItemBatch, err error) {
	u, err := ur.Client.ItemBatch.
		Query().
		Where(itembatch.CompanyUUID(companyUUID)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	itembatches = make([]domain.ItemBatch, 0)

	for _, v := range u {
		itembatches = append(itembatches, ur.databaseToEntity(v))
	}

	return itembatches, nil
}

// returns the itembatch with the given UUID
func (ur *itemBatchDBEntImplem) GetByUUID(ctx context.Context, uuid uuid.UUID) (domain.ItemBatch, error) {
	u, err := ur.Client.ItemBatch.
		Query().
		Where(itembatch.ID(uuid)).
		First(ctx)
	if err != nil {
		return domain.ItemBatch{}, err
	}

	return ur.databaseToEntity(u), nil
}

func (ur *itemBatchDBEntImplem) databaseToEntity(ib *ent.ItemBatch) domain.ItemBatch {
	return domain.ItemBatch{
		UUID:        ib.ID,
		ItemNumber:  ib.ItemNumber,
		Description: ib.Description,
		CompanyUUID: ib.CompanyUUID,
	}
}
