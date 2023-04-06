package usecase_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/samber/do"
	"github.com/stretchr/testify/assert"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase/repo"
)

// Reuseable user usecase test setup
func usecaseUserTestSetup(t *testing.T, fun func(i *do.Injector) (usecase.UserRepo, error)) (*usecase.User, error) {
	injector := do.New()

	// The usecase requires the logging interface to be fulfilled but
	// during unit tests we do not want any logs, so we are using a noop logger
	do.Provide(injector, func(i *do.Injector) (usecase.LoggerRepo, error) {
		return repo.LoggerNoopBuilder().New()
	})
	do.Provide(injector, usecase.NewLogger)
	do.Provide(injector, fun)
	do.Provide(injector, usecase.NewUser)

	return do.MustInvoke[*usecase.User](injector), nil
}

func TestUsecaseUserFindByEmail(t *testing.T) {
	tests := []struct {
		mock func(i *do.Injector) (usecase.UserRepo, error)
		call func(*usecase.User)
	}{
		{ // Testing
			mock: func(i *do.Injector) (usecase.UserRepo, error) {
				ctrl := gomock.NewController(t)
				mur := NewMockUserRepo(ctrl)
				mur.
					EXPECT().
					GetByEmail(context.Background(), "test@test.com").
					Return(domain.User{
						Email: "test@test.com",
					}, nil)

				return mur, nil
			},
			call: func(u *usecase.User) {
				foundUser, err := u.FindByEmail(context.Background(), "test@test.com")

				assert.NoError(t, err)
				assert.Equal(t, foundUser.Email, "test@test.com")
			},
		},
	}

	for _, v := range tests {
		user, err := usecaseUserTestSetup(t, v.mock)
		assert.NoError(t, err)

		v.call(user)
	}
}
