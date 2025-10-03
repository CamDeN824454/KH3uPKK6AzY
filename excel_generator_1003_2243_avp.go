// 代码生成时间: 2025-10-03 22:43:50
package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "path/filepath"
    "strconv"

    "github.com/tealeg/xlsx"
)

// ExcelGenerator 结构体封装Excel生成器
type ExcelGenerator struct {
    File     *xlsx.File
    Sheet    *xlsx.Sheet
    CurRow   int
    CurCol   int
}

// NewExcelGenerator 创建一个新的Excel生成器
func NewExcelGenerator() *ExcelGenerator {
    return &ExcelGenerator{
        File: xlsx.NewFile(),
        Sheet: xlsx.NewSheet("Sheet1"),
        CurRow: 0,
        CurCol: 0,
    }
}

// AddRow 添加一行数据到Excel表格
func (e *ExcelGenerator) AddRow(data []string) error {
    e.Sheet = e.File.AddSheet("Sheet1")
    if e.CurRow >= 1000 { // Excel限制行数为1048576
        return fmt.Errorf("row limit exceeded")
    }
    for i, value := range data {
        e.Sheet.AddRow().SetCellStr(e.CurCol, e.CurRow, value)
        e.CurCol++
        if e.CurCol > 16384 { // Excel限制列数为16384
            return fmt.Errorf("column limit exceeded")
        }
    }
    e.CurCol = 0 // 重置列索引
    e.CurRow++ // 行索引加1
    return nil
}

// SaveExcel 保存Excel文件
func (e *ExcelGenerator) SaveExcel(filename string) error {
    if err := e.File.Save(filename); err != nil {
        return fmt.Errorf("failed to save file: %w", err)
    }
    return nil
}

func main() {
    // 创建Excel生成器
    generator := NewExcelGenerator()
    // 添加数据
    if err := generator.AddRow([]string{"Name", "Age", "City"}); err != nil {
        fmt.Printf("Error adding row: %s
", err)
        return
    }
    if err := generator.AddRow([]string{"Alice", "30", "New York"}); err != nil {
        fmt.Printf("Error adding row: %s
", err)
        return
    }
    // 保存Excel文件
    filename := "example.xlsx"
    if err := generator.SaveExcel(filename); err != nil {
        fmt.Printf("Error saving Excel file: %s
", err)
        return
    }
    fmt.Printf("Excel file saved as: %%s
", filename)
}
