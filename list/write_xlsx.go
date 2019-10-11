package list

import (
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func WriteXLSX(xlsxFilename string, list *List) error {
	xlsx := excelize.NewFile()

	// Sheet information
	xlsx.MergeCell("Sheet1", "A1", "O1")
	xlsx.SetCellValue("Sheet1", "A1", "Effective Date: "+list.EffectiveDate)

	// Column header
	xlsx.SetCellValue("Sheet1", "A2", "CS Code")
	xlsx.SetCellValue("Sheet1", "B2", "Status")
	xlsx.SetCellValue("Sheet1", "C2", "Product Name")
	xlsx.SetCellValue("Sheet1", "D2", "Category")
	xlsx.SetCellValue("Sheet1", "E2", "Category Description")
	xlsx.SetCellValue("Sheet1", "F2", "Size")
	xlsx.SetCellValue("Sheet1", "G2", "Case Pack")
	xlsx.SetCellValue("Sheet1", "H2", "Cost/Ounce")
	xlsx.SetCellValue("Sheet1", "I2", "Current Retail")
	xlsx.SetCellValue("Sheet1", "J2", "New Retail")
	xlsx.SetCellValue("Sheet1", "K2", "Price Change")
	xlsx.SetCellValue("Sheet1", "L2", "Comment")
	xlsx.SetCellValue("Sheet1", "M2", "Inventory")
	xlsx.SetCellValue("Sheet1", "N2", "Warehouse")
	xlsx.SetCellValue("Sheet1", "O2", "On Order")

	var ii string
	for i, p := range list.Products {
		ii = strconv.Itoa(i + 3)
		xlsx.SetCellValue("Sheet1", "A"+ii, p.CSCode)
		xlsx.SetCellValue("Sheet1", "B"+ii, p.StatusCode)
		xlsx.SetCellValue("Sheet1", "C"+ii, p.Name)
		xlsx.SetCellValue("Sheet1", "D"+ii, p.CategoryCode)
		xlsx.SetCellValue("Sheet1", "E"+ii, p.CategoryDescription)
		xlsx.SetCellValue("Sheet1", "F"+ii, p.Size)
		xlsx.SetCellValue("Sheet1", "G"+ii, p.CasePack)
		xlsx.SetCellValue("Sheet1", "H"+ii, p.CostPerOunce)
		xlsx.SetCellValue("Sheet1", "I"+ii, p.CurrentRetail)
		xlsx.SetCellValue("Sheet1", "J"+ii, p.NewRetail)
		xlsx.SetCellFormula("Sheet1", "K"+ii, "=IF(J"+ii+"=0,J"+ii+",J"+ii+"-I"+ii+")")
		xlsx.SetCellValue("Sheet1", "L"+ii, p.Comment)
		xlsx.SetCellValue("Sheet1", "M"+ii, p.StoresInventory)
		xlsx.SetCellValue("Sheet1", "N"+ii, p.WarehouseInventory)
		xlsx.SetCellValue("Sheet1", "O"+ii, p.WarehouseOnOrder)
	}

	// Style formatting the cells
	style, _ := xlsx.NewStyle(`{"number_format":1}`)
	xlsx.SetCellStyle("Sheet1", "F3", "G"+ii, style)
	style, _ = xlsx.NewStyle(`{"number_format":2}`)
	xlsx.SetCellStyle("Sheet1", "H3", "K"+ii, style)

	// Set column widths
	xlsx.SetColWidth("Sheet1", "A", "A", 9)
	xlsx.SetColWidth("Sheet1", "B", "B", 7.5)
	xlsx.SetColWidth("Sheet1", "C", "C", 40)
	xlsx.SetColWidth("Sheet1", "D", "D", 9.67)
	xlsx.SetColWidth("Sheet1", "E", "E", 27.83)
	xlsx.SetColWidth("Sheet1", "F", "F", 5.83)
	xlsx.SetColWidth("Sheet1", "G", "G", 10.17)
	xlsx.SetColWidth("Sheet1", "H", "H", 12)
	xlsx.SetColWidth("Sheet1", "I", "I", 13.67)
	xlsx.SetColWidth("Sheet1", "J", "J", 11.17)
	xlsx.SetColWidth("Sheet1", "K", "K", 12.5)
	xlsx.SetColWidth("Sheet1", "L", "L", 10.5)
	xlsx.SetColWidth("Sheet1", "M", "M", 10.17)
	xlsx.SetColWidth("Sheet1", "N", "N", 11.33)
	xlsx.SetColWidth("Sheet1", "O", "O", 9.67)

	// Set up auto filter on the column header
	xlsx.AutoFilter("Sheet1", "A2", "O2", "")
	// Freeze the header to allow it to stay on the top
	xlsx.SetPanes("Sheet1", `{"freeze":true,"split":false,"x_split":0,"y_split":2,"top_left_cell":"A3","active_pane":"bottomLeft","panes":[{"sqref":"A3","active_cell":"A3","pane":"bottomLeft"}]}`)

	return xlsx.SaveAs(xlsxFilename)
}
