package utils

import "cmp"

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Min[T cmp.Ordered](values ...T) T {
	if len(values) == 0 {
		return *new(T)
	}
	m := values[0]
	for _, v := range values[1:] {
		if v < m {
			m = v
		}
	}
	return m
}

func Max[T cmp.Ordered](values ...T) T {
	if len(values) == 0 {
		return *new(T)
	}
	m := values[0]
	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}
	return m
}
