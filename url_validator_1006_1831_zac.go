// 代码生成时间: 2025-10-06 18:31:40
package main

import (
    "fmt"
    "net/url"
    "strings"
    "github.com/astaxie/beego"
)

// URLValidator is a struct that holds the url to be validated
type URLValidator struct {
    beego.Controller
}

// CheckURL checks if the given URL is valid and returns an appropriate response
func (u *URLValidator) CheckURL() {
    // Get the URL from the request query parameter
    urlStr := u.GetString("url")

    // Validate the URL
    valid, err := validateURL(urlStr)
    if err != nil {
        // Return the error if validation fails
        u.Data["json"] = map[string]string{"error": err.Error()}
        u.ServeJSON()
        return
    }

    // Return the validation result
    u.Data["json"] = map[string]bool{"valid": valid}
    u.ServeJSON()
}

// validateURL checks if the URL is valid
// Returns true if valid, false otherwise, and any error encountered
func validateURL(urlStr string) (bool, error) {
    // Parse the URL
    parsedURL, err := url.ParseRequestURI(urlStr)
    if err != nil {
        return false, err
    }

    // Check if the scheme is valid (e.g., http, https)
    if !strings.HasPrefix(parsedURL.Scheme, "http") {
        return false, fmt.Errorf("invalid scheme: %s", parsedURL.Scheme)
    }

    // Additional validation can be added here (e.g., check host, path)

    return true, nil
}

func main() {
    // Set up the Beego router
    beego.Router("/check-url", &URLValidator{})

    // Start the Beego server
    beego.Run()
}