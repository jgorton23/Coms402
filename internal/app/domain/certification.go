package domain

import "github.com/google/uuid"

type Certification struct {
	UUID             uuid.UUID
	PrimaryAttribute string
	ImageUUID        uuid.UUID
	ItemBatchUUID    uuid.UUID
	TemplateUUID     uuid.UUID
	CompanyUUID      uuid.UUID
}
