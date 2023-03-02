package repo

import (
	"context"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/samber/do"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/internal/app/usecase"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/session"
)

// Pattern to verify userDBEntImplem conforms to the required interfaces
var (
	_assertSCSImplem           = &scsDBEntImplem{}
	_                scs.Store = _assertSCSImplem
)

// NewUserRepo -.
func NewSCSRepo(i *do.Injector) (scs.Store, error) {
	implem := &scsDBEntImplem{
		db:     do.MustInvoke[*DatabaseConnection](i).Client(),
		logger: do.MustInvoke[*usecase.Logger](i),
	}

	go implem.runCleanup(10 * time.Minute) // TODO load from config or run in a worker system

	return implem, nil
}

type scsDBEntImplem struct {
	db          *ent.Client
	stopCleanup chan bool
	logger      *usecase.Logger
}

func (s scsDBEntImplem) Shutdown() error {
	s.cancelCleanup()
	logger := s.logger.WithSubsystem("session reaper")

	logger.Info("session reaper shutdown")
	return nil
}

func (s scsDBEntImplem) Find(token string) (b []byte, found bool, err error) {
	ses, err := s.db.Session.
		Query().
		// Select(session.FieldData).
		Where(session.And(
			session.Token(token),
			session.ExpiryGT(time.Now()),
		)).
		Only(context.Background())

	if ent.IsNotFound(err) {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}
	return ses.Data, true, nil
}

func (s scsDBEntImplem) Commit(token string, b []byte, expiry time.Time) (err error) {
	err = s.db.Session.
		Create().
		SetToken(token).
		SetData(b).
		SetExpiry(expiry).
		OnConflictColumns(session.FieldToken).
		UpdateNewValues().
		Exec(context.Background())

	if err != nil {
		return err
	}

	return nil
}

func (s scsDBEntImplem) Delete(token string) (err error) {
	_, err = s.db.Session.
		Delete().
		Where(
			session.Token(token),
		).
		Exec(context.Background())

	return err
}

func (s scsDBEntImplem) runCleanup(interval time.Duration) {
	logger := s.logger.WithSubsystem("session reaper")

	logger.Info("session reaper started")
	s.stopCleanup = make(chan bool)
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			logger.Info("reaping expired sessions")
			err := s.deleteExpired()
			if err != nil {
				logger.Warn("Failed to delete expired sessions")
			}
		case <-s.stopCleanup:
			ticker.Stop()
			logger.Info("stopping session reaper")
			return
		}
	}
}

func (s scsDBEntImplem) cancelCleanup() {
	if s.stopCleanup != nil {
		s.stopCleanup <- true
	}
}

func (s scsDBEntImplem) deleteExpired() error {
	_, err := s.db.Session.
		Delete().
		Where(
			session.ExpiryLT(time.Now()),
		).
		Exec(context.Background())

	return err
}
