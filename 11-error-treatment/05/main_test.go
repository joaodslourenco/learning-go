package math

import "testing"

func TestIsEven(t *testing.T) {
	testCase := 2
	v = IsEven(testCase)
	if v == false {
		t.Error("Expected to be true")
	}
}
