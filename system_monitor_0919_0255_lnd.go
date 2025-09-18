// 代码生成时间: 2025-09-19 02:55:38
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "os/exec"
    "time"

    "github.com/astaxie/beego"
)

// SystemInfo 存储系统性能信息
type SystemInfo struct {
    CpuUsage float64 `json:"cpu_usage"`
    MemoryUsage float64 `json:"memory_usage"`
    DiskUsage float64 `json:"disk_usage"`
}

// GetSystemInfo 获取系统性能信息
func GetSystemInfo() (*SystemInfo, error) {
    var info SystemInfo
    var err error

    // 获取CPU使用率
    info.CpuUsage, err = getCpuUsage()
    if err != nil {
        return nil, err
    }

    // 获取内存使用率
    info.MemoryUsage, err = getMemoryUsage()
    if err != nil {
        return nil, err
    }

    // 获取磁盘使用率
    info.DiskUsage, err = getDiskUsage()
    if err != nil {
        return nil, err
    }

    return &info, nil
}

// getCpuUsage 获取CPU使用率
func getCpuUsage() (float64, error) {
    cmd := exec.Command("sh", "-c", "top -bn1 | grep load | awk '{printf "%.2f", $(NF-2)}'")
    var out bytes.Buffer
    cmd.Stdout = &out
    if err := cmd.Run(); err != nil {
        return 0, err
    }
    return strconv.ParseFloat(strings.TrimSpace(out.String()), 64)
}

// getMemoryUsage 获取内存使用率
func getMemoryUsage() (float64, error) {
    cmd := exec.Command("sh", "-c", "free -m | awk 'NR==2{printf "%.2f", $3/$2 * 100.0}'")
    var out bytes.Buffer
    cmd.Stdout = &out
    if err := cmd.Run(); err != nil {
        return 0, err
    }
    return strconv.ParseFloat(strings.TrimSpace(out.String()), 64)
}

// getDiskUsage 获取磁盘使用率
func getDiskUsage() (float64, error) {
    cmd := exec.Command("sh", "-c", "df -h | awk 'NR==2{printf "%.2f", $5}'")
    var out bytes.Buffer
    cmd.Stdout = &out
    if err := cmd.Run(); err != nil {
        return 0, err
    }
    return strconv.ParseFloat(strings.TrimSpace(out.String()), 64)
}

func main() {
    beego.Router("/monitor", &SystemMonitorController{})
    beego.Run()
}

// SystemMonitorController 控制器
type SystemMonitorController struct {
    beego.Controller
}

// Get 获取系统性能监控数据
func (c *SystemMonitorController) Get() {
    info, err := GetSystemInfo()
    if err != nil {
        c.Data["json"] = map[string]interface{}{"error": err.Error()}
        c.ServeJSON()
    } else {
        c.Data["json"] = info
        c.ServeJSON()
    }
}
