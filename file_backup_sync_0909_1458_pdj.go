// 代码生成时间: 2025-09-09 14:58:10
package main

import (
    "beego/logs"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "path/filepath"
    "strings"
    "time"
)

// BackupAndSync 结构体用于存储备份和同步的配置
type BackupAndSync struct {
    SourceDir  string // 源目录
    TargetDir  string // 目标目录
    SyncPeriod time.Duration // 同步周期
}

// NewBackupAndSync 创建一个新的BackupAndSync实例
func NewBackupAndSync(source, target string, period time.Duration) *BackupAndSync {
    return &BackupAndSync{
        SourceDir:  source,
        TargetDir:  target,
        SyncPeriod: period,
    }
}

// Sync 同步源目录和目标目录
func (bs *BackupAndSync) Sync() error {
    logs.Info("Starting synchronization...")
    defer logs.Info("Synchronization completed.")

    // 获取源目录文件列表
    files, err := ioutil.ReadDir(bs.SourceDir)
    if err != nil {
        return err
    }

    // 遍历文件并同步
    for _, file := range files {
        srcPath := filepath.Join(bs.SourceDir, file.Name())
        dstPath := filepath.Join(bs.TargetDir, file.Name())

        // 如果目标文件存在，则比较文件内容，如果文件内容不同，则进行同步
        if _, err := os.Stat(dstPath); err == nil {
            srcContent, err := ioutil.ReadFile(srcPath)
            if err != nil {
                return err
            }
            dstContent, err := ioutil.ReadFile(dstPath)
            if err != nil {
                return err
            }
            if !strings.EqualFold(string(srcContent), string(dstContent)) {
                if err := os.WriteFile(dstPath, srcContent, file.Mode()); err != nil {
                    return err
                }
            }
        } else {
            // 如果目标文件不存在，则直接复制文件
            if err := copyFile(srcPath, dstPath, file.Mode()); err != nil {
                return err
            }
        }
    }

    return nil
}

// copyFile 复制文件
func copyFile(src, dst string, perm os.FileMode) error {
    sourceFile, err := os.Open(src)
    if err != nil {
        return err
    }
    defer sourceFile.Close()

    destFile, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer destFile.Close()

    _, err = io.Copy(destFile, sourceFile)
    if err != nil {
        return err
    }
    if err := destFile.Sync(); err != nil {
        return err
    }

    return os.Chmod(dst, perm)
}

func main() {
    beego.AppConfig.SetDefault(string(http.MethodGet), "127.0.0.1:8080")
    beego.Router("/sync", &BackupAndSync{"./source", "./destination", 1 * time.Hour})
    beego.Run()
}
