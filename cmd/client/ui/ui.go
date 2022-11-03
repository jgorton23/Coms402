package ui

import (
	"fmt"
	"net/http"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/mitchellh/go-wordwrap"

	"github.com/MatthewBehnke/exampleGoApi/cmd/client/ui/common"
	"github.com/MatthewBehnke/exampleGoApi/cmd/client/ui/landing"
	"github.com/MatthewBehnke/exampleGoApi/cmd/client/ui/login"
	"github.com/MatthewBehnke/exampleGoApi/cmd/client/ui/register"
	"github.com/MatthewBehnke/exampleGoApi/cmd/client/ui/restClient"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/indent"
)

func NewProgram() (*tea.Program, error) {
	m, err := newModel()
	if err != nil {
		return nil, err
	}
	return tea.NewProgram(m, tea.WithAltScreen()), nil
}

type status int

const (
	statusInit status = iota
	statusLanding
	statusRegister
	statusRegistering
	statusLogin
	statusLoggingIn
	statusDisplayResults
	statusQuitting
	statusError
)

func (s status) String() string {
	return [...]string{
		"init",
		"landing",
		"register",
		"registering",
		"login",
		"display",
		"quitting",
		"error",
	}[s]
}

func newModel() (*model, error) {
	c := http.Client{
		Timeout: 10 * time.Second,
	}

	s := spinner.New()
	s.Spinner = common.Spinner

	m := &model{
		status:   statusInit,
		register: register.NewModel(),
		landing:  landing.NewModel(),
		login:    login.NewModel(),

		client:  c,
		spinner: s,
	}

	return m, nil
}

type model struct {
	status status
	err    error

	// winsize is used to keep track of current windows size, it is used to set the size for other models that are initialized in status (e.g. the importlist).
	winsize tea.WindowSizeMsg

	spinner spinner.Model

	register register.Model
	landing  landing.Model
	login    login.Model

	client       http.Client
	clientStatus int
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		func() tea.Msg {
			return restClient.CheckServerConnection(&m.client)
		},
		spinner.Tick,
	)

}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// if _, ok := msg.(spinner.TickMsg); !ok {
	// 	log.Printf("[UI] STATUS: %s | MSG: %#v\n", m.status, msg)
	// }

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.winsize = msg
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			m.status = statusQuitting
			return m, tea.Quit
		}
		switch m.status {
		case statusLanding:
			var cmd tea.Cmd
			m.landing, cmd = m.landing.Update(msg)
			return m, cmd
		case statusRegister:
			var cmd tea.Cmd
			m.register, cmd = m.register.Update(msg)
			m.err = nil
			return m, cmd
		case statusLogin:
			var cmd tea.Cmd
			m.login, cmd = m.login.Update(msg)
			m.err = nil
			return m, cmd
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case common.LandingMsg:
		m.status = statusLanding
	case common.RegisterLandingMsg:
		m.status = statusRegister
		m.err = msg.Err //temp set
		return m, nil
	case common.LoginLandingMsg:
		m.status = statusLogin
		m.err = msg.Err //temp set
		return m, nil
	case common.ServerAvailableMsg:
		m.status = statusLanding
		return m, nil
	case common.RegisterMsg:
		m.status = statusRegistering
		return m, func() tea.Msg {
			return restClient.Register(&m.client, msg.Email, msg.Password, msg.PasswordConfirm)
		}
	case common.LoginMsg:
		m.status = statusLoggingIn
		return m, func() tea.Msg {
			return restClient.Login(&m.client, msg.Email, msg.Password)
		}
	case common.StatusMsg:
		m.status = statusDisplayResults
		m.clientStatus = int(msg)
		return m, nil
	case common.ErrMsg:
		m.status = statusError
		m.err = msg.Err
		return m, nil
	}

	return m, nil
}

func (m model) View() string {
	s := m.logoView()

	switch m.status {
	case statusInit:
		s += m.spinner.View() + "Initializing..."
	case statusLanding:
		s += m.landing.View()
	case statusRegister:
		if m.err != nil {
			s += fmt.Sprintf("Failed to register try again\n")
		}
		s += m.register.View()
	case statusLogin:
		if m.err != nil {
			s += fmt.Sprintf("Failed to register try again\n")
		}
		s += m.login.View()
	case statusRegistering:
		s += m.spinner.View() + "Registering..."
	case statusLoggingIn:
		s += m.spinner.View() + "Logging in..."
	case statusDisplayResults:
		s += fmt.Sprintf("%d %s", m.status, http.StatusText(m.clientStatus))
	case statusError:
		s += errorView(m)
	}

	return indent.String(s, 2)
}

func (m model) logoView() string {
	return "\n" + common.TitleStyle.Render(" CDC messaging ") + "\n\n"
}

func errorView(m model) string {
	return common.ErrorMsgStyle.Render(wordwrap.WrapString(m.err.Error(), uint(m.winsize.Width-2)))
}
