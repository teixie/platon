package paginator

import (
	"math"
)

var (
	DefaultPage    int64 = 1
	DefaultPerPage int64 = 10
)

type Pagination struct {
	Total       int64       `json:"total"`
	PerPage     int64       `json:"per_page"`
	CurrentPage int64       `json:"current_page"`
	LastPage    int64       `json:"last_page"`
	Data        interface{} `json:"data"`
}

func Page(page int64) int64 {
	if page > 0 {
		return page
	}

	return DefaultPage
}

func PerPage(perPage int64) int64 {
	if perPage > 0 {
		return perPage
	}

	return DefaultPerPage
}

func Paginate(items interface{}, total int64, perPage int64, currentPage int64) *Pagination {
	return &Pagination{
		Total:       total,
		PerPage:     perPage,
		CurrentPage: currentPage,
		LastPage:    int64(math.Ceil(float64(total) / float64(perPage))),
		Data:        items,
	}
}
