package tests

import "testing"

func Add(x, y int) int {
	return x + y
}

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Add(2, 3) = %d; want 5", result)
	}
}
