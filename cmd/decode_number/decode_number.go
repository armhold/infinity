package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/armhold/infinity"
	"io/ioutil"
	"log"
)

var (
	number string
	inFile string
	outFile string
)

func init() {
	var usage = func() {
		u := `Usage: decode_number [-n number] [-o outfile] [file]

If a file is given, the number to decode is read as text from that file.
Else if [-n number] is given, the specified number is used.

If an output file is specified, the decoded bytes are written to that file.
Else they are printed to stdout.

`
		fmt.Fprintf(os.Stderr, u)
	}

	flag.Usage = usage

	flag.StringVar(&number, "n", "", "specifies a number to decode")
	flag.StringVar(&outFile, "o", "", "specifies an output file in which to write the decoded bytes")
	flag.Parse()

	// remaining args after flags are parsed
	args := flag.Args()

	if number == "" && len(args) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	if len(args) > 0 {
		inFile = args[0]
	}
}

// decode_number decodes the given number into its bytes, converts those bytes into a string, and
// prints the string to stdout.
func main() {
	if inFile != "" {
		s, err := ioutil.ReadFile(inFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading file: %s: %s ", inFile, err)
			os.Exit(1)
		}
		number = string(s)
	}

	bytes := infinity.IntStringToBytes(number)

	if outFile != "" {
		err := ioutil.WriteFile(outFile, bytes, 0644)
		if err != nil {
			log.Fatalf("error writing output file: %s", err)
		}
	} else {
		s := string(bytes)
		fmt.Fprintf(os.Stdout, "%s\n", s)
	}
}
