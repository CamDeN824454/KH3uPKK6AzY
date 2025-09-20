// 代码生成时间: 2025-09-21 02:39:33
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strings"
    "golang.org/x/net/html"
    "github.com/astaxie/beego"
)

// ScrapeData defines the structure to hold scraped data
type ScrapeData struct {
    Title   string `json:"title"`
    Content string `json:"content"`
}

// Scraper is a function that takes a URL, fetches the HTML content, and extracts the title and content
func Scraper(url string) (*ScrapeData, error) {
    // Create an HTTP GET request
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    // Parse HTML
    node, err := html.Parse(resp.Body)
    if err != nil {
        return nil, err
    }

    // Walk the HTML tree to find title and content
    var title, content string
    var scraper func(*html.Node)
    scraper = func(n *html.Node) {
        switch n.Type {
        case html.ElementType:
            switch n.Data {
            case "title":
                if len(n.FirstChild) > 0 {
                    title = strings.TrimSpace(n.FirstChild.Data)
                }
            case "div":
                if len(n.Attr) > 0 && n.Attr[0].Val == "content" {
                    if len(n.FirstChild) > 0 {
                        content = strings.TrimSpace(n.FirstChild.Data)
                    }
                }
            }
        }
        // Continue to walk through the HTML tree
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            scraper(c)
        }
    }
    scraper(node)

    // Return scraped data
    return &ScrapeData{Title: title, Content: content}, nil
}

// ScrapeHandler is the Beego controller to handle scraping requests
type ScrapeHandler struct {
    beego.Controller
}

// Get method to handle GET requests and scrape the provided URL
func (sh *ScrapeHandler) Get() {
    url := sh.GetString("url")
    if url == "" {
        sh.Data[{"json"}] = map[string]string{"error": "URL parameter is missing"}
        sh.ServeJSON()
        return
    }

    data, err := Scraper(url)
    if err != nil {
        sh.Data[{"json"}] = map[string]string{"error": err.Error()}
    } else {
        sh.Data[{"json"}] = data
    }
    sh.ServeJSON()
}

func main() {
    // Set Beego to run in debug mode
    beego.BConfig.RunMode = "dev"

    // Register the ScrapeHandler to handle GET requests at /scrape URL
    beego.Router("/scrape", &ScrapeHandler{}, "get:Get")

    // Start the Beego application
    beego.Run()
}
