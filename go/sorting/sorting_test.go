package sorting_test

import (
	"slices"
	"testing"

	"github.com/zoumas/dsa/go/sorting"
)

func TestMergeSort(t *testing.T) {
	given := []int{2, 6, 9, 5, 0, 1, 5, 3}
	want := []int{0, 1, 2, 3, 5, 5, 6, 9}

	got := sorting.MergeSort(given)

	if !slices.Equal(got, want) {
		t.Errorf("\ngot:\n%v\nwant:\n%v\ngiven:\n%v", got, want, given)
	}
}
