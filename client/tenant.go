package client

import (
	"github.com/datag8r/xerogo/accountingAPI/accounts"
	"github.com/datag8r/xerogo/filter"
)

type Tenant struct {
	c              *client `json:"-"`
	TenantID       string  `json:"tenantId"`
	Name           string  `json:"tenantName"`
	CreatedDateUTC string  `json:"createdDateUtc"`
	UpdatedDateUTC string  `json:"updatedDateUtc"`
}

// Accounts

func (t *Tenant) GetAccounts(where *filter.Filter) (accs []accounts.Account, err error) {
	err = t.c.Refresh()
	if err != nil {
		return
	}
	return accounts.GetAccounts(t.TenantID, t.c.AccessToken, where)
}

func (t *Tenant) GetAccount(accountID string) (acc accounts.Account, err error) {
	err = t.c.Refresh()
	if err != nil {
		return
	}
	return accounts.GetAccount(accountID, t.TenantID, t.c.AccessToken)
}

func (t *Tenant) CreateAccount(account accounts.Account) (acc accounts.Account, err error) {
	err = t.c.Refresh()
	if err != nil {
		return
	}
	return accounts.CreateAccount(account, t.TenantID, t.c.AccessToken)
}

func (t *Tenant) UpdateAccount(account accounts.Account) (err error) {
	err = t.c.Refresh()
	if err != nil {
		return
	}
	return accounts.UpdateAccount(account, t.TenantID, t.c.AccessToken)
}
