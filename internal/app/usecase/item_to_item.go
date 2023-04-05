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
		roles:          do.MustInvoke[*UserToCompany](i),
		logger:         do.MustInvoke[*Logger](i),
	}

	return c, nil
}

type ItemToItem struct {
	itemToItemRepo ItemToItemRepo
	roles          *UserToCompany
	logger         *Logger
}

func (u ItemToItem) Create(ctx context.Context, iti domain.ItemToItem, userUUID uuid.UUID, companyUUID uuid.UUID) error {
	return nil
}
