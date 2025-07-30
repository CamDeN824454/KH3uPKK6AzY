// 代码生成时间: 2025-07-30 21:45:22
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "os"
    "strings"

    "github.com/astaxie/beego"
)

// HashCalculatorController is the controller to handle hash calculation requests
type HashCalculatorController struct {
    beego.Controller
}

// Post handles the POST request to calculate the hash of the input string
func (c *HashCalculatorController) Post() {
    var input string
    if err := c.GetString(&input, "input"); err != nil {
        c.CustomAbort(400, "Invalid input parameter")
        return
    }

    if input == "" {
        c.CustomAbort(400, "Input string cannot be empty")
        return
    }

    hash := sha256.Sum256([]byte(input))
    hashString := hex.EncodeToString(hash[:])

    // Set the response data and type
    c.Data["json"] = map[string]string{
        "status": "success",
        "hash": hashString,
    }
    c.ServeJSON()
}

// main is the entry point of the application
func main() {
    // Initialize the Beego framework
    beego.AddFuncMap("html", "htmlescape", `{{html "{{.}}"}}`)
    // Register the controller
    beego.Router("/hash", &HashCalculatorController{})
    // Start the HTTP server
    beego.Run()
}
