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

func (ur *UserRepo) GetById(ctx context.Context, id int) (entity.User, error) {
	user, err := ur.Client.User.
		Query().
		Where(user.ID(id)).
		First(ctx)

	if err != nil {
		return entity.User{}, err
	}

	return ur.databaseUserToEntityUser(user), nil
}

func (ur *UserRepo) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	user, err := ur.Client.User.
		Query().
		Where(user.Email(email)).
		First(ctx)

	if err != nil {
		return entity.User{}, err
	}

	return ur.databaseUserToEntityUser(user), nil
}

// Exists -.
func (ur *UserRepo) Exists(ctx context.Context, u string) (bool, error) {
	exists, err := ur.Client.User.
		Query().
		Where(user.Email(u)).
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
		// SetConfirmSelector(u.ConfirmSelector).
		// SetConfirmVerifier(u.ConfirmVerifier).
		// SetConfirmed(u.Confirmed).
		SetAttemptCount(u.AttemptCount).
		SetLastAttempt(u.LastAttempt).
		SetLocked(u.Locked).
		Save(ctx)

	if err != nil {
		return entity.User{}, err
	}

	return ur.databaseUserToEntityUser(user), nil
}

// Create -.
func (ur *UserRepo) Update(ctx context.Context, u entity.User) error {
	_, err := ur.Client.User.
		Update().
		SetEmail(u.Email).
		SetPasswordHash(u.PasswordHash).
		// SetConfirmSelector(u.ConfirmSelector).
		// SetConfirmVerifier(u.ConfirmVerifier).
		// SetConfirmed(u.Confirmed).
		SetAttemptCount(u.AttemptCount).
		SetLastAttempt(u.LastAttempt).
		SetLocked(u.Locked).
		Save(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepo) databaseUserToEntityUser(u *ent.User) entity.User {
	return entity.User{
		CreatedAt:            u.CreatedAt,
		UpdatedAt:            u.UpdatedAt,
		Email:                u.Email,
		PasswordHash:         u.PasswordHash,
		Password:             "",
		PasswordConfirmation: "",
	}
}
