package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/thimalw/exn/exchange"
	"github.com/thimalw/number"
)

func main() {
	// Display usage message on help flag.
	if len(os.Args) > 1 && os.Args[1] == "-help" {
		showUsage()
		os.Exit(0)
	}

	if len(os.Args) != 4 {
		showUsage()
		os.Exit(1)
	}

	value, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		showUsage()
		os.Exit(1)
	}
	from := strings.ToUpper(os.Args[2])
	to := strings.ToUpper(os.Args[3])

	converted, err := exchange.Convert(value, from, to)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	formattedValue, err := number.CommaFormat(fmt.Sprintf("%.2f", value))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	formattedConverted, err := number.CommaFormat(fmt.Sprintf("%.2f", converted))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s %s = %s %s\n", formattedValue, from, formattedConverted, to)
}

// showUsage prints the usage message.
func showUsage() {
	fmt.Println(`Usage:
exn {value} {from} {to}

Example:
exn 10 USD AUD`)
}
