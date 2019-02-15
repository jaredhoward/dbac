package list

import (
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func WriteXLSX(xlsxFilename string, list *List) error {
	if xlsxFilename == "" {
		xlsxFilename = DefaultOutputAlphaPriceList
	}

	xlsx := excelize.NewFile()

	xlsx.MergeCell("Sheet1", "A1", "L1")
	xlsx.SetCellValue("Sheet1", "A1", "Effective Date: "+list.EffectiveDate)

	xlsx.SetCellValue("Sheet1", "A2", "CS Code")
	xlsx.SetCellValue("Sheet1", "B2", "Status")
	xlsx.SetCellValue("Sheet1", "C2", "Product Name")
	xlsx.SetCellValue("Sheet1", "D2", "Category")
	xlsx.SetCellValue("Sheet1", "E2", "Description")
	xlsx.SetCellValue("Sheet1", "F2", "Size")
	xlsx.SetCellValue("Sheet1", "G2", "Case Pack")
	xlsx.SetCellValue("Sheet1", "H2", "Cost/Ounce")
	xlsx.SetCellValue("Sheet1", "I2", "Current Retail")
	xlsx.SetCellValue("Sheet1", "J2", "New Retail")
	xlsx.SetCellValue("Sheet1", "K2", "Price Change")
	xlsx.SetCellValue("Sheet1", "L2", "Comment")

	var ii string
	for i, p := range list.Products {
		ii = strconv.Itoa(i + 3)
		xlsx.SetCellValue("Sheet1", "A"+ii, p.CSCode)
		xlsx.SetCellValue("Sheet1", "B"+ii, p.Status)
		xlsx.SetCellValue("Sheet1", "C"+ii, p.Name)
		xlsx.SetCellValue("Sheet1", "D"+ii, p.Category)
		xlsx.SetCellValue("Sheet1", "E"+ii, p.Description)
		xlsx.SetCellValue("Sheet1", "F"+ii, p.Size)
		xlsx.SetCellValue("Sheet1", "G"+ii, p.CasePack)
		xlsx.SetCellValue("Sheet1", "H"+ii, p.CostPerOunce)
		xlsx.SetCellValue("Sheet1", "I"+ii, p.CurrentRetail)
		xlsx.SetCellValue("Sheet1", "J"+ii, p.NewRetail)
		xlsx.SetCellFormula("Sheet1", "K"+ii, "=IF(J"+ii+"=0,J"+ii+",J"+ii+"-I"+ii+")")
		xlsx.SetCellValue("Sheet1", "L"+ii, p.Comment)
	}

	style, _ := xlsx.NewStyle(`{"number_format":1}`)
	xlsx.SetCellStyle("Sheet1", "F3", "G"+ii, style)
	style, _ = xlsx.NewStyle(`{"number_format":2}`)
	xlsx.SetCellStyle("Sheet1", "H3", "K"+ii, style)

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
	xlsx.SetColWidth("Sheet1", "K", "L", 10.5)

	xlsx.AutoFilter("Sheet1", "A2", "L2", "")
	xlsx.SetPanes("Sheet1", `{"freeze":true,"split":false,"x_split":0,"y_split":2,"top_left_cell":"A3","active_pane":"bottomLeft","panes":[{"sqref":"A3","active_cell":"A3","pane":"bottomLeft"}]}`)

	return xlsx.SaveAs(xlsxFilename)
}