package list

import (
	"os"
	"os/user"
	"path"
	"strings"

	"github.com/jaredhoward/dbac/product"
)

const (
	// Default input constants
	DefaultInputAlphaPriceList         = "https://abc.utah.gov/products/documents/AlphaPriceList.pdf"
	DefaultInputAlphaPriceListPrevious = "https://abc.utah.gov/products/documents/previousMonthAlphaPriceList.pdf"

	// Default output constants
	DefaultOutputFolder                         = "./"
	DefaultOutputAlphaPriceListFilename         = "AlphaPriceList.xlsx"
	DefaultOutputAlphaPriceListPreviousFilename = "previousMonthAlphaPriceList.xlsx"
	DefaultOutputAlphaPriceList                 = DefaultOutputFolder + DefaultOutputAlphaPriceListFilename
	DefaultOutputAlphaPriceListPrevious         = DefaultOutputFolder + DefaultOutputAlphaPriceListPreviousFilename
)

type List struct {
	EffectiveDate string
	Products      []*product.Product
}

func (l *List) AddProduct(p *product.Product) {
	l.Products = append(l.Products, p)
}

func GetList(input string) (*List, error) {
	if strings.ToLower(input[0:7]) == "http://" || strings.ToLower(input[0:8]) == "https://" {
		body, err := GetAlphaPriceList(input)
		if err != nil {
			return nil, err
		}
		return GetListFromBytes(body)
	} else {
		return GetListFromPDFFile(input)
	}
}

func CorrectOutputFilename(filename string, previous bool) string {
	if len(filename) == 0 {
		if previous {
			filename = DefaultOutputAlphaPriceListPrevious
		} else {
			filename = DefaultOutputAlphaPriceList
		}
	} else if previous && filename == DefaultOutputAlphaPriceList {
		// This is using the default setting for output but wanting the
		// previous naming.
		filename = DefaultOutputAlphaPriceListPrevious
	}

	// If the OS shell did not expand the `~` from the command argument.
	if filename[0] == '~' {
		usr, _ := user.Current()
		filename = usr.HomeDir + filename[1:]
	}

	var isDir bool
	fi, err := os.Lstat(filename)
	// Expecting an error here if a new file will be created.
	if err != nil {
		if os.IsPathSeparator(filename[len(filename)-1]) {
			isDir = true
		}
	} else {
		if fi.Mode().IsDir() {
			isDir = true
		}
	}
	// If the given filename is a directory, a default filename will need to be
	// appended onto it.
	if isDir {
		if previous {
			filename = path.Join(filename, DefaultOutputAlphaPriceListPreviousFilename)
		} else {
			filename = path.Join(filename, DefaultOutputAlphaPriceListFilename)
		}
	}

	return filename
}
