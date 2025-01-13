package pagination

type PaginationData struct {
	Page      uint `json:"page"`
	PageSize  uint `json:"pageSize"`
	PageCount uint `json:"pageCount"`
	ItemCount uint `json:"itemCount"`
}

var DefaultPageSize uint = 100
var maxPageSize uint = 1000

var CustomPageSize uint = 100

// 1 - 1000
func SetPageSize(size uint) {
	CustomPageSize = min(size, maxPageSize)
}

func IsDefaultPageSize() bool {
	return CustomPageSize == DefaultPageSize
}
