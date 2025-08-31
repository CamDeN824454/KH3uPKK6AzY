// 代码生成时间: 2025-09-01 03:10:17
package main

import (
    "bytes"
    "fmt"
    "net/http"
    "strings"

    "github.com/astaxie/beego"
    "github.com/astaxie/beego/context"
    "golang.org/x/net/html"
)

// escapeHTML escapes HTML special characters in a string.
func escapeHTML(s string) string {
    var buf bytes.Buffer
    html.Escape(&buf, []byte(s))
    return buf.String()
}

// xssFilter is a middleware that escapes HTML to prevent XSS attacks.
func xssFilter(ctx *context.Context) {
    // Before the actual request processing
    ctx.Input.SetData("XSS Filter", true)
    ctx.ResponseWriter.Header().Set("X-XSS-Protection", "1; mode=block")

    // Write the original response
    err := ctx.ResponseWriter.Write(ctx.ResponseWriter.Body)
    if err != nil {
        beego.Error("Error writing response: ", err)
        ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
        return
    }

    // Replace the body with the escaped version
    ctx.ResponseWriter.Body = []byte(escapeHTML(string(ctx.ResponseWriter.Body)))
}

func main() {
    beego.InsertFilter("*", beego.BeforeRouter, xssFilter)
    beego.Router("/", &MainController{})
    beego.Run()
}

// MainController handles the main page request.
type MainController struct {
    beego.Controller
}

// Get method for the main page.
func (c *MainController) Get() {
    c.Data["Website"] = "Beego Framework"
    c.Data["Title"] = "XSS Protection Example"
    c.TplNames = "index.tpl"
}

// index.tpl is the template for the main page.
// Please note that this template should be placed in the views folder and should
// include the necessary logic to escape the output to prevent XSS.
// Here is a sample content for the template:

// {{define "index"}}
// <html>
// <head>
//     <title>{{.Title}}</title>
// </head>
// <body>
//     <h1>Welcome to {{.Website}}</h1>
//     <p>This is an example of XSS protection in action.</p>
// </body>
// </html>
// {{end}}
