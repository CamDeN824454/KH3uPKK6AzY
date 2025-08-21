// 代码生成时间: 2025-08-21 15:12:49
package main

import (
    "beego/logs"
    "fmt"
    "time"
)

// Task 定义一个任务类型，包含任务执行的函数
type Task func() error

// Scheduler 定时任务调度器结构体
type Scheduler struct {
    tasks     map[string]Task
    ticker    *time.Ticker
    stopChan  chan bool
}

// NewScheduler 创建一个新的调度器
func NewScheduler() *Scheduler {
    return &Scheduler{
        tasks: make(map[string]Task),
        stopChan: make(chan bool),
    }
}

// AddTask 添加一个新任务到调度器中
func (s *Scheduler) AddTask(name string, task Task, interval time.Duration) {
    if _, exists := s.tasks[name]; exists {
        logs.Error("Task with name %s already exists", name)
        return
    }
    s.tasks[name] = task
    go s.runTask(name, interval)
}

// runTask 在一个独立的goroutine中运行任务
func (s *Scheduler) runTask(name string, interval time.Duration) {
    s.ticker = time.NewTicker(interval)
    for {
        select {
        case <-s.stopChan:
            s.ticker.Stop()
            return
        case <-s.ticker.C:
            if err := s.tasks[name](); err != nil {
                logs.Error("Error executing task %s: %v", name, err)
            }
        }
    }
}

// Stop 停止调度器，不再运行任何任务
func (s *Scheduler) Stop() {
    s.stopChan <- true
    close(s.stopChan)
}

// ExampleTask 一个示例任务函数
func ExampleTask() error {
    fmt.Println("Executing example task...")
    return nil
}

func main() {
    // 创建调度器
    scheduler := NewScheduler()
    defer scheduler.Stop()

    // 添加一个定时任务，每5秒执行一次ExampleTask
    scheduler.AddTask("example", ExampleTask, 5*time.Second)

    // 运行调度器，直到程序被停止
    select {}
}
