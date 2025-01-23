package accounts

import (
	"github.com/datag8r/xerogo/errors"
	"github.com/datag8r/xerogo/filter"
)

type accountsEndpoint struct {
	tenantId          string
	accessToken       string
	rateLimitCallback func()
}

func NewAccountsEndpoint(tenantId, accessToken string, rateLimitCallback func()) *accountsEndpoint {
	return &accountsEndpoint{
		tenantId:          tenantId,
		accessToken:       accessToken,
		rateLimitCallback: rateLimitCallback,
	}
}

func (e *accountsEndpoint) RateLimitCallback() {
	if e.rateLimitCallback != nil {
		e.rateLimitCallback()
	}
}

func (e *accountsEndpoint) GetOne(id string) (acc Account, err error) {
	e.RateLimitCallback()
	return GetAccount(id, e.tenantId, e.accessToken)
}

func (e *accountsEndpoint) GetMulti(where *filter.Filter) (accs []Account, err error) {
	e.RateLimitCallback()
	return GetAccounts(e.tenantId, e.accessToken, where)
}

func (e *accountsEndpoint) CreateOne(account Account) (acc Account, err error) {
	e.RateLimitCallback()
	return CreateAccount(account, e.tenantId, e.accessToken)
}

// Not Supported
func (e *accountsEndpoint) CreateMulti(accounts []Account) (accs []Account, err error) {
	return nil, errors.ErrEndpointCallNotSupported
}

func (e *accountsEndpoint) UpdateOne(account Account) (err error) {
	e.RateLimitCallback()
	return UpdateAccount(account, e.tenantId, e.accessToken)
}

// Not Supported
func (e *accountsEndpoint) UpdateMulti(accounts []Account) (err error) {
	return errors.ErrEndpointCallNotSupported
}

func (e *accountsEndpoint) ArchiveOne(id string) (err error) {
	e.RateLimitCallback()
	return ArchiveAccount(id, e.tenantId, e.accessToken)
}

// Not Supported
func (e *accountsEndpoint) ArchiveMulti(ids []string) (err error) {
	return errors.ErrEndpointCallNotSupported
}

func (e *accountsEndpoint) DeleteOne(id string) (err error) {
	e.RateLimitCallback()
	return DeleteAccount(id, e.tenantId, e.accessToken)
}

// Not Supported
func (e *accountsEndpoint) DeleteMulti(ids []string) (err error) {
	return errors.ErrEndpointCallNotSupported
}
