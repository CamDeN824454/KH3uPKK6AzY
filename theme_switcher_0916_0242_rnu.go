// 代码生成时间: 2025-09-16 02:42:57
package main

import (
    "encoding/json"
    "github.com/astaxie/beego"
    "net/http"
)

// ThemeSwitcher defines the structure for theme switching
type ThemeSwitcher struct {
    beego.Controller
}

// SwitchTheme handles the theme switching by setting a cookie.
// It expects a POST request with a JSON body containing the desired theme.
func (t *ThemeSwitcher) SwitchTheme() {
    // Parse the JSON from the request body
    var themeData struct {
        Theme string `json:"theme"`
    }
    if err := json.Unmarshal(t.Ctx.Input.RequestBody, &themeData); err != nil {
        t.Ctx.WriteString("Invalid JSON in request body.")
        t.Ctx.SetStatus(http.StatusBadRequest)
        return
    }

    // Check if the theme is valid
    if !isValidTheme(themeData.Theme) {
        t.Ctx.WriteString("Invalid theme provided.")
        t.Ctx.SetStatus(http.StatusBadRequest)
        return
    }

    // Set the theme as a cookie
    t.Ctx.SetCookie("theme", themeData.Theme, 86400, "/", "", false, true)
    t.Ctx.WriteString("Theme switched to " + themeData.Theme)
}

// isValidTheme checks if a theme is valid.
// It should be updated to include all available themes.
func isValidTheme(theme string) bool {
    validThemes := []string{"light", "dark", "colorful"}
    for _, validTheme := range validThemes {
        if theme == validTheme {
            return true
        }
    }
    return false
}

func main() {
    // Initialize the Beego application
    beego.Application.Run()
}
