// 代码生成时间: 2025-09-08 06:41:52
package main

import (
    "fmt"
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/astaxie/beego/orm"
)

// DatabaseConfig 定义数据库配置结构
type DatabaseConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
}

// NewDatabasePool 创建并返回数据库连接池
func NewDatabasePool(config DatabaseConfig) (*sql.DB, error) {
    // 构建DSN（Data Source Name）
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        config.User, config.Password, config.Host, config.Port, config.DBName)

    // 打开数据库连接
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }

    // 设置数据库连接池的参数
    db.SetMaxOpenConns(100) // 设置最大打开连接数
    db.SetMaxIdleConns(50)  // 设置最大闲置连接数
    db.SetConnMaxLifetime(3600) // 设置连接的最大存活时间，单位秒

    // 测试数据库连接
    if err := db.Ping(); err != nil {
        return nil, err
    }

    return db, nil
}

func main() {
    // 数据库配置
    config := DatabaseConfig{
        Host:     "localhost",
        Port:     3306,
        User:     "root",
        Password: "password",
        DBName:   "example_db",
    }

    // 创建数据库连接池
    db, err := NewDatabasePool(config)
    if err != nil {
        log.Fatalf("Failed to create database pool: %v", err)
    }
    defer db.Close()

    // 使用orm进行数据库操作
    o := orm.NewOrm(db)
    // 这里可以添加具体的数据库操作, 例如查询、插入、更新、删除等

    fmt.Println("Database connection pool is initialized successfully.")
}
