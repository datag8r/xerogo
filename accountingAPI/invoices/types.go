package invoices

type invoiceType string

type invoiceStatus string

type lineAmountType string

type taxCalculationtype string

type Item struct {
	ItemID string
	Name   string
	Code   string
}
