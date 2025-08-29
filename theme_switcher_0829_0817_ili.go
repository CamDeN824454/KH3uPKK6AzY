// 代码生成时间: 2025-08-29 08:17:45
// theme_switcher.go
// 该程序使用BEEGO框架实现主题切换功能。

package main

import (
    "encoding/json"
    "github.com/astaxie/beego"
    "net/http"
)

// ThemeService 定义所有与主题相关的操作
type ThemeService struct{}

// SwitchTheme 切换主题
// @Title Theme Switching
// @Description Switches the theme
// @Success 200 {string} string "{status: 'success'}"
// @Failure 400 {string} string "{status: 'error'}"
// @router /switchTheme [post]
func (t *ThemeService) SwitchTheme() string {
    theme := beego.AppConfig.String("theme")
    if theme == "" {
        return "{"status": "error"}"
    }

    // 设置新的theme配置
    beego.SetStaticPath("theme