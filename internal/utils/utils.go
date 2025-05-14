package utils

func GetNumberOfElementInSlice[S []E, E comparable](s S, e E) int {
	n := 0

	for _, elem := range s {
		if elem == e {
			n++
		}
	}

	return n
}
