package sorting

import "cmp"

// O(n^2)
func BubbleSort[T cmp.Ordered](s []T) []T {
	sorted := make([]T, len(s))
	copy(sorted, s)

	swapping := true
	for swapping {
		swapping = false
		for i := 1; i < len(sorted); i++ {
			if sorted[i-1] > sorted[i] {
				sorted[i-1], sorted[i] = sorted[i], sorted[i-1]
				swapping = true
			}
		}
	}

	return sorted
}

// O(n*log(n))
func MergeSort[T cmp.Ordered](s []T) []T {
	l := len(s)
	if l < 2 {
		return s
	}

	mid := l / 2
	return merge(MergeSort(s[:mid]), MergeSort(s[mid:]))
}

func merge[T cmp.Ordered](left, right []T) []T {
	merged := []T{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	for i < len(left) {
		merged = append(merged, left[i])
		i++
	}

	for j < len(right) {
		merged = append(merged, right[j])
		j++
	}

	return merged
}

// O(n^2)
func InsertionSort[T cmp.Ordered](s []T) []T {
	sorted := []T{}
	copy(sorted, s)

	for i := range sorted {
		j := i
		for j > 0 && sorted[j-1] > sorted[j] {
			sorted[j], sorted[j-1] = sorted[j-1], sorted[j]
			j--
		}
	}

	return sorted
}

func QuickSort[T cmp.Ordered](s []T) []T {
	return quicksort(s, 0, len(s)-1)
}

func quicksort[T cmp.Ordered](s []T, low, high int) []T {
	if low < high {
		var p int
		s, p = partition(s, low, high)
		s = quicksort(s, low, p-1)
		s = quicksort(s, p+1, high)
	}
	return s
}

func partition[T cmp.Ordered](s []T, low, high int) ([]T, int) {
	pivot := s[high]

	i := low
	for j := low; j < high; j++ {
		if s[j] < pivot {
			s[i], s[j] = s[j], s[i]
			i++
		}
	}
	s[i], s[high] = s[high], s[i]
	return s, i
}
