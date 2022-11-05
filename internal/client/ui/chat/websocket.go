package chat

import (
	"errors"
	"net/http"
	"os"
	"os/signal"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gorilla/websocket"

	"github.com/MatthewBehnke/apis/internal/client/ui/common"
)

func Chat(httpClient *http.Client, m *Model) error {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	websocket.DefaultDialer.Jar = httpClient.Jar
	c, resp, err := websocket.DefaultDialer.Dial("ws://"+os.Getenv("APP_HOST")+"/v1/ws", nil)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusForbidden ||
		resp.StatusCode == http.StatusUnauthorized {
		return errors.New("not authenticated")
	}

	m.websocketConn = c
	return nil
}

func ListenForChatUpdates(sub chan common.WsUpdate, m *Model) tea.Cmd {
	return func() tea.Msg {
		for {
			_, message, err := m.websocketConn.ReadMessage()
			if err != nil {
				return func() tea.Msg {
					return common.ErrMsg{
						Err: err,
					}
				}
			}

			//m.messages = append(m.messages, string(message))
			sub <- common.WsUpdate{
				Message: string(message),
			}
		}
	}
}

func WaitForActivity(sub chan common.WsUpdate) tea.Cmd {
	return func() tea.Msg {
		return common.ChatUpdateMsg(<-sub)
	}
}
