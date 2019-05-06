package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/thimalw/exn/exchange"
)

// usage message to display on arg errors
const usage = `Usage:
exn {value} {from} {to}

Example:
exn 10 USD AUD`

func main() {
	if len(os.Args) != 4 {
		fmt.Println(usage)
		os.Exit(1)
	}

	value, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println(usage)
		os.Exit(1)
	}
	from := strings.ToUpper(os.Args[2])
	to := strings.ToUpper(os.Args[3])

	converted, err := exchange.Convert(value, from, to)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	fmt.Printf("%.2f %s = %.2f %s\n", value, from, converted, to)
}
