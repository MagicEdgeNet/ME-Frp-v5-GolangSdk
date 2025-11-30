package mefrp

// Response is the common response structure for API requests
type Response[T any] struct {
	Code    int    `json:"code"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}

// UserInfo represents the user information
type UserInfo struct {
	UserID        int    `json:"userId"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	Group         string `json:"group"`
	IsRealname    bool   `json:"isRealname"`
	InBound       int    `json:"inBound"`
	MaxProxies    int    `json:"maxProxies"`
	OutBound      int    `json:"outBound"`
	RegTime       int64  `json:"regTime"`
	Status        int    `json:"status"` // 0-正常 1-封禁 2-流量超限
	TodaySigned   bool   `json:"todaySigned"`
	Traffic       int64  `json:"traffic"`
	VipExpireTime int64  `json:"vipExpireTime,omitempty"`
	UsedProxies   int    `json:"usedProxies"`
	FriendlyGroup string `json:"friendlyGroup"`
}

// Tunnel represents a proxy tunnel configuration
type Tunnel struct {
	ProxyID              int    `json:"proxyId"`
	Username             string `json:"username"`
	ProxyName            string `json:"proxyName"`
	ProxyType            string `json:"proxyType"`
	IsBanned             bool   `json:"isBanned"`
	IsDisabled           bool   `json:"isDisabled"`
	LocalIP              string `json:"localIp"`
	LocalPort            int    `json:"localPort"`
	RemotePort           int    `json:"remotePort"`
	NodeID               int    `json:"nodeId"`
	RunID                string `json:"runId"`
	IsOnline             bool   `json:"isOnline"`
	Domain               string `json:"domain"`
	LastStartTime        int64  `json:"lastStartTime"`
	LastCloseTime        int64  `json:"lastCloseTime"`
	ClientVersion        string `json:"clientVersion"`
	ProxyProtocolVersion string `json:"proxyProtocolVersion"`
	UseEncryption        bool   `json:"useEncryption"`
	UseCompression       bool   `json:"useCompression"`
	Location             string `json:"location"`
	AccessKey            string `json:"accessKey"`
	HostHeaderRewrite    string `json:"hostHeaderRewrite"`
	HeaderXFromWhere     string `json:"headerXFromWhere"`
}

// Node represents a server node
type Node struct {
	NodeID          int    `json:"nodeId"`
	Name            string `json:"name"`
	Hostname        string `json:"hostname"`
	Description     string `json:"description"`
	Token           string `json:"token"`
	ServicePort     int    `json:"servicePort"`
	AdminPort       int    `json:"adminPort"`
	AdminPass       string `json:"adminPass"`
	AllowGroup      string `json:"allowGroup"`
	AllowPort       string `json:"allowPort"`
	AllowType       string `json:"allowType"`
	Region          string `json:"region"`
	Bandwidth       string `json:"bandwidth"`
	IsOnline        bool   `json:"isOnline"`
	IsDisabled      bool   `json:"isDisabled"`
	TotalTrafficIn  int64  `json:"totalTrafficIn"`
	TotalTrafficOut int64  `json:"totalTrafficOut"`
	UpTime          int64  `json:"upTime"`
	Version         string `json:"version"`
}

// NodeStatus represents the status of a node
type NodeStatus struct {
	NodeID          int    `json:"nodeId"`
	Name            string `json:"name"`
	TotalTrafficIn  int64  `json:"totalTrafficIn"`
	TotalTrafficOut int64  `json:"totalTrafficOut"`
	OnlineClient    int    `json:"onlineClient"`
	OnlineProxy     int    `json:"onlineProxy"`
	IsOnline        bool   `json:"isOnline"`
	Version         string `json:"version"`
	Uptime          int64  `json:"uptime"`
	CurConns        int    `json:"curConns"`
	LoadPercent     int    `json:"loadPercent"`
}

// Statistics represents public statistics
type Statistics struct {
	Users   int   `json:"users"`
	Nodes   int   `json:"nodes"`
	Proxies int   `json:"proxies"`
	Traffic int64 `json:"traffic"`
}

// UserGroup represents a user group
type UserGroup struct {
	Name         string `json:"name"`
	FriendlyName string `json:"friendlyName"`
	MaxProxies   int    `json:"maxProxies"`
	BaseTraffic  int64  `json:"baseTraffic"`
	OutBound     int    `json:"outBound"`
	InBound      int    `json:"inBound"`
}

// FrpToken represents the user's frp token
type FrpToken struct {
Token string `json:"token"`
}

// TunnelConfig represents a single tunnel configuration
type TunnelConfig struct {
Config string `json:"config"`
Type   string `json:"type"`
}

// CreateTunnelRequest represents the request to create a tunnel
type CreateTunnelRequest struct {
ProxyName            string `json:"proxyName"`
ProxyType            string `json:"proxyType"`
LocalIP              string `json:"localIp"`
LocalPort            int    `json:"localPort"`
RemotePort           int    `json:"remotePort"`
NodeID               int    `json:"nodeId"`
Domain               string `json:"domain,omitempty"`
ProxyProtocolVersion string `json:"proxyProtocolVersion,omitempty"`
UseEncryption        bool   `json:"useEncryption"`
UseCompression       bool   `json:"useCompression"`
HostHeaderRewrite    string `json:"hostHeaderRewrite,omitempty"`
HeaderXFromWhere     string `json:"headerXFromWhere,omitempty"`
}

// ToggleTunnelRequest represents the request to toggle a tunnel
type ToggleTunnelRequest struct {
ProxyID    int  `json:"proxyId"`
IsDisabled bool `json:"isDisabled"`
}

// IDRequest represents a request with just a proxy ID (Delete, Kick, Config)
type IDRequest struct {
ProxyID int `json:"proxyId"`
}

// ConfigRequest represents a request for tunnel config
type ConfigRequest struct {
ProxyID int    `json:"proxyId"`
Format  string `json:"format"` // "toml", "json", "yml", "ini"
}

// CaptchaRequest represents a request requiring a captcha token
type CaptchaRequest struct {
CaptchaToken string `json:"captchaToken"`
}

// ResetTokenResponse represents the response for resetting access key
type ResetTokenResponse struct {
NewToken string `json:"newToken"`
}

// StoreItem represents an item in the store
type StoreItem struct {
Type                     string  `json:"type"`
Name                     string  `json:"name"`
Price                    float64 `json:"price"`
Unit                     string  `json:"unit"`
Description              string  `json:"description"`
Enabled                  bool    `json:"enabled"`
DiscountEnabled          bool    `json:"discountEnabled"`
DiscountPrice            float64 `json:"discountPrice"`
DiscountStartTime        int64   `json:"discountStartTime"`
DiscountEndTime          int64   `json:"discountEndTime"`
CurrentPrice             float64 `json:"currentPrice"`
IsDiscountActive         bool    `json:"isDiscountActive"`
DiscountRemainingSeconds int64   `json:"discountRemainingSeconds"`
}

// OperationLog represents a user operation log entry
type OperationLog struct {
LogID     int    `json:"logId"`
Category  string `json:"category"`
Details   string `json:"details"`
IPAddress string `json:"ipAddress"`
Status    string `json:"status"`
CreatedAt string `json:"createdAt"`
}

// OperationLogList represents the response for operation logs
type OperationLogList struct {
Data       []OperationLog `json:"data"`
Total      int            `json:"total"`
Page       int            `json:"page"`
PageSize   int            `json:"pageSize"`
TotalPages int            `json:"totalPages"`
}

// SystemStatus represents the system status
type SystemStatus struct {
Status int    `json:"status"` // 0 正常 1 降级 2 离线
Remark string `json:"remark"`
}

// NodeToken represents the node token response
type NodeToken struct {
ServerPort int    `json:"serverPort"`
Token      string `json:"token"`
}

// MultipleConfigResponse represents the response for multiple tunnel configs
type MultipleConfigResponse struct {
Config string `json:"config"`
Type   string `json:"type"`
}

// MultipleConfigRequest represents the request for multiple tunnel configs
type MultipleConfigRequest struct {
ProxyIDs []int  `json:"proxyIds"`
Format   string `json:"format"`
}

// UserGroupsResponse represents the response for user groups
type UserGroupsResponse struct {
Groups []UserGroup `json:"groups"`
}

// UserLogStats represents the user log statistics
type UserLogStats struct {
	MonthCount int `json:"monthCount"`
	TodayCount int `json:"todayCount"`
	TotalCount int `json:"totalCount"`
	WeekCount  int `json:"weekCount"`
}

// NodeConnection represents a node connection address
type NodeConnection struct {
	NodeID   int    `json:"nodeId"`
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
}
