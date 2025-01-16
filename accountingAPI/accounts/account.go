package accounts

import (
	"github.com/datag8r/xerogo/accountingAPI/endpoints"
	"github.com/datag8r/xerogo/filter"
	"github.com/datag8r/xerogo/helpers"
)

type Account struct {
	Code                    string            // Required For Creation // maxlen = 10
	Name                    string            // Required For Creation // maxlen = 150
	Type                    accountType       // Required For Creation
	BankAccountNumber       *string           `json:",omitempty"` // Required For Creation If Type == Bank
	Status                  accountStatusCode `json:"-"`
	Description             *string           `json:",omitempty"`
	BankAccountType         *bankAccountType  `json:",omitempty"`
	CurrencyCode            *string           `json:",omitempty"`
	TaxType                 string            // Will Be a taxType type when i make them
	EnablePaymentsToAccount bool
	ShowInExpenseClaims     bool
	AccountID               string `json:",omitempty"`
	Class                   accountClassType
	SystemAccount           *systemAccountType `json:",omitempty"`
	ReportingCode           string             `json:",omitempty"`
	ReportingCodeName       string             `json:",omitempty"`
	HasAttachments          bool
	UpdatedDateUTC          string
	AddToWatchlist          bool
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
	if accountID == "" {
		err = ErrInvalidAccountID
		return
	}
	url := endpoints.EndpointAccounts + "/" + accountID
	request, err := helpers.BuildRequest("GET", url, nil, nil, nil)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	body, err := helpers.DoRequest(request, 200)
	if err != nil {
		return
	}
	err = helpers.UnmarshalJson(body, &acc)
	return
}

func CreateAccount(account Account, tenantId string, accessToken string) (acc Account, err error) {
	url := endpoints.EndpointAccounts
	if !account.validForCreation() {
		err = ErrInvalidAccountForCreation
		return
	}
	buf, err := helpers.MarshallJsonToBuffer(account.toCreate())
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
	url := endpoints.EndpointAccounts
	if !account.validForUpdate() {
		err = ErrInvalidAccountForUpdating
		return
	}
	buf, err := helpers.MarshallJsonToBuffer(account.toUpdate())
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

func ArchiveAccount(accountID string, tenantId, accessToken string) (err error) {
	if accountID == "" {
		err = ErrInvalidAccountID
		return
	}
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
	if accountID == "" {
		err = ErrInvalidAccountID
		return
	}
	url := endpoints.EndpointAccounts + "/" + accountID
	request, err := helpers.BuildRequest("DELETE", url, nil, nil, nil)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	_, err = helpers.DoRequest(request, 200)
	return
}
