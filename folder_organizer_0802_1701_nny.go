// 代码生成时间: 2025-08-02 17:01:52
package main

import (
    "beego/logs"
    "flag"
    "fmt"
    "os"
    "path/filepath"
    "sort"
    "strings"
)

// FolderOrganizer defines the main structure for folder organizer
type FolderOrganizer struct {
    Path string
}

// NewFolderOrganizer creates a new FolderOrganizer instance
func NewFolderOrganizer(path string) *FolderOrganizer {
    return &FolderOrganizer{Path: path}
}

// OrganizeFolders is the main function that organizes the folders
func (f *FolderOrganizer) OrganizeFolders() error {
    files, err := os.ReadDir(f.Path)
    if err != nil {
        return err
    }

    for _, file := range files {
        if file.IsDir() {
            continue
        }

        filename := file.Name()
        fileExtension := strings.TrimPrefix(filepath.Ext(filename), ".")
        dirName := fmt.Sprintf("%s.%s", strings.TrimSuffix(filename, filepath.Ext(filename)), fileExtension)
        dirPath := filepath.Join(f.Path, dirName)

        if _, err := os.Stat(dirPath); os.IsNotExist(err) {
            if err := os.MkdirAll(dirPath, 0755); err != nil {
                return err
            }
        }

        if err := os.Rename(filepath.Join(f.Path, filename), filepath.Join(dirPath, filename)); err != nil {
            return err
        }
    }

    return nil
}

func main() {
    var path string
    flag.StringVar(&path, "path", ".", "The path to organize")
    flag.Parse()

    organizer := NewFolderOrganizer(path)
    if err := organizer.OrganizeFolders(); err != nil {
        logs.Error("Error organizing folders: %v", err)
    } else {
        fmt.Println("Folders organized successfully")
    }
}