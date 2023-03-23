package usecase_test

// import (
// 	"testing"

// 	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/domain"
// 	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
// 	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase/repo"
// 	"github.com/golang/mock/gomock"
// 	"github.com/samber/do"
// )

// func TestUsecaseUserFindByEmail(t *testing.T) {
// 	injector := do.New()
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	do.ProvideValue(injector, &domain.Config{
// 		LOG: domain.LOG{
// 			Format:  "json",
// 			Level:   "info",
// 			NoColor: true,
// 		},
// 	})
// 	do.Provide(injector, repo.NewLoggerRepo)
// 	do.Provide(injector, usecase.NewLogger)

// 	do.Provide(injector, func(i *do.Injector) (usecase.UserRepo, error) {
// 		return NewMockUserRepo(ctrl), nil
// 	})
// 	do.Provide(injector, usecase.NewUser)

// 	_ = do.MustInvoke[*usecase.User](injector)
// 	// ur := do.MustInvoke[usecase.UserRepo](injector)


// 	m := NewMockUserRepo(ctrl)

// 	m.EXPECT()

// 	// _, err := user.FindByEmail(context.Background(), "test@test.com")

// 	// if assert.Error(t, err) {
// 	// 	t.Fail()
// 	// }

// }
