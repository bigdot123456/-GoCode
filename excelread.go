package main

import (
	"fmt"
	"os"
	"github.com/Luxurioust/excelize"
)

func main() {
	//xlsx := excelize.CreateFile()
	xlsx := excelize.NewFile()
	// Create a new sheet.
	xlsx.NewSheet("Sheet2")
	// Set value of a cell.
	xlsx.SetCellValue("Sheet2", "A2", "Hello world.")
	xlsx.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	xlsx.SetActiveSheet(2)
	// Save xlsx file by the given path.
	err:= xlsx.SaveAs("./Book3.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Print("ok!")
	}

}
