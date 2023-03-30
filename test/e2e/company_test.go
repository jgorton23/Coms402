package e2e

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gentleman.v2"

	client2 "git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/httpclient"
)

// TestCreateCompany e2e test to verify creating a company
func TestCreateCompany(t *testing.T) {
	cli := gentleman.New()
	cli.URL(basePath)
	cli.CookieJar()

	// First create a user to create the company
	email := gofakeit.Email()
	password := gofakeit.Password(true, true, true, true, false, 10)

	res, err := client2.AuthRegister(cli, email, password)
	assert.NoError(t, err, "Failed registering user")
	if res.StatusCode != http.StatusTemporaryRedirect {
		t.Errorf("bad status code %d", res.StatusCode)
	}

	cookies := res.Cookies

	// Second create the company
	companyName := gofakeit.FirstName()
	res, err = client2.CompanyCreate(cli, cookies, companyName)
	assert.NoError(t, err, "Failed creating company")
	if res.StatusCode != http.StatusCreated {
		t.Errorf("bad status code %d", res.StatusCode)
	}

	type Response struct {
		UUID string
		Name string
	}

	// snippet only
	var result Response
	if err := json.Unmarshal([]byte(res.String()), &result); err != nil {
		t.Errorf("Can not unmarshal JSON")
	}

	assert.Equal(t, companyName, result.Name)

	res, err = client2.CompanyGet(cli, cookies, result.UUID)
	assert.NoError(t, err, "Failed getting company")
	if res.StatusCode != http.StatusOK {
		t.Errorf("bad status code %d", res.StatusCode)
	}

	if err := json.Unmarshal([]byte(res.String()), &result); err != nil {
		t.Errorf("Can not unmarshal JSON")
	}

	assert.Equal(t, companyName, result.Name)
}

// TestCreateCompanyDuplicate e2e test to verify duplicate company names are not allowed
func TestCreateCompanyDuplicate(t *testing.T) {
	cli := gentleman.New()
	cli.URL(basePath)
	cli.CookieJar()

	// First create a user to create the company
	email := gofakeit.Email()
	password := gofakeit.Password(true, true, true, true, false, 10)

	res, err := client2.AuthRegister(cli, email, password)
	assert.NoError(t, err, "Failed registering user")
	if res.StatusCode != http.StatusTemporaryRedirect {
		t.Errorf("bad status code %d", res.StatusCode)
	}

	cookies := res.Cookies

	// Second create the company
	companyName := gofakeit.FirstName()
	res, err = client2.CompanyCreate(cli, cookies, companyName)
	assert.NoError(t, err, "Failed creating company")
	if res.StatusCode != http.StatusCreated {
		t.Errorf("bad status code %d", res.StatusCode)
	}

	// Third try and fail to create a duplicate company
	res, _ = client2.CompanyCreate(cli, cookies, companyName)

	// assert.NoError(t, err, "Failed creating company")
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("bad status code %d", res.StatusCode)
	}
}

// TestUpdateCompany e2e test to verify updating works as expected
func TestUpdateCompany(t *testing.T) {
	cli := gentleman.New()
	cli.URL(basePath)
	cli.CookieJar()

	// First create a user to create the company
	email := gofakeit.Email()
	password := gofakeit.Password(true, true, true, true, false, 10)

	res, err := client2.AuthRegister(cli, email, password)
	assert.NoError(t, err, "Failed registering user")
	if res.StatusCode != http.StatusTemporaryRedirect {
		t.Errorf("bad status code %d", res.StatusCode)
	}

	cookies := res.Cookies

	// Second create the company
	companyName := gofakeit.FirstName()
	res, err = client2.CompanyCreate(cli, cookies, companyName)
	assert.NoError(t, err, "Failed creating company")
	if res.StatusCode != http.StatusCreated {
		t.Errorf("bad status code %d", res.StatusCode)
	}

	type Response struct {
		UUID string
		Name string
	}

	// snippet only
	var result Response
	if err := json.Unmarshal([]byte(res.String()), &result); err != nil {
		t.Errorf("Can not unmarshal JSON")
	}

	// Third try and update the company
	companyName2 := gofakeit.FirstName()
	res, err = client2.CompanyUpdate(cli, cookies, companyName2, result.UUID)
	assert.NoError(t, err, "Failed updating company")
	if res.StatusCode != http.StatusOK {
		t.Errorf("bad status code %d", res.StatusCode)
	}

	// snippet only
	var result2 Response
	if err := json.Unmarshal([]byte(res.String()), &result2); err != nil {
		t.Errorf("Can not unmarshal JSON")
	}

	assert.Equal(t, result2.UUID, result.UUID)
	assert.Equal(t, result2.Name, companyName2)
}

// TestUpdateCompany2 e2e test to verify updating non existing companies does not work
func TestUpdateCompany2(t *testing.T) {
	cli := gentleman.New()
	cli.URL(basePath)
	cli.CookieJar()

	// First create a user to create the company
	email := gofakeit.Email()
	password := gofakeit.Password(true, true, true, true, false, 10)

	res, err := client2.AuthRegister(cli, email, password)
	assert.NoError(t, err, "Failed registering user")
	if res.StatusCode != http.StatusTemporaryRedirect {
		t.Errorf("bad status code %d", res.StatusCode)
	}

	cookies := res.Cookies

	type test struct {
		name           string
		uuid           string
		expectedStatus int
	}

	tests := []test{
		{name: gofakeit.FirstName(), uuid: "f3db1bab-dc39-45d7-893a-0000000000", expectedStatus: http.StatusBadRequest}, // Invalid uuid length
		{name: gofakeit.FirstName(), uuid: "f3db1bab-dc39-45d7-893a-00000000", expectedStatus: http.StatusBadRequest},   // non valid uuid
	}

	for _, tc := range tests {
		// Third try and update the company
		res, _ = client2.CompanyUpdate(cli, cookies, tc.name, tc.uuid)

		if res.StatusCode != tc.expectedStatus {
			t.Errorf("Failed")
		}
	}

}
