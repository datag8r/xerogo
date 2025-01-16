package banktransfers

type bankTransferForCreate struct {
	FromBankAccount  bankAccountDetails // required for creation
	ToBankAccount    bankAccountDetails // required for creation
	Amount           float64            `json:",string"`    // required for creation
	Date             string             `json:",omitempty"` // optional for creation
	FromIsReconciled *bool              `json:",omitempty"` // optional for creation // default false probs
	ToIsReconciled   *bool              `json:",omitempty"` // optional for creation // default false probs
	Reference        string             `json:",omitempty"` // optional for creation
}

type bankAccountDetails struct {
	Code      string // this or accountId required to create bank transfer
	AccountID string // this or code required to create bank transfer
	Name      string // not required
}
