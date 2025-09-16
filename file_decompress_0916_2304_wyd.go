// 代码生成时间: 2025-09-16 23:04:07
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "archive/zip"
    "log"
)

// DecompressTool 压缩文件解压工具
type DecompressTool struct{}

// Unzip 解压ZIP文件
// @param zipFilePath 压缩文件的路径
// @param destPath 解压目标路径
// @return error 如果解压失败则返回错误
func (tool *DecompressTool) Unzip(zipFilePath, destPath string) error {
    // 读取ZIP文件
    src, err := zip.OpenReader(zipFilePath)
    if err != nil {
        return fmt.Errorf("failed to open zip file: %w", err)
    }
    defer src.Close()

    // 确保目标路径存在
    if _, err := os.Stat(destPath); os.IsNotExist(err) {
        if err := os.MkdirAll(destPath, 0755); err != nil {
            return fmt.Errorf("failed to create destination directory: %w", err)
        }
    }

    // 解压文件
    for _, file := range src.File {
        targetPath := filepath.Join(destPath, file.Name)
        // 创建文件
        if file.FileInfo().IsDir() {
            os.MkdirAll(targetPath, 0755)
            continue
        }
        
        fileReader, err := file.Open()
        if err != nil {
            return fmt.Errorf("failed to open file in zip: %w", err)
        }
        defer fileReader.Close()
        
        // 创建文件或覆盖文件
        targetFile, err := os.Create(targetPath)
        if err != nil {
            return fmt.Errorf("failed to create file on disk: %w", err)
        }
        defer targetFile.Close()
        
        // 复制文件内容
        if _, err := io.Copy(targetFile, fileReader); err != nil {
            return fmt.Errorf("failed to copy file content: %w", err)
        }
    }

    return nil
}

func main() {
    var tool DecompressTool
    zipFilePath := "example.zip"
    destPath := "extracted"
    if err := tool.Unzip(zipFilePath, destPath); err != nil {
        log.Fatalf("an error occurred while unzipping: %s", err)
    } else {
        fmt.Println("File decompressed successfully!")
    }
}
