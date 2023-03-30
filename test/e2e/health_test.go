package e2e

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gentleman.v2"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/httpclient"
)

func TestHealth(t *testing.T) {
	cli := gentleman.New()
	cli.URL(basePath)
	cli.CookieJar()

	err := httpclient.Healthz(cli)
	assert.NoError(t, err, "Failed health testing")
}
