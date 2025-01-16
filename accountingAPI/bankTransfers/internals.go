package banktransfers

func (b BankTransfer) toCreate() (c bankTransferForCreate) {
	c.FromBankAccount = b.FromBankAccount
	c.ToBankAccount = b.ToBankAccount
	c.Amount = b.Amount
	c.Date = b.Date
	c.FromIsReconciled = b.FromIsReconciled
	c.ToIsReconciled = b.ToIsReconciled
	c.Reference = b.Reference
	return
}

func (b BankTransfer) validForCreation() bool {
	if b.FromBankAccount.Code == "" && b.FromBankAccount.AccountID == "" {
		return false
	}
	if b.ToBankAccount.Code == "" && b.ToBankAccount.AccountID == "" {
		return false
	}
	if b.Amount == 0.0 {
		return false
	}
	return true
}
