// 代码生成时间: 2025-09-06 19:24:02
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "beego/logs"
    "beego/toolbox"
)

// FileBackupSync is a struct that holds the source and destination directories
type FileBackupSync struct {
    SourceDir string
    TargetDir string
}

// NewFileBackupSync creates a new instance of FileBackupSync
func NewFileBackupSync(source, target string) *FileBackupSync {
    return &FileBackupSync{SourceDir: source, TargetDir: target}
}

// BackupAndSync performs the backup and sync operation
func (fbs *FileBackupSync) BackupAndSync() error {
    // Check if source directory exists
    if _, err := os.Stat(fbs.SourceDir); os.IsNotExist(err) {
        return fmt.Errorf("source directory '%s' does not exist", fbs.SourceDir)
    }

    // Create target directory if it doesn't exist
    if _, err := os.Stat(fbs.TargetDir); os.IsNotExist(err) {
        if err := os.MkdirAll(fbs.TargetDir, 0755); err != nil {
            return fmt.Errorf("failed to create target directory: %s", err)
        }
    }

    // Walk through the source directory
    err := filepath.Walk(fbs.SourceDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // Skip directories
        if info.IsDir() {
            return nil
        }

        // Construct the target file path
        relPath, err := filepath.Rel(fbs.SourceDir, path)
        if err != nil {
            return err
        }
        targetPath := filepath.Join(fbs.TargetDir, relPath)

        // Copy file from source to target
        if err := toolbox.CopyFile(path, targetPath); err != nil {
            logs.Error("Failed to copy file from %s to %s: %v", path, targetPath, err)
            return err
        }

        return nil
    })

    if err != nil {
        return err
    }

    return nil
}

func main() {
    // Set up logging
    logs.SetLogger(logs.AdapterFile, `{"filename":"backup.log"}`)
    defer logs.Flush()

    // Create a new FileBackupSync instance
    fbs := NewFileBackupSync("/path/to/source", "/path/to/destination")

    // Perform backup and sync
    if err := fbs.BackupAndSync(); err != nil {
        fmt.Printf("Backup and sync failed: %s
", err)
    } else {
        fmt.Println("Backup and sync completed successfully")
    }
}
