// 代码生成时间: 2025-09-21 14:04:33
package main

import (
    "fmt"
    "os"
    "path/filepath"

    "github.com/360EntSecGroup-Skylar/excelize/v2"
)

// ExcelAutoGenerator 结构体定义了Excel自动生成器的基本属性
type ExcelAutoGenerator struct {
    // FilePath 是Excel文件的存储路径
    FilePath string
    // SheetName 是Excel工作表的名称
    SheetName string
}

// NewExcelAutoGenerator 创建一个新的ExcelAutoGenerator实例
func NewExcelAutoGenerator(filePath, sheetName string) *ExcelAutoGenerator {
    return &ExcelAutoGenerator{
        FilePath: filePath,
        SheetName: sheetName,
    }
}

// Generate 方法用于生成Excel文件
func (e *ExcelAutoGenerator) Generate(data [][]string) error {
    // 创建一个新的Excel文件
    f := excelize.NewFile()
    // 设置工作表的名称
    index := f.NewSheet(e.SheetName)
    // 如果创建失败，返回错误
    if index == 0 {
        return fmt.Errorf("failed to create sheet: %s", e.SheetName)
    }

    // 设置工作表的标题行
    for i, row := range data {
        for j, cell := range row {
            // 将数据写入对应的单元格
            f.SetCellValue(e.SheetName, excelize.CoordinatesFromCell(i, j+1), cell)
        }
    }

    // 保存Excel文件
    err := f.SaveAs(e.FilePath)
    if err != nil {
        return fmt.Errorf("failed to save excel file: %s", err)
    }
    return nil
}

// main 函数是程序的入口点
func main() {
    basePath := filepath.Join("./", "test.xlsx")
    sheetName := "Sheet1"
    data := [][]string{{"Name", "Age"}, {"Alice", "30"}, {"Bob", "25"}}

    // 创建Excel自动生成器实例
    generator := NewExcelAutoGenerator(basePath, sheetName)

    // 调用Generate方法生成Excel文件
    if err := generator.Generate(data); err != nil {
        fmt.Println("Error: ", err)
    } else {
        fmt.Println("Excel file generated successfully!")
    }
}
