// 代码生成时间: 2025-08-18 20:16:15
package main

import (
    "os"
    "log"
    "github.com/tealeg/xlsx"
    "github.com/astaxie/beego"
)

// ExcelGenerator 结构体，用于生成Excel文件
type ExcelGenerator struct {
    // 文件名
    FileName string
    // 工作簿
    Workbook *xlsx.Workbook
}

// NewExcelGenerator 创建一个ExcelGenerator实例
func NewExcelGenerator(fileName string) *ExcelGenerator {
    return &ExcelGenerator{
        FileName: fileName,
        Workbook: xlsx.NewFile(),
    }
}

// AddSheet 添加一个新的工作表
func (e *ExcelGenerator) AddSheet(title string) *xlsx.Worksheet {
    sheet, err := e.Workbook.AddSheet(title)
    if err != nil {
        beego.Error("Failed to add sheet: ", err)
        return nil
    }
    return sheet
}

// WriteData 向工作表写入数据
func (e *ExcelGenerator) WriteData(sheetIndex int, data [][]string) error {
    sheet := e.Workbook.Sheets[sheetIndex]
    for _, row := range data {
        for _, cell := range row {
            sheet.AddRow()
            sheet.Rows[len(sheet.Rows)-1].AddCell().Value = cell
        }
    }
    return nil
}

// Save 保存Excel文件
func (e *ExcelGenerator) Save() error {
    file, err := os.Create(e.FileName)
    if err != nil {
        beego.Error("Failed to create file: ", err)
        return err
    }
    defer file.Close()
    if err := e.Workbook.Write(file); err != nil {
        beego.Error("Failed to write to file: ", err)
        return err
    }
    return nil
}

func main() {
    // 创建ExcelGenerator实例
    generator := NewExcelGenerator("example.xlsx")
    defer generator.Save()

    // 添加工作表
    sheet := generator.AddSheet("Sheet1")
    if sheet == nil {
        log.Fatal("Failed to add sheet")
    }

    // 写入数据
    data := [][]string{{"Header1", "Header2"}, {"Data1", "Data2"}}
    if err := generator.WriteData(0, data); err != nil {
        log.Fatal("Failed to write data: ", err)
    }
}
