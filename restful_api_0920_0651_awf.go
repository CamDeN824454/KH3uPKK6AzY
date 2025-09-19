// 代码生成时间: 2025-09-20 06:51:59
package main

import (
    "encoding/json"
    "github.com/astaxie/beego"
    "net/http"
)

// 定义一个结构体来表示用户
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func main() {
    // 初始化Beego框架
    beego.BeeApp.Handlers.Clear()

    // 注册路由
    user := &UserController{}
    beego.Router("/user/:id", &user, "get:Get; put:Update; delete:Delete")

    // 启动服务
    beego.Run()
}

// UserController 定义用户控制器
type UserController struct {
    beego.Controller
}

// Get 获取指定ID的用户
func (c *UserController) Get() {
    id := c.Ctx.Input.Param(":id")
    user, err := GetUserByID(id)
    if err != nil {
        c.CustomAbort(404, "User not found")
        return
    }
    c.Data["json"] = user
    c.ServeJSON()
}

// Update 更新指定ID的用户
func (c *UserController) Update() {
    id := c.Ctx.Input.Param(":id\)
    var user User
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &user); err != nil {
        c.CustomAbort(400, "Invalid request")
        return
    }
    if err := UpdateUserByID(id, user); err != nil {
        c.CustomAbort(500, "Update failed")
        return
    }
    c.Data["json"] = user
    c.ServeJSON()
}

// Delete 删除指定ID的用户
func (c *UserController) Delete() {
    id := c.Ctx.Input.Param(":id\)
    if err := DeleteUserByID(id); err != nil {
        c.CustomAbort(500, "Delete failed")
        return
    }
    c.Data["json"] = map[string]string{
        "message": "User deleted successfully",
    }
    c.ServeJSON()
}

// 模拟数据库操作
var users = map[int]User{
    1: {ID: 1, Name: "John Doe"},
    2: {ID: 2, Name: "Jane Doe"},
}

// GetUserByID 获取指定ID的用户
func GetUserByID(id string) (User, error) {
    if user, exists := users[parseInt(id)]; exists {
        return user, nil
    }
    return User{}, nil
}

// UpdateUserByID 更新指定ID的用户
func UpdateUserByID(id string, user User) error {
    if _, exists := users[parseInt(id)]; !exists {
        return nil
    }
    users[parseInt(id)] = user
    return nil
}

// DeleteUserByID 删除指定ID的用户
func DeleteUserByID(id string) error {
    if _, exists := users[parseInt(id)]; !exists {
        return nil
    }
    delete(users, parseInt(id))
    return nil
}

// parseInt 将字符串转换为整数
func parseInt(s string) int {
    v, err := strconv.Atoi(s)
    if err != nil {
        panic(err)
    }
    return v
}