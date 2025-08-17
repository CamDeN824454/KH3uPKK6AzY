// 代码生成时间: 2025-08-18 03:01:22
package controllers

import (
	"beego/validation"
	"github.com/astaxie/beego"
)

// ThemeController handles theme switching.
type ThemeController struct {
	beego.Controller
}

// SwitchTheme changes the current theme.
func (c *ThemeController) SwitchTheme() {
	// Parse the theme parameter from the request.
	theme := c.GetString("theme")
	
	// Validate the theme parameter.
	validThemes := []string{"light", "dark", "colorful"}
	isValid := false
	for _, validTheme := range validThemes {
		if theme == validTheme {
			isValid = true
			break
		}
	}
	
	if !isValid {
		// If the theme is not valid, return an error.
		c.CustomAbort(400, "Invalid theme parameter.")
		return
	}
	
	// Set the theme in the session.
	c.SetSession("currentTheme\,theme)
	
	// Redirect to the previous page or a default page.
	referer := c.Ctx.Request.Referer()
	if referer == "" {
		referer = "/"
	}
	c.Redirect(302, referer)
}
