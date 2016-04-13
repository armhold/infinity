package main

import (
	"fmt"
	"os"

	"github.com/armhold/infinity"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s [binary]\n", os.Args)
		os.Exit(1)
	}
}

// Simple program to aid debugging binary -> decimal conversions.
// Note on most Unix systems you can simply use: $ echo 'ibase=2;100000000' | bc
func main() {
	decimal, err := infinity.BinaryToDecimal(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing binary string %s: \n", err)
		os.Exit(1)
	}

	fmt.Printf("%d\n", decimal)
}
