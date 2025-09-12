// 代码生成时间: 2025-09-12 08:45:09
package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"
)

// 定义批量重命名工具的结构体
type BulkRenamer struct {
    rootDir string
    pattern string
    newPrefix string
}

// NewBulkRenamer 创建一个新的批量重命名工具
func NewBulkRenamer(rootDir, pattern, newPrefix string) *BulkRenamer {
    return &BulkRenamer{
        rootDir: rootDir,
        pattern: pattern,
        newPrefix: newPrefix,
    }
}

// RenameFiles 执行批量重命名操作
func (br *BulkRenamer) RenameFiles() error {
    err := filepath.Walk(br.rootDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.IsDir() {
            return nil
        }

        // 检查文件名是否符合指定的模式
        if matched, _ := filepath.Match(br.pattern, filepath.Base(path)); matched {
            newFileName := br.newPrefix + filepath.Ext(path)
            newFilePath := filepath.Join(filepath.Dir(path), newFileName)

            // 检查新文件名是否已存在
            if _, err := os.Stat(newFilePath); err == nil {
                return fmt.Errorf("file %s already exists", newFilePath)
            }

            // 重命名文件
            if err := os.Rename(path, newFilePath); err != nil {
                return err
            }
            fmt.Printf("Renamed %s to %s
", path, newFilePath)
        }

        return nil
    })
    return err
}

func main() {
    rootDirPtr := flag.String("rootDir", "./", "Root directory to search for files")
    patternPtr := flag.String("pattern", "*.txt", "Pattern to match file names")
    newPrefixPtr := flag.String("newPrefix", "new_", "New prefix for file names")

    // 解析命令行参数
    flag.Parse()

    // 创建批量重命名工具实例
    renamer := NewBulkRenamer(*rootDirPtr, *patternPtr, *newPrefixPtr)

    // 执行重命名操作
    if err := renamer.RenameFiles(); err != nil {
        log.Fatalf("Error during file renaming: %v", err)
    }
}
