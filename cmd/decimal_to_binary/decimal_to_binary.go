package main

import (
	"fmt"
	"os"

	"github.com/armhold/infinity"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s [decimal]\n", os.Args)
		os.Exit(1)
	}
}

// Simple program to aid debugging decimal -> binary conversions.
// Note that on most Unix systems you can simply use: $ echo 'obase=2;256' | bc
func main() {
	decimal, err := infinity.DecimalToBinary(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing decimal: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", decimal)
}
