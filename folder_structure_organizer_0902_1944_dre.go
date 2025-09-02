// 代码生成时间: 2025-09-02 19:44:51
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "time"

    "github.com/astaxie/beego"
)

// FolderItem represents an item in the folder structure
type FolderItem struct {
    Name     string    `json:"name"`
    Path     string    `json:"path"`
    IsFolder bool      `json:"is_folder"`
    Modified time.Time `json:"modified"`
    Children []FolderItem `json:"children,omitempty"`
}

// OrganizeFolderStructure organizes the folder structure starting from the given rootPath
func OrganizeFolderStructure(rootPath string) ([]FolderItem, error) {
    var items []FolderItem
    filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        relPath, _ := filepath.Rel(rootPath, path)
        fi := FolderItem{
            Name:     strings.TrimPrefix(relPath, string(filepath.Separator)),
            Path:     path,
            IsFolder: info.IsDir(),
            Modified: info.ModTime(),
        }

        if fi.IsFolder {
            fi.Children, _ = OrganizeFolderStructure(path)
        }

        items = append(items, fi)
        return nil
    })

    return items, nil
}

func main() {
    beego.BeeLogger.SetLevel(beego.LevelDebug)
    beego.BeeLogger.EnableFuncCallDepth(true)
    beego.BeeLogger.SetLogFuncCallDepth(2)

    // Example usage: Organize the structure of the folder at the given path
    rootPath := "/path/to/your/folder"
    organizedItems, err := OrganizeFolderStructure(rootPath)
    if err != nil {
        beego.BeeLogger.Error("Error organizing folder structure: %s", err)
        return
    }

    // Print the organized folder structure (for demonstration purposes)
    for _, item := range organizedItems {
        fmt.Printf("Folder: %s
", item.Path)
        if len(item.Children) > 0 {
            for _, child := range item.Children {
                fmt.Printf("  Child: %s
", child.Path)
            }
        }
    }
}
