package mefrp

import "fmt"

// GetUserInfo retrieves the current user's information
func (c *Client) GetUserInfo() (*UserInfo, error) {
	var resp Response[UserInfo]
	err := c.request("GET", "/auth/user/info", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return &resp.Data, nil
}

// Sign performs the daily sign-in
func (c *Client) Sign(captchaToken string) error {
	req := CaptchaRequest{CaptchaToken: captchaToken}
	var resp Response[any]
	err := c.request("POST", "/auth/user/sign", req, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// GetUserFrpToken retrieves the user's frp token
func (c *Client) GetUserFrpToken() (string, error) {
var resp Response[struct {
Token string `json:"token"`
}]
err := c.request("GET", "/auth/user/frpToken", nil, &resp)
if err != nil {
return "", err
}

if resp.Code != 200 {
return "", fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
}

return resp.Data.Token, nil
}

// GetUserGroups retrieves the user groups information
func (c *Client) GetUserGroups() ([]UserGroup, error) {
var resp Response[UserGroupsResponse]
err := c.request("GET", "/auth/user/groups", nil, &resp)
if err != nil {
return nil, err
}

if resp.Code != 200 {
return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
}

return resp.Data.Groups, nil
}

// ResetAccessKey resets the user's access key
func (c *Client) ResetAccessKey(captchaToken string) (string, error) {
	req := CaptchaRequest{CaptchaToken: captchaToken}
	var resp Response[ResetTokenResponse]
	err := c.request("POST", "/auth/user/tokenReset", req, &resp)
	if err != nil {
		return "", err
	}

	if resp.Code != 200 {
		return "", fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return resp.Data.NewToken, nil
}

// GetUserLogs retrieves the user's operation logs
func (c *Client) GetUserLogs(page, pageSize int, startTime, endTime string) (*OperationLogList, error) {
path := fmt.Sprintf("/auth/operationLog/list?page=%d&pageSize=%d", page, pageSize)
if startTime != "" {
path += "&startTime=" + startTime
}
if endTime != "" {
path += "&endTime=" + endTime
}

var resp Response[struct {
Data OperationLogList `json:"data"`
}]
err := c.request("GET", path, nil, &resp)
if err != nil {
return nil, err
}

if resp.Code != 200 {
return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
}

return &resp.Data.Data, nil
}

// GetUserLogStats retrieves the user's log statistics
func (c *Client) GetUserLogStats() (*UserLogStats, error) {
var resp Response[UserLogStats]
err := c.request("GET", "/auth/operationLog/stats", nil, &resp)
if err != nil {
return nil, err
}

if resp.Code != 200 {
return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
}

return &resp.Data, nil
}
