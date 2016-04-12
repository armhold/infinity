package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/armhold/infinity"
)

var (
	width int
	file string
)


func init() {
	var usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-w width] file\n", os.Args)
		flag.PrintDefaults()
	}

	flag.Usage = usage

	flag.IntVar(&width, "w", 0, "breaks the output into lines of the given width")
	flag.Parse()

	// remaining args after flags are parsed
	args := flag.Args()

	if len(args) != 1 {
		flag.Usage()
		os.Exit(1)
	}
	file = args[0]
}


// find_number interprets the bytes of the given file as a single large integer, and prints that integer to stdout.
// NB: it reads the entire file into a byte array in memory, so it might not be appropriate for large files.
func main() {
	b, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %s %s:\n", file, err)
		os.Exit(1)
	}

	decimalString := infinity.BytesToIntString(b)

	if width == 0 {
		// use full width- don't split the string at all
		fmt.Printf("%s\n", decimalString)
	} else {
		for i, r := range decimalString {
			fmt.Printf("%c", r)
			if (i+1) % width == 0 {
				fmt.Println()
			}
		}

		fmt.Println()
	}
}
