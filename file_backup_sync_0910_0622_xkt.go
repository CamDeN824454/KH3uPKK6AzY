// 代码生成时间: 2025-09-10 06:22:21
package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"
    "github.com/astaxie/beego"
)

// Config stores the configuration for the backup and sync tool
type Config struct {
    SourceDirectory string `json:"source_directory"`
    DestinationDirectory string `json:"destination_directory"`
    SyncInterval time.Duration `json:"sync_interval"`
}

// BackupSyncTool is the main struct for the backup and sync tool
type BackupSyncTool struct {
    config *Config
}

// NewBackupSyncTool creates a new backup and sync tool with the given configuration
func NewBackupSyncTool(config *Config) *BackupSyncTool {
    return &BackupSyncTool{config: config}
}

// StartBackupSync starts the backup and sync process
func (b *BackupSyncTool) StartBackupSync() {
    // Check if the source directory exists
    if _, err := os.Stat(b.config.SourceDirectory); os.IsNotExist(err) {
        log.Fatalf("Source directory does not exist: %s", b.config.SourceDirectory)
    }

    // Check if the destination directory exists, create if not
    if _, err := os.Stat(b.config.DestinationDirectory); os.IsNotExist(err) {
        log.Printf("Destination directory does not exist, creating: %s", b.config.DestinationDirectory)
        os.MkdirAll(b.config.DestinationDirectory, 0755)
    }

    // Start the backup and sync loop
    for {
        err := b.backupAndSync()
        if err != nil {
            log.Printf("Error during backup and sync: %s", err)
        }

        // Wait for the sync interval before starting the next loop
        time.Sleep(b.config.SyncInterval)
    }
}

// backupAndSync performs the backup and sync operation
func (b *BackupSyncTool) backupAndSync() error {
    // Walk through the source directory and sync files
    err := filepath.Walk(b.config.SourceDirectory, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // Skip directories
        if info.IsDir() {
            return nil
        }

        // Construct the destination file path
        relativePath, err := filepath.Rel(b.config.SourceDirectory, path)
        if err != nil {
            return err
        }
        destPath := filepath.Join(b.config.DestinationDirectory, relativePath)

        // Check if the file exists in the destination directory
        if _, err := os.Stat(destPath); os.IsNotExist(err) {
            // Copy the file to the destination directory
            sourceFile, err := os.Open(path)
            if err != nil {
                return err
            }
            defer sourceFile.Close()

            destFile, err := os.Create(destPath)
            if err != nil {
                return err
            }
            defer destFile.Close()

            _, err = io.Copy(destFile, sourceFile)
            return err
        } else if err != nil {
            return err
        }

        // Check if the source file is newer than the destination file
        sourceModTime := info.ModTime()
        destFileInfo, err := os.Stat(destPath)
        if err != nil {
            return err
        }
        destModTime := destFileInfo.ModTime()

        if sourceModTime.After(destModTime) {
            // Copy the file to the destination directory
            sourceFile, err := os.Open(path)
            if err != nil {
                return err
            }
            defer sourceFile.Close()

            destFile, err := os.Create(destPath)
            if err != nil {
                return err
            }
            defer destFile.Close()

            _, err = io.Copy(destFile, sourceFile)
            return err
        }

        return nil
    })

    return err
}

func main() {
    // Load the configuration from a JSON file
    configData, err := ioutil.ReadFile("config.json")
    if err != nil {
        log.Fatalf("Error reading configuration file: %s", err)
    }

    var config Config
    err = json.Unmarshal(configData, &config)
    if err != nil {
        log.Fatalf("Error parsing configuration file: %s", err)
    }

    // Create a new backup and sync tool with the loaded configuration
    backupSyncTool := NewBackupSyncTool(&config)

    // Start the backup and sync process
    backupSyncTool.StartBackupSync()
}
