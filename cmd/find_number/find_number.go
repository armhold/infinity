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
	indent int
	text string
	file string
)


func init() {
	var usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-w width] [-t text] [file]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Usage = usage

	flag.IntVar(&width, "w", 0, "breaks the output into lines of the given width")
	flag.IntVar(&indent, "i", 0, "indents the output")
	flag.StringVar(&text, "t", "", "search given text")
	flag.Parse()

	// remaining args after flags are parsed
	args := flag.Args()

	if text == "" && len(args) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	if len(args) > 0 {
		file = args[0]
	}
}


// find_number interprets the bytes of the given string or file as a single large integer,
// and prints that integer to stdout.
//
// NB: it reads the entire file into a byte array in memory, so it might not be appropriate for large files.
func main() {
	var bytes []byte

	if text != "" {
		bytes = []byte(text)
	} else {
		b, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading file: %s %s:\n", file, err)
			os.Exit(1)
		}
		bytes = b
	}

	decimalString := infinity.BytesToIntString(bytes)

	indentString := ""
	for i := 0; i < indent; i++ {
		indentString += " "
	}

	if width == 0 {
		// use full width- don't split the string at all
		fmt.Printf("%s%s\n", indentString, decimalString)
	} else {
		for i, r := range decimalString {
			if i % width == 0 {
				fmt.Printf("%s%c", indentString, r)
			} else if (i+1) % width == 0 {
				fmt.Printf("%c", r)
				fmt.Println()
			} else {
				fmt.Printf("%c", r)
			}
		}
		fmt.Println()
	}
}
