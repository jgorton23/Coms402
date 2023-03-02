package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/volatiletech/authboss/v3"
)

// This pattern is useful in real code to ensure that
// we've got the right interfaces implemented.
var (
	_assertUser = &User{}

	_ authboss.User         = _assertUser
	_ authboss.AuthableUser = _assertUser
	// _ authboss.ConfirmableUser = _assertUser
	// _ authboss.LockableUser = _assertUser
	// _ authboss.RecoverableUser = _assertUser
	// _ authboss.ArbitraryUser = _assertUser
)

type User struct {
	UUID         uuid.UUID
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`

	Password             string `json:"-"`
	PasswordConfirmation string `json:"-"`

	// // Confirm
	// ConfirmSelector string
	// ConfirmVerifier string
	// Confirmed       bool

	// Lock
	AttemptCount int
	LastAttempt  time.Time
	Locked       time.Time

	// // Recover
	// RecoverSelector    string
	// RecoverVerifier    string
	// RecoverTokenExpiry time.Time

	Role string
}

// GetPID from user
func (u User) GetPID() string { return u.Email }

// PutPID into user
func (u *User) PutPID(pid string) { u.Email = pid }

// GetPassword from user
func (u User) GetPassword() string { return u.PasswordHash }

// PutPassword into user
func (u *User) PutPassword(password string) { u.PasswordHash = password }

// GetEmail from user
func (u User) GetEmail() string { return u.Email }

// PutEmail into user
func (u *User) PutEmail(email string) { u.Email = email }

// GetLocked from user
func (u User) GetLocked() time.Time { return u.Locked }

// PutLocked into user
func (u *User) PutLocked(locked time.Time) { u.Locked = locked }

// GetAttemptCount from user
func (u User) GetAttemptCount() int { return u.AttemptCount }

// PutAttemptCount into user
func (u *User) PutAttemptCount(attempts int) { u.AttemptCount = attempts }

// GetLastAttempt from user
func (u User) GetLastAttempt() time.Time { return u.LastAttempt }

// PutLastAttempt into user
func (u *User) PutLastAttempt(last time.Time) { u.LastAttempt = last }

// // GetConfirmSelector from user
// func (u User) GetConfirmSelector() string { return u.ConfirmSelector }

// // PutConfirmSelector into user
// func (u *User) PutConfirmSelector(confirmSelector string) { u.ConfirmSelector = confirmSelector }

// // GetConfirmVerifier from user
// func (u User) GetConfirmVerifier() string { return u.ConfirmVerifier }

// // PutConfirmVerifier into user
// func (u *User) PutConfirmVerifier(confirmVerifier string) { u.ConfirmVerifier = confirmVerifier }

// // GetConfirmed from user
// func (u User) GetConfirmed() bool { return u.Confirmed }

// // PutConfirmed into user
// func (u *User) PutConfirmed(confirmed bool) { u.Confirmed = confirmed }
