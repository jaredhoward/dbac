# Alpha Price List

Alpha Price List application works with the Utah Department of Alcohol Beverage
Control's price list named AlphaPriceList.pdf. The DABC provides a PDF version
of the available products. However, a PDF does not allow filtering or sorting
the products. This application can parse the PDF and use that data as needed;
initially by providing an Excel document as the output.

This app provides tools to:

* retrieve the current or previous AlphaPriceList.pdf that is provided on the
  website
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
  -previous
    	retrieve the previous month's price list
  -version
    	display the version information
```

### Mac Usage

The Mac version of this application if intended to be ran from a command
prompt. This ensures that the user is aware of what the application is doing.

Note: The application can be ran from a Finder window by double clicking;
this is fine, however, not ideal. When ran this way, a new terminal window
is opened, the application runs with the default flags and the outputted file
is typically placed in the user's root folder.

### Windows Usage

The Windows executable would typically ran by double clicking the application.
If it is placed on the user's desktop, the outputted file would also be placed
on the desktop.

The application can also be called from a command prompt. This would allow the
user to call the command flags.

