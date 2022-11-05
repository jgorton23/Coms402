package chat

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gorilla/websocket"

	"github.com/MatthewBehnke/apis/internal/client/ui/common"
)

type (
	errMsg error
)

type Model struct {
	viewport    viewport.Model
	Messages    []string
	textarea    textarea.Model
	senderStyle lipgloss.Style
	err         error

	websocketConn *websocket.Conn
	Sub           chan common.WsUpdate // where we'll receive activity notifications
	User          string
}

func NewModel() Model {
	ta := textarea.New()
	ta.Placeholder = "Send a message..."
	ta.Focus()

	ta.Prompt = "┃ "
	ta.CharLimit = 280

	ta.SetWidth(30)
	ta.SetHeight(3)

	// Remove cursor line styling
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	ta.ShowLineNumbers = false

	vp := viewport.New(30, 5)
	vp.SetContent(`Welcome to the chat room!
Type a message and press Enter to send.`)

	ta.KeyMap.InsertNewline.SetEnabled(false)

	m := Model{
		textarea:    ta,
		Messages:    []string{},
		viewport:    vp,
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		err:         nil,
		Sub:         make(chan common.WsUpdate),
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return textarea.Blink
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
	)

	m.textarea, tiCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			fmt.Println(m.textarea.Value())
			m.websocketConn.Close()
			return m, tea.Quit
		case tea.KeyEnter:
			err := m.websocketConn.WriteMessage(websocket.TextMessage, []byte(m.senderStyle.Render(m.User+": ")+m.textarea.Value()))
			if err != nil {
				return m, func() tea.Msg {
					return common.ErrMsg{
						Err: err,
					}
				}
			}
			m.textarea.Reset()
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		m.websocketConn.Close()
		return m, nil
	}

	return m, tea.Batch(tiCmd, vpCmd)
}

func (m Model) View() string {
	m.viewport.SetContent(strings.Join(m.Messages, "\n"))
	m.viewport.GotoBottom()

	return fmt.Sprintf(
		"%s\n\n%s",
		m.viewport.View(),
		m.textarea.View(),
	) + "\n\n"
}
