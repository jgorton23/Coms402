package restClient

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/MatthewBehnke/exampleGoApi/cmd/client/ui/common"
	tea "github.com/charmbracelet/bubbletea"
)

type registerResponse struct {
	//Errors struct {
	//	Error []string `json:""`
	//} `json:"errors"`
	//Modules struct {
	//	Auth     bool `json:"auth"`
	//	Logout   bool `json:"logout"`
	//	Register bool `json:"register"`
	//} `json:"modules"`
	//Preserve struct {
	//} `json:"preserve"`
	Status string `json:"status"`
}

func CheckServerConnection(c *http.Client) tea.Msg {

	res, err := c.Get("http://localhost:8082")
	if err != nil {
		return common.ErrMsg{Err: err}
	}
	defer res.Body.Close()

	return common.ServerAvailableMsg{}
	//return common.StatusMsg(res.StatusCode)
}

func Register(c *http.Client, email, password, passwordConfirm string) tea.Msg {

	httpposturl := "http://localhost:8082/auth/register"

	values := map[string]string{
		"email":            email,
		"password":         password,
		"confirm_password": passwordConfirm,
	}

	jsonData, _ := json.Marshal(values)

	request, err := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	response, err := c.Do(request)
	if err != nil {
		return common.ErrMsg{Err: err}
	}

	var bod registerResponse
	err = json.NewDecoder(response.Body).Decode(&bod)
	if err != nil {
		return common.ErrMsg{Err: err}
	}

	if bod.Status == "failure" {
		return common.RegisterLandingMsg{Err: errors.New("failed to register user")}
	}

	defer response.Body.Close()
	return common.ServerAvailableMsg{}
	//return common.StatusMsg(res.StatusCode)
}

func Login(c *http.Client, email, password string) tea.Msg {
	httpposturl := "http://localhost:8082/auth/login"

	values := map[string]string{
		"email":    email,
		"password": password,
	}

	jsonData, _ := json.Marshal(values)

	request, err := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	response, err := c.Do(request)
	if err != nil {
		return common.ErrMsg{Err: err}
	}

	var bod registerResponse
	err = json.NewDecoder(response.Body).Decode(&bod)
	if err != nil {
		return common.ErrMsg{Err: err}
	}

	if bod.Status == "failure" {
		return common.RegisterLandingMsg{Err: errors.New("failed to login user")}
	}

	defer response.Body.Close()
	return common.ServerAvailableMsg{}

}
