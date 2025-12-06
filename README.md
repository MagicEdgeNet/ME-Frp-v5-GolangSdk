# ME Frp Golang SDK

ME Frp 5.0 官方 Go 语言 SDK，提供完整的 API 封装。

## 安装

```bash
go get github.com/MagicEdgeNet/MEFrp-v5-Golang-Sdk
```

## 快速开始

### 1. 注册与登录

```go
package main

import (
"fmt"
"log"

"github.com/MagicEdgeNet/MEFrp-v5-Golang-Sdk"
)

func main() {
	// 创建客户端（无需 token 用于注册/登录）
	client := mefrp.NewClient("")

	// 获取注册验证码（需人机验证 token）
	err := client.GetRegisterEmailCode("your@email.com", "captcha_token_here")
	if err != nil {
		log.Fatal(err)
	}

	// 注册账户
	err = client.Register(mefrp.RegisterRequest{
Username:  "myusername",
Email:     "your@email.com",
EmailCode: "123456",
Password:  "MyPass123",
})
	if err != nil {
		log.Fatal(err)
	}

	// 登录获取 token（需人机验证 token）
	token, err := client.Login(mefrp.LoginRequest{
Username:     "myusername",
Password:     "MyPass123",
CaptchaToken: "captcha_token_here",
})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Login successful, token: %s\n", token)
}
```

### 2. 用户信息与签到

```go
package main

import (
"fmt"
"log"

"github.com/MagicEdgeNet/MEFrp-v5-Golang-Sdk"
)

func main() {
	client := mefrp.NewClient("YOUR_API_TOKEN")

	// 获取用户信息
	user, err := client.GetUserInfo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User: %s, Traffic: %d MB\n", user.Username, user.Traffic/1024)

	// 每日签到（需人机验证 token）
	err = client.Sign("captcha_token_here")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sign-in successful!")

	// 获取用户组信息
	groups, err := client.GetUserGroups()
	if err != nil {
		log.Fatal(err)
	}
	for _, g := range groups {
		fmt.Printf("Group: %s, Max Proxies: %d\n", g.FriendlyName, g.MaxProxies)
	}
}
```

### 3. 隧道管理

```go
package main

import (
"fmt"
"log"

"github.com/MagicEdgeNet/MEFrp-v5-Golang-Sdk"
)

func main() {
	client := mefrp.NewClient("YOUR_API_TOKEN")

	// 获取隧道列表
	proxies, err := client.GetProxyList()
	if err != nil {
		log.Fatal(err)
	}
	for _, t := range proxies {
		status := "Offline"
		if t.IsOnline {
			status = "Online"
		}
		fmt.Printf("- [%s] %s (ID: %d, Type: %s)\n", status, t.ProxyName, t.ProxyID, t.ProxyType)
	}

	// 创建隧道
	err = client.CreateProxy(mefrp.CreateProxyRequest{
ProxyName:      "MyProxy",
ProxyType:      "tcp",
LocalIP:        "127.0.0.1",
LocalPort:      8080,
RemotePort:     10080,
NodeID:         1,
UseEncryption:  false,
UseCompression: false,
})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Proxy created successfully!")

	// 获取隧道配置（用于启动 frpc）
	config, err := client.GetProxyConfig(12345, "toml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Config:\n%s\n", config.Config)

	// 启用/禁用隧道
	err = client.ToggleProxy(12345, true) // true = 禁用
	if err != nil {
		log.Fatal(err)
	}

	// 删除隧道
	err = client.DeleteProxy(12345)
	if err != nil {
		log.Fatal(err)
	}
}
```

### 4. 节点信息

```go
package main

import (
"fmt"
"log"

"github.com/MagicEdgeNet/MEFrp-v5-Golang-Sdk"
)

func main() {
	client := mefrp.NewClient("YOUR_API_TOKEN")

	// 获取节点列表
	nodes, err := client.GetNodeList()
	if err != nil {
		log.Fatal(err)
	}
	for _, n := range nodes {
		fmt.Printf("Node: %s (ID: %d, Region: %s, Online: %v)\n", 
n.Name, n.NodeID, n.Region, n.IsOnline)
	}

	// 获取节点状态
	statuses, err := client.GetNodeStatus()
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range statuses {
		fmt.Printf("Node %s: %d clients, %d proxies, Load: %d%%\n",
s.Name, s.OnlineClient, s.OnlineProxy, s.LoadPercent)
	}
}
```

### 5. 公共信息

```go
package main

import (
"fmt"
"log"

"github.com/MagicEdgeNet/MEFrp-v5-Golang-Sdk"
)

func main() {
	client := mefrp.NewClient("")

	// 获取公共统计（无需 token）
	stats, err := client.GetStatistics()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Users: %d, Nodes: %d, Proxies: %d\n", 
stats.Users, stats.Nodes, stats.Proxies)

	// 获取商城商品（无需 token）
	items, err := client.GetStoreItems()
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		fmt.Printf("%s: %.2f CNY/%s\n", item.Name, item.CurrentPrice, item.Unit)
	}
}
```

## API 列表

### 公共接口（无需认证）
- `GetStatistics()` - 获取平台统计信息
- `GetStoreItems()` - 获取商城商品列表
- `GetRegisterEmailCode(email, captchaToken)` - 获取注册验证码
- `Register(req)` - 注册账户
- `Login(req)` - 密码登录
- `RecoverAccount(req)` - 找回账户

### 用户相关
- `GetUserInfo()` - 获取用户信息
- `Sign(captchaToken)` - 每日签到
- `GetUserFrpToken()` - 获取 FRP 启动 Token
- `GetUserGroups()` - 获取用户组信息
- `ResetAccessKey(captchaToken)` - 重置访问密钥
- `ChangePassword(req)` - 修改密码
- `GetUserLogs(page, pageSize, startTime, endTime)` - 获取操作日志
- `GetUserLogStats()` - 获取日志统计

### 隧道相关
- `GetProxyList()` - 获取隧道列表
- `CreateProxy(req)` - 创建隧道
- `UpdateProxy(req)` - 更新隧道
- `DeleteProxy(proxyID)` - 删除隧道
- `KickProxy(proxyID)` - 强制下线隧道
- `ToggleProxy(proxyID, isDisabled)` - 启用/禁用隧道
- `GetProxyConfig(proxyID, format)` - 获取单一隧道配置
- `GetMultipleProxyConfigs(proxyIDs, format)` - 获取多个隧道配置

### 节点相关
- `GetNodeList()` - 获取节点列表
- `GetNodeStatus()` - 获取节点状态
- `GetNodeToken(nodeID)` - 获取节点 Token
- `GetNodeConnectionList()` - 获取节点连接地址

### 系统相关
- `GetSystemStatus()` - 获取系统状态
- `GetPopupNotice()` - 获取重要公告

## 配置选项

```go
import "time"

client := mefrp.NewClient("token",
mefrp.WithTimeout(30 * time.Second),
mefrp.WithBaseURL("https://api.mefrp.com/api"),
mefrp.WithUserAgent("MyApp/1.0"),
)
```

## 注意事项

### 人机验证
以下接口需要人机验证 Token（`captchaToken`）：
- `GetRegisterEmailCode` - 获取注册验证码
- `Login` - 密码登录
- `Sign` - 每日签到
- `ResetAccessKey` - 重置访问密钥

您需要在前端或其他地方完成人机验证后获取 Token 并传入。

### User-Agent 要求
根据官方文档，建议使用自定义 User-Agent，格式为：
```
客户端名称/版本号 联系方式
```

例如：`MyFrpClient/1.0.0 admin@example.com`

不要使用 `MEFrp-Client` 开头的 User-Agent。

### 更新隧道接口说明
`UpdateProxy` 接口的路径在官方文档中显示为 `/auth/proxy/create`，但这可能是文档错误。SDK 暂时使用 `/auth/proxy/update` 作为路径。如果调用失败，请尝试使用 `CreateProxy` 方法。

### 密码修改警告
调用 `ChangePassword` 会**重置启动令牌和访问密钥**，请谨慎操作。

## 目录结构

```
.
├── README.md           # 本文档
├── auth.go             # 注册、登录、密码管理
├── client.go           # 核心 HTTP 客户端
├── examples/           # 使用示例
│   └── demo/
│       └── main.go
├── go.mod              # Go 模块定义
├── node.go             # 节点相关接口
├── public.go           # 公共信息接口
├── system.go           # 系统状态接口
├── proxy.go           # 隧道管理接口
├── types.go            # 数据结构定义
└── user.go             # 用户信息接口
```

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License
