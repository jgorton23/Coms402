package httpclient

import (
	"fmt"
	"net/http"

	"gopkg.in/h2non/gentleman.v2"
)

func Healthz(cli *gentleman.Client) error {
	req := cli.Request()
	req.Path("/healthz")
	req.Method("GET")

	res, err := req.Send()

	if err != nil {
		return err
	}
	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return fmt.Errorf("invalid server response: %d", res.StatusCode)
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status code %d", res.StatusCode)
	}

	return nil
}
