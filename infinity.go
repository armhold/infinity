package infinity

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
)

func BinaryToDecimal(binary string) (int64, error) {
	decimal, err := strconv.ParseInt(os.Args[1], 2, 63)
	if err != nil {
		return 0, fmt.Errorf("error parsing binary string %s: \n", err)
	}

	return decimal, nil
}

func DecimalToBinary(decimal string) (string, error) {
	d, err := strconv.ParseInt(decimal, 10, 64)
	if err != nil {
		return "", fmt.Errorf("error parsing decimal: %s\n", err)
	}

	return fmt.Sprintf("%0b", d), nil
}

// BytesToIntString interprets a slice of bytes as a single giant integer, and returns that integer as a decimal string.
// e.g. [65 66 67] => "4276803"
func BytesToIntString(b []byte) string {
	i := &big.Int{}
	i.SetBytes(b)

	return i.String()
}

// IntStringToBytes takes a decimal integer in string format and returns the byte representation of the integer.
// e.g. "4276803" => [65 66 67]
func IntStringToBytes(s string) []byte {
	i := &big.Int{}
	_, ok := i.SetString(s, 10)
	if !ok {
		log.Fatalf("error creating big.Int from string: %s", s)
	}

	return i.Bytes()
}
