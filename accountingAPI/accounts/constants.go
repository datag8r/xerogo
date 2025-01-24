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

	SystemAccountTypeAccountsReceivable      systemAccountType = "DEBTORS"
	SystemAccountTypeAccountsPayable         systemAccountType = "CREDITORS"
	SystemAccountTypeBankRevaluations        systemAccountType = "BANKCURRENCYGAIN"
	SystemAccountTypeCISAssets               systemAccountType = "CISASSETS"        // UK Only
	SystemAccountTypeCISLabourExpense        systemAccountType = "CISLABOUREXPENSE" // UK Only
	SystemAccountTypeCISLabourIncome         systemAccountType = "CISLABOURINCOME"  // UK Only
	SystemAccountTypeCISLiability            systemAccountType = "CISLIABILITY"     // UK Only
	SystemAccountTypeCISMaterials            systemAccountType = "CISMATERIALS"     // UK Only
	SystemAccountTypeGST                     systemAccountType = "GST"
	SystemAccountTypeGSTOnImports            systemAccountType = "GSTONIMPORTS"
	SystemAccountTypeHistoricalAdjustment    systemAccountType = "HISTORICAL"
	SystemAccountTypeRealisedCurrencyGains   systemAccountType = "REALISEDCURRENCYGAIN"
	SystemAccountTypeRetainedEarnings        systemAccountType = "RETAINEDEARNINGS"
	SystemAccountTypeRounding                systemAccountType = "ROUNDING"
	SystemAccountTypeTrackingTransfers       systemAccountType = "TRACKINGTRANSFERS"
	SystemAccountTypeUnpaidExpenseClaims     systemAccountType = "UNPAIDEXPCLM"
	SystemAccountTypeUnrealisedCurrencyGains systemAccountType = "UNREALISEDCURRENCYGAIN"
	SystemAccountTypeWagesPayable            systemAccountType = "WAGEPAYABLES"
)
