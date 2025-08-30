// 代码生成时间: 2025-08-30 21:38:45
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "sort"
)

// FolderOrganizer 结构体，用于存储文件夹信息
type FolderOrganizer struct {
    RootPath string
}

// NewFolderOrganizer 创建 FolderOrganizer 实例
func NewFolderOrganizer(rootPath string) *FolderOrganizer {
    return &FolderOrganizer{
        RootPath: rootPath,
    }
}

// Organize 整理文件夹结构
func (f *FolderOrganizer) Organize() error {
    // 检查根路径是否存在
    if _, err := os.Stat(f.RootPath); os.IsNotExist(err) {
        return fmt.Errorf("root path does not exist: %s", f.RootPath)
    }

    // 获取根路径下的所有文件和文件夹
    files, err := os.ReadDir(f.RootPath)
    if err != nil {
        return fmt.Errorf("failed to read root path: %s", err)
    }

    // 对文件和文件夹进行排序
    sort.Slice(files, func(i, j int) bool {
        return strings.ToLower(files[i].Name()) < strings.ToLower(files[j].Name())
    })

    // 遍历文件和文件夹
    for _, file := range files {
        // 创建新的路径
        newPath := filepath.Join(f.RootPath, file.Name())

        // 如果是文件夹，则递归整理
        if file.IsDir() {
            if err := f.OrganizeFolder(newPath); err != nil {
                return err
            }
        } else {
            // 如果是文件，则输出文件信息
            fmt.Printf("File: %s
", newPath)
        }
    }
    return nil
}

// OrganizeFolder 递归整理文件夹
func (f *FolderOrganizer) OrganizeFolder(folderPath string) error {
    // 获取文件夹下的所有文件和子文件夹
    files, err := os.ReadDir(folderPath)
    if err != nil {
        return fmt.Errorf("failed to read folder: %s", err)
    }

    // 对文件和子文件夹进行排序
    sort.Slice(files, func(i, j int) bool {
        return strings.ToLower(files[i].Name()) < strings.ToLower(files[j].Name())
    })

    // 遍历文件和子文件夹
    for _, file := range files {
        // 创建新的路径
        newPath := filepath.Join(folderPath, file.Name())

        // 如果是子文件夹，则递归整理
        if file.IsDir() {
            if err := f.OrganizeFolder(newPath); err != nil {
                return err
            }
        } else {
            // 如果是文件，则输出文件信息
            fmt.Printf("File: %s
", newPath)
        }
    }
    return nil
}

func main() {
    // 创建 FolderOrganizer 实例
    organizer := NewFolderOrganizer("./")

    // 整理文件夹结构
    if err := organizer.Organize(); err != nil {
        fmt.Printf("Error organizing folder: %s
", err)
    } else {
        fmt.Println("Folder organized successfully")
    }
}