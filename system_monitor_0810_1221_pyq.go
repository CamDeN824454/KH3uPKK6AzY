// 代码生成时间: 2025-08-10 12:21:20
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "os/exec"
    "strings"
    "time"

    "github.com/astaxie/beego"
)

// SystemInfo contains the system performance data.
type SystemInfo struct {
    CpuUsage   string `json:"cpu_usage"`
    MemoryInfo string `json:"memory_info"`
    DiskInfo   string `json:"disk_info"`
}

// GetSystemInfo retrieves the system performance data.
func GetSystemInfo() (*SystemInfo, error) {
    // Retrieve CPU usage.
    cpuUsageCmd := "top -bn1 | grep load | awk '{printf "%.2f%%", $(NF-2)}'"
    cpuOut, err := exec.Command("/bin/sh", "-c", cpuUsageCmd).Output()
    if err != nil {
        return nil, err
    }
    cpuUsage := strings.TrimSpace(string(cpuOut))

    // Retrieve memory information.
    memInfoCmd := "free -m"
    memOut, err := exec.Command("/bin/sh", "-c", memInfoCmd).Output()
    if err != nil {
        return nil, err
    }
    memoryInfo := strings.TrimSpace(string(memOut))

    // Retrieve disk information.
    diskInfoCmd := "df -h"
    diskOut, err := exec.Command("/bin/sh", "-c", diskInfoCmd).Output()
    if err != nil {
        return nil, err
    }
    diskInfo := strings.TrimSpace(string(diskOut))

    return &SystemInfo{
        CpuUsage:   cpuUsage,
        MemoryInfo: memoryInfo,
        DiskInfo:   diskInfo,
    }, nil
}

// SystemMonitorController handles the HTTP requests.
type SystemMonitorController struct {
    beego.Controller
}

// Get method handles the GET request to get system performance data.
func (c *SystemMonitorController) Get() {
    info, err := GetSystemInfo()
    if err != nil {
        log.Printf("Error retrieving system info: %v", err)
        c.Data["json"] = map[string]interface{}{"error": "Unable to retrieve system information"}
        c.ServeJSON()
        return
    }
    c.Data["json"] = info
    c.ServeJSON()
}

func main() {
    beego.Router("/systemInfo", &SystemMonitorController{})
    beego.Run()
}
