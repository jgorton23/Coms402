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
	ErrItemBatchFound    = errors.New("ItemBatch found")
	ErrItemBatchNotFound = errors.New("ItemBatch not found")
)

func NewItemBatch(i *do.Injector) (*ItemBatch, error) {
	c := &ItemBatch{
		itemBatchRepo: do.MustInvoke[ItemBatchRepo](i),
		roles:         do.MustInvoke[*UserToCompany](i),
		logger:        do.MustInvoke[*Logger](i),
	}

	return c, nil
}

type ItemBatch struct {
	itemBatchRepo ItemBatchRepo
	roles         *UserToCompany
	logger        *Logger
}

// Create
// creates a new itemBatch
func (u ItemBatch) Create(ctx context.Context, ib domain.ItemBatch, userUUID uuid.UUID) (domain.ItemBatch, error) {
	// verify a user is allowed to create an item batch user a company
	ok, err := u.roles.AllowedToEditData(ctx, userUUID, ib.CompanyUUID)

	if err != nil {
		return domain.ItemBatch{}, err
	}

	if !ok {
		return domain.ItemBatch{}, ErrUnauthorized
	}

	exists, err := u.itemBatchRepo.Exists(ctx, ib.ItemNumber)

	if err != nil {
		return domain.ItemBatch{}, err
	}

	if exists {
		u.logger.Info("Item Batch item number already exists")
		return domain.ItemBatch{}, ErrItemBatchFound
	}

	repoItemBatch, err := u.itemBatchRepo.Create(ctx, ib)

	if err != nil {
		return domain.ItemBatch{}, err
	}

	u.logger.Info(fmt.Sprintf("Item Batch %v created in database", repoItemBatch.ItemNumber))

	return repoItemBatch, nil
}

// Update
// updates the given itemBatch
func (u ItemBatch) Update(ctx context.Context, ib domain.ItemBatch, userUUID uuid.UUID) error {
	// verify a user is allowed to create an item batch user a company
	ok, err := u.roles.AllowedToEditData(ctx, userUUID, ib.CompanyUUID)

	if err != nil {
		return err
	}

	if !ok {
		return ErrUnauthorized
	}

	exists, err := u.itemBatchRepo.Exists(ctx, ib.ItemNumber)

	if err != nil {
		return err
	}

	if !exists {
		u.logger.Info("Item Batch item number already exists")
		return ErrItemBatchNotFound
	}

	err = u.itemBatchRepo.Update(ctx, ib)

	if err != nil {
		return err
	}

	u.logger.Info(fmt.Sprintf("Item Batch %v updated in database", ib.ItemNumber))

	return nil
}

// GetByCompanyUUID
// returns all itembatches belonging to a given company
func (u ItemBatch) GetByCompanyUUID(ctx context.Context, companyUUID uuid.UUID, userUUID uuid.UUID) ([]domain.ItemBatch, error) {
	//TODO switch to using the RBAC system / change how roles work
	ok, err := u.roles.AllowedToEditData(ctx, userUUID, companyUUID)

	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, ErrUnauthorized
	}

	itemBatches, err := u.itemBatchRepo.GetByCompanyUUID(ctx, companyUUID)

	if err != nil {
		return nil, err
	}

	return itemBatches, nil
}

// GetByUUID
// returns the itemBatch with the given UUID
func (u ItemBatch) GetByUUID(ctx context.Context, companyUUID uuid.UUID, userUUID uuid.UUID, uuid uuid.UUID) (domain.ItemBatch, error) {
	//TODO switch to using the RBAC system / change how roles work
	ok, err := u.roles.AllowedToEditData(ctx, userUUID, companyUUID)

	if err != nil {
		return domain.ItemBatch{}, err
	}

	if !ok {
		return domain.ItemBatch{}, ErrUnauthorized
	}

	itemBatches, err := u.itemBatchRepo.GetByUUID(ctx, uuid)

	if err != nil {
		return domain.ItemBatch{}, err
	}

	return itemBatches, nil
}
