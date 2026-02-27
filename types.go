package mefrpApi

// Response is the common response structure for API requests
type Response[T any] struct {
	Code    int    `json:"code"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}

// UserInfo represents the user information
type UserInfo struct {
	UserID        int64  `json:"userId"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	Group         string `json:"group"`
	IsRealname    bool   `json:"isRealname"`
	InBound       int64  `json:"inBound"`
	MaxProxies    int64  `json:"maxProxies"`
	OutBound      int64  `json:"outBound"`
	RegTime       int64  `json:"regTime"`
	Status        int    `json:"status"` // 0-正常 1-封禁 2-流量超限
	TodaySigned   bool   `json:"todaySigned"`
	Traffic       int64  `json:"traffic"`
	VipExpireTime int64  `json:"vipExpireTime,omitempty"`
	UsedProxies   int    `json:"usedProxies"`
	FriendlyGroup string `json:"friendlyGroup"`
	RealnameTimes int    `json:"realnameTimes"`
}

// Proxy represents a proxy proxy configuration
type Proxy struct {
	ProxyID              int64  `json:"proxyId"`
	Username             string `json:"username"`
	ProxyName            string `json:"proxyName"`
	ProxyType            string `json:"proxyType"`
	IsBanned             bool   `json:"isBanned"`
	IsDisabled           bool   `json:"isDisabled"`
	LocalIP              string `json:"localIp"`
	LocalPort            int32  `json:"localPort"`
	RemotePort           int32  `json:"remotePort"`
	NodeID               int64  `json:"nodeId"`
	RunID                string `json:"runId"`
	IsOnline             bool   `json:"isOnline"`
	Domain               string `json:"domain"`
	LastStartTime        int64  `json:"lastStartTime"`
	LastCloseTime        int64  `json:"lastCloseTime"`
	ClientVersion        string `json:"clientVersion"`
	ProxyProtocolVersion string `json:"proxyProtocolVersion"`
	UseEncryption        bool   `json:"useEncryption"`
	UseCompression       bool   `json:"useCompression"`
	Locations            string `json:"locations"`
	AccessKey            string `json:"accessKey"`
	HostHeaderRewrite    string `json:"hostHeaderRewrite"`
	HttpPlugin           string `json:"httpPlugin"`
	CrtPath              string `json:"crtPath"`
	KeyPath              string `json:"keyPath"`
	RequestHeaders       string `json:"requestHeaders"`
	ResponseHeaders      string `json:"responseHeaders"`
	HTTPUser             string `json:"httpUser"`
	HTTPPassword         string `json:"httpPassword"`
	TransportProtocol    string `json:"transportProtocol"`
}

// ProxyListResponse represents the response for proxy list
type ProxyListResponse struct {
	Proxies []Proxy          `json:"proxies"`
	Nodes   []NodeConnection `json:"nodes"`
}

// Node represents a server node
type Node struct {
	NodeID          int64  `json:"nodeId"`
	Name            string `json:"name"`
	Hostname        string `json:"hostname"`
	Description     string `json:"description"`
	Token           string `json:"token"`
	ServicePort     int32  `json:"servicePort"`
	AdminPort       int32  `json:"adminPort"`
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
	DonateID        int64  `json:"donateId"`
	DonateUser      string `json:"donateUser"`
}

// NodeStatus represents the status of a node
type NodeStatus struct {
	NodeID          int64  `json:"nodeId"`
	Name            string `json:"name"`
	TotalTrafficIn  int64  `json:"totalTrafficIn"`
	TotalTrafficOut int64  `json:"totalTrafficOut"`
	OnlineClient    int64  `json:"onlineClient"`
	OnlineProxy     int64  `json:"onlineProxy"`
	IsOnline        bool   `json:"isOnline"`
	Version         string `json:"version"`
	Uptime          int64  `json:"uptime"`
	CurConns        int64  `json:"curConns"`
	LoadPercent     int64  `json:"loadPercent"`
}

// Statistics represents public statistics
type Statistics struct {
	Users   int64 `json:"users"`
	Nodes   int64 `json:"nodes"`
	Proxies int64 `json:"proxies"`
	Traffic int64 `json:"traffic"`
}

// UserGroup represents a user group
type UserGroup struct {
	Name         string `json:"name"`
	FriendlyName string `json:"friendlyName"`
	MaxProxies   int64  `json:"maxProxies"`
	BaseTraffic  int64  `json:"baseTraffic"`
	OutBound     int64  `json:"outBound"`
	InBound      int64  `json:"inBound"`
}

// FrpToken represents the user's frp token
type FrpToken struct {
	Token string `json:"token"`
}

// ProxyConfig represents a single proxy configuration
type ProxyConfig struct {
	Config string `json:"config"`
	Type   string `json:"type"`
}

// CreateProxyRequest represents the request to create a proxy
type CreateProxyRequest struct {
	ProxyName            string            `json:"proxyName"`
	ProxyType            string            `json:"proxyType"`
	LocalIP              string            `json:"localIp"`
	LocalPort            int32             `json:"localPort"`
	RemotePort           int32             `json:"remotePort"`
	NodeID               int64             `json:"nodeId"`
	Domain               string            `json:"domain"`
	Locations            string            `json:"locations"`
	AccessKey            string            `json:"accessKey"`
	HostHeaderRewrite    string            `json:"hostHeaderRewrite"`
	UseEncryption        bool              `json:"useEncryption"`
	UseCompression       bool              `json:"useCompression"`
	ProxyProtocolVersion string            `json:"proxyProtocolVersion"`
	HttpPlugin           string            `json:"httpPlugin"`
	CrtPath              string            `json:"crtPath,omitempty"`
	KeyPath              string            `json:"keyPath,omitempty"`
	RequestHeaders       map[string]string `json:"requestHeaders"`
	ResponseHeaders      map[string]string `json:"responseHeaders"`
	HTTPUser             string            `json:"httpUser"`
	HTTPPassword         string            `json:"httpPassword"`
	TransportProtocol    string            `json:"transportProtocol"`
}

// UpdateProxyRequest represents the request to update a proxy
type UpdateProxyRequest struct {
	ProxyID              int64             `json:"proxyId"`
	ProxyName            string            `json:"proxyName"`
	LocalIP              string            `json:"localIp"`
	LocalPort            int32             `json:"localPort"`
	RemotePort           int32             `json:"remotePort"`
	NodeID               int64             `json:"nodeId"`
	Domain               string            `json:"domain"`
	ProxyProtocolVersion string            `json:"proxyProtocolVersion"`
	UseEncryption        bool              `json:"useEncryption"`
	UseCompression       bool              `json:"useCompression"`
	HostHeaderRewrite    string            `json:"hostHeaderRewrite"`
	Locations            string            `json:"locations,omitempty"`
	AccessKey            string            `json:"accessKey,omitempty"`
	HttpPlugin           string            `json:"httpPlugin,omitempty"`
	CrtPath              string            `json:"crtPath,omitempty"`
	KeyPath              string            `json:"keyPath,omitempty"`
	RequestHeaders       map[string]string `json:"requestHeaders,omitempty"`
	ResponseHeaders      map[string]string `json:"responseHeaders,omitempty"`
	HTTPUser             string            `json:"httpUser,omitempty"`
	HTTPPassword         string            `json:"httpPassword,omitempty"`
	TransportProtocol    string            `json:"transportProtocol,omitempty"`
}

// ToggleProxyRequest represents the request to toggle a proxy
type ToggleProxyRequest struct {
	ProxyID    int64 `json:"proxyId"`
	IsDisabled bool  `json:"isDisabled"`
}

// IDRequest represents a request with just a proxy ID (Delete, Kick, Config)
type IDRequest struct {
	ProxyID int64 `json:"proxyId"`
}

// ConfigRequest represents a request for proxy config
type ConfigRequest struct {
	ProxyID int64  `json:"proxyId"`
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
	Total      int64          `json:"total"`
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
	ServerPort int32  `json:"serverPort"`
	Token      string `json:"token"`
}

// MultipleConfigResponse represents the response for multiple proxy configs
type MultipleConfigResponse struct {
	Config string `json:"config"`
	Type   string `json:"type"`
}

// MultipleConfigRequest represents the request for multiple proxy configs
type MultipleConfigRequest struct {
	ProxyIDs []int64 `json:"proxyIds"`
	Format   string  `json:"format"`
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
	NodeID   int64  `json:"nodeId"`
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
}

// Ads represents an advertisement
type Ads struct {
	AdsID          int64   `json:"adsId"`
	AdsOwner       string  `json:"adsOwner"`
	AdsURL         string  `json:"adsUrl"`
	AdsType        string  `json:"adsType"`
	AdsContent     string  `json:"adsContent"`
	AdsImageUrl    string  `json:"adsImageUrl"`
	AdsStartTime   int64   `json:"adsStartTime"`
	AdsExpire      int64   `json:"adsExpire"`
	RenewalPrice   float64 `json:"renewalPrice"`
	AdsPlacement   string  `json:"adsPlacement"`
	AdsClick       int64   `json:"adsClick"`
	AdsImpression  int64   `json:"adsImpression"`
	AdsStatus      int     `json:"adsStatus"` // 0=pending,1=approved,2=rejected,3=returned
	AdsReviewNote  string  `json:"adsReviewNote"`
	AdsReviewer    string  `json:"adsReviewer"`
	AdsReviewTime  int64   `json:"adsReviewTime"`
	AdsSlotID      int64   `json:"adsSlotId"`
	AdsCreatedTime int64   `json:"adsCreatedTime"`
}

// AdsStats represents advertisement statistics
type AdsStats struct {
	Total         int64 `json:"total"`
	PendingCount  int64 `json:"pendingCount"`
	ApprovedCount int64 `json:"approvedCount"`
	RejectedCount int64 `json:"rejectedCount"`
	ReturnedCount int64 `json:"returnedCount"`
	TotalClicks   int64 `json:"totalClicks"`
}

// AdSlot represents an advertisement slot
type AdSlot struct {
	SlotID       int64   `json:"slotId"`
	Name         string  `json:"name"`
	Placement    string  `json:"placement"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	ValidityDays int     `json:"validityDays"`
	MaxQuota     int     `json:"maxQuota"`
	Enabled      bool    `json:"enabled"`
	CreatedTime  int64   `json:"createdTime"`
}

// AdSlotWithUsage represents an advertisement slot with usage info
type AdSlotWithUsage struct {
	AdSlot        AdSlot `json:"adSlot"`
	ApprovedCount int64  `json:"approvedCount"`
	Remaining     int64  `json:"remaining"`
}

// AdCredit represents advertisement credits
type AdCredit struct {
	CreditID   int64  `json:"creditId"`
	UserID     int64  `json:"userId"`
	Username   string `json:"username"`
	SlotID     int64  `json:"slotId"`
	SlotName   string `json:"slotName"`
	Total      int    `json:"total"`
	Used       int    `json:"used"`
	UpdateTime int64  `json:"updateTime"`
	ExpireTime int64  `json:"expireTime"`
}

// Coupon represents a discount coupon
type Coupon struct {
	CouponID       int64   `json:"couponId"`
	Code           string  `json:"code"`
	DiscountType   string  `json:"discountType"`
	DiscountValue  float64 `json:"discountValue"`
	MinOrderAmount float64 `json:"minOrderAmount"`
	ProductTypes   string  `json:"productTypes"`
	MaxUsage       int     `json:"maxUsage"`
	PerUserLimit   int     `json:"perUserLimit"`
	UsedCount      int     `json:"usedCount"`
	StartTime      int64   `json:"startTime"`
	ExpireTime     int64   `json:"expireTime"`
	Enabled        bool    `json:"enabled"`
	TargetType     string  `json:"targetType"`
	TargetValue    string  `json:"targetValue"`
	CreatedTime    int64   `json:"createdTime"`
}

// NodeDonate represents a node donation application
type NodeDonate struct {
	DonateID     int64  `json:"donateId"`
	Username     string `json:"username"`
	NodeName     string `json:"nodeName"`
	Hostname     string `json:"hostname"`
	Description  string `json:"description"`
	ServicePort  int32  `json:"servicePort"`
	AdminPort    int32  `json:"adminPort"`
	AdminPass    string `json:"adminPass"`
	AllowGroup   string `json:"allowGroup"`
	AllowPort    string `json:"allowPort"`
	AllowType    string `json:"allowType"`
	Region       string `json:"region"`
	Bandwidth    string `json:"bandwidth"`
	Status       int    `json:"status"` // 0: 待审核, 1: 已通过, 2: 已拒绝, 3: 打回重改
	RejectReason string `json:"rejectReason"`
	ApplyTime    int64  `json:"applyTime"`
	ReviewTime   int64  `json:"reviewTime"`
	NodeID       int64  `json:"nodeId"`
}

// NodeDeleteRequest represents a request to delete a node
type NodeDeleteRequest struct {
	RequestID    int64  `json:"requestId"`
	NodeID       int64  `json:"nodeId"`
	Username     string `json:"username"`
	Reason       string `json:"reason"`
	Status       int    `json:"status"`
	RejectReason string `json:"rejectReason"`
	ApplyTime    int64  `json:"applyTime"`
	ReviewTime   int64  `json:"reviewTime"`
}

// NodeEditRequest represents a request to edit a node
type NodeEditRequest struct {
	RequestID    int64  `json:"requestId"`
	NodeID       int64  `json:"nodeId"`
	Username     string `json:"username"`
	NodeName     string `json:"nodeName"`
	Hostname     string `json:"hostname"`
	Description  string `json:"description"`
	ServicePort  int32  `json:"servicePort"`
	AdminPort    int32  `json:"adminPort"`
	AdminPass    string `json:"adminPass"`
	AllowGroup   string `json:"allowGroup"`
	AllowPort    string `json:"allowPort"`
	AllowType    string `json:"allowType"`
	Region       string `json:"region"`
	Bandwidth    string `json:"bandwidth"`
	Reason       string `json:"reason"`
	Status       int    `json:"status"`
	RejectReason string `json:"rejectReason"`
	ApplyTime    int64  `json:"applyTime"`
	ReviewTime   int64  `json:"reviewTime"`
}

// IcpDomain represents an ICP registered domain
type IcpDomain struct {
	Domain     string `json:"domain"`
	NatureName string `json:"natureName"`
	Username   string `json:"username"`
	IcpID      string `json:"icpId"`
	UnitName   string `json:"unitName"`
}

// Product represents a software product
type Product struct {
	ProductID string `json:"productId"`
	System    string `json:"system"`
	Arch      string `json:"arch"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	Path      string `json:"path"`
	Version   string `json:"version"`
	IsPublic  bool   `json:"isPublic"`
}

// DownloadSource represents a download source
type DownloadSource struct {
	ID   int64  `json:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
}

// UserDailyTrafficResponse represents the daily traffic stats
type UserDailyTrafficResponse struct {
	Dates        []string `json:"dates"`
	TrafficIn    []int64  `json:"trafficIn"`
	TrafficOut   []int64  `json:"trafficOut"`
	TotalTraffic []int64  `json:"totalTraffic"`
}

// Order represents a financial order
type Order struct {
	OrderID    string  `json:"orderId"`
	UserID     int64   `json:"userId"`
	Type       string  `json:"type"`
	Amount     int     `json:"amount"`
	Months     int     `json:"months"`
	Money      float64 `json:"money"`
	Status     int     `json:"status"`
	PayType    string  `json:"payType"`
	PayURL     string  `json:"payURL"`
	PayInfo    string  `json:"payInfo"`
	PayHTML    string  `json:"payHTML"`
	PayQRCode  string  `json:"payQRCode"`
	TradeNo    string  `json:"tradeNo"`
	CouponCode string  `json:"couponCode"`
	AdSlotType string  `json:"adSlotType"`
	CreateTime int64   `json:"createTime"`
	UpdateTime int64   `json:"updateTime"`
}

// SubmitOrderRequest represents the request to submit an order
type SubmitOrderRequest struct {
	Type         string `json:"type"`
	Amount       int    `json:"amount"`
	Months       int    `json:"months"`
	PayMethod    string `json:"payMethod"`
	CaptchaToken string `json:"captchaToken"`
	CouponCode   string `json:"couponCode"`
	AdSlotType   string `json:"adSlotType"`
}

// SubmitOrderResponse represents the response after submitting an order
type SubmitOrderResponse struct {
	OrderID   string `json:"orderId"`
	TradeNo   string `json:"tradeNo,omitempty"`
	PayType   string `json:"payType"`
	PayURL    string `json:"payURL,omitempty"`
	PayInfo   string `json:"payInfo,omitempty"`
	PayHTML   string `json:"payHTML,omitempty"`
	PayQRCode string `json:"payQRCode,omitempty"`
}

// QueryOrderResponse represents the order query response
type QueryOrderResponse struct {
	OrderID string  `json:"orderId"`
	Status  int     `json:"status"`
	Money   float64 `json:"money"`
	PayType string  `json:"payType"`
	PayItem string  `json:"payItem"`
	PayTime string  `json:"payTime"`
}

// RealnameInfoResponse represents the user's realname status
type RealnameInfoResponse struct {
	RealnameTime   int64 `json:"realnameTime"`
	IsRealnamed    bool  `json:"isRealnamed"`
	AvailableTimes int   `json:"availableTimes"`
}

// CheckUpdateResponse represents the update check result
type CheckUpdateResponse struct {
	ProductID      string `json:"productId"`
	CurrentVersion string `json:"currentVersion"`
	LatestVersion  string `json:"latestVersion"`
	HasUpdate      bool   `json:"hasUpdate"`
	Product        struct {
		Name    string `json:"name"`
		Desc    string `json:"desc"`
		Path    string `json:"path"`
		System  string `json:"system"`
		Arch    string `json:"arch"`
		Version string `json:"version"`
	} `json:"product"`
}

// NodeInstallScript represents a node installation script
type NodeInstallScript struct {
	ScriptID   string `json:"scriptId"`
	NodeID     int64  `json:"nodeId"`
	Username   string `json:"username"`
	System     string `json:"system"`
	Arch       string `json:"arch"`
	Script     string `json:"script"`
	CreateTime int64  `json:"createTime"`
	ExpireTime int64  `json:"expireTime"`
}

// GetInstallScriptResponse represents the response for getting an installation script
type GetInstallScriptResponse struct {
	ScriptID    string `json:"scriptId"`
	DownloadURL string `json:"downloadUrl"`
	ExpireTime  int64  `json:"expireTime"`
}

// EasyStartProxy represents the configuration for easy startup
type EasyStartProxy struct {
	ProxyID              int64             `json:"proxyId"`
	Username             string            `json:"username"`
	ProxyName            string            `json:"proxyName"`
	ProxyType            string            `json:"proxyType"`
	IsBanned             bool              `json:"isBanned"`
	IsDisabled           bool              `json:"isDisabled"`
	LocalIP              string            `json:"localIp"`
	LocalPort            int32             `json:"localPort"`
	RemotePort           int32             `json:"remotePort"`
	RunID                string            `json:"runId"`
	IsOnline             bool              `json:"isOnline"`
	Domain               string            `json:"domain"`
	LastStartTime        int64             `json:"lastStartTime"`
	LastCloseTime        int64             `json:"lastCloseTime"`
	ClientVersion        string            `json:"clientVersion"`
	ProxyProtocolVersion string            `json:"proxyProtocolVersion"`
	UseEncryption        bool              `json:"useEncryption"`
	UseCompression       bool              `json:"useCompression"`
	Locations            string            `json:"locations"`
	AccessKey            string            `json:"accessKey"`
	HostHeaderRewrite    string            `json:"hostHeaderRewrite"`
	HttpPlugin           string            `json:"httpPlugin"`
	CrtPath              string            `json:"crtPath"`
	KeyPath              string            `json:"keyPath"`
	RequestHeaders       map[string]string `json:"requestHeaders"`
	HTTPUser             string            `json:"httpUser"`
	HTTPPassword         string            `json:"httpPassword"`
	NodeAddr             string            `json:"nodeAddr"`
	NodePort             int32             `json:"nodePort"`
	NodeToken            string            `json:"nodeToken"`
}

// NodeWithLoad represents a node with its current load percentage
type NodeWithLoad struct {
	Node
	LoadPercent int64 `json:"loadPercent"`
}

// CreateProxyData represents the data needed to create a proxy
type CreateProxyData struct {
	Nodes        []NodeWithLoad `json:"nodes"`
	Groups       []UserGroup    `json:"groups"`
	CurrentGroup string         `json:"currentGroup"`
}

// PurchaseStatus represents the user's purchase status and limits
type PurchaseStatus struct {
	RealnameTimes int   `json:"realnameTimes"`
	UnpaidOrders  int64 `json:"unpaidOrders"`
	UnpaidLimit   int   `json:"unpaidLimit"`
}

// OperationLogCategory represents a category for operation logs
type OperationLogCategory struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

// UserOperationLogFilter represents the filter for user operation logs
type UserOperationLogFilter struct {
	Page      int    `json:"page"`
	PageSize  int    `json:"pageSize"`
	Category  string `json:"category"`
	Status    string `json:"status"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

// ValidateCouponResponse represents the response for coupon validation
type ValidateCouponResponse struct {
	Coupon          Coupon  `json:"coupon"`
	DiscountedPrice float64 `json:"discountedPrice"`
}

// CDKUsageLog represents a log entry for CDK usage
type CDKUsageLog struct {
	LogID     int64  `json:"logId"`
	Code      string `json:"code"`
	Username  string `json:"username"`
	Type      string `json:"type"`
	Value     int64  `json:"value"`
	UseTime   int64  `json:"useTime"`
	ClientIP  string `json:"clientIp"`
	UserAgent string `json:"userAgent"`
}

// CDKUsageLogList represents a list of CDK usage logs with pagination
type CDKUsageLogList struct {
	Logs       []CDKUsageLog `json:"logs"`
	Total      int64         `json:"total"`
	Page       int           `json:"page"`
	PageSize   int           `json:"pageSize"`
	TotalPages int           `json:"totalPages"`
}

// RedeemCDKResponse represents the response after redeeming a CDK
type RedeemCDKResponse struct {
	Type  string `json:"type"`
	Value int64  `json:"value"`
}
