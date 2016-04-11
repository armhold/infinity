package infinity

import "testing"

func TestFoo(t *testing.T) {

	i := addBig(1)
	actual := i.String()

	expected := "1234567890123456789012345678901234567891"
	if actual != expected {
		t.Fatalf("expected: %s, got: %s", expected, actual)
	}


}
