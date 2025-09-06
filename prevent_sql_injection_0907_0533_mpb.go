// 代码生成时间: 2025-09-07 05:33:29
package main

import (
    "fmt"
    "github.com/astaxie/beego/orm"
)

// 初始化数据库连接
func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)
    orm.RegisterDataBase("default", "mysql", "username:password@tcp(127.0.0.1:3306)/dbname?charset=utf8")
}

// User struct 定义用户模型
type User struct {
    Id   int    `orm:"auto"`
    Name string `orm:"size(100)"`
}

// GetUserByUsername 防止SQL注入，通过ORM的参数化查询
func GetUserByUsername(username string) (*User, error) {
    o := orm.NewOrm()
    var user User
    // 使用参数化查询来防止SQL注入
    err := o.QueryTable("user").Filter("name", username).One(&user)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func main() {
    // 模拟获取用户
    user, err := GetUserByUsername("testUser")
    if err != nil {
        fmt.Println("Error: ", err)
    } else {
        fmt.Printf("User Found: %+v
", user)
    }
}
