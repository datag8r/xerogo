package pagination

import "cmp"

type PaginationData struct {
	Page      uint `json:"page"`
	PageSize  uint `json:"pageSize"`
	PageCount uint `json:"pageCount"`
	ItemCount uint `json:"itemCount"`
}

const (
	DefaultPageSize uint = 100
	maxPageSize     uint = 1000
)

var CustomPageSize uint = 100

// 1 - 1000
func SetPageSize(size uint) {
	CustomPageSize = clamp(1, size, maxPageSize)
}

func IsDefaultPageSize() bool {
	return CustomPageSize == DefaultPageSize
}

func clamp[T cmp.Ordered](minimum, actual, maximum T) T {
	return max(minimum, min(actual, maximum))
}
