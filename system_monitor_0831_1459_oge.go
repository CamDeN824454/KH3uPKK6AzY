// 代码生成时间: 2025-08-31 14:59:48
package main

import (
    "fmt"
    "log"
    "os"
    "os/exec"
    "time"
    "beego"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/load"
)

// 监控系统性能结构体
type SystemMonitor struct {
    CpuUsage float64 // CPU使用率
    MemUsage float64 // 内存使用率
    Load1   float64 // 1分钟平均负载
    Load5   float64 // 5分钟平均负载
    Load15  float64 // 15分钟平均负载
}

// 获取CPU使用率
func getCPUUsage() (float64, error) {
    cpuPercentage, err := cpu.Percent(0, false)
    if err != nil {
        log.Fatalf("Failed to get CPU usage: %v", err)
    }
    return cpuPercentage[0], nil
}

// 获取内存使用率
func getMemUsage() (float64, error) {
    virtualMemory, err := mem.VirtualMemory()
    if err != nil {
        log.Fatalf("Failed to get memory usage: %v", err)
    }
    return virtualMemory.UsedPercent, nil
}

// 获取系统负载
func getSystemLoad() (float64, float64, float64, error) {
    load, err := load.Avg()
    if err != nil {
        log.Fatalf("Failed to get system load: %v", err)
    }
    return load.Load1, load.Load5, load.Load15, nil
}

// 获取系统性能数据
func getSystemPerformance() (*SystemMonitor, error) {
    cpuUsage, err := getCPUUsage()
    if err != nil {
        return nil, err
    }
    memUsage, err := getMemUsage()
    if err != nil {
        return nil, err
    }
    load1, load5, load15, err := getSystemLoad()
    if err != nil {
        return nil, err
    }

    return &SystemMonitor{
        CpuUsage: cpuUsage,
        MemUsage: memUsage,
        Load1:   load1,
        Load5:   load5,
        Load15:  load15,
    }, nil
}

// 监控系统性能控制器
type SystemMonitorController struct {
    beego.Controller
}

// 获取系统性能
func (c *SystemMonitorController) Get() {
    monitor, err := getSystemPerformance()
    if err != nil {
        c.Data["json"] = map[string]string{
            "error": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{
        "cpu_usage":    monitor.CpuUsage,
        "mem_usage":    monitor.MemUsage,
        "load_1":       monitor.Load1,
        "load_5":       monitor.Load5,
        "load_15":      monitor.Load15,
    }
    c.ServeJSON()
}

func main() {
    // 初始化Beego框架
    beego.Router("/monitor", &SystemMonitorController{})
    if err := beego.Run(); err != nil {
        log.Fatalf("Failed to start Beego: %v", err)
    }
}