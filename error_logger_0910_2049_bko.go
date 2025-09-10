// 代码生成时间: 2025-09-10 20:49:21
package main

import (
    "beego/logs"
    "fmt"
    "os"
    "time"
)

// ErrorLogger 结构体用于封装日志配置
type ErrorLogger struct {
    logger *logs.BeeLogger
}

// NewErrorLogger 创建一个新的 ErrorLogger 实例
func NewErrorLogger() *ErrorLogger {
    logger := logs.NewLogger()
    // 设置日志文件名
    logger.SetLogger("file", fmt.Sprintf("error_%s.log", time.Now().Format(\