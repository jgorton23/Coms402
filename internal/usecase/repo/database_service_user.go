package repo

import (
	"context"

	"github.com/samber/do"

	"github.com/MatthewBehnke/exampleGoApi/internal/entity"
	"github.com/MatthewBehnke/exampleGoApi/pkg/database/ent"
	"github.com/MatthewBehnke/exampleGoApi/pkg/database/ent/user"
)

// NewDataBaseServiceUser -.
func NewDataBaseServiceUser(i *do.Injector) (*DataBaseServiceUser, error) {
	dbServiceUser := &DataBaseServiceUser{}
	dbServiceUser.Client = do.MustInvoke[*DatabaseService](i).Client()

	return dbServiceUser, nil
}

type DataBaseServiceUser struct {
	*ent.Client
}

func (ur *DataBaseServiceUser) Get(ctx context.Context) ([]entity.User, error) {
	us, err := ur.Client.User.
		Query().
		All(ctx)

	if err != nil {
		return nil, err
	}

	users := make([]entity.User, len(us))

	for i := 0; i < len(us); i++ {
		users[i] = ur.databaseUserToEntityUser(us[i])
	}

	return users, nil
}

func (ur *DataBaseServiceUser) GetById(ctx context.Context, id int) (entity.User, error) {
	user, err := ur.Client.User.
		Query().
		Where(user.ID(id)).
		First(ctx)

	if err != nil {
		return entity.User{}, err
	}

	return ur.databaseUserToEntityUser(user), nil
}

func (ur *DataBaseServiceUser) GetByEmail(ctx context.Context, email string) (entity.User, error) {
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
func (ur *DataBaseServiceUser) Exists(ctx context.Context, u string) (bool, error) {
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
func (ur *DataBaseServiceUser) Create(ctx context.Context, u entity.User) (entity.User, error) {
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
func (ur *DataBaseServiceUser) Update(ctx context.Context, u entity.User) error {
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

func (ur *DataBaseServiceUser) databaseUserToEntityUser(u *ent.User) entity.User {
	return entity.User{
		CreatedAt:            u.CreatedAt,
		UpdatedAt:            u.UpdatedAt,
		Email:                u.Email,
		PasswordHash:         u.PasswordHash,
		Password:             "",
		PasswordConfirmation: "",
	}
}
