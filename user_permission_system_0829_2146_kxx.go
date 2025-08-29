// 代码生成时间: 2025-08-29 21:46:17
// user_permission_system.go
// 这个程序是一个用户权限管理系统，使用GOLANG和BEEGO框架实现

package main

import (
    "encoding/json"
    "github.com/astaxie/beego"
    "net/http"
)

// User define user struct
type User struct {
    ID       int    "json:"id"`
    Username string "json:"username"`
    Role     string "json:"role"`
}

// RolePermission define role and its permissions
type RolePermission struct {
    Role       string   "json:"role"`
    Permissions []string "json:"permissions"`
}

// userPermissionsMap holds the user permissions
var userPermissionsMap = map[int]RolePermission{
    1: {Role: "admin", Permissions: []string{"create", "read", "update", "delete"}},
    2: {Role: "editor", Permissions: []string{"create", "read", "update"}},
    3: {Role: "viewer", Permissions: []string{"read"}},
}

// CheckPermission checks if user has the permission
func CheckPermission(userID int, permission string) bool {
    rolePermission, exists := userPermissionsMap[userID]
    if !exists {
        return false
    }
    for _, perm := range rolePermission.Permissions {
        if perm == permission {
            return true
        }
    }
    return false
}

// GetUserPermissions returns user permissions as JSON
func GetUserPermissions(w http.ResponseWriter, r *http.Request) {
    userIDStr := r.URL.Query().Get(":userId")
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        beego.Error("Error getting user ID: ", err)
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }
    userRolePermissions, exists := userPermissionsMap[userID]
    if !exists {
        beego.Error("User not found")
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    response, err := json.Marshal(userRolePermissions)
    if err != nil {
        beego.Error("Error marshalling user permissions: ", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(response)
}

func main() {
    // Set up Beego
    beego.AddFuncMap("json", func(v interface{}) string {
        bytes, _ := json.Marshal(v)
        return string(bytes)
    })

    // Register the handler
    beego.Router("/user/permissions/:userId", &UserController{}, "get:GetUserPermissions")

    // Run the server
    beego.Run()
}
