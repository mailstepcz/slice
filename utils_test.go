package slice

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicates(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		req := require.New(t)
		res := RemoveDuplicates([]string{})
		req.Equal([]string{}, res)
	})

	t.Run("simple", func(t *testing.T) {
		req := require.New(t)
		res := RemoveDuplicates([]string{"aa", "bb", "cc", "bb"})
		req.Equal([]string{"aa", "bb", "cc"}, res)
	})

	t.Run("uuid", func(t *testing.T) {
		req := require.New(t)
		res := RemoveDuplicates([]int{1, 2, 3, 4, 4, 3, 2, 5})
		req.Equal([]int{1, 2, 3, 4, 5}, res)
	})
}

func TestChunkSliceInts(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		chunkSize int
		expected  [][]int
	}{
		{
			name:      "basic chunking",
			input:     []int{1, 2, 3, 4, 5, 6},
			chunkSize: 2,
			expected:  [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			name:      "chunk size larger than slice",
			input:     []int{1, 2, 3},
			chunkSize: 5,
			expected:  [][]int{{1, 2, 3}},
		},
		{
			name:      "chunk size equal to slice length",
			input:     []int{1, 2, 3, 4},
			chunkSize: 4,
			expected:  [][]int{{1, 2, 3, 4}},
		},
		{
			name:      "chunk size one",
			input:     []int{1, 2, 3},
			chunkSize: 1,
			expected:  [][]int{{1}, {2}, {3}},
		},
		{
			name:      "empty slice",
			input:     []int{},
			chunkSize: 2,
			expected:  nil,
		},
		{
			name:      "nil slice",
			input:     nil,
			chunkSize: 2,
			expected:  nil,
		},
		{
			name:      "uneven division",
			input:     []int{1, 2, 3, 4, 5, 6, 7},
			chunkSize: 3,
			expected:  [][]int{{1, 2, 3}, {4, 5, 6}, {7}},
		},
		{
			name:      "single element",
			input:     []int{42},
			chunkSize: 1,
			expected:  [][]int{{42}},
		},
		{
			name:      "single element with larger chunk",
			input:     []int{42},
			chunkSize: 5,
			expected:  [][]int{{42}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := require.New(t)
			result := ChunkSlice(tt.input, tt.chunkSize)
			req.Equal(tt.expected, result)
		})
	}
}

func TestChunkSliceStrings(t *testing.T) {
	tests := []struct {
		name      string
		input     []string
		chunkSize int
		expected  [][]string
	}{
		{
			name:      "basic string chunking",
			input:     []string{"a", "b", "c", "d", "e"},
			chunkSize: 3,
			expected:  [][]string{{"a", "b", "c"}, {"d", "e"}},
		},
		{
			name:      "empty strings",
			input:     []string{"", "", ""},
			chunkSize: 2,
			expected:  [][]string{{"", ""}, {""}},
		},
		{
			name:      "unicode strings",
			input:     []string{"🚀", "🎉", "💡", "🔥"},
			chunkSize: 2,
			expected:  [][]string{{"🚀", "🎉"}, {"💡", "🔥"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := require.New(t)
			result := ChunkSlice(tt.input, tt.chunkSize)
			req.Equal(tt.expected, result)
		})
	}
}

func TestChunkSliceEdgeCases(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		chunkSize int
		expectNil bool
	}{
		{
			name:      "zero chunk size",
			input:     []int{1, 2, 3},
			chunkSize: 0,
			expectNil: true,
		},
		{
			name:      "negative chunk size",
			input:     []int{1, 2, 3},
			chunkSize: -1,
			expectNil: true,
		},
		{
			name:      "negative chunk size with empty slice",
			input:     []int{},
			chunkSize: -5,
			expectNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := require.New(t)
			result := ChunkSlice(tt.input, tt.chunkSize)
			if tt.expectNil {
				req.Nil(result)
			} else {
				req.NotNil(result)
			}
		})
	}
}

func TestChunkSliceCustomTypes(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	tests := []struct {
		name      string
		input     []Person
		chunkSize int
		expected  [][]Person
	}{
		{
			name: "person structs",
			input: []Person{
				{"Alice", 25},
				{"Bob", 30},
				{"Charlie", 35},
				{"Diana", 40},
				{"Eve", 45},
			},
			chunkSize: 2,
			expected: [][]Person{
				{{"Alice", 25}, {"Bob", 30}},
				{{"Charlie", 35}, {"Diana", 40}},
				{{"Eve", 45}},
			},
		},
		{
			name: "single person",
			input: []Person{
				{"John", 28},
			},
			chunkSize: 3,
			expected: [][]Person{
				{{"John", 28}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := require.New(t)
			result := ChunkSlice(tt.input, tt.chunkSize)
			req.Equal(tt.expected, result)
		})
	}
}

func TestChunkSlicePointers(t *testing.T) {
	values := []int{1, 2, 3, 4}
	pointers := make([]*int, len(values))
	for i := range values {
		pointers[i] = &values[i]
	}

	tests := []struct {
		name      string
		input     []*int
		chunkSize int
		verify    func(t *testing.T, result [][]*int)
	}{
		{
			name:      "pointer chunking",
			input:     pointers,
			chunkSize: 2,
			verify: func(t *testing.T, result [][]*int) {
				req := require.New(t)
				req.Len(result, 2)
				req.Len(result[0], 2)
				req.Len(result[1], 2)

				req.Equal(1, *result[0][0])
				req.Equal(2, *result[0][1])
				req.Equal(3, *result[1][0])
				req.Equal(4, *result[1][1])
			},
		},
		{
			name:      "nil pointers",
			input:     []*int{nil, nil, nil},
			chunkSize: 2,
			verify: func(t *testing.T, result [][]*int) {
				req := require.New(t)
				req.Len(result, 2)
				req.Len(result[0], 2)
				req.Len(result[1], 1)

				req.Nil(result[0][0])
				req.Nil(result[0][1])
				req.Nil(result[1][0])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ChunkSlice(tt.input, tt.chunkSize)
			tt.verify(t, result)
		})
	}
}

func TestChunkSliceLargeData(t *testing.T) {
	tests := []struct {
		name           string
		sliceSize      int
		chunkSize      int
		expectedChunks int
		lastChunkSize  int
	}{
		{
			name:           "1000 elements, 100 chunks",
			sliceSize:      1000,
			chunkSize:      100,
			expectedChunks: 10,
			lastChunkSize:  100,
		},
		{
			name:           "1000 elements, 99 chunks",
			sliceSize:      1000,
			chunkSize:      99,
			expectedChunks: 11,
			lastChunkSize:  10,
		},
		{
			name:           "small slice, large chunks",
			sliceSize:      5,
			chunkSize:      100,
			expectedChunks: 1,
			lastChunkSize:  5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := require.New(t)

			input := make([]int, tt.sliceSize)
			for i := range input {
				input[i] = i
			}

			result := ChunkSlice(input, tt.chunkSize)

			req.Len(result, tt.expectedChunks)

			if len(result) > 0 {
				req.Len(result[len(result)-1], tt.lastChunkSize)
			}

			for i := 0; i < len(result)-1; i++ {
				req.Len(result[i], tt.chunkSize)
			}

			totalElements := 0
			for _, chunk := range result {
				totalElements += len(chunk)
			}
			req.Equal(tt.sliceSize, totalElements)

			if tt.sliceSize > 0 {
				req.Equal(0, result[0][0])
				lastChunk := result[len(result)-1]
				lastElement := lastChunk[len(lastChunk)-1]
				req.Equal(tt.sliceSize-1, lastElement)
			}
		})
	}
}
