package infinity

import (
	"testing"
	"math/big"
)

func TestAdd(t *testing.T) {
	i := addBig(1)
	actual := i.String()

	expected := "1234567890123456789012345678901234567891"
	if actual != expected {
		t.Fatalf("expected: %s, got: %s", expected, actual)
	}
}

func TestReadFile(t *testing.T) {
	// NB: no newline in the file; use "od -t u1 samples/abc.txt" to examine it
	actual := ReadFile("samples/ABC.txt").String()
	expected := "4276803"

	// a == 65 == 01000001
	// b == 66 == 01000010
	// c == 67 == 01000011
	//
	// 01000001 01000010 01000011  => 10000010100001001000011 => 4,276,803

	if actual != expected {
		t.Fatalf("expected: %s, got: %s", expected, actual)
	}
}

func TestSetBytes(t *testing.T) {

	i := big.Int{}
	i.SetBytes([]byte { 255, 254 })

	actual := i.String()

	expected := "65534"
	if actual != expected {
		t.Fatalf("expected: %s, got: %s", expected, actual)
	}
}

func TestRoundTrip(t *testing.T) {
	// test bytes -> int
	//
	expected := "4276803"
	actual := BytesToIntString([]byte("ABC"))

	if actual != expected {
		t.Fatalf("expected: %s, got: %s", expected, actual)
	}

	// test int -> bytes
	//
	expected = "ABC"
	actual = string(IntStringToBytes("4276803"))

	if actual != expected {
		t.Fatalf("expected: %s, got: %s", expected, actual)
	}
}
