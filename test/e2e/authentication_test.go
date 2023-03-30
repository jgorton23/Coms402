package e2e

import (
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gentleman.v2"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/httpclient"
)

// TestAuthRegister e2e to test registration
func TestAuthRegister(t *testing.T) {
	email := gofakeit.Email()
	password := gofakeit.Password(true, true, true, true, false, 10)

	t.Log(email)
	t.Log(password)

	cli := gentleman.New()
	cli.URL(basePath)
	cli.CookieJar()

	res, err := httpclient.AuthRegister(cli, email, password)
	assert.NoError(t, err, "Failed registering user")

	if res.StatusCode != http.StatusTemporaryRedirect {
		t.Errorf("bad status code %d", res.StatusCode)
	}
}

// TestAuthLogin e2e to test logging in
func TestAuthLogin(t *testing.T) {
	email := gofakeit.Email()
	password := gofakeit.Password(true, true, true, true, false, 10)

	t.Log(email)
	t.Log(password)

	cli := gentleman.New()
	cli.URL(basePath)
	cli.CookieJar()

	res, err := httpclient.AuthRegister(cli, email, password)
	assert.NoError(t, err, "Failed registering user")
	if res.StatusCode != http.StatusTemporaryRedirect {
		t.Errorf("bad status code %d", res.StatusCode)
	}

	res, err = httpclient.AuthLogin(cli, email, password)
	assert.NoError(t, err, "Failed logging in user")

	if res.StatusCode != http.StatusTemporaryRedirect {
		t.Errorf("bad status code %d", res.StatusCode)
	}
}
