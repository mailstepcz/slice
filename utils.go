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

// ChunkSlice generic helper to split a slice into chunks of given size.
func ChunkSlice[T any](items []T, chunkSize int) [][]T {
	if chunkSize <= 0 {
		return nil
	}

	if len(items) == 0 {
		return nil
	}

	var chunks [][]T
	for start := 0; start < len(items); start += chunkSize {
		end := start + chunkSize
		if end > len(items) {
			end = len(items)
		}
		chunks = append(chunks, items[start:end])
	}
	return chunks
}
