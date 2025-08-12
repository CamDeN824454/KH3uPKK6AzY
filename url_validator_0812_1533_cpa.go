// 代码生成时间: 2025-08-12 15:33:38
package main

import (
    "bytes"
    "fmt"
    "net/http"
    "net/url"
    "os"
    "regexp"

    beego "github.com/astaxie/beego"
)

// ValidateURL checks if the given URL is valid and reachable.
func ValidateURL(rawurl string) (bool, error) {
    // Parse the URL
    u, err := url.ParseRequestURI(rawurl)
    if err != nil {
        return false, err
    }

    // Validate URL format using a regular expression
    re := regexp.MustCompile(`^(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]*[-A-Za-z0-9+&@#/%=~_|]$`)
    if !re.MatchString(rawurl) {
        return false, fmt.Errorf("invalid URL format")
    }

    // Check if the URL is reachable
    resp, err := http.Head(u.String())
    if err != nil {
        return false, err
    }
    defer resp.Body.Close()

    // Check the status code of the response
    if resp.StatusCode != http.StatusOK {
        return false, fmt.Errorf("URL is not reachable or returned an error status: %d", resp.StatusCode)
    }

    return true, nil
}

func main() {
    // Initialize Beego
    beego.BeeLogger.SetLevel(beego.LevelDebug)

    // Define the URL to validate
    urlToTest := "http://example.com"

    // Perform the validation
    valid, err := ValidateURL(urlToTest)
    if err != nil {
        fmt.Printf("Error validating URL: %s
", err)
        os.Exit(1)
    }

    if valid {
        fmt.Printf("The URL '%s' is valid and reachable.
", urlToTest)
    } else {
        fmt.Printf("The URL '%s' is not valid or not reachable.
", urlToTest)
    }
}
