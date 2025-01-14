package accounts

type accountStatusCode string
type accountClassType string
type accountType string

type bankAccountType string
type systemAccountType string

var (
	allAccountTypes = []accountType{
		AccountTypeBank,
		AccountTypeCurrent,
		AccountTypeCurrentLiability,
		AccountTypeDepreciation,
		AccountTypeDirectCosts,
		AccountTypeEquity,
		AccountTypeExpense,
		AccountTypeFixedAsset,
		AccountTypeInventoryAsset,
		AccountTypeLiability,
		AccountTypeNonCurrentAsset,
		AccountTypeOtherIncome,
		AccountTypeOverHeads,
		AccountTypePrepayment,
		AccountTypeRevenue,
		AccountTypeSales,
		AccountTypeNonCurrentLiability,
	}
)

func validateAccountType(accType accountType) bool {
	for _, aT := range allAccountTypes {
		if aT == accType {
			return true
		}
	}
	return false
}

type accountForCreate struct {
	Code              string      // Required For Creation // maxlen = 10
	Name              string      // Required For Creation // maxlen = 150
	Type              accountType // Required For Creation
	BankAccountNumber *string     `json:",omitempty"` // Required For Creation If Type == Bank
}

type accountForUpdate struct {
	Code              string           // Required For Creation // maxlen = 10
	Name              string           // Required For Creation // maxlen = 150
	Type              accountType      // Required For Creation
	BankAccountNumber *string          `json:",omitempty"` // Required For Creation If Type == Bank
	Description       *string          `json:",omitempty"`
	BankAccountType   *bankAccountType `json:",omitempty"`
	CurrencyCode      *string          `json:",omitempty"`
	AccountID         string
	Class             accountClassType
	SystemAccount     *systemAccountType `json:",omitempty"`
	ReportingCode     string
	ReportingCodeName string
	HasAttachments    bool
	UpdatedDateUTC    string
	AddToWatchlist    bool
}
