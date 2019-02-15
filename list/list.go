package list

import (
	"fmt"
	"os"
	"strings"

	"github.com/jaredhoward/dbac/product"
)

const (
	DefaultInputAlphaPriceList  = "https://abc.utah.gov/products/documents/AlphaPriceList.pdf"
	DefaultOutputAlphaPriceList = "./AlphaPriceList.xlsx"
)

type List struct {
	EffectiveDate string
	Products      []product.Product
}

func (l *List) AddProduct(p product.Product) {
	l.Products = append(l.Products, p)
}

func GetList(input string) (*List, error) {
	if strings.ToLower(input[0:7]) == "http://" || strings.ToLower(input[0:8]) == "https://" {
		body, err := GetAlphaPriceList(input)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		return GetListFromBytes(body)
	} else {
		return GetListFromPDFFile(input)
	}
}
