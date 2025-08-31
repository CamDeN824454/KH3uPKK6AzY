// 代码生成时间: 2025-08-31 08:21:13
package main

import (
    "fmt"
    "github.com/astaxie/beego"
    "log"
    "runtime"
    "time"
)

// MemoryUsage struct to store memory usage data
type MemoryUsage struct {
    Time    time.Time
    Memory  runtime.MemStats
    CLIArgs []string
}

// StartMemoryUsageAnalysis starts the memory usage analysis
func StartMemoryUsageAnalysis(cliArgs []string) {
    beego.Info("Starting memory usage analysis...")
    // Capture initial memory usage
    initialMemoryUsage()
    // Run memory usage checks at regular intervals
    ticker := time.NewTicker(5 * time.Second)
    defer ticker.Stop()
    for {
        select {
        case <-ticker.C:
            analyzeMemoryUsage(cliArgs)
        }
    }
}

// analyzeMemoryUsage captures memory usage and logs it along with CLI arguments
func analyzeMemoryUsage(cliArgs []string) {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    usage := MemoryUsage{
        Time:    time.Now(),
        Memory:  m,
        CLIArgs: cliArgs,
    }
    fmt.Printf("Time: %v
", usage.Time)
    fmt.Printf("Alloc = %v MiB
", m.Alloc/1024/1024)
    fmt.Printf("TotalAlloc = %v MiB
", m.TotalAlloc/1024/1024)
    fmt.Printf("Sys = %v MiB
", m.Sys/1024/1024)
    fmt.Printf("Mallocs = %v
", m.Mallocs)
    fmt.Printf("Frees = %v
", m.Frees)
    fmt.Printf("CLI Args: %v
", usage.CLIArgs)
}

// initialMemoryUsage captures initial memory usage
func initialMemoryUsage() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Println("Initial memory usage: ", m)
}

func main() {
    // Parse CLI arguments
    cliArgs := []string{fmt.Sprintf("%s", beego.BConfig.Listen.HTTPAddr), fmt.Sprintf("%d", beego.BConfig.Listen.HTTPPort)}
    // Start memory usage analysis
    StartMemoryUsageAnalysis(cliArgs)
}
