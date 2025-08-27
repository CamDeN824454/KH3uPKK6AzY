// 代码生成时间: 2025-08-27 13:34:55
package main

import (
    "archive/zip"
    "bufio"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"

    "github.com/astaxie/beego" // Import Beego framework
)

// DecompressFile解压指定的zip文件到目标目录
func DecompressFile(src, dest string) error {
    // Open the zip archive for reading.
    reader, err := zip.OpenReader(src)
    if err != nil {
        return fmt.Errorf("failed to open zip file: %w", err)
    }
    defer reader.Close()

    // Iterate through the files in the zip archive.
    for _, file := range reader.File {
        // Create the full path to the file in the destination directory.
        fpath := filepath.Join(dest, file.Name)
        if strings.HasPrefix(file.Name, "/") {
            fpath = filepath.Join(dest, file.Name[1:])
        }

        // Check if the file is a directory.
        if file.FileInfo().IsDir() {
            // Create the directory if it doesn't exist.
            if err := os.MkdirAll(fpath, os.ModePerm); err != nil {
                return fmt.Errorf("failed to create directory: %w", err)
            }
            continue
        }

        // Create the directory structure.
        if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
            return fmt.Errorf("failed to create directory: %w", err)
        }

        // Open the file in the zip archive.
        fileReader, err := file.Open()
        if err != nil {
            return fmt.Errorf("failed to open file in zip: %w", err)
        }
        defer fileReader.Close()

        // Create the new file.
        newFile, err := os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, file.Mode())
        if err != nil {
            return fmt.Errorf("failed to create file: %w", err)
        }
        defer newFile.Close()

        // Copy the contents of the file in the zip archive to the new file.
        _, err = io.Copy(newFile, fileReader)
        if err != nil {
            return fmt.Errorf("failed to copy file contents: %w", err)
        }
    }
    return nil
}

func main() {
    beego.BeeLogger.SetLevel(beego.LevelDebug) // Set log level
    src := "example.zip" // Source zip file path
    dest := "destination" // Destination directory path

    // Call the DecompressFile function and handle potential errors.
    if err := DecompressFile(src, dest); err != nil {
        log.Fatalf("Error decompressing file: %s", err)
    } else {
        fmt.Println("File decompressed successfully.")
    }
}
