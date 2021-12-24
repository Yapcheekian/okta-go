package okta

import (
	"context"
	"fmt"
)

const (
	USERS_API = "/api/v1/users"
)

type UsersService struct {
	c *Client
}

func NewUsersService(client *Client) *UsersService {
	return &UsersService{c: client}
}

type Profile struct {
	Login             string `json:"login,omitempty"`
	FirstName         string `json:"firstName,omitempty"`
	LastName          string `json:"lastName,omitempty"`
	NickName          string `json:"nickName,omitempty"`
	DisplayName       string `json:"displayName,omitempty"`
	Email             string `json:"email,omitempty"`
	SecondEmail       string `json:"secondEmail,omitempty"`
	ProfileUrl        string `json:"profileUrl,omitempty"`
	PreferredLanguage string `json:"preferredLanguage,omitempty"`
	UserType          string `json:"userType,omitempty"`
	Organization      string `json:"organization,omitempty"`
	Title             string `json:"title,omitempty"`
	Division          string `json:"division,omitempty"`
	Department        string `json:"department,omitempty"`
	CostCenter        string `json:"costCenter,omitempty"`
	MobilePhone       string `json:"mobilePhone,omitempty"`
}

type CreateUserParameters struct {
	Activate bool    `json:"activate,omitempty"`
	Provider bool    `json:"provider,omitempty"`
	Profile  Profile `json:"profile,omitempty"`
}

type UserResponse struct {
	Id        string
	Status    string
	Created   string
	Activated string
	Profile   Profile
}

func NewCreateUserParameters(firstName, lastName, email, login, mobilePhone string) *CreateUserParameters {
	return &CreateUserParameters{
		Profile: Profile{
			FirstName:   firstName,
			LastName:    lastName,
			Email:       email,
			Login:       login,
			MobilePhone: mobilePhone,
		},
	}
}

func (s *UsersService) Create(ctx context.Context, params *CreateUserParameters, debug bool) (*UserResponse, error) {
	resp, err := postRequest(ctx, s.c.Client, s.c.OktaDomain, USERS_API, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("response status code %d", resp.StatusCode)
	}
	var response UserResponse
	if err := parseResponse(resp.Body, &response, debug); err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *UsersService) Get(ctx context.Context, userId string, debug bool) (*UserResponse, error) {
	var response UserResponse
	resp, err := getRequest(ctx, s.c.Client, s.c.OktaDomain, USERS_API+"/"+userId)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := parseResponse(resp.Body, &response, debug); err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *UsersService) List(ctx context.Context) {
	// TODO
	getRequest(ctx, s.c.Client, s.c.OktaDomain, USERS_API)
}

func (s *UsersService) Delete(ctx context.Context, params CreateUserParameters) {
	// TODO
	postRequest(ctx, s.c.Client, s.c.OktaDomain, USERS_API, params)
}
