// 代码生成时间: 2025-10-08 02:42:26
package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
    "time"
)

// LatencyMonitor 用于监控网络延迟
type LatencyMonitor struct {
    Targets []string // 监控目标列表
}

// NewLatencyMonitor 创建一个新的 LatencyMonitor 实例
func NewLatencyMonitor(targets []string) *LatencyMonitor {
    return &LatencyMonitor{
        Targets: targets,
    }
}

// MonitorLatency 监控网络延迟
func (lm *LatencyMonitor) MonitorLatency() error {
    for _, target := range lm.Targets {
        err := lm.ping(target)
        if err != nil {
            return fmt.Errorf("monitoring latency for %s failed: %w", target, err)
        }
    }
    return nil
}

// ping 发送 ICMP Echo 请求并测量延迟
func (lm *LatencyMonitor) ping(target string) error {
    conn, err := net.Dial("ip:icmp", target)
    if err != nil {
        return fmt.Errorf("failed to create ICMP connection: %w", err)
    }
    defer conn.Close()

    start := time.Now()
    _, err = conn.Write([]byte("ping"))
    if err != nil {
        return fmt.Errorf("failed to send ICMP request: %w", err)
    }
    buf := make([]byte, 1024)
    _, err = conn.Read(buf)
    if err != nil {
        return fmt.Errorf("failed to read ICMP response: %w", err)
    }

    duration := time.Since(start)
    fmt.Printf("Ping to %s took %v
", target, duration)
    return nil
}

func main() {
    // 定义监控目标
    targets := []string{"8.8.8.8", "1.1.1.1"}
    monitor := NewLatencyMonitor(targets)

    // 监控网络延迟
    if err := monitor.MonitorLatency(); err != nil {
        fmt.Printf("Error monitoring latency: %v
", err)
    } else {
        fmt.Println("Latency monitoring completed successfully.")
    }
}
