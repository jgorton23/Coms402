package common

type RegisterLandingMsg struct {
	Err error
}

type LoginLandingMsg struct {
	Err error
}

type LandingMsg struct{}

type ErrMsg struct{ Err error }

func (e ErrMsg) Error() string {
	return e.Error()
}

type StatusMsg int

type ServerAvailableMsg struct{}

type RegisterMsg struct {
	Email           string
	Password        string
	PasswordConfirm string
}

type LoginMsg struct {
	Email    string
	Password string
}
