package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/samber/do"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/itembatchtoitembatch"
)

// Pattern to verify certificationDBEntImplem conforms to the required interfaces
var (
	_assertItemToItemImplem                        = &itemToItemDBEntImplem{}
	_                       usecase.ItemToItemRepo = _assertItemToItemImplem
)

// NewCompanyRepo -.
func NewItemToItemRepo(i *do.Injector) (usecase.ItemToItemRepo, error) {
	implem := &itemToItemDBEntImplem{
		do.MustInvoke[*DatabaseConnection](i).Client(),
	}

	return implem, nil
}

type itemToItemDBEntImplem struct {
	*ent.Client
}

func (ur itemToItemDBEntImplem) Create(ctx context.Context, parentUUID uuid.UUID, childUUID uuid.UUID) (domain.ItemToItem, error) {
	utc, err := ur.Client.ItemBatchToItemBatch.
		Create().
		SetChildUUID(childUUID).
		SetParentUUID(parentUUID).
		Save(ctx)
	if err != nil {
		return domain.ItemToItem{}, err
	}

	return ur.databaseToEntity(utc), nil
}

func (ur itemToItemDBEntImplem) Delete(ctx context.Context, parentUUID uuid.UUID, childUUID uuid.UUID) (int, error) {
	deleted, err := ur.Client.ItemBatchToItemBatch.
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

func (ur itemToItemDBEntImplem) Exists(ctx context.Context, parentUUID uuid.UUID, childUUID uuid.UUID) (bool, error) {
	exists, err := ur.Client.ItemBatchToItemBatch.
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

func (ur itemToItemDBEntImplem) FindByParentUUID(ctx context.Context, parentUUID uuid.UUID) ([]domain.ItemToItem, error) {
	u, err := ur.Client.ItemBatchToItemBatch.
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

func (ur *itemToItemDBEntImplem) databaseToEntity(iti *ent.ItemBatchToItemBatch) domain.ItemToItem {
	return domain.ItemToItem{
		UUID:       iti.ID,
		ParentUUID: iti.ParentUUID,
		ChildUUID:  iti.ChildUUID,
	}
}
