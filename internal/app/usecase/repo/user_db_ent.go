package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/samber/do"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/user"
)

// Pattern to verify userDBEntImplem conforms to the required interfaces
var (
	_assertUserImplem                  = &userDBEntImplem{}
	_                 usecase.UserRepo = _assertUserImplem
)

// NewUserRepo -.
func NewUserRepo(i *do.Injector) (usecase.UserRepo, error) {
	implem := &userDBEntImplem{
		do.MustInvoke[*DatabaseConnection](i).Client(),
	}

	return implem, nil
}

type userDBEntImplem struct {
	*ent.Client
}

// Get
// returns all users
func (ur *userDBEntImplem) Get(ctx context.Context) ([]domain.User, error) {
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

// GetByUUID
// returns the user with the give UUID
func (ur *userDBEntImplem) GetByUUID(ctx context.Context, uuid uuid.UUID) (domain.User, error) {
	u, err := ur.Client.User.
		Query().
		Where(user.ID(uuid)).
		First(ctx)
	if err != nil {
		return domain.User{}, err
	}

	return ur.databaseUserToEntityUser(u), nil
}

// GetByEmail
// returns the user with the given email
func (ur *userDBEntImplem) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := ur.Client.User.
		Query().
		Where(user.Email(email)).
		First(ctx)
	if err != nil {
		return domain.User{}, err
	}

	return ur.databaseUserToEntityUser(u), nil
}

// Exists
// returns if a user with the given email exists
func (ur *userDBEntImplem) Exists(ctx context.Context, u string) (bool, error) {
	exists, err := ur.Client.User.
		Query().
		Where(user.Email(u)).
		Exist(ctx)
	if err != nil {
		return false, err
	}

	return exists, nil
}

// Create
// creates a new user
func (ur *userDBEntImplem) Create(ctx context.Context, usr domain.User) (domain.User, error) {
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

// Update
// updates the given user
func (ur *userDBEntImplem) Update(ctx context.Context, u domain.User) error {
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

func (ur *userDBEntImplem) databaseUserToEntityUser(u *ent.User) domain.User {
	return domain.User{
		UUID:                 u.ID,
		CreatedAt:            u.CreatedAt,
		UpdatedAt:            u.UpdatedAt,
		Email:                u.Email,
		PasswordHash:         u.PasswordHash,
		Password:             "",
		PasswordConfirmation: "",
		Role:                 u.Role,
	}
}
