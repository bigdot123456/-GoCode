package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

func main() {
	xlsx, err := excelize.OpenFile("./plan.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	cell := xlsx.GetCellValue("Sheet1", "B2")
	a := xlsx.SheetCount
	fmt.Print(a)
	fmt.Println(cell)
	for a > 0 {
		fmt.Print("Sheet" + strconv.Itoa(a) + ":\n")
		fmt.Print(xlsx.GetSheetName(a) + "\n")
		a--
	}

	for index, name := range xlsx.GetSheetMap() {
		fmt.Println("**********")
		fmt.Println(index, name)

		// Get all the rows in the Sheet1.
		rows := xlsx.GetRows(name)
		for _, row := range rows {
			for _, colCell := range row {
				fmt.Print(colCell, "\t")
			}
			fmt.Println()
		}
	}
}
