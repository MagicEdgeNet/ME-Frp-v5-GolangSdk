package mefrp

import "fmt"

// GetNodeList retrieves the list of nodes
func (c *Client) GetNodeList() ([]Node, error) {
	var resp Response[[]Node]
	err := c.request("GET", "/auth/node/list", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return resp.Data, nil
}

// GetNodeStatus retrieves the status of nodes
func (c *Client) GetNodeStatus() ([]NodeStatus, error) {
	var resp Response[[]NodeStatus]
	err := c.request("GET", "/auth/node/status", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return resp.Data, nil
}

// GetNodeToken retrieves the token for a specific node
func (c *Client) GetNodeToken(nodeID int) (*NodeToken, error) {
	// Note: The API doc says request body with nodeId, but it's a GET request?
// Doc: /auth/node/secret GET
// Request Body: { nodeId: number }
// GET requests with body are non-standard but possible.
// However, usually it's a query param for GET.
	// Let's try query param first, or maybe it's POST?
	// The doc explicitly says "get".
	// Let's try sending body with GET request. Go's http.NewRequest allows body in GET.
	
	req := struct {
		NodeID int `json:"nodeId"`
	}{NodeID: nodeID}
	
	var resp Response[NodeToken]
	// Using GET with body as per doc, though unusual.
	err := c.request("GET", "/auth/node/secret", req, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return &resp.Data, nil
}

// GetNodeConnectionList retrieves the connection addresses for nodes (only for created tunnels)
func (c *Client) GetNodeConnectionList() ([]NodeConnection, error) {
	var resp Response[[]NodeConnection]
	err := c.request("GET", "/auth/node/nameList", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("api error: %s (code: %d)", resp.Message, resp.Code)
	}

	return resp.Data, nil
}
