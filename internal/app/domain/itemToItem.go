package domain

import "github.com/google/uuid"

type ItemToItem struct {
	UUID       uuid.UUID
	ParentUUID uuid.UUID
	ChildUUID  uuid.UUID
}
