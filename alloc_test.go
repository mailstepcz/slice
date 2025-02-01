package slice

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAlloc(t *testing.T) {
	req := require.New(t)

	sl := Alloc[int](10)

	req.Len(sl, 10)

	for i := 0; i < len(sl); i++ {
		sl[i] = i + 1
	}
	req.Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, sl)
}
