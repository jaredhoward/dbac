package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/jaredhoward/dbac/list"
)

var (
	input   = flag.String("input", list.DefaultInputAlphaPriceList, "input filename or URL of AlphaPriceList.pdf")
	output  = flag.String("output", list.DefaultOutputAlphaPriceList, "output filename of the converted Excel file")
	version = flag.Bool("version", false, "display the version information")

	GitCommit string
	BuildDate string
)

func main() {
	flag.Parse()

	if *version {
		fmt.Fprintf(os.Stderr, "Git commit: %s\nBuild date: %s\nGo version: %s\n", GitCommit, BuildDate, runtime.Version())
		os.Exit(0)
	}

	l, err := list.GetList(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	err = list.WriteXLSX(*output, l)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
