package repo

import (
	"context"

	"github.com/google/uuid"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/itembatch"
)

// Pattern to verify itemBatchDBEntImplem conforms to the required interfaces
var (
	_assertDatabaseEntImplemItemBatch                       = &databaseEntImplemItemBatch{}
	_                                 usecase.ItemBatchRepo = _assertDatabaseEntImplemItemBatch
)

func (implem *DatabaseEnt) RepoItemBatch() usecase.ItemBatchRepo {
	return &databaseEntImplemItemBatch{
		client: implem.client,
	}
}

type databaseEntImplemItemBatch struct {
	client *ent.Client
}

// Exists
// returns if the given itembatch exists
func (ur *databaseEntImplemItemBatch) Exists(ctx context.Context, itemNumber string) (bool, error) {

	exists, err := ur.client.ItemBatch.
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
func (ur *databaseEntImplemItemBatch) Create(ctx context.Context, itemBatch domain.ItemBatch) (domain.ItemBatch, error) {
	utc, err := ur.client.ItemBatch.
		Create().
		SetID(uuid.New()).
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
func (ur *databaseEntImplemItemBatch) Update(ctx context.Context, ib domain.ItemBatch) error {
	found, err := ur.getByUUID(ctx, ib.UUID)

	if err != nil {
		return err
	}

	found.
		Update().
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
func (ur *databaseEntImplemItemBatch) GetByCompanyUUID(ctx context.Context, companyUUID uuid.UUID) (itembatches []domain.ItemBatch, err error) {
	u, err := ur.client.ItemBatch.
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
func (ur *databaseEntImplemItemBatch) GetByUUID(ctx context.Context, uuid uuid.UUID) (domain.ItemBatch, error) {
	u, err := ur.getByUUID(ctx, uuid)
	if err != nil {
		return domain.ItemBatch{}, err
	}

	return ur.databaseToEntity(u), nil
}

// returns the itembatch with the given UUID
func (ur *databaseEntImplemItemBatch) getByUUID(ctx context.Context, uuid uuid.UUID) (*ent.ItemBatch, error) {
	return ur.client.ItemBatch.
		Query().
		Where(itembatch.ID(uuid)).
		First(ctx)
}

func (ur *databaseEntImplemItemBatch) databaseToEntity(ib *ent.ItemBatch) domain.ItemBatch {
	return domain.ItemBatch{
		UUID:        ib.ID,
		ItemNumber:  ib.ItemNumber,
		Description: ib.Description,
		CompanyUUID: ib.CompanyUUID,
	}
}
