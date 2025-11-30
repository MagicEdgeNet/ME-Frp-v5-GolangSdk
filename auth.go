package mefrp

import "fmt"

// GetRegisterEmailCode requests an email verification code for registration
// Requires captcha token for human verification
func (c *Client) GetRegisterEmailCode(email, captchaToken string) error {
	req := struct {
		Email        string `json:"email"`
		CaptchaToken string `json:"captchaToken"`
	}{Email: email, CaptchaToken: captchaToken}

	var resp Response[any]
	err := c.request("POST", "/public/register/emailCode", req, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// RegisterRequest represents the user registration request
type RegisterRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	EmailCode string `json:"emailCode"`
	Password  string `json:"password"`
}

// Register creates a new user account
func (c *Client) Register(req RegisterRequest) error {
	var resp Response[any]
	err := c.request("POST", "/public/register", req, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// LoginRequest represents the login request
type LoginRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	CaptchaToken string `json:"captchaToken"`
}

// LoginResponse represents the login response with token
type LoginResponse struct {
	Token string `json:"token"`
}

// Login authenticates a user and returns a token
// Requires captcha token for human verification
func (c *Client) Login(req LoginRequest) (string, error) {
	var resp Response[struct {
		Token string `json:"token"`
	}]
	err := c.request("POST", "/public/login", req, &resp)
	if err != nil {
		return "", err
	}

	if resp.Code != 200 {
		return "", fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	// Update client token for subsequent requests
	c.token = resp.Data.Token
	return resp.Data.Token, nil
}

// RecoverAccountRequest represents the account recovery request
type RecoverAccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RecoverAccount recovers/resets an account
func (c *Client) RecoverAccount(req RecoverAccountRequest) error {
	var resp Response[any]
	err := c.request("POST", "/public/iforgot", req, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// ChangePasswordRequest represents the password change request
type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

// ChangePassword changes the user's password
// Warning: This will reset the frp token and access key
func (c *Client) ChangePassword(req ChangePasswordRequest) error {
var resp Response[any]
err := c.request("POST", "/auth/user/passwordReset", req, &resp)
if err != nil {
return err
}

if resp.Code != 200 {
return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
}

return nil
}
