package sorting

import "cmp"

func MergeSort[T cmp.Ordered](s []T) []T {
	l := len(s)
	if l < 2 {
		return s
	}

	mid := l / 2
	return merge(MergeSort(s[:mid]), MergeSort(s[mid:]))
}

func merge[T cmp.Ordered](xs, ys []T) []T {
	final := []T{}
	lx := len(xs)
	ly := len(ys)

	x, y := 0, 0
	for x < lx && y < ly {
		if xs[x] <= ys[y] {
			final = append(final, xs[x])
			x++
		} else {
			final = append(final, ys[y])
			y++
		}
	}

	for x < lx {
		final = append(final, xs[x])
		x++
	}
	for y < ly {
		final = append(final, ys[y])
		y++
	}

	return final
}
