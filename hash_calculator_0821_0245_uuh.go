// 代码生成时间: 2025-08-21 02:45:22
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "os"
    "strings"
    "beego"
)

// HashCalculator 定义哈希值计算工具
type HashCalculator struct {
    // 这里可以添加更多的属性
}

// GenerateHash 生成哈希值
func (h *HashCalculator) GenerateHash(input string) (string, error) {
    // 使用SHA256算法生成哈希值
    hash := sha256.Sum256([]byte(input))
    return hex.EncodeToString(hash[:]), nil
}

func main() {
    beego.Router("/hash", &HashCalculator{})
    beego.Run()
}

// controller的实现
func (h *HashCalculator) Get() {
    var input string
    // 从请求中获取输入值
    if ctx := beego.BeeApp.Handlers; ctx.Request.Method == "GET" {
        input = ctx.Request.URL.Query().Get("input")
    }

    // 检查输入是否为空
    if input == "" {
        // 响应错误信息
        beego.Controller.Ctx.ResponseWriter.WriteHeader(400)
        fmt.Fprintf(beego.Controller.Ctx.ResponseWriter, "Input cannot be empty")
        return
    }

    // 生成哈希值
    hash, err := h.GenerateHash(input)
    if err != nil {
        // 响应错误信息
        beego.Controller.Ctx.ResponseWriter.WriteHeader(500)
        fmt.Fprintf(beego.Controller.Ctx.ResponseWriter, "Error generating hash: %s", err)
        return
    }

    // 响应哈希值
    beego.Controller.Data["json"] = map[string]string{
        "hash": hash,
    }
    beego.Controller.ServeJSON()
}