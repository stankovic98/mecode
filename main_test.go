package main

import (
	"testing"
)

func TestFib(t *testing.T) {
	fib := Fib(6)

	if fib != 8 {
		t.Errorf("Fib was incorrect, got: %d, want: %d.", fib, 8)
	}
}
