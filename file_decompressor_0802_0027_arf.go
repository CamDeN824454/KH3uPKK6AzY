// 代码生成时间: 2025-08-02 00:27:45
package main

import (
    "compress/gzip"
    "archive/zip"
    "io"
    "io/ioutil"
    "os"
    "path/filepath"
    "fmt"
    "log"
    "os/exec"
)

// DecompressGzip decompresses a gzip file to the specified destination directory
func DecompressGzip(src, dest string) error {
    r, err := gzip.Open(src)
    if err != nil {
        return err
    }
    defer r.Close()

    out, err := os.Create(dest)
    if err != nil {
        return err
    }
    defer out.Close()

    _, err = io.Copy(out, r)
    return err
}

// DecompressZip decompresses a zip file to the specified destination directory
func DecompressZip(src, dest string) error {
    r, err := zip.OpenReader(src)
    if err != nil {
        return err
    }
    defer r.Close()

    for _, f := range r.File {
        fpath := filepath.Join(dest, f.Name)
        if f.FileInfo().IsDir() {
            os.MkdirAll(fpath, os.ModePerm)
        } else {
            if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
                return err
            }
            in, err := f.Open()
            if err != nil {
                return err
            }
            defer in.Close()
            out, err := os.Create(fpath)
            if err != nil {
                return err
            }
            defer out.Close()
            if _, err := io.Copy(out, in); err != nil {
                return err
            }
        }
    }
    return nil
}

// DecompressTarGz decompresses a tar.gz file to the specified destination directory
func DecompressTarGz(src, dest string) error {
    cmd := exec.Command("tar", "-zxf", src, "-C", dest)
    return cmd.Run()
}

// Decompress any supported file type
func DecompressFile(filePath, destination string) error {
    switch filepath.Ext(filePath) {
    case ".gz":
        return DecompressGzip(filePath, destination)
    case ".zip":
        return DecompressZip(filePath, destination)
    case ".tar.gz":
        return DecompressTarGz(filePath, destination)
    default:
        return fmt.Errorf("unsupported file type: %s", filepath.Ext(filePath))
    }
}

func main() {
    // Example usage of the DecompressFile function
    src := "example.tar.gz"
    dest := "destination_directory"
    err := DecompressFile(src, dest)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Decompression successful. Files are in: %s
", dest)
}