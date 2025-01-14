package accounts

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
