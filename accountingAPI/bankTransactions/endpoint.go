package banktransactions

import (
	"github.com/datag8r/xerogo/errors"
	"github.com/datag8r/xerogo/filter"
)

type bankTransactionsEndpoint struct {
	tenantId          string
	accessToken       string
	rateLimitCallback func()
}

func NewBankTransactionsEndpoint(tenantId, accessToken string, rateLimitCallback func()) *bankTransactionsEndpoint {
	return &bankTransactionsEndpoint{
		tenantId:          tenantId,
		accessToken:       accessToken,
		rateLimitCallback: rateLimitCallback,
	}
}

func (e *bankTransactionsEndpoint) RateLimitCallback() {
	if e.rateLimitCallback != nil {
		e.rateLimitCallback()
	}
}

func (e *bankTransactionsEndpoint) GetOne(id string) (trns BankTransaction, err error) {
	e.RateLimitCallback()
	return GetBankTransaction(id, e.tenantId, e.accessToken)
}

func (e *bankTransactionsEndpoint) GetMulti(where *filter.Filter) (transactions []BankTransaction, err error) {
	var currentPage uint = 1
	for {
		// rate limit handling
		e.RateLimitCallback()
		// actual fetching
		t, pData, err := GetBankTransactions(e.tenantId, e.accessToken, &currentPage, where)
		if err != nil {
			return nil, err
		}
		if pData == nil {
			return nil, errors.ErrNoPageDataReturned
		}
		transactions = append(transactions, t...)
		if pData.PageCount == currentPage {
			break
		}
		currentPage++
	}
	return
}

func (e *bankTransactionsEndpoint) CreateOne(bankTransaction BankTransaction) (trns BankTransaction, err error) {
	e.RateLimitCallback()
	return CreateBankTransaction(e.tenantId, e.accessToken, bankTransaction)
}

func (e *bankTransactionsEndpoint) CreateMulti(bankTransactions []BankTransaction) (transactions []BankTransaction, err error) {
	e.RateLimitCallback()
	return CreateBankTransactions(e.tenantId, e.accessToken, bankTransactions)
}

func (e *bankTransactionsEndpoint) UpdateOne(bankTransaction BankTransaction) (err error) {
	e.RateLimitCallback()
	return UpdateBankTransaction(e.tenantId, e.accessToken, bankTransaction)
}

func (e *bankTransactionsEndpoint) UpdateMulti(bankTransactions []BankTransaction) (err error) {
	e.RateLimitCallback()
	return UpdateBankTransactions(e.tenantId, e.accessToken, bankTransactions)
}

// Not Supported
func (e *bankTransactionsEndpoint) ArchiveOne(id string) (err error) {
	return errors.ErrEndpointCallNotSupported
}

// Not Supported
func (e *bankTransactionsEndpoint) ArchiveMulti(ids []string) (err error) {
	return errors.ErrEndpointCallNotSupported
}

func (e *bankTransactionsEndpoint) DeleteOne(id string) (err error) {
	e.RateLimitCallback()
	return DeleteBankTransaction(e.tenantId, e.accessToken, id)
}

func (e *bankTransactionsEndpoint) DeleteMulti(ids []string) (err error) {
	e.RateLimitCallback()
	return DeleteBankTransactions(e.tenantId, e.accessToken, ids)
}
