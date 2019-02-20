# Alpha Price List

Alpha Price List application works with the Utah Department of Alcohol Beverage
Control's price list named AlphaPriceList.pdf. The DABC provides a PDF version
of the available products. However, a PDF does not allow filtering or sorting
the products. This application can parse the PDF and use that data as needed;
initially by providing an Excel document as the output.

This app provides tools to:

* retrieve the current AlphaPriceList.pdf that is provided on the website
* read a saved price list
* convert the PDF price list to an Excel XLSX version that includes:
  * automatic filters
  * cost difference

## Usage

```
./dbacAlphaPriceList -h

Usage of ./dbacAlphaPriceList:
  -input string
        input filename or URL of AlphaPriceList.pdf (default "https://abc.utah.gov/products/documents/AlphaPriceList.pdf")
  -output string
        output filename of the converted Excel file (default "./AlphaPriceList.xlsx")
  -version
        display the version information
```
