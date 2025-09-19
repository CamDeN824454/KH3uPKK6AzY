// 代码生成时间: 2025-09-19 08:14:01
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/astaxie/beego"
    "github.com/robfig/cron/v3"
)

// Task 定义了定时任务的结构体
type Task struct {
    Name    string
    CronExp string
    Command func() error
}

// NewTask 创建一个新的任务
func NewTask(name, cronExp string, command func() error) *Task {
    return &Task{name, cronExp, command}
}

// RunTask 运行指定的任务
func (t *Task) RunTask() {
    cronExpression := t.CronExp
    taskName := t.Name
    cmd := t.Command

    // 创建一个新的cron调度器
    c := cron.New(cron.WithSeconds())
    _, err := c.AddFunc(cronExpression, func() { cmd() })
    if err != nil {
        log.Fatalf("Failed to add task: %s", err)
    }

    // 打印任务调度信息
    fmt.Printf("Task %s has been scheduled to run every %s
", taskName, cronExpression)
    fmt.Printf("Press Ctrl+C to exit
")

    // 运行调度器
    c.Start()

    // 等待中断信号来停止调度器
    sigCh := make(chan os.Signal, 1)
    signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
    <-sigCh
    fmt.Println("Shutting down...
")
    c.Stop()
}

func main() {
    // 初始化beego框架
    beego.Run()

    // 定义任务
    tasks := []*Task{
        NewTask("TestTask", "*/10 * * * *", testTask),
    }

    // 运行任务
    for _, task := range tasks {
        go task.RunTask()
    }
}

// testTask 是一个示例任务
func testTask() error {
    fmt.Println("Executing test task")
    return nil
}
