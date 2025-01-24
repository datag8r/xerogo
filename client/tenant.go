package client

import (
	"sync"
	"time"

	"github.com/datag8r/xerogo/accountingAPI/accounts"
	banktransactions "github.com/datag8r/xerogo/accountingAPI/bankTransactions"
	banktransfers "github.com/datag8r/xerogo/accountingAPI/bankTransfers"
	batchpayments "github.com/datag8r/xerogo/accountingAPI/batchPayments"
	contactgroups "github.com/datag8r/xerogo/accountingAPI/contactGroups"
	"github.com/datag8r/xerogo/accountingAPI/contacts"
	"github.com/datag8r/xerogo/accountingAPI/currencies"
	"github.com/datag8r/xerogo/accountingAPI/employees"
	"github.com/datag8r/xerogo/accountingAPI/items"
	"github.com/datag8r/xerogo/endpoints"
)

type Tenant struct {
	*tenantEndpoints `json:"-"`
	*Client          `json:"-"`
	TenantID         string        `json:"tenantId"`
	Name             string        `json:"tenantName"`
	CreatedDateUTC   string        `json:"createdDateUtc"`
	UpdatedDateUTC   string        `json:"updatedDateUtc"`
	rateLimit        time.Duration `json:"-"`
	lastCall         time.Time     `json:"-"`
	rateMutex        sync.Mutex    `json:"-"`
}

type tenantEndpoints struct {
	Accounts         endpoints.ResourceEndpoint[accounts.Account]
	BankTransfers    endpoints.ResourceEndpoint[banktransfers.BankTransfer]
	BankTransactions endpoints.ResourceEndpoint[banktransactions.BankTransaction]
	BatchPayments    endpoints.ResourceEndpoint[batchpayments.BatchPayment]
	Contacts         endpoints.ResourceEndpoint[contacts.Contact]
	Items            endpoints.ResourceEndpoint[items.Item]
	Employees        endpoints.ResourceEndpoint[employees.Employee]
	Currencies       endpoints.ResourceEndpoint[currencies.Currency]
	ContactGroups    endpoints.ResourceEndpoint[contactgroups.ContactGroup]
	// etc
}

func NewTenantEndpoints(t *Tenant) *tenantEndpoints {
	f := func() {
		t.Call()
		if err := t.Refresh(); err != nil {
			panic(err) // not sure whether to panic here or not
		}
	}
	return &tenantEndpoints{
		Items:            items.NewItemsEndpoint(t.TenantID, t.AccessToken, f),
		Accounts:         accounts.NewAccountsEndpoint(t.TenantID, t.AccessToken, f),
		Contacts:         contacts.NewContactsEndpoint(t.TenantID, t.AccessToken, f),
		Employees:        employees.NewEmployeesEndpoint(t.TenantID, t.AccessToken, f),
		Currencies:       currencies.NewCurrencysEndpoint(t.TenantID, t.AccessToken, f),
		ContactGroups:    contactgroups.NewContactGroupsEndpoint(t.TenantID, t.AccessToken, f),
		BankTransfers:    banktransfers.NewBankTransfersEndpoint(t.TenantID, t.AccessToken, f),
		BatchPayments:    batchpayments.NewBatchPaymentsEndpoint(t.TenantID, t.AccessToken, f),
		BankTransactions: banktransactions.NewBankTransactionsEndpoint(t.TenantID, t.AccessToken, f),
	}
}

// This Function is Used For Rate Limiting, To avoid being rate lmited you should call this before every request
// Built In (*Tenant), (*client) and (ResourceEndpoint) Methods Call This For you
func (t *Tenant) Call() {
	t.rateMutex.Lock()
	t.Client.Call()                     // Call Client Rate Limiter
	next := t.lastCall.Add(t.rateLimit) // Min Time Of Next Call
	current := time.Now()
	toWait := next.Sub(current)
	<-time.After(toWait)
	t.lastCall = time.Now()
	t.rateMutex.Unlock()
}
