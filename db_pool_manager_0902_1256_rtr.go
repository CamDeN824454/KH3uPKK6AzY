// 代码生成时间: 2025-09-02 12:56:24
package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql" // MySQL driver
    "log"
    "os"
    "time"
)

// DBConfig holds the configuration for the database connection
type DBConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
}

// DBPool represents the database connection pool
type DBPool struct {
    *sql.DB
}

// NewDBPool creates and returns a new database connection pool
func NewDBPool(cfg DBConfig) (*DBPool, error) {
    // Construct the DSN (Data Source Name)
    dsn := cfg.User + ":" + cfg.Password + "@tcp(" + cfg.Host + ":" + strconv.Itoa(cfg.Port) + ")/" + cfg.DBName + "?parseTime=True"

    // Open the database connection
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }

    // Set the maximum number of connections in the idle connection pool
    db.SetMaxIdleConns(10)

    // Set the maximum number of open connections to the database
    db.SetMaxOpenConns(100)

    // Set the time to wait for a connection to be available
    db.SetConnMaxLifetime(30 * time.Minute)

    // Ping the database to check the connection
    if err = db.Ping(); err != nil {
        db.Close()
        return nil, err
    }

    return &DBPool{db}, nil
}

func main() {
    cfg := DBConfig{
        Host:     "localhost",
        Port:     3306,
        User:     "your_username",
        Password: "your_password",
        DBName:   "your_dbname",
    }

    dbPool, err := NewDBPool(cfg)
    if err != nil {
        log.Fatalf("Failed to create DB pool: %v", err)
    }
    defer dbPool.Close()

    // Use the dbPool to perform database operations...
    // Example:
    // result := dbPool.Query("SELECT * FROM your_table")
    // ... handle the result ...
}
