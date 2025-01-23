package endpoints

import (
	"github.com/datag8r/xerogo/filter"
)

type Resource any

type ResourceEndpoint[T Resource] interface {
	GetOne(id string) (result T, err error)
	GetMulti(where *filter.Filter) (result []T, err error)

	CreateOne(data T) (result T, err error)
	CreateMulti(data []T) (result []T, err error)

	UpdateOne(data T) (err error)
	UpdateMulti(data []T) (err error)

	ArchiveOne(id string) (err error)
	ArchiveMulti(ids []string) (err error)

	DeleteOne(id string) (err error)
	DeleteMulti(ids []string) (err error)

	RateLimitCallback()
}
