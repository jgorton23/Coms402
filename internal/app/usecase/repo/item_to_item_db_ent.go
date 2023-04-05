package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/samber/do"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent"
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

func (ur itemToItemDBEntImplem) Create(ctx context.Context, parentUUID uuid.UUID, childUUID uuid.UUID) error {
	return nil
}
