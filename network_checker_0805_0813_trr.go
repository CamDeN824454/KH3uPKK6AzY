// 代码生成时间: 2025-08-05 08:13:56
package main

import (
    "context"
    "fmt"
    "net"
    "time"

    "github.com/astaxie/beego"
)

// NetworkChecker 结构体封装了网络连接检查的相关参数
type NetworkChecker struct {
    // Host 是需要检查的主机地址
    Host string
    // Port 是需要检查的端口号
    Port int
    // Timeout 是连接超时时间
    Timeout time.Duration
}

// NewNetworkChecker 创建一个新的 NetworkChecker 实例
func NewNetworkChecker(host string, port int, timeout time.Duration) *NetworkChecker {
    return &NetworkChecker{
        Host:   host,
        Port:   port,
        Timeout: timeout,
    }
}

// Check 检查网络连接状态
func (n *NetworkChecker) Check() (bool, error) {
    // 使用 context 包来实现超时控制
    ctx, cancel := context.WithTimeout(context.Background(), n.Timeout)
    defer cancel()

    // 构建网络地址
    address := fmt.Sprintf("%s:%d", n.Host, n.Port)

    // 尝试建立网络连接
    conn, err := net.Dial("tcp", address)
    if err != nil {
        return false, err
    }
    defer conn.Close()

    // 如果连接成功，返回 true
    return true, nil
}

func main() {
    beego.Router("/check", &NetworkCheckerController{})
    beego.Run()
}

// NetworkCheckerController 控制器处理网络连接状态检查请求
type NetworkCheckerController struct {
    beego.Controller
}

// Get 处理 GET 请求，检查网络连接状态
func (c *NetworkCheckerController) Get() {
    // 获取请求参数
    host := c.GetString("host")
    port := c.GetInt("port")
    timeout := c.GetInt64("timeout")

    // 将超时时间从毫秒转换为秒
    timeoutDuration := time.Duration(timeout) * time.Millisecond

    // 创建 NetworkChecker 实例
    nc := NewNetworkChecker(host, port, timeoutDuration)

    // 检查网络连接状态
    connected, err := nc.Check()

    // 错误处理
    if err != nil {
        c.Ctx.WriteString(fmt.Sprintf("Error checking network connection: %s", err))
        c.Ctx.ResponseWriter.WriteHeader(500)
        return
    }

    // 根据连接状态返回结果
    if connected {
        c.Ctx.WriteString("Connected")
    } else {
        c.Ctx.WriteString("Not connected")
    }
    c.Ctx.ResponseWriter.WriteHeader(200)
}
