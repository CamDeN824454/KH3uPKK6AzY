// 代码生成时间: 2025-09-09 04:38:31
 * This program provides functions to compress a directory into a ZIP file and decompress a ZIP file into a directory.
 * It follows GoLang best practices for code maintainability and extensibility.
 */

package main

import (
    "archive/zip"
    "fmt"
    "io"
    "io/ioutil
    "log"
    "os"
    "path/filepath"
    "strings"
)

// Compress compresses the specified directory into a ZIP file.
func Compress(srcDir, destZip string) error {
    // Create destination file and a zip writer.
    destFile, err := os.Create(destZip)
    if err != nil {
        return err
    }
    defer destFile.Close()

    zipWriter := zip.NewWriter(destFile)
    defer zipWriter.Close()

    // Walk the source directory.
    err = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.Mode().IsRegular() {
            return nil // Exclude directories.
        }

        relPath, err := filepath.Rel(srcDir, path)
        if err != nil {
            return err
        }

        zipFile, err := zipWriter.Create(relPath)
        if err != nil {
            return err
        }

        file, err := os.Open(path)
        if err != nil {
            return err
        }
        defer file.Close()

        _, err = io.Copy(zipFile, file)
        return err
    })
    if err != nil {
        return err
    }
    return nil
}

// Decompress decompresses the specified ZIP file into a directory.
func Decompress(srcZip, destDir string) error {
    // Open the ZIP file.
    src, err := zip.OpenReader(srcZip)
    if err != nil {
        return err
    }
    defer src.Close()

    // Iterate through the ZIP file entries.
    for _, f := range src.File {
        rc, err := f.Open()
        if err != nil {
            return err
        }
        defer rc.Close()

        // Create the directory structure.
        fpath := filepath.Join(destDir, f.Name)
        if f.FileInfo().IsDir() {
            os.MkdirAll(fpath, os.ModePerm)
        } else {
            // Create the file.
            if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
                return err
            }
            file, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
            if err != nil {
                return err
            }
            defer file.Close()

            _, err = io.Copy(file, rc)
            if err != nil {
                return err
            }
        }
    }
    return nil
}

func main() {
    // Example usage of the Compress and Decompress functions.
    srcDir := "path/to/source/directory"
    destZip := "path/to/destination/zipfile.zip"
    err := Compress(srcDir, destZip)
    if err != nil {
        log.Fatalf("Error compressing: %s", err)
    }

    srcZip := "path/to/source/zipfile.zip"
    destDir := "path/to/destination/directory"
    err = Decompress(srcZip, destDir)
    if err != nil {
        log.Fatalf("Error decompressing: %s", err)
    }
    fmt.Println("Compressed and decompressed successfully.")
}