package main

import (
	"testing"
)

func TestSoma(t *testing.T) {
	total := Soma(20, 20)
	r := 40
	if total != r {
		t.Errorf("Invalid result: expected: %d, actual: %d", r, total)
	}

}
