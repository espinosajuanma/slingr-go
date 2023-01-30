package slingr

import (
	"encoding/json"
)

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	App                            string `json:"app"`
	PhoneNumber                    string `json:"phoneNumber"`
	AdminUserEmail                 string `json:"adminUserEmail"`
	TwoFactorAuthenticationChecked bool   `json:"twoFactorAuthenticationChecked"`
	IP                             string `json:"ip"`
	GroupRequire2FA                bool   `json:"groupRequire2FA"`
	TwoFactorAuthentication        bool   `json:"twoFactorAuthentication"`
	UserEmail                      string `json:"userEmail"`
	UserName                       string `json:"userName"`
	UserID                         string `json:"userId"`
	Token                          string `json:"token"`
}

func (c *App) Login(email, password string) (*LoginResponse, error) {
	payload := &LoginPayload{Email: email, Password: password}
	response, err := c.Post("/auth/login", payload, nil)
	if err != nil {
		return nil, err
	}

	var loginResponse LoginResponse
	err = json.Unmarshal(response, &loginResponse)
	if err != nil {
		return nil, err
	}

	c.Email = loginResponse.UserEmail
	c.Token = loginResponse.Token
	return &loginResponse, nil
}

func (c *App) Logout() ([]byte, error) {
	return c.Post("/auth/logout", nil, nil)
}
