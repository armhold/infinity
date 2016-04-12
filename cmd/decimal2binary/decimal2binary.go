package main

import (
	"os"
	"fmt"
	"github.com/armhold/infinity"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s [decimal]\n", os.Args)
		os.Exit(1)
	}
}

func main() {
	decimal, err := infinity.DecimalToBinary(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing decimal: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", decimal)
}
