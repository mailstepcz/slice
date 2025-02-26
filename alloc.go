package slice

import (
	"unsafe"

	"github.com/fealsamh/go-utils/mem"
)

// Alloc allocates a slice if type T and the given size.
func Alloc[T any](size int) []T {
	arr, err := mem.Alloc(size * int(unsafe.Sizeof(*(*T)(nil))))
	if err != nil {
		panic(err)
	}
	return unsafe.Slice((*T)(unsafe.Pointer(unsafe.SliceData(arr))), size)
}
