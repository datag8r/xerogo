package banktransfers

import (
	"github.com/datag8r/xerogo/errors"
	"github.com/datag8r/xerogo/filter"
)

type bankTransfersEndpoint struct {
	tenantId          string
	accessToken       string
	rateLimitCallback func()
}

func NewBankTransfersEndpoint(tenantId, accessToken string, rateLimitCallback func()) *bankTransfersEndpoint {
	return &bankTransfersEndpoint{
		tenantId:          tenantId,
		accessToken:       accessToken,
		rateLimitCallback: rateLimitCallback,
	}
}

func (e *bankTransfersEndpoint) RateLimitCallback() {
	if e.rateLimitCallback != nil {
		e.rateLimitCallback()
	}
}

func (e *bankTransfersEndpoint) GetOne(id string) (transfer BankTransfer, err error) {
	e.RateLimitCallback()
	return GetBankTransfer(id, e.tenantId, e.accessToken)
}

func (e *bankTransfersEndpoint) GetMulti(where *filter.Filter) (transfers []BankTransfer, err error) {
	e.RateLimitCallback()
	return GetBankTransfers(e.tenantId, e.accessToken, where)
}

func (e *bankTransfersEndpoint) CreateOne(bt BankTransfer) (transfer BankTransfer, err error) {
	e.RateLimitCallback()
	return CreateBankTransfer(e.tenantId, e.accessToken, bt)
}

func (e *bankTransfersEndpoint) CreateMulti(bts []BankTransfer) (transfers []BankTransfer, err error) {
	e.RateLimitCallback()
	return CreateBankTransfers(e.tenantId, e.accessToken, bts)
}

// Not Supported
func (e *bankTransfersEndpoint) UpdateOne(bt BankTransfer) (err error) {
	return errors.ErrEndpointCallNotSupported
}

// Not Supported
func (e *bankTransfersEndpoint) UpdateMulti(accounts []BankTransfer) (err error) {
	return errors.ErrEndpointCallNotSupported
}

// Not Supported
func (e *bankTransfersEndpoint) ArchiveOne(id string) (err error) {
	return errors.ErrEndpointCallNotSupported
}

// Not Supported
func (e *bankTransfersEndpoint) ArchiveMulti(ids []string) (err error) {
	return errors.ErrEndpointCallNotSupported
}

// Not Supported
func (e *bankTransfersEndpoint) DeleteOne(id string) (err error) {
	return errors.ErrEndpointCallNotSupported
}

// Not Supported
func (e *bankTransfersEndpoint) DeleteMulti(ids []string) (err error) {
	return errors.ErrEndpointCallNotSupported
}
