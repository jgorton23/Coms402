package domain

import "github.com/google/uuid"

type Role string

const (
	RolePrimaryOwner Role = "primary owner"
	RoleOwner        Role = "owners"
)

type UserToCompany struct {
	UUID        uuid.UUID
	CompanyUUID uuid.UUID
	UserID      int
	RoleType    Role
	Approved    bool
}
