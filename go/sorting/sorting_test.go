package sorting_test

import (
	"cmp"
	"math/rand"
	"slices"
	"testing"

	"github.com/zoumas/dsa/go/sorting"
)

type testCase[T cmp.Ordered] struct {
	name     string
	sortFunc func(s []T) []T
}

var testCases = []testCase[int]{
	{
		name:     "Bubble Sort",
		sortFunc: sorting.BubbleSort[int],
	},
	{
		name:     "Merge Sort",
		sortFunc: sorting.MergeSort[int],
	},
	{
		name:     "Insertion Sort",
		sortFunc: sorting.InsertionSort[int],
	},
	{
		name:     "Quick Sort",
		sortFunc: sorting.QuickSort[int],
	},
}

func TestSort(t *testing.T) {
	s := make([]int, 20)
	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(len(s))
	}

	for _, cs := range testCases {
		t.Run(cs.name, func(t *testing.T) {
			assertIsSorted(t, cs.sortFunc(s))
		})
	}
}

func assertIsSorted[T cmp.Ordered](t testing.TB, s []T) {
	t.Helper()

	if !slices.IsSorted(s) {
		t.Errorf("\nSlice is not sorted\n%v", s)
	}
}
