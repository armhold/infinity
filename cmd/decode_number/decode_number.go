package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/armhold/infinity"
	"io/ioutil"
)

var (
	number string
	file   string
)

func init() {
	var usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-n number] [file]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Usage = usage

	flag.StringVar(&number, "n", "", "decode given number")
	flag.Parse()

	// remaining args after flags are parsed
	args := flag.Args()

	if number == "" && len(args) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	if len(args) > 0 {
		file = args[0]
	}
}

// decode_number decodes the given number into its bytes, converts those bytes into a string, and
// prints the string to stdout.
func main() {
	if file != "" {
		s, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading file: %s: %s ", file, err)
			os.Exit(1)
		}
		number = string(s)
	}

	bytes := infinity.IntStringToBytes(number)
	s := string(bytes)

	fmt.Fprintf(os.Stdout, "%s\n", s)
}
