package e2e

import (
	"net/http"
	"testing"

	"github.com/Eun/go-hit"
)

const (
	basePath  = "http://localhost:8082"
	apiV1Path = "/api/v1"
	authPath  = "/auth"
)

func TestHealth(t *testing.T) {
	err := hit.Do(
		hit.Get(basePath+"/healthz"),
		hit.Expect().Status().Equal(http.StatusOK),
	)

	if err != nil {
		t.Errorf(err.Error())
	}
}
