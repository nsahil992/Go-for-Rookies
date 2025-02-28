package main

import "testing"

// Test function should always start with Test word itself and with capital T

func TestCheckEven(t *testing.T) {
	i := 10
	expected := "Even"

	res := checkEven(i)
	if expected != res {
		t.Errorf("Expected %s, got %s", expected, res)
	}
}
