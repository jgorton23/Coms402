package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/samber/do"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
)

var (
	ErrItemToItemFound    = errors.New("ItemToItem found")
	ErrItemToItemNotFound = errors.New("ItemToItem not found")
)

func NewItemToItem(i *do.Injector) (*ItemToItem, error) {
	c := &ItemToItem{
		itemToItemRepo: do.MustInvoke[ItemToItemRepo](i),
		itemBatchRepo:  do.MustInvoke[ItemBatchRepo](i),
		roles:          do.MustInvoke[*UserToCompany](i),
		logger:         do.MustInvoke[*Logger](i),
	}

	return c, nil
}

type ItemToItem struct {
	itemToItemRepo ItemToItemRepo
	itemBatchRepo  ItemBatchRepo
	roles          *UserToCompany
	logger         *Logger
}

func (u ItemToItem) CreateAll(ctx context.Context, userUUID uuid.UUID, parentUUID uuid.UUID, children []uuid.UUID) ([]domain.ItemToItem, error) {

	parent, err := u.itemBatchRepo.GetByUUID(ctx, parentUUID)

	if err != nil {
		return nil, err
	}

	allowed, err := u.roles.AllowedToEditData(ctx, userUUID, parent.CompanyUUID)

	if err != nil {
		return nil, err
	}

	if !allowed {
		return nil, ErrUnauthorized
	}

	var created []domain.ItemToItem

	for _, childUUID := range children {
		exists, err := u.itemToItemRepo.Exists(ctx, parentUUID, childUUID)

		if err != nil {
			return created, err
		}

		if !exists {

			child, err := u.itemToItemRepo.Create(ctx, parentUUID, childUUID)

			if err != nil {
				return created, err
			}

			created = append(created, child)
		}
	}

	return created, nil
}

func (u ItemToItem) DeleteAll(ctx context.Context, userUUID uuid.UUID, parentUUID uuid.UUID, children []uuid.UUID) (int, error) {
	parent, err := u.itemBatchRepo.GetByUUID(ctx, parentUUID)

	if err != nil {
		return 0, err
	}

	allowed, err := u.roles.AllowedToEditData(ctx, userUUID, parent.CompanyUUID)

	if err != nil {
		return 0, err
	}

	if !allowed {
		return 0, ErrUnauthorized
	}

	deleted := 0

	for _, childUUID := range children {
		numDeleted, err := u.itemToItemRepo.Delete(ctx, parentUUID, childUUID)

		deleted += numDeleted

		if err != nil {
			return deleted, err
		}
	}

	return deleted, nil
}

func (u ItemToItem) FindByParentUUID(ctx context.Context, userUUID uuid.UUID, parentUUID uuid.UUID) ([]domain.ItemToItem, error) {

	parent, err := u.itemBatchRepo.GetByUUID(ctx, parentUUID)

	if err != nil {
		return nil, err
	}

	allowed, err := u.roles.AllowedToEditData(ctx, userUUID, parent.CompanyUUID)

	if err != nil {
		return nil, err
	}

	if !allowed {
		return nil, ErrUnauthorized
	}

	itemToItems, err := u.itemToItemRepo.FindByParentUUID(ctx, parentUUID)

	if err != nil {
		return nil, err
	}

	return itemToItems, nil
}
