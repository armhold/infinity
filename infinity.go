package infinity

import (
	"os"
	"fmt"
	"strconv"
	"math/big"
	"log"
	"io/ioutil"
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

func addBig(n int64) *big.Int{
	i := &big.Int{}
	_, success := i.SetString("1234567890123456789012345678901234567890", 10)
	if !success {
		log.Fatal("failed to SetString")
	}

	nInt := &big.Int{}
	nInt.SetInt64(n)
	i.Add(i, nInt)

	return i
}


func ReadFile(file string) *big.Int {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("error reading file %s: %s", file, err)
	}
	i := &big.Int{}

	return i.SetBytes(b)
}
