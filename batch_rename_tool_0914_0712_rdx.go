// 代码生成时间: 2025-09-14 07:12:13
package main

import (
    "bufio"
    "flag"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
)

// BatchRenameTool 结构体包含批量重命名需要的信息
type BatchRenameTool struct {
    SourceDir string // 源目录
    RenameMap map[string]string // 重命名映射
}

// NewBatchRenameTool 初始化BatchRenameTool结构体
func NewBatchRenameTool(sourceDir string) *BatchRenameTool {
    return &BatchRenameTool{
        SourceDir: sourceDir,
        RenameMap: make(map[string]string),
    }
}

// AddRename 添加重命名规则
func (b *BatchRenameTool) AddRename(oldName, newName string) {
    b.RenameMap[oldName] = newName
}

// RenameFiles 执行批量文件重命名
func (b *BatchRenameTool) RenameFiles() error {
    // 获取源目录下所有文件
    files, err := ioutil.ReadDir(b.SourceDir)
    if err != nil {
        return fmt.Errorf("failed to read directory: %w", err)
    }
    for _, file := range files {
        if !file.IsDir() { // 忽略目录
            oldPath := filepath.Join(b.SourceDir, file.Name())
            newFileName := file.Name()
            // 检查是否有重命名规则
            if newName, ok := b.RenameMap[newFileName]; ok {
                newFilePath := filepath.Join(b.SourceDir, newName)
                // 执行重命名操作
                if err := os.Rename(oldPath, newFilePath); err != nil {
                    return fmt.Errorf("failed to rename file %s to %s: %w", oldPath, newFilePath, err)
                }
                fmt.Printf("Renamed %s to %s
", oldPath, newFilePath)
            }
        }
    }
    return nil
}

func main() {
    var sourceDir string
    flag.StringVar(&sourceDir, "dir", ".", "source directory for renaming files")
    flag.Parse()

    // 创建批量重命名工具实例
    renameTool := NewBatchRenameTool(sourceDir)

    // 添加重命名规则
    renameTool.AddRename("oldfile1.txt", "newfile1.txt")
    renameTool.AddRename("oldfile2.txt", "newfile2.txt")
    // 更多的重命名规则可以根据需要添加

    // 执行批量文件重命名
    if err := renameTool.RenameFiles(); err != nil {
        fmt.Printf("Error: %s
", err)
    } else {
        fmt.Println("All files have been renamed successfully.")
    }
}
