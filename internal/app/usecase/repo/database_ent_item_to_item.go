package repo

import (
	"context"

	"github.com/google/uuid"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/itembatchtoitembatch"
)

// Pattern to verify certificationDBEntImplem conforms to the required interfaces
var (
	_assertItemToItemImplem                        = &databaseEntImplemItemToItem{}
	_                       usecase.ItemToItemRepo = _assertItemToItemImplem
)

// NewCompanyRepo -.
func (implem *DatabaseEnt) RepoItemToItem() usecase.ItemToItemRepo {
	return &databaseEntImplemItemToItem{
		client: implem.client,
	}
}

type databaseEntImplemItemToItem struct {
	client *ent.Client
}

func (ur databaseEntImplemItemToItem) Create(ctx context.Context, parentUUID uuid.UUID, childUUID uuid.UUID) (domain.ItemToItem, error) {
	utc, err := ur.client.ItemBatchToItemBatch.
		Create().
		SetChildUUID(childUUID).
		SetParentUUID(parentUUID).
		Save(ctx)
	if err != nil {
		return domain.ItemToItem{}, err
	}

	return ur.databaseToEntity(utc), nil
}

func (ur databaseEntImplemItemToItem) Delete(ctx context.Context, parentUUID uuid.UUID, childUUID uuid.UUID) (int, error) {
	deleted, err := ur.client.ItemBatchToItemBatch.
		Delete().
		Where(
			itembatchtoitembatch.And(
				itembatchtoitembatch.ParentUUID(parentUUID),
				itembatchtoitembatch.ChildUUID(childUUID),
			),
		).Exec(ctx)

	if err != nil {
		return 0, err
	}

	return deleted, nil
}

func (ur databaseEntImplemItemToItem) Exists(ctx context.Context, parentUUID uuid.UUID, childUUID uuid.UUID) (bool, error) {
	exists, err := ur.client.ItemBatchToItemBatch.
		Query().
		Where(
			itembatchtoitembatch.And(
				itembatchtoitembatch.ParentUUID(parentUUID),
				itembatchtoitembatch.ChildUUID(childUUID),
			),
		).
		Exist(ctx)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (ur databaseEntImplemItemToItem) FindByParentUUID(ctx context.Context, parentUUID uuid.UUID) ([]domain.ItemToItem, error) {
	u, err := ur.client.ItemBatchToItemBatch.
		Query().
		Where(
			itembatchtoitembatch.ParentUUID(parentUUID),
		).All(ctx)

	if err != nil {
		return nil, err
	}
	itemToItems := make([]domain.ItemToItem, 0)

	for _, v := range u {
		itemToItems = append(itemToItems, ur.databaseToEntity(v))
	}

	return itemToItems, nil
}

func (ur *databaseEntImplemItemToItem) databaseToEntity(iti *ent.ItemBatchToItemBatch) domain.ItemToItem {
	return domain.ItemToItem{
		UUID:       iti.ID,
		ParentUUID: iti.ParentUUID,
		ChildUUID:  iti.ChildUUID,
	}
}
