// 代码生成时间: 2025-08-19 14:22:10
package main

import (
    "beego"
    "beego/context"
    "encoding/json"
    "fmt"
    "os/exec"
    "strings"
)

// ProcessManager is a handler to manage processes (start, stop, restart, etc.)
type ProcessManager struct{
    // contains filtered or unexported fields
}

// StartProcess starts a new process with the given command
func (pm *ProcessManager) StartProcess(ctx *context.Context) {
    cmd := ctx.Input.Param(":cmd")
    if cmd == "" {
        ctx.Output.JSON(200, beego.NSMap{"error": "No command provided"})
        return
    }

    // Run the command as a new process
    output, err := exec.Command("/bin/sh", "-c", cmd).CombinedOutput()
    if err != nil {
        ctx.Output.JSON(500, beego.NSMap{"error": err.Error(), "output": string(output)})
    } else {
        ctx.Output.JSON(200, beego.NSMap{"success": "Process started successfully", "output": string(output)})
    }
}

// StopProcess stops the process with the given process ID
func (pm *ProcessManager) StopProcess(ctx *context.Context) {
    pid := ctx.Input.Param(":pid")
    if pid == "" {
        ctx.Output.JSON(200, beego.NSMap{"error": "No process ID provided"})
        return
    }

    // Convert the PID to an integer
    pidInt, err := strconv.Atoi(pid)
    if err != nil {
        ctx.Output.JSON(400, beego.NSMap{"error": "Invalid process ID format"})
        return
    }

    // Send SIGTERM to the process to stop it gracefully
    err = exec.Command("kill", "SIGTERM", pid).Run()
    if err != nil {
        ctx.Output.JSON(500, beego.NSMap{"error": err.Error()})
    } else {
        ctx.Output.JSON(200, beego.
            NSMap{"success": "Process stopped successfully"})
    }
}

func main() {
    // Create a new Beego application
    beego.Router("/process/start", &ProcessManager{}, "get:StartProcess")
    beego.Router("/process/stop", &ProcessManager{}, "get:StopProcess")

    // Run the application
    beego.Run()
}
