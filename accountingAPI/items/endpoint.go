package items

import (
	"github.com/datag8r/xerogo/errors"
	"github.com/datag8r/xerogo/filter"
)

type itemsEndpoint struct {
	tenantId          string
	accessToken       string
	rateLimitCallback func()
}

func NewItemsEndpoint(tenantId, accessToken string, rateLimitCallback func()) *itemsEndpoint {
	return &itemsEndpoint{
		tenantId:          tenantId,
		accessToken:       accessToken,
		rateLimitCallback: rateLimitCallback,
	}
}

func (e *itemsEndpoint) RateLimitCallback() {
	if e.rateLimitCallback != nil {
		e.rateLimitCallback()
	}
}

func (e *itemsEndpoint) GetOne(id string) (item Item, err error) {
	e.RateLimitCallback()
	return GetItem(e.tenantId, e.accessToken, id)
}

func (e *itemsEndpoint) GetMulti(where *filter.Filter) (items []Item, err error) {
	e.RateLimitCallback()
	return GetItems(e.tenantId, e.accessToken, where)
}

func (e *itemsEndpoint) CreateOne(itemToCreate Item) (item Item, err error) {
	e.RateLimitCallback()
	return CreateItem(e.tenantId, e.accessToken, itemToCreate)
}

func (e *itemsEndpoint) CreateMulti(itemsToCreate []Item) (items []Item, err error) {
	e.RateLimitCallback()
	return CreateItems(e.tenantId, e.accessToken, itemsToCreate)
}

func (e *itemsEndpoint) UpdateOne(item Item) (err error) {
	e.RateLimitCallback()
	return UpdateItem(e.tenantId, e.accessToken, item)
}

func (e *itemsEndpoint) UpdateMulti(items []Item) (err error) {
	e.RateLimitCallback()
	return UpdateItems(e.tenantId, e.accessToken, items)
}

// Not Supported
func (e *itemsEndpoint) ArchiveOne(id string) (err error) {
	return errors.ErrEndpointCallNotSupported
}

// Not Supported
func (e *itemsEndpoint) ArchiveMulti(ids []string) (err error) {
	return errors.ErrEndpointCallNotSupported
}

func (e *itemsEndpoint) DeleteOne(id string) (err error) {
	e.RateLimitCallback()
	return DeleteItem(id, e.tenantId, e.accessToken)
}

// Not Supported
func (e *itemsEndpoint) DeleteMulti(ids []string) (err error) {
	return errors.ErrEndpointCallNotSupported
}
