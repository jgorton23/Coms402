package httpclient

import (
	"fmt"

	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
)

const (
	authPath = "/api/auth"
)

func AuthRegister(cli *gentleman.Client, email string, password string) (*gentleman.Response, error) {
	req := cli.Request()
	req.Path(authPath + "/register")
	req.Method("POST")
	req.Use(body.JSON(
		map[string]interface{}{
			"email":            email,
			"password":         password,
			"confirm_password": password,
		}),
	)

	res, err := req.Send()

	if err != nil {
		return nil, err
	}

	if !res.Ok {
		return nil, fmt.Errorf("invalid server response: %d", res.StatusCode)
	}

	return res, nil
}

func AuthLogin(cli *gentleman.Client, email string, password string) (*gentleman.Response, error) {
	req := cli.Request()
	req.Path(authPath + "/login")
	req.Method("POST")
	req.Use(body.JSON(
		map[string]interface{}{
			"email":    email,
			"password": password,
		}),
	)

	res, err := req.Send()

	if err != nil {
		return nil, err
	}

	if !res.Ok {
		return nil, fmt.Errorf("invalid server response: %d", res.StatusCode)
	}

	return res, nil
}
