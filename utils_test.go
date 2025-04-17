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
		res := RemoveDuplicates([]int{1,2,3,4,4,3,2,5})
		req.Equal([]int{1,2,3,4,5}, res)
	})
}
