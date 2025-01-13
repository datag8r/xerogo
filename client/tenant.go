package client

import (
	"sync"
	"time"

	"github.com/datag8r/xerogo/accountingAPI/accounts"
	"github.com/datag8r/xerogo/accountingAPI/items"
	"github.com/datag8r/xerogo/filter"
)

type tenant struct {
	c              *client       `json:"-"`
	TenantID       string        `json:"tenantId"`
	Name           string        `json:"tenantName"`
	CreatedDateUTC string        `json:"createdDateUtc"`
	UpdatedDateUTC string        `json:"updatedDateUtc"`
	rateLimit      time.Duration `json:"-"`
	lastCall       time.Time     `json:"-"`
	rateMutex      sync.Mutex    `json:"-"`
}

// This Function is Used For Rate Limiting, To avoid being rate lmited you should call this before every request
// Built In (*tenant) and (*client) Methods Call This For you
func (t *tenant) Call() {
	t.rateMutex.Lock()
	t.c.Call()                          // Call Client Rate Limiter
	next := t.lastCall.Add(t.rateLimit) // Min Time Of Next Call
	current := time.Now()
	toWait := next.Sub(current)
	<-time.After(toWait)
	t.lastCall = time.Now()
	t.rateMutex.Unlock()
}

// Accounts

func (t *tenant) GetAccounts(where *filter.Filter) (accs []accounts.Account, err error) {
	err = t.c.Refresh()
	if err != nil {
		return
	}
	t.Call()
	return accounts.GetAccounts(t.TenantID, t.c.AccessToken, where)
}

func (t *tenant) GetAccount(accountID string) (acc accounts.Account, err error) {
	err = t.c.Refresh()
	if err != nil {
		return
	}
	t.Call()
	return accounts.GetAccount(accountID, t.TenantID, t.c.AccessToken)
}

func (t *tenant) CreateAccount(account accounts.Account) (acc accounts.Account, err error) {
	err = t.c.Refresh()
	if err != nil {
		return
	}
	t.Call()
	return accounts.CreateAccount(account, t.TenantID, t.c.AccessToken)
}

func (t *tenant) UpdateAccount(account accounts.Account) (err error) {
	err = t.c.Refresh()
	if err != nil {
		return
	}
	t.Call()
	return accounts.UpdateAccount(account, t.TenantID, t.c.AccessToken)
}

func (t *tenant) ArchiveAccount(account accounts.Account) (err error) {
	err = t.c.Refresh()
	if err != nil {
		return
	}
	t.Call()
	return accounts.ArchiveAccount(account.AccountID, t.TenantID, t.c.AccessToken)
}

func (t *tenant) DeleteAccount(account accounts.Account) (err error) {
	err = t.c.Refresh()
	if err != nil {
		return
	}
	t.Call()
	return accounts.DeleteAccount(account.AccountID, t.TenantID, t.c.AccessToken)
}

// Items

func (t *tenant) GetItems(where *filter.Filter) (itemList []items.Item, err error) {
	err = t.c.Refresh()
	if err != nil {
		return
	}
	t.Call()
	return items.GetItems(t.TenantID, t.c.AccessToken, where)
}

func (t *tenant) GetItem(itemIdOrCode string) (item items.Item, err error) {
	err = t.c.Refresh()
	if err != nil {
		return
	}
	t.Call()
	return items.GetItem(itemIdOrCode, t.TenantID, t.c.AccessToken)
}

func (t *tenant) GetItemHistory(itemIdOrCode string) (history []items.ItemHistory, err error) {
	err = t.c.Refresh()
	if err != nil {
		return
	}
	t.Call()
	return items.GetItemHistory(itemIdOrCode, t.TenantID, t.c.AccessToken)
}

func (t *tenant) CreateItem(itemToCreate items.Item) (item items.Item, err error) {
	err = t.c.Refresh()
	if err != nil {
		return
	}
	t.Call()
	return items.CreateItem(itemToCreate, t.TenantID, t.c.AccessToken)
}

func (t *tenant) UpdateItem(item items.Item) (err error) {
	err = t.c.Refresh()
	if err != nil {
		return
	}
	t.Call()
	return items.UpdateItem(item, t.TenantID, t.c.AccessToken)
}

func (t *tenant) DeleteItem(itemIdOrCode string) (err error) {
	err = t.c.Refresh()
	if err != nil {
		return
	}
	t.Call()
	return items.DeleteItem(itemIdOrCode, t.TenantID, t.c.AccessToken)
}

func (t *tenant) AddNoteToItem(itemIdOrCode, note string) (err error) {
	err = t.c.Refresh()
	if err != nil {
		return
	}
	t.Call()
	return items.AddNoteToItem(itemIdOrCode, note, t.TenantID, t.c.AccessToken)
}

// Contacts

// Template For Tenant Funcs (bc lazy)
func (t *tenant) _() {}
