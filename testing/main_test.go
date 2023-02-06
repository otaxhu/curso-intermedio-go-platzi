package main

import "testing"

func TestSum(t *testing.T) {
	/*
		total := Sum(5, 5)

		if total != 11 {
			t.Errorf("Sum was incorrect, got %d expected %d", total, 11)
		}
	*/

	// Code Coverage del 100% ya que la funcion "Sum()" solo tiene un
	// statement if y en todos los casos retorna su valor exitosamente

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

/*

// Code Coverage del 75% ya que todos los casos que se presentan solo retornan
// el primer statement pero nunca el segundo

func TestGetMax(t *testing.T) {
	tables := []struct {
		a      int
		b      int
		result int
	}{
		{3, 2, 3},
		{4, 2, 4},
		{31, 30, 31},
		{4, 3, 4},
	}
	for _, item := range tables {
		max := GetMax(item.a, item.b)
		if max != item.result {
			t.Errorf("GetMax was incorrect, got %d expected %d", max, item.result)
		}
	}
}
*/

// Code Coverage del 100% ya que en todos los casos la funcion "GetMax()"
// logra retornar sus dos statements exitosamente
func TestGetMax(t *testing.T) {
	tables := []struct {
		a      int
		b      int
		result int
	}{
		{3, 4, 4},
		{4, 2, 4},
		{31, 30, 31},
		{4, 3, 4},
	}
	for _, item := range tables {
		max := GetMax(item.a, item.b)
		if max != item.result {
			t.Errorf("GetMax was incorrect, got %d expected %d", max, item.result)
		}
	}
}
