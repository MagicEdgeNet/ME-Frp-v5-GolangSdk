package mefrp

import "fmt"

// GetStatistics retrieves public statistics
func (c *Client) GetStatistics() (*Statistics, error) {
	var resp Response[Statistics]
	err := c.request("GET", "/public/statistics", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return &resp.Data, nil
}

// GetStoreItems retrieves items from the store
func (c *Client) GetStoreItems() ([]StoreItem, error) {
	var resp Response[[]StoreItem]
	err := c.request("GET", "/public/store/products", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return resp.Data, nil
}
