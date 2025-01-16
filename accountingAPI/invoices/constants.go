package invoices

const (
	InvoiceStatusDraft      invoiceStatus = "DRAFT"
	InvoiceStatusSubmitted  invoiceStatus = "SUBMITTED"
	InvoiceStatusDeleted    invoiceStatus = "DELETED"
	InvoiceStatusAuthorised invoiceStatus = "AUTHORISED"
	InvoiceStatusPaid       invoiceStatus = "PAID"
	InvoiceStatusVoided     invoiceStatus = "VOIDED"
)

const (
	InvoiceTypeAccountsPayable    invoiceType = "ACCPAY" // Bill
	InvoiceTypeAccountsReceivable invoiceType = "ACCREC" // Sales Invoice
)

const (
	LineAmountTypeExclusive lineAmountType = "Exclusive"
	LineAmountTypeInclusive lineAmountType = "Inclusive"
	LineAmountTypeNoTax     lineAmountType = "NoTax"
)

const (
	TaxCalculationTypeTaxCalc taxCalculationtype = "TAXCALC"
	TaxCalculationTypeAuto    taxCalculationtype = "AUTO" // ? shows as "TAXCALC/AUTO" in docs
)
