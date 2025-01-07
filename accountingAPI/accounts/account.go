package accounts

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/datag8r/xerogo/accountingAPI/endpoints"
)

type Account struct {
	Code                    string
	Name                    string
	Type                    accountType
	BankAccountNumber       *string
	Status                  accountStatusCode
	Description             *string
	BankAccountType         *bankAccountType
	CurrencyCode            *string
	TaxType                 string // Will Be a taxType type when i make them
	EnablePaymentsToAccount bool
	ShowInExpenseClaims     bool
	AccountID               string
	Class                   accountClassType
	SystemAccount           *systemAccountType
	ReportingCode           string
	ReportingCodeName       string
	HasAttachments          bool
	UpdatedDateUTC          string
	AddToWatchlist          bool
}

type accountStatusCode = string
type accountClassType = string
type accountType = string

type bankAccountType = string
type systemAccountType = string

// source: https://developer.xero.com/documentation/api/accounting/types#accounts
const (
	AccountClassAsset     accountClassType = "ASSET"
	AccountClassEquity    accountClassType = "EQUITY"
	AccountClassExpense   accountClassType = "EXPENSE"
	AccountClassLiability accountClassType = "LIABILITY"
	AccountClassRevenue   accountClassType = "REVENUE"

	AccountTypeBank                accountType = "BANK"
	AccountTypeCurrent             accountType = "CURRENT"
	AccountTypeCurrentLiability    accountType = "CURRLIAB"
	AccountTypeDepreciation        accountType = "DEPRECIATN"
	AccountTypeDirectCosts         accountType = "DIRECTCOSTS"
	AccountTypeEquity              accountType = "EQUITY"
	AccountTypeExpense             accountType = "EXPENSE"
	AccountTypeFixedAsset          accountType = "FIXED"
	AccountTypeInventoryAsset      accountType = "INVENTORY"
	AccountTypeLiability           accountType = "LIABILITY"
	AccountTypeNonCurrentAsset     accountType = "NONCURRENT"
	AccountTypeOtherIncome         accountType = "OTHERINCOME"
	AccountTypeOverHeads           accountType = "OVERHEADS"
	AccountTypePrepayment          accountType = "PREPAYMENT"
	AccountTypeRevenue             accountType = "REVENUE"
	AccountTypeSales               accountType = "SALES"
	AccountTypeNonCurrentLiability accountType = "TERMLIAB"

	AccountStatusCodeActive   accountStatusCode = "ACTIVE"
	AccountStatusCodeArchived accountStatusCode = "ARCHIVED"

	BankAccountTypeBank       bankAccountType = "BANK"
	BankAccountTypeCreditCard bankAccountType = "CREDITCARD"
	BankAccountTypePaypal     bankAccountType = "PAYPAL"

	SystemAccountTypeAccountsReceivable systemAccountType = "DEBTORS"
	// More of these
)

// TODO: add filters for where clause, modified after, account ID, orderBy stuff also
func GetAccounts(tenantID string, accessToken string) (accounts []Account, err error) {
	url := endpoints.EndpointAccounts
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	request.Header.Add("Authorization", "Bearer "+accessToken)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("xero-tenant-id", tenantID)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	type accountResponseBody struct {
		Accounts []Account
	}
	var body accountResponseBody
	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &body)
	accounts = body.Accounts
	return
}
