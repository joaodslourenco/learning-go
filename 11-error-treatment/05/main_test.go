package main

import "testing"

func TestIsEven(t *testing.T) {
	testCase := 2
	v := isEven(testCase)
	if v == false {
		t.Error("Expected to be true")
	}
}
