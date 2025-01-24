package batchpayments

import (
	"github.com/datag8r/xerogo/errors"
	"github.com/datag8r/xerogo/filter"
)

type batchPaymentsEndpoint struct {
	tenantId          string
	accessToken       string
	rateLimitCallback func()
}

func NewBatchPaymentsEndpoint(tenantId, accessToken string, rateLimitCallback func()) *batchPaymentsEndpoint {
	return &batchPaymentsEndpoint{
		tenantId:          tenantId,
		accessToken:       accessToken,
		rateLimitCallback: rateLimitCallback,
	}
}

func (e *batchPaymentsEndpoint) RateLimitCallback() {
	if e.rateLimitCallback != nil {
		e.rateLimitCallback()
	}
}

func (e *batchPaymentsEndpoint) GetOne(id string) (acc BatchPayment, err error) {
	e.RateLimitCallback()
	return GetBatchPayment(id, e.tenantId, e.accessToken)
}

func (e *batchPaymentsEndpoint) GetMulti(where *filter.Filter) (batchPayments []BatchPayment, err error) {
	// rate limit handling
	e.RateLimitCallback()
	// actual fetching
	return GetBatchPayments(e.tenantId, e.accessToken, where)
}

func (e *batchPaymentsEndpoint) CreateOne(batchPaymentToCreate BatchPayment) (batchPayment BatchPayment, err error) {
	e.RateLimitCallback()
	return CreateBatchPayment(e.tenantId, e.accessToken, batchPaymentToCreate)
}

func (e *batchPaymentsEndpoint) CreateMulti(batchPaymentsToCreate []BatchPayment) (batchPayments []BatchPayment, err error) {
	e.RateLimitCallback()
	return CreateBatchPayments(e.tenantId, e.accessToken, batchPaymentsToCreate)
}

// Not Supported
func (e *batchPaymentsEndpoint) UpdateOne(account BatchPayment) (err error) {
	return errors.ErrEndpointCallNotSupported
}

// Not Supported
func (e *batchPaymentsEndpoint) UpdateMulti(batchPayments []BatchPayment) (err error) {
	return errors.ErrEndpointCallNotSupported
}

// Not Supported
func (e *batchPaymentsEndpoint) ArchiveOne(id string) (err error) {
	return errors.ErrEndpointCallNotSupported
}

// Not Supported
func (e *batchPaymentsEndpoint) ArchiveMulti(ids []string) (err error) {
	return errors.ErrEndpointCallNotSupported
}

func (e *batchPaymentsEndpoint) DeleteOne(id string) (err error) {
	e.RateLimitCallback()
	return DeleteBatchPayment(e.tenantId, e.accessToken, id)
}

func (e *batchPaymentsEndpoint) DeleteMulti(ids []string) (err error) {
	e.RateLimitCallback()
	return DeleteBatchPayments(e.tenantId, e.accessToken, ids)
}
