package infinity

import (
	"os"
	"fmt"
	"strconv"
)

func Binary2decimal(binary string) (int64, error) {
	decimal, err := strconv.ParseInt(os.Args[1], 2, 63)
	if err != nil {
		return 0, fmt.Errorf("error parsing binary string %s: \n", err)
	}

	return decimal, nil
}

func Decimal2binary(decimal string) (string, error) {
	d, err := strconv.ParseInt(decimal, 10, 64)
	if err != nil {
		return "", fmt.Errorf("error parsing decimal: %s\n", err)
	}

	return fmt.Sprintf("%0b", d), nil
}
