package repo

import (
	"context"
	"iseage/bank/internal/entity"
	"iseage/bank/pkg/database"
	"iseage/bank/pkg/database/ent"
	"iseage/bank/pkg/database/ent/user"
)

type UserRepo struct {
	*database.Database
}

// NewUserRepo -.
func NewUserRepo(db *database.Database) *UserRepo {
	return &UserRepo{db}
}

// Create -.
func (ur *UserRepo) Exists(ctx context.Context, u entity.User) (bool, error) {
	exists, err := ur.Client.User.
		Query().
		Where(user.Email(u.Email)).
		Exist(ctx)

	if err != nil {
		return false, err
	}

	return exists, nil
}

// Create -.
func (ur *UserRepo) Create(ctx context.Context, u entity.User) (entity.User, error) {
	user, err := ur.Client.User.
		Create().
		SetEmail(u.Email).
		SetPasswordHash(u.PasswordHash).
		Save(ctx)

	if err != nil {
		return entity.User{}, err
	}

	return ur.databaseUserToEntityUser(user), nil
}

func (ur *UserRepo) databaseUserToEntityUser(u *ent.User) entity.User {
	return entity.User{
		ID:                   u.ID,
		CreatedAt:            u.CreatedAt,
		UpdatedAt:            u.UpdatedAt,
		Email:                u.Email,
		PasswordHash:         u.PasswordHash,
		Password:             "",
		PasswordConfirmation: "",
	}
}
