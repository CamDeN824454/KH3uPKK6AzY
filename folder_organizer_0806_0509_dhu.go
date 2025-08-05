// 代码生成时间: 2025-08-06 05:09:46
package main

import (
    "beego"
    "os"
    "path/filepath"
    "strings"
)

// FolderOrganizer 结构体，用于存储文件夹的路径信息
type FolderOrganizer struct {
    Path string
}

// NewFolderOrganizer 创建并返回一个新的 FolderOrganizer 实例
func NewFolderOrganizer(path string) *FolderOrganizer {
    return &FolderOrganizer{Path: path}
}

// Organize 整理指定文件夹，将文件按照类型分类存放
func (fo *FolderOrganizer) Organize() error {
    // 获取文件夹内所有子目录和文件
    files, err := os.ReadDir(fo.Path)
    if err != nil {
        return err
    }

    for _, file := range files {
        if file.IsDir() {
            continue
        }

        // 获取文件扩展名
        extension := strings.TrimPrefix(filepath.Ext(file.Name()), ".")
        if extension == "" {
            continue
        }

        // 创建新的文件夹路径
        newFolderPath := filepath.Join(fo.Path, extension)
        if _, err := os.Stat(newFolderPath); os.IsNotExist(err) {
            if err := os.MkdirAll(newFolderPath, 0755); err != nil {
                return err
            }
        }

        // 移动文件到新的文件夹
        srcPath := filepath.Join(fo.Path, file.Name())
        destPath := filepath.Join(newFolderPath, file.Name())
        if err := os.Rename(srcPath, destPath); err != nil {
            return err
        }
    }
    return nil
}

func main() {
    // 使用 Beego 框架的日志系统
    beego.Info("Starting folder organizer...")

    // 创建 FolderOrganizer 实例
    organizer := NewFolderOrganizer("./")

    // 执行文件夹整理
    if err := organizer.Organize(); err != nil {
        beego.Error("Error organizing folder: ", err)
    } else {
        beego.Info("Folder organized successfully.")
    }
}
