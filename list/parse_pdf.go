package list

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"

	pdfcontent "github.com/unidoc/unidoc/pdf/contentstream"
	pdf "github.com/unidoc/unidoc/pdf/model"

	"github.com/jaredhoward/dbac/product"
)

var reListEffectiveDate = regexp.MustCompile(`(?is)NEW PRICES EFFECTIVE: (.+?)\n`)
var reProductsStart = regexp.MustCompile(`(?s)\A.*CS.*?Code.*?Size.*?Case.*?Pack.*?Product Name.*?Status.*?New.*?Retail.*Comment.*?Category / DescriptionCurrent.*?RetailCost/.*?Ounce\n?`)
var reProductsEnd = regexp.MustCompile(`(?s)\n?\*  Indicates Available by Split`)
var reProductLine = regexp.MustCompile(`(?s)` +
	`(?P<csCode>\d{6})` +
	`\n` +
	`(?P<size>\d+)` +
	`\n` +
	`(?P<productName>[^\n]+)` +
	`\n` +
	`(?P<statusCode>[1DSLXNAUTR]+)` +
	`[\n ]*?` +
	`(?P<costPerOunce>\d+\.\d+)` +
	`(?:` +
	`[\n ]*?` +
	`(?P<newRetail>\d+\.\d+)` +
	`)?` +
	`[\n ]*?` +
	`(?P<casePack>\d+)` +
	`[\n ]*?` +
	`(?P<currentRetail>\d+\.\d+)` +
	`(?:` +
	`\n` +
	`(?P<comment>[^\n]+)` +
	`)?` +
	`\n` +
	`(?P<categoryCode>[A-Z]{3})` +
	`[\n\t]+` +
	`(?P<categoryDescription>[^\n]+)`)

func GetListFromPDFFile(inputPath string) (*List, error) {
	f, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return parseAlphaPriceListPDF(f)
}

func parseAlphaPriceListPDF(rs io.ReadSeeker) (*List, error) {
	pdfReader, err := pdf.NewPdfReader(rs)
	if err != nil {
		return nil, err
	}

	isEncrypted, err := pdfReader.IsEncrypted()
	if err != nil {
		return nil, err
	}
	if isEncrypted {
		_, err = pdfReader.Decrypt([]byte(""))
		if err != nil {
			return nil, err
		}
	}

	list := &List{}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return nil, err
	}
	for i := 0; i < numPages; i++ {
		pageNum := i + 1

		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			return list, err
		}

		contentStreams, err := page.GetContentStreams()
		if err != nil {
			return list, err
		}

		// If the value is an array, the effect shall be as if all of the streams in the array were concatenated,
		// in order, to form a single stream.
		pageContentStr := ""
		for _, cstream := range contentStreams {
			pageContentStr += cstream
		}

		cstreamParser := pdfcontent.NewContentStreamParser(pageContentStr)
		txt, err := cstreamParser.ExtractText()
		if err != nil {
			return list, err
		}

		parseAlphaPriceListInfo(list, txt)

		txt, err = reduceAlphaPriceListProductPage(list, txt)
		if err != nil {
			// Some pages, mainly the last page, will error here.
			// This is expected. This is not a page that will be
			// parsed for product information.
			continue
		}

		err = parseAlphaPriceListProducts(list, txt)
		if err != nil {
			fmt.Printf("Error on page %d\n%s\n", pageNum, err)
			return list, err
		}
	}

	return list, nil
}

func parseAlphaPriceListInfo(list *List, txt string) error {
	if list.EffectiveDate == "" {
		effective := reListEffectiveDate.FindStringSubmatch(txt)
		list.EffectiveDate = effective[1]
	}
	return nil
}

func reduceAlphaPriceListProductPage(list *List, txt string) (string, error) {
	ints := reProductsStart.FindStringIndex(txt)
	if len(ints) != 2 {
		return txt, fmt.Errorf("Failed parsing beginning of page")
	}
	txt = txt[ints[1]:]

	ints = reProductsEnd.FindStringIndex(txt)
	if len(ints) != 2 {
		return txt, fmt.Errorf("Failed parsing end of page")
	}
	txt = txt[:ints[0]]

	return txt, nil
}

func parseAlphaPriceListProducts(list *List, txt string) error {
	var err error

	for _, line := range reProductLine.FindAllStringSubmatch(txt, -1) {
		result := make(map[string]string)
		for i, name := range reProductLine.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = line[i]
			}
		}

		p := &product.Product{
			CSCode:              result["csCode"],
			Name:                result["productName"],
			StatusCode:          result["statusCode"],
			CategoryCode:        result["categoryCode"],
			CategoryDescription: result["categoryDescription"],
			Comment:             result["comment"],
		}

		p.Size, err = strconv.Atoi(result["size"])
		if err != nil {
			return err
		}
		p.CasePack, err = strconv.Atoi(result["casePack"])
		if err != nil {
			return err
		}
		p.CostPerOunce, err = strconv.ParseFloat(result["costPerOunce"], 64)
		if err != nil {
			return err
		}
		if result["newRetail"] != "" {
			p.NewRetail, err = strconv.ParseFloat(result["newRetail"], 64)
			if err != nil {
				return err
			}
		}
		p.CurrentRetail, err = strconv.ParseFloat(result["currentRetail"], 64)
		if err != nil {
			return err
		}

		list.AddProduct(p)
	}
	return nil
}
