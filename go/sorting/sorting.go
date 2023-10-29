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
