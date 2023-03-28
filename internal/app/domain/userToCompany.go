package domain

import "github.com/google/uuid"

type Role string

const (
	RolePrimaryOwner Role = "primary owner"
	RoleOwner        Role = "owners"
	RoleUser         Role = "user"
)

type UserToCompany struct {
	UUID        uuid.UUID
	CompanyUUID uuid.UUID
	UserUUID    uuid.UUID
	RoleType    Role
	Approved    bool
}
