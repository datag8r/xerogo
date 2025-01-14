package accounts

func (a Account) toUpdate() accountForUpdate {
	return accountForUpdate{
		Code:              a.Code,
		Name:              a.Name,
		Type:              a.Type,
		BankAccountNumber: a.BankAccountNumber,
		Description:       a.Description,
		BankAccountType:   a.BankAccountType,
		CurrencyCode:      a.CurrencyCode,
		AccountID:         a.AccountID,
		Class:             a.Class,
		SystemAccount:     a.SystemAccount,
		ReportingCode:     a.ReportingCode,
		ReportingCodeName: a.ReportingCodeName,
		HasAttachments:    a.HasAttachments,
		UpdatedDateUTC:    a.UpdatedDateUTC,
		AddToWatchlist:    a.AddToWatchlist,
	}
}

func (a Account) toCreate() accountForCreate {
	return accountForCreate{
		Code:              a.Code,
		Name:              a.Name,
		Type:              a.Type,
		BankAccountNumber: a.BankAccountNumber,
	}
}

func (a Account) validForCreation() bool {
	// validate Code
	if a.Code == "" || len(a.Code) > 10 {
		return false
	}
	// validate Name
	if a.Name == "" || len(a.Name) > 150 {
		return false
	}
	// validate Type
	if !validateAccountType(a.Type) {
		return false
	}
	if a.Type == AccountTypeBank {
		if a.BankAccountNumber == nil || len(*a.BankAccountNumber) == 0 {
			return false
		}
	}
	return true
}

func (a Account) validForUpdate() bool {
	// validate AccountID
	return a.AccountID != ""
}
