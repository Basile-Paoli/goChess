package main

func filter[T any](tab []T, f func(T) bool) []T {
	var res []T
	for _, v := range tab {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}
