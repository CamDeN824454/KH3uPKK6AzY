// 代码生成时间: 2025-10-11 01:36:25
package main

import (
    "encoding/json"
    "fmt"
    "github.com/astaxie/beego"
)

// LearningResource represents the structure of a learning resource.
type LearningResource struct {
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Content   string `json:"content"`
    Category  string `json:"category"`
    CreatedAt string `json:"createdAt"`
    UpdatedAt string `json:"updatedAt"`
}

// Controller for learning resources.
type LearningResourceController struct {
    beego.Controller
}

// Get all learning resources.
func (c *LearningResourceController) GetAll() {
    resources := []LearningResource{
        {ID: 1, Title: "Go Basics", Content: "Learn the basics of Go.", Category: "Programming", CreatedAt: "2023-04-01", UpdatedAt: "2023-04-01"},
        {ID: 2, Title: "Beego Framework", Content: "Dive into the Beego framework.", Category: "Web Development", CreatedAt: "2023-04-02", UpdatedAt: "2023-04-02"},
    }
    c.Data["json"] = resources
    c.ServeJSON()
}

// Get a single learning resource by ID.
func (c *LearningResourceController) GetByID() {
    id := c.Ctx.Input.Param(":id")
    resource := LearningResource{ID: 1, Title: "Go Basics", Content: "Learn the basics of Go.", Category: "Programming", CreatedAt: "2023-04-01", UpdatedAt: "2023-04-01"}

    // Check if the ID is valid and resource exists.
    if id != "" && id == "1" {
        c.Data["json"] = resource
        c.ServeJSON()
    } else {
        c.Ctx.Output.SetStatus(404)
        c.Data["json"] = map[string]string{"error": "Resource not found"}
        c.ServeJSON()
    }
}

func main() {
    // Initialize Beego framework.
    beego.AddFuncMap("test", func() string {
        return "Hello, World!"
    })

    // Register the learning resource controller with the Beego framework.
    beego.Router("/resources", &LearningResourceController{}, "*:GetAll")
    beego.Router("/resources/:id", &LearningResourceController{}, "*:GetByID")

    // Start the Beego web server.
    fmt.Println("Server is running on http://localhost:8080")
    beego.Run()
}