package main

import "testing"

func TestAdd(t *testing.T) {
	total := Add(2, 5)
	if total != 7 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 7)
	}
}
