package repo

import (
	"context"
	"github.com/MatthewBehnke/exampleGoApi/internal/usecase"

	"github.com/samber/do"

	"github.com/MatthewBehnke/exampleGoApi/internal/domain"
	"github.com/MatthewBehnke/exampleGoApi/pkg/database/ent"
	"github.com/MatthewBehnke/exampleGoApi/pkg/database/ent/user"
)

// Pattern to verify userDbEntImplem conforms to the required interfaces
var (
	assertUserImplem                  = &userDbEntImplem{}
	_                usecase.UserRepo = assertUserImplem
)

// NewUserRepo -.
func NewUserRepo(i *do.Injector) (usecase.UserRepo, error) {
	implem := &userDbEntImplem{
		do.MustInvoke[*DatabaseConnection](i).Client(),
	}
	return implem, nil
}

type userDbEntImplem struct {
	*ent.Client
}

func (ur *userDbEntImplem) Get(ctx context.Context) ([]domain.User, error) {
	us, err := ur.Client.User.
		Query().
		All(ctx)

	if err != nil {
		return nil, err
	}

	users := make([]domain.User, len(us))

	for i := 0; i < len(us); i++ {
		users[i] = ur.databaseUserToEntityUser(us[i])
	}

	return users, nil
}

func (ur *userDbEntImplem) GetById(ctx context.Context, id int) (domain.User, error) {
	u, err := ur.Client.User.
		Query().
		Where(user.ID(id)).
		First(ctx)

	if err != nil {
		return domain.User{}, err
	}

	return ur.databaseUserToEntityUser(u), nil
}

func (ur *userDbEntImplem) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := ur.Client.User.
		Query().
		Where(user.Email(email)).
		First(ctx)

	if err != nil {
		return domain.User{}, err
	}

	return ur.databaseUserToEntityUser(u), nil
}

// Exists -.
func (ur *userDbEntImplem) Exists(ctx context.Context, u string) (bool, error) {
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
func (ur *userDbEntImplem) Create(ctx context.Context, usr domain.User) (domain.User, error) {
	u, err := ur.Client.User.
		Create().
		SetEmail(usr.Email).
		SetPasswordHash(usr.PasswordHash).
		// SetConfirmSelector(u.ConfirmSelector).
		// SetConfirmVerifier(u.ConfirmVerifier).
		// SetConfirmed(u.Confirmed).
		SetAttemptCount(usr.AttemptCount).
		SetLastAttempt(usr.LastAttempt).
		SetLocked(usr.Locked).
		Save(ctx)

	if err != nil {
		return domain.User{}, err
	}

	return ur.databaseUserToEntityUser(u), nil
}

// Update -.
func (ur *userDbEntImplem) Update(ctx context.Context, u domain.User) error {
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

func (ur *userDbEntImplem) databaseUserToEntityUser(u *ent.User) domain.User {
	return domain.User{
		CreatedAt:            u.CreatedAt,
		UpdatedAt:            u.UpdatedAt,
		Email:                u.Email,
		PasswordHash:         u.PasswordHash,
		Password:             "",
		PasswordConfirmation: "",
		Role:                 u.Role,
	}
}
