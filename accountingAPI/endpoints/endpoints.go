package endpoints

const (
	accountingAPIBase        string = "https://api.xero.com/api.xro/2.0"
	EndpointItems                   = accountingAPIBase + "/Items"
	EndpointUsers                   = accountingAPIBase + "/Users"
	EndpointAccounts                = accountingAPIBase + "/Accounts"
	EndpointInvoices                = accountingAPIBase + "/Invoices"
	EndpointContacts                = accountingAPIBase + "/Contacts"
	EndpointCurrencies              = accountingAPIBase + "/Currencies"
	EndpointContactGroups           = accountingAPIBase + "/ContactGroups"
	EndpointBankTransfers           = accountingAPIBase + "/BankTransfers"
	EndpointBankTransactions        = accountingAPIBase + "/BankTransactions"
)
