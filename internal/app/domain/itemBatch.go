package domain

import "github.com/google/uuid"

type ItemBatch struct {
	UUID        uuid.UUID
	ItemNumber  string
	Description string
	CompanyUUID uuid.UUID
}
