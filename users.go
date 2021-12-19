package okta

import "context"

const (
	USERS_API = "/api/v1/users"
)

type Profile struct {
	Login             string
	FirstName         string
	LastName          string
	NickName          string
	DisplayName       string
	Email             string
	SecondEmail       string
	ProfileUrl        string
	PreferredLanguage string
}

type CreateUserParameters struct {
	Activate bool
	Provider bool
	Profile  Profile
}

func (c *Client) CreateUser(ctx context.Context, params CreateUserParameters) {
	c.postRequest(ctx, USERS_API, params)
}

func (c *Client) GetUser(ctx context.Context, userId string) {
	c.getRequest(ctx, USERS_API+"/"+userId)
}

func (c *Client) ListUsers(ctx context.Context) {
	c.getRequest(ctx, USERS_API)
}

func (c *Client) DeleteUser(ctx context.Context, params CreateUserParameters) {
	c.postRequest(ctx, USERS_API, params)
}
