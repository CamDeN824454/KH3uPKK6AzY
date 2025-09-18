// 代码生成时间: 2025-09-18 08:43:42
package main

import (
    "log"
    "time"

    "github.com/astaxie/beego"
    "github.com/robfig/cron/v3"
)

// TaskScheduler represents a simple task scheduler.
type TaskScheduler struct {
    cron *cron.Cron
}

// NewTaskScheduler creates a new instance of TaskScheduler with a default scheduler.
func NewTaskScheduler() *TaskScheduler {
    scheduler := new(TaskScheduler)
    scheduler.cron = cron.New(cron.WithSeconds())
    return scheduler
}

// AddTask adds a new task to the scheduler.
// The task will be executed at the specified time using the cron syntax.
func (s *TaskScheduler) AddTask(cronExpr string, task func()) error {
    if _, err := s.cron.AddFunc(cronExpr, task); err != nil {
        return err
    }
    return nil
}

// Start starts the scheduler.
func (s *TaskScheduler) Start() {
    s.cron.Start()
}

// Stop stops the scheduler.
func (s *TaskScheduler) Stop() {
    s.cron.Stop()
}

// An example task that logs a message.
func exampleTask() {
    log.Println("Executing example task...")
}

func main() {
    beego.SetLogger(beego.NewConsoleLogger(10000))

    // Create a new task scheduler.
    scheduler := NewTaskScheduler()

    // Add an example task to the scheduler to run every 5 seconds.
    if err := scheduler.AddTask("*/5 * * * * *", exampleTask); err != nil {
        log.Printf("Failed to add task: %s", err)
        return
    }

    // Start the scheduler.
    scheduler.Start()

    // Block the main goroutine to keep the application running.
    // In a real application, you may want to handle this differently.
    select{}
}
