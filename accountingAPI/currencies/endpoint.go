package currencies

import (
	"github.com/datag8r/xerogo/errors"
	"github.com/datag8r/xerogo/filter"
)

type currenciesEndpoint struct {
	tenantId          string
	accessToken       string
	rateLimitCallback func()
}

func NewCurrencysEndpoint(tenantId, accessToken string, rateLimitCallback func()) *currenciesEndpoint {
	return &currenciesEndpoint{
		tenantId:          tenantId,
		accessToken:       accessToken,
		rateLimitCallback: rateLimitCallback,
	}
}

func (e *currenciesEndpoint) RateLimitCallback() {
	if e.rateLimitCallback != nil {
		e.rateLimitCallback()
	}
}

// Not Supported
func (e *currenciesEndpoint) GetOne(id string) (acc Currency, err error) {
	return Currency{}, errors.ErrEndpointCallNotSupported
}

func (e *currenciesEndpoint) GetMulti(where *filter.Filter) (accs []Currency, err error) {
	e.RateLimitCallback()
	return GetCurrencies(e.tenantId, e.accessToken, where)
}

func (e *currenciesEndpoint) CreateOne(currencyToCreate Currency) (currency Currency, err error) {
	e.RateLimitCallback()
	return CreateCurrency(e.tenantId, e.accessToken, currencyToCreate)
}

func (e *currenciesEndpoint) CreateMulti(currenciesToCreate []Currency) (currencies []Currency, err error) {
	return CreateCurrencies(e.tenantId, e.accessToken, currenciesToCreate)
}

// Not Supported
func (e *currenciesEndpoint) UpdateOne(currency Currency) (err error) {
	return errors.ErrEndpointCallNotSupported
}

// Not Supported
func (e *currenciesEndpoint) UpdateMulti(currencies []Currency) (err error) {
	return errors.ErrEndpointCallNotSupported
}

// Not Supported
func (e *currenciesEndpoint) ArchiveOne(id string) (err error) {
	return errors.ErrEndpointCallNotSupported
}

// Not Supported
func (e *currenciesEndpoint) ArchiveMulti(ids []string) (err error) {
	return errors.ErrEndpointCallNotSupported
}

// Not Supported
func (e *currenciesEndpoint) DeleteOne(id string) (err error) {
	return errors.ErrEndpointCallNotSupported
}

// Not Supported
func (e *currenciesEndpoint) DeleteMulti(ids []string) (err error) {
	return errors.ErrEndpointCallNotSupported
}
