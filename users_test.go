package okta_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Yapcheekian/okta-go"
	"github.com/stretchr/testify/assert"
)

func TestUsersCreate(t *testing.T) {
	response := okta.UserResponse{
		Id:        "1",
		Status:    "active",
		Created:   "2021-12-12",
		Activated: "true",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/users", func(w http.ResponseWriter, r *http.Request) {
		content, _ := json.Marshal(response)
		fmt.Fprint(w, string(content))
	})

	svr := httptest.NewServer(mux)
	defer svr.Close()

	userService := okta.NewUsersService(&okta.Client{
		OktaDomain: svr.URL,
		Client:     http.DefaultClient,
	})

	param := okta.NewCreateUserParameters("yap", "cheekian", "yapcheekian@gmail.com", "yapcheekian@gmail.com", "123456789")

	resp, err := userService.Create(context.Background(), param, false)

	assert.Nil(t, err)
	assert.Equal(t, response.Id, resp.Id)
	assert.Equal(t, response.Status, resp.Status)
	assert.Equal(t, response.Created, resp.Created)
	assert.Equal(t, response.Activated, resp.Activated)
}
