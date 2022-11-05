package restClient

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"

	"github.com/MatthewBehnke/exampleGoApi/internal/client/ui/common"
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

	res, err := c.Get("http://" + os.Getenv("APP_HOST"))
	if err != nil {
		return common.ErrMsg{Err: err}
	}
	defer res.Body.Close()

	return common.ServerAvailableMsg{}
	//return common.StatusMsg(res.StatusCode)
}

func Register(c *http.Client, email, password, passwordConfirm string) tea.Msg {
	u, err := url.Parse("http://" + os.Getenv("APP_HOST") + "/auth/register")
	if err != nil {
		return common.ErrMsg{Err: err}
	}

	values := map[string]string{
		"email":            email,
		"password":         password,
		"confirm_password": passwordConfirm,
	}

	jsonData, _ := json.Marshal(values)

	request, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	response, err := c.Do(request)
	if err != nil {
		return common.ErrMsg{Err: err}
	}

	c.Jar.SetCookies(u, response.Cookies())

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
}

func Login(c *http.Client, email, password string) tea.Msg {
	u, err := url.Parse("http://" + os.Getenv("APP_HOST") + "/auth/login")
	if err != nil {
		return common.ErrMsg{Err: err}
	}

	values := map[string]string{
		"email":    email,
		"password": password,
	}

	jsonData, _ := json.Marshal(values)

	request, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	response, err := c.Do(request)
	if err != nil {
		return common.ErrMsg{Err: err}
	}

	c.Jar.SetCookies(u, response.Cookies())

	var bod registerResponse
	err = json.NewDecoder(response.Body).Decode(&bod)
	if err != nil {
		return common.ErrMsg{Err: err}
	}

	if bod.Status == "failure" {
		return common.LoginLandingMsg{Err: errors.New("failed to login user")}
	}

	defer response.Body.Close()
	return common.ServerAvailableMsg{}

}
