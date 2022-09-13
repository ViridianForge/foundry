package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any length", func(t *testing.T) {
		numbers := []int{4, 10, 2}

		got := Sum(numbers)
		expected := 16

		if got != expected {
			t.Errorf("got %d, expected %d", got, expected)
		}
	})
}

func TestSumAll(t *testing.T) {

	t.Run("summing multiple collections", func(t *testing.T) {
		slice1 := []int{10, 10, 3}
		slice2 := []int{2, 5, 3}

		got := SumAll(slice1, slice2)
		expected := []int{23, 10}

		checkSums(t, got, expected)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		slice1 := []int{}
		slice2 := []int{3, 3, 6}

		got := SumAll(slice1, slice2)
		expected := []int{0, 12}

		checkSums(t, got, expected)
	})
}

func TestSumAllTails(t *testing.T) {

	t.Run("summing tails of multiple colletions", func(t *testing.T) {
		slice1 := []int{5, 7, 3}
		slice2 := []int{3, 4, 1}

		got := SumAllTails(slice1, slice2)
		expected := []int{10, 5}

		checkSums(t, got, expected)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		slice1 := []int{5, 7, 3}
		slice2 := []int{}

		got := SumAllTails(slice1, slice2)
		expected := []int{10, 0}

		checkSums(t, got, expected)
	})
}

func checkSums(t testing.TB, got, expected []int) {
	t.Helper()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v, expected %v", got, expected)
	}
}
