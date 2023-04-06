package main

import "testing"

func TestSum(t *testing.T) {
	test := sum(3, 2, 1)

	expectedResult := 6
	if test != expectedResult {
		t.Error("Expected", expectedResult, "Got:", test)
	}
}

// this will fail
func TestSum2(t *testing.T) {
	test := sum(3, 2, 1)

	expectedResult := 7
	if test != expectedResult {
		t.Error("Expected", expectedResult, "Got:", test)
	}
}

func TestMultiply(t *testing.T) {
	test := multiply(10, 10)

	expectedResult := 100
	if test != expectedResult {
		t.Error("Expected", expectedResult, "Got:", test)
	}
}
