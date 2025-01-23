package accounts

import (
	"github.com/datag8r/xerogo/accountingAPI/endpoints"
	"github.com/datag8r/xerogo/filter"
	"github.com/datag8r/xerogo/helpers"
	"github.com/datag8r/xerogo/utils"
)

type Account struct {
	Code                    string          `xero:"create,*update,*id"` // Required For Creation // maxlen = 10
	Name                    string          `xero:"create,*update"`     // Required For Creation // maxlen = 150
	Type                    accountType     `xero:"create,*update"`     // Required For Creation
	BankAccountNumber       string          `xero:"*create,*update"`    // Required For Creation if Type == BANK
	Description             string          `xero:"*create,*update"`
	BankAccountType         bankAccountType `xero:"*create,*update"`
	CurrencyCode            string          `xero:"*create,*update"`
	EnablePaymentsToAccount bool            `xero:"*create,*update"`
	ShowInExpenseClaims     bool            `xero:"*create,*update"`
	AddToWatchlist          bool            `xero:"*create,*update"`
	AccountID               string          `xero:"update,*id"`
	TaxType                 string
	Status                  accountStatusCode
	Class                   accountClassType
	SystemAccount           systemAccountType
	ReportingCode           string
	ReportingCodeName       string
	HasAttachments          bool
	UpdatedDateUTC          string
}

func GetAccounts(tenantId, accessToken string, where *filter.Filter) (accounts []Account, err error) {
	url := endpoints.EndpointAccounts
	request, err := helpers.BuildRequest("GET", url, nil, where, nil)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	body, err := helpers.DoRequest(request, 200)
	if err != nil {
		return
	}
	var responseBody struct {
		Accounts []Account
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	accounts = responseBody.Accounts
	return
}

func GetAccount(accountID string, tenantId, accessToken string) (acc Account, err error) {
	url := endpoints.EndpointAccounts + "/" + accountID
	request, err := helpers.BuildRequest("GET", url, nil, nil, nil)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	b, err := helpers.DoRequest(request, 200)
	if err != nil {
		return
	}
	var responseBody struct {
		Accounts []Account
	}
	err = helpers.UnmarshalJson(b, &responseBody)
	if len(responseBody.Accounts) == 1 {
		acc = responseBody.Accounts[0]
	}
	return
}

func CreateAccount(account Account, tenantId string, accessToken string) (acc Account, err error) {
	url := endpoints.EndpointAccounts
	inter, err := utils.XeroCustomMarshal(account, "create")
	if err != nil {
		return
	}
	buf, err := helpers.MarshallJsonToBuffer(inter)
	if err != nil {
		return
	}
	request, err := helpers.BuildRequest("PUT", url, nil, nil, buf)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	body, err := helpers.DoRequest(request, 200)
	if err != nil {
		return
	}
	var responseBody struct {
		Accounts []Account
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	if err != nil {
		return
	}
	if len(responseBody.Accounts) == 1 {
		acc = responseBody.Accounts[0]
	}
	return
}

func UpdateAccount(account Account, tenantId string, accessToken string) (err error) {
	url := endpoints.EndpointAccounts + "/" + account.AccountID
	inter, err := utils.XeroCustomMarshal(account, "update")
	if err != nil {
		return
	}
	buf, err := helpers.MarshallJsonToBuffer(inter)
	if err != nil {
		return
	}
	request, err := helpers.BuildRequest("POST", url, nil, nil, buf)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	_, err = helpers.DoRequest(request, 200)
	return
}

func ArchiveAccount(accountID string, tenantId, accessToken string) (err error) {
	url := endpoints.EndpointAccounts + "/" + accountID
	var requestBody struct {
		AccountID string
		Status    string
	}
	requestBody.AccountID = accountID
	requestBody.Status = string(AccountStatusCodeArchived)
	buf, err := helpers.MarshallJsonToBuffer(requestBody)
	if err != nil {
		return
	}
	request, err := helpers.BuildRequest("PUT", url, nil, nil, buf)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	_, err = helpers.DoRequest(request, 200)
	return
}

// System accounts and accounts used on transactions can not be deleted using the delete method.
// If an account is not able to be deleted you can update the status to ARCHIVED using the accounts.ArchiveAccount Function
func DeleteAccount(accountID string, tenantId, accessToken string) (err error) {
	url := endpoints.EndpointAccounts + "/" + accountID
	request, err := helpers.BuildRequest("DELETE", url, nil, nil, nil)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	_, err = helpers.DoRequest(request, 200)
	return
}
