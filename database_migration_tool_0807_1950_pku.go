// 代码生成时间: 2025-08-07 19:50:42
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "github.com/astaxie/beego/migration"
)

// DatabaseMigrationTool is a tool to manage database migrations.
type DatabaseMigrationTool struct {
    // Up will apply all available migrations.
    Up func() error
    // Down will revert the last migration.
    Down func() error
    // Reset will revert all migrations and re-apply them.
    Reset func() error
    // MigratePath is the path to the migration files.
    MigratePath string
}

// NewDatabaseMigrationTool creates a new instance of DatabaseMigrationTool.
func NewDatabaseMigrationTool(migratePath string) *DatabaseMigrationTool {
    return &DatabaseMigrationTool{
        Up: func() error {
            return migration.UpAll()
        },
        Down: func() error {
            return migration.DownAll()
        },
        Reset: func() error {
            if err := migration.DownAll(); err != nil {
                return err
            }
            return migration.UpAll()
        },
        MigratePath: migratePath,
    }
}

// ExecuteMigrations applies the migrations based on the command provided.
func (dmt *DatabaseMigrationTool) ExecuteMigrations(command string) error {
    switch command {
    case "up":
        return dmt.Up()
    case "down":
        return dmt.Down()
    case "reset":
        return dmt.Reset()
    default:
        return fmt.Errorf("unknown command: %s", command)
    }
}

// RunMigrations sets up the migration tool and runs the specified command.
func RunMigrations(command string, migratePath string) {
    if _, err := os.Stat(migratePath); os.IsNotExist(err) {
        fmt.Printf("Migration directory does not exist: %s
", migratePath)
        os.Exit(1)
    }

    dmt := NewDatabaseMigrationTool(migratePath)
    if err := dmt.ExecuteMigrations(command); err != nil {
        fmt.Printf("Migration error: %s
", err)
        os.Exit(1)
    }
    fmt.Println("Migration completed successfully.")
}

// SetMigratePath sets the path to the migration files.
func SetMigratePath(path string) {
    migration.SetMigrationSourcePath(path)
}

func main() {
    migratePath := "./migrations" // Default migration directory path.
    command := "up" // Default command to apply migrations.

    // Parse command line arguments for custom migration path and command.
    if len(os.Args) > 1 {
        if strings.HasPrefix(os.Args[1], "--path=") {
            migratePath = strings.TrimPrefix(os.Args[1], "--path=")
            if !strings.HasSuffix(migratePath, string(filepath.Separator)) {
                migratePath += string(filepath.Separator)
            }
        } else if os.Args[1] == "down" || os.Args[1] == "reset" {
            command = os.Args[1]
        }
    }

    SetMigratePath(migratePath)
    RunMigrations(command, migratePath)
}
