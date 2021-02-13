package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	if calculate(2) != 3 {
		t.Error("Expected 2 + 2 to equal 3")
	}
}
