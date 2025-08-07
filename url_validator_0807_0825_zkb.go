// 代码生成时间: 2025-08-07 08:25:11
package main

import (
    "net/url"
    "strings"
    "github.com/astaxie/beego"
)

// URLValidator checks if a given URL is valid
type URLValidator struct {
    beego.Controller
}

// CheckValidURL checks the validity of a URL
// @Title Check URL Validity
// @Description Validate a URL
// @Param url query string true "The URL to validate"
// @Success 200 {string} string "URL is valid"
// @Failure 400 {string} string "Invalid URL"
// @router /validate [get]
func (u *URLValidator) CheckValidURL() {
    // Get the URL from the query parameter
    inputURL := strings.TrimSpace(u.GetString("url"))

    // Check if the input URL is empty
    if inputURL == "" {
        u.Data["json"] = map[string]string{"error": "URL cannot be empty"}
        u.ServeJSON()
        return
    }

    // Parse the URL to check its validity
    parsedURL, err := url.ParseRequestURI(inputURL)
    if err != nil {
        u.Data["json"] = map[string]string{"error": "Invalid URL format"}
        u.ServeJSON()
        return
    }

    // Check if the scheme and host are both set
    if parsedURL.Scheme == "" || parsedURL.Host == "" {
        u.Data["json"] = map[string]string{"error": "URL must have a scheme and a host"}
        u.ServeJSON()
        return
    }

    // If all checks pass, return a success message
    u.Data["json"] = map[string]string{"message": "URL is valid"}
    u.ServeJSON()
}

func main() {
    // Set the Beego framework to run in development mode
    beego.BConfig.AppName = "urlvalidator"
    beego.BConfig.RunMode = "dev"
    beego.BConfig.WebConfig.Session.SessionOn = false

    // Enable output of debug messages
    beego.SetLogger("console", `{"level": "info"}`)

    // Register the URLValidator controller
    beego.Router("/validate", &URLValidator{})

    // Start the Beego server
    beego.Run()
}
