package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func main() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// 创建一个工作表
	sheetName := "Sheet1"
	index, _ := f.NewSheet(sheetName)

	//// 隐藏名称为 Sheet1 的工作表中的 D 至 F 列：
	//f.SetColVisible("Sheet1", "D:F", false)
	//
	//// 根据给定的工作表名称（大小写敏感）、列范围和宽度值设置单个或多个列的宽度。
	//// 例如设置名为 Sheet1 工作表上 A 到 D 列的宽度为 20：
	//f.SetColWidth("Sheet1", "A", "D", 20)
	//
	//// 根据给定的工作表名称（大小写敏感）、行号和高度值设置单行高度。
	//// 例如设置名为 Sheet1 工作表第二行行的高度为 50：
	//f.SetRowHeight("Sheet1", 2, 50)
	//
	//// 设置单元格的值
	//f.SetCellValue("Sheet1", "A2", "100")

	//// 根据给定的工作表名（大小写敏感）和单元格坐标区域合并单元格。例如，合并名为 Sheet1 的工作表上 D3:E9 区域内的单元格：
	//f.MergeCell("Sheet1", "D3", "D4")
	//f.SetCellValue("Sheet1", "D3", 1000)
	//f.SetCellValue("Sheet1", "D5", "hello")

	f.SetCellValue(sheetName, "A1", "Name")
	f.SetCellValue(sheetName, "B1", "Age")
	f.SetCellValue(sheetName, "A2", "张三")
	f.SetCellValue(sheetName, "B2", 30)
	f.SetCellValue(sheetName, "A3", "李四")
	f.SetCellValue(sheetName, "B3", 32)

	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)

	// 根据指定路径保存文件
	if err := f.SaveAs("/Users/kankan/go/src/go-test/excelize/Book1.xlsx"); err != nil {
		println(err.Error())
	}
}
