package invoices

func (i Invoice) toCreate() invoiceForCreation {
	c := invoiceForCreation{
		Type:      i.Type,
		Contact:   i.Contact,
		LineItems: i.LineItems,
	}
	if i.Reference != "" && i.Type == InvoiceTypeAccountsReceivable {
		c.Reference = &i.Reference
	}
	if i.DateString != "" {
		c.DateString = &i.DateString
	}
	if i.DueDate != "" {
		c.DueDateString = &i.DueDateString
	}
	if i.Status != "" {
		c.Status = &i.Status
	}
	if i.LineAmountTypes != "" {
		c.LineAmountTypes = &i.LineAmountTypes
	}
	if i.InvoiceNumber != "" {
		c.InvoiceNumber = &i.InvoiceNumber
	}
	if i.CurrencyCode != "" {
		c.CurrencyCode = &i.CurrencyCode
	}
	if i.CurrencyRate != "" {
		c.CurrencyRate = &i.CurrencyRate
	}
	if i.BrandingThemeID != "" {
		c.BrandingThemeID = &i.BrandingThemeID
	}
	if i.Url != "" {
		c.Url = &i.Url
	}
	if i.SentToContact {
		c.SentToContact = &i.SentToContact
	}
	if i.ExpectedPaymentDate != "" {
		c.ExpectedPaymentDate = &i.ExpectedPaymentDate
	}
	if i.PlannedPaymentDate != "" {
		c.PlannedPaymentDate = &i.PlannedPaymentDate
	}
	return c
}

func (i Invoice) toUpdate() invoiceForUpdating {
	c := invoiceForUpdating{
		Contact:   i.Contact,
		LineItems: i.LineItems,
	}
	if i.Reference != "" && i.Type == InvoiceTypeAccountsReceivable {
		c.Reference = &i.Reference
	}
	if i.DateString != "" {
		c.DateString = &i.DateString
	}
	if i.DueDateString != "" {
		c.DueDateString = &i.DueDateString
	}
	if i.Status != "" {
		c.Status = &i.Status
	}
	if i.LineAmountTypes != "" {
		c.LineAmountTypes = &i.LineAmountTypes
	}
	if i.InvoiceNumber != "" {
		c.InvoiceNumber = &i.InvoiceNumber
	}
	if i.CurrencyCode != "" {
		c.CurrencyCode = &i.CurrencyCode
	}
	if i.CurrencyRate != "" {
		c.CurrencyRate = &i.CurrencyRate
	}
	if i.BrandingThemeID != "" {
		c.BrandingThemeID = &i.BrandingThemeID
	}
	if i.Url != "" {
		c.Url = &i.Url
	}
	if i.SentToContact {
		c.SentToContact = &i.SentToContact
	}
	if i.ExpectedPaymentDate != "" {
		c.ExpectedPaymentDate = &i.ExpectedPaymentDate
	}
	if i.PlannedPaymentDate != "" {
		c.PlannedPaymentDate = &i.PlannedPaymentDate
	}
	return c
}

func (i Invoice) validForCreation() bool {
	if i.Type == "" || len(i.LineItems) == 0 || i.Contact.IsZero() {
		return false
	}
	return true
}

func (i Invoice) validForUpdating() bool {
	if i.InvoiceID == "" {
		return false
	}
	if i.Type == InvoiceTypeAccountsPayable && i.Status == InvoiceStatusPaid {
		return false
	}
	return true
}
