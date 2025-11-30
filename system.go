package mefrp

import "fmt"

// GetSystemStatus retrieves the system status
func (c *Client) GetSystemStatus() (*SystemStatus, error) {
	var resp Response[SystemStatus]
	err := c.request("GET", "/auth/system/status", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return &resp.Data, nil
}

// GetPopupNotice retrieves the important popup notice
func (c *Client) GetPopupNotice() (string, error) {
	var resp Response[string]
	err := c.request("GET", "/auth/popupNotice", nil, &resp)
	if err != nil {
		return "", err
	}

	if resp.Code != 200 {
		return "", fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return resp.Data, nil
}
