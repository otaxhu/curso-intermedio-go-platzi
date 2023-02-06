package main

import "testing"

func TestSum(t *testing.T) {
	/*
		total := Sum(5, 5)

		if total != 11 {
			t.Errorf("Sum was incorrect, got %d expected %d", total, 11)
		}
	*/
	tables := []struct {
		a      int
		b      int
		result int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{4, 5, 9},
		{20, 21, 41},
	}
	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.result {
			t.Errorf("Sum was incorrect, got %d expected %d", total, item.result)
		}
	}
}
