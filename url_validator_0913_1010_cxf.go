// 代码生成时间: 2025-09-13 10:10:59
package main

import (
    "fmt"
    "net/http"
    "net/url"
    "beego"
)

// URLValidator defines a struct for URL validation
type URLValidator struct {
    // nothing required for this simple example
}

// ValidateURL checks if the provided URL is valid
func (u *URLValidator) ValidateURL(URL string) (bool, error) {
    // Parse the URL to check its validity
    parsedURL, err := url.ParseRequestURI(URL)
    if err != nil {
        return false, err
    }

    // Check if the scheme is valid (http or https)
    if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
        return false, fmt.Errorf("invalid URL scheme: %s", parsedURL.Scheme)
    }

    // Check if the host is not empty
    if parsedURL.Host == "" {
        return false, fmt.Errorf("URL host is empty")
    }

    // If no errors, return true indicating the URL is valid
    return true, nil
}

func main() {
    // Register the URLValidator controller
    beego.Router("/validate", &URLValidator{})

    // Start the Beego application
    beego.Run()
}

// URLValidationController handles HTTP requests for validating URLs
func (u *URLValidator) Get() {
    var URL string
    // Bind the query parameter to the URL variable
    if err := beego.Input().Bind(&URL, "url"); err != nil {
        beego.Error(err)
        beego.ReturnJSON(&beego.Controller{Ctx: beego.Ctx{}}, `{"valid":false,"error":"Failed to bind URL parameter"}`)
        return
    }

    // Validate the URL
    valid, err := u.ValidateURL(URL)
    if err != nil {
        beego.Error(err)
        beego.ReturnJSON(&beego.Controller{Ctx: beego.Ctx{}}, `{"valid":false,"error":"Invalid URL"}`)
        return
    }

    // Return the validation result as JSON
    beego.ReturnJSON(&beego.Controller{Ctx: beego.Ctx{}}, fmt.Sprintf(`{"valid":%v}`, valid))
}