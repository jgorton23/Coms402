package httpclient

import (
	"fmt"
	"net/http"

	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
)

func CompanyCreate(cli *gentleman.Client, cookies []*http.Cookie, name string) (*gentleman.Response, error) {
	req := cli.Request()
	req.AddCookies(cookies)
	req.Path("/api/v1/company")
	req.Method("POST")
	req.Use(body.JSON(
		map[string]interface{}{
			"name": name,
		}),
	)

	res, err := req.Send()

	if err != nil {
		return nil, err
	}

	// if !res.Ok {
	// 	return nil, fmt.Errorf("invalid server response: %d", res.StatusCode)
	// }

	return res, nil
}

func CompanyGet(cli *gentleman.Client, cookies []*http.Cookie, uuid string) (*gentleman.Response, error) {
	req := cli.Request()
	req.AddCookies(cookies)
	req.Path("/api/v1/company/" + uuid)
	req.Method("GET")

	res, err := req.Send()

	if err != nil {
		return nil, err
	}

	if !res.Ok {
		return nil, fmt.Errorf("invalid server response: %d", res.StatusCode)
	}

	return res, nil
}

func CompanyUpdate(cli *gentleman.Client, cookies []*http.Cookie, name string, uuid string) (*gentleman.Response, error) {
	req := cli.Request()
	req.AddCookies(cookies)
	req.Path("/api/v1/company")
	req.Method("PUT")
	req.Use(body.JSON(
		map[string]interface{}{
			"name": name,
			"uuid": uuid,
		}),
	)

	res, err := req.Send()

	if err != nil {
		return nil, err
	}

	return res, nil
}
