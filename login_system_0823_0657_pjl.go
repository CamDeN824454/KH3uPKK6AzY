// 代码生成时间: 2025-08-23 06:57:20
package main

import (
    "beego/adapter/template"
    "encoding/json"
    "github.com/astaxie/beego"
    "net/http"
    "strings"
)

// User 用户结构体
type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// LoginHandler 处理登录请求
func LoginHandler() beego.Controller {
    this := controller{
        resp: new(http.ResponseWriter),
        req: new(http.Request),
    }
    return &this
}

// 用户登录控制器
type controller struct {
    resp *http.ResponseWriter
    req  *http.Request
}

// Post 方法处理POST请求
func (c *controller) Post() {
    var u User
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &u); err != nil {
        c.resp.WriteHeader(http.StatusBadRequest)
        c.resp.Write([]byte("Invalid request body"))
        return
    }

    if u.Username == "admin" && u.Password == "password" {
        c.resp.WriteHeader(http.StatusOK)
        c.resp.Write([]byte("Login successful"))
    } else {
        c.resp.WriteHeader(http.StatusUnauthorized)
        c.resp.Write([]byte("Invalid username or password"))
    }
}

func main() {
    // 注册路由
    beego.Router("/login", &controller{}, "post:Get;post:Post")
    
    // 运行服务器
    beego.Run()
}