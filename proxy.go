package mefrp

import "fmt"

// GetProxyList retrieves the list of proxies for the current user
func (c *Client) GetProxyList() ([]Proxy, error) {
	var resp Response[[]Proxy]
	err := c.request("GET", "/auth/proxy/list", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return resp.Data, nil
}

// CreateProxy creates a new proxy
func (c *Client) CreateProxy(req CreateProxyRequest) error {
	var resp Response[any]
	err := c.request("POST", "/auth/proxy/create", req, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// DeleteProxy deletes a proxy by ID
func (c *Client) DeleteProxy(proxyID int) error {
	req := IDRequest{ProxyID: proxyID}
	var resp Response[any]
	err := c.request("POST", "/auth/proxy/delete", req, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// UpdateProxy updates an existing proxy
func (c *Client) UpdateProxy(req CreateProxyRequest) error {
	var resp Response[any]
	err := c.request("POST", "/auth/proxy/update", req, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// KickProxy kicks a proxy offline
func (c *Client) KickProxy(proxyID int) error {
	req := IDRequest{ProxyID: proxyID}
	var resp Response[any]
	err := c.request("POST", "/auth/proxy/kick", req, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// ToggleProxy enables or disables a proxy
func (c *Client) ToggleProxy(proxyID int, isDisabled bool) error {
	req := ToggleProxyRequest{ProxyID: proxyID, IsDisabled: isDisabled}
	var resp Response[any]
	err := c.request("POST", "/auth/proxy/toggle", req, &resp)
	if err != nil {
		return err
	}

	if resp.Code != 200 {
		return fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return nil
}

// GetProxyConfig retrieves the configuration for a single proxy
func (c *Client) GetProxyConfig(proxyID int, format string) (*ProxyConfig, error) {
	req := ConfigRequest{ProxyID: proxyID, Format: format}
	var resp Response[ProxyConfig]
	err := c.request("POST", "/auth/proxy/config", req, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return &resp.Data, nil
}

// GetMultipleProxyConfigs retrieves configurations for multiple proxies
func (c *Client) GetMultipleProxyConfigs(proxyIDs []int, format string) (*ProxyConfig, error) {
	req := MultipleConfigRequest{ProxyIDs: proxyIDs, Format: format}
	var resp Response[ProxyConfig]
	err := c.request("POST", "/auth/proxy/config/multiple", req, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return &resp.Data, nil
}
