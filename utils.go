package slice

// RemoveDuplicates removes duplicates from slice.
func RemoveDuplicates[T comparable](in []T) []T {
	set := make(map[T]struct{}, len(in))
	unique := []T{}

	for _, val := range in {
		if _, ok := set[val]; !ok {
			set[val] = struct{}{}
			unique = append(unique, val)
		}
	}

	return unique
}
