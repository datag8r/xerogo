package client

import (
	"errors"
	"sync"
	"time"

	"github.com/datag8r/xerogo/accountingAPI/accounts"
	"github.com/datag8r/xerogo/accountingAPI/contacts"
	"github.com/datag8r/xerogo/accountingAPI/endpoints"
	"github.com/datag8r/xerogo/accountingAPI/history"
	"github.com/datag8r/xerogo/accountingAPI/items"
	"github.com/datag8r/xerogo/accountingAPI/pagination"
	"github.com/datag8r/xerogo/accountingAPI/users"
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

func (t *tenant) GetItemHistory(itemIdOrCode string) (historyList []history.History, err error) {
	err = t.c.Refresh()
	if err != nil {
		return
	}
	t.Call()
	return history.GetResourceHistory(endpoints.EndpointItems, itemIdOrCode, t.TenantID, t.c.AccessToken)
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
	return history.AddNoteToResource(endpoints.EndpointItems, itemIdOrCode, note, t.TenantID, t.c.AccessToken)
}

// Contacts

func (t *tenant) GetContacts(where *filter.Filter, usePagination bool) (contactList []contacts.Contact, err error) {
	if usePagination {
		var currentPage uint = 1
		for {
			err = t.c.Refresh()
			if err != nil {
				return
			}
			t.Call()
			var conList []contacts.Contact
			var pData *pagination.PaginationData
			conList, pData, err = contacts.GetContacts(t.TenantID, t.c.AccessToken, &currentPage, nil)
			if err != nil {
				return
			}
			if pData == nil {
				err = errors.New("no page data returned from paginated call")
				return
			}
			contactList = append(contactList, conList...)
			if pData.PageCount == currentPage {
				break
			}
			currentPage++
		}
		return
	} else {
		err = t.c.Refresh()
		if err != nil {
			return
		}
		t.Call()
		contactList, _, err = contacts.GetContacts(t.TenantID, t.c.AccessToken, nil, nil)
		return
	}
}

func (t *tenant) GetContact(contactId string) (contact contacts.Contact, err error) {
	err = t.c.Refresh()
	if err != nil {
		return
	}
	t.Call()
	return contacts.GetContact(t.TenantID, t.c.AccessToken, contactId)
}

func (t *tenant) CreateContact(contactToCreate contacts.Contact) (contact contacts.Contact, err error) {
	err = t.c.Refresh()
	if err != nil {
		return
	}
	t.Call()
	return contacts.CreateContact(contactToCreate, t.TenantID, t.c.AccessToken)
}

func (t *tenant) UpdateContact(contact contacts.Contact) (err error) {
	err = t.c.Refresh()
	if err != nil {
		return
	}
	t.Call()
	return contacts.UpdateContact(contact, t.TenantID, t.c.AccessToken)
}

func (t *tenant) ArchiveContact(contactId string) (err error) {
	err = t.c.Refresh()
	if err != nil {
		return
	}
	t.Call()
	return contacts.ArchiveContact(contactId, t.TenantID, t.c.AccessToken)
}

// Users

func (t *tenant) GetUsers(where *filter.Filter) (userList []users.User, err error) {
	err = t.c.Refresh()
	if err != nil {
		return
	}
	t.Call()
	return users.GetUsers(t.TenantID, t.c.AccessToken, where)
}

func (t *tenant) GetUser(userId string) (user users.User, err error) {
	err = t.c.Refresh()
	if err != nil {
		return
	}
	t.Call()
	return users.GetUser(t.TenantID, t.c.AccessToken, userId)
}

// Template For Tenant Funcs (bc lazy)
func (t *tenant) _() (err error) {
	err = t.c.Refresh()
	if err != nil {
		return
	}
	t.Call()
	return
}
