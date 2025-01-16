package invoices

import (
	"github.com/datag8r/xerogo/accountingAPI/contacts"
	"github.com/datag8r/xerogo/accountingAPI/currencies"
)

type invoiceType string

type invoiceStatus string

type lineAmountType string

type taxCalculationtype string

type Tracking struct {
	Name               string
	TrackingCategoryID string
	Option             string
}

type Item struct {
	ItemID string
	Name   string
	Code   string
}

type invoiceForCreation struct {
	Type      invoiceType       // required
	Contact   contacts.Contact  // required// only needs id field
	LineItems []InvoiceLineItem // required

	Reference           *string                  `json:",omitempty"` // optional // ACCREC only
	DateString          *string                  `json:",omitempty"` // optional
	DueDateString       *string                  `json:",omitempty"` // optional
	Status              *invoiceStatus           `json:",omitempty"` // optional
	LineAmountTypes     *lineAmountType          `json:",omitempty"` // optional // default exclusive
	InvoiceNumber       *string                  `json:",omitempty"` // optional
	CurrencyCode        *currencies.CurrencyCode `json:",omitempty"` // optional
	CurrencyRate        *string                  `json:",omitempty"` // optional
	BrandingThemeID     *string                  `json:",omitempty"` // optional
	Url                 *string                  `json:",omitempty"` // optional
	SentToContact       *bool                    `json:",omitempty"` // optional
	ExpectedPaymentDate *string                  `json:",omitempty"` // optional
	PlannedPaymentDate  *string                  `json:",omitempty"` // optional
}

type invoiceForUpdating struct {
	Contact   contacts.Contact  // required // only needs id field
	LineItems []InvoiceLineItem // required

	Reference           *string                  `json:",omitempty"` // optional // ACCREC only
	DateString          *string                  `json:",omitempty"` // optional // YYYY-MM-DD
	DueDateString       *string                  `json:",omitempty"` // optional // YYYY-MM-DD
	Status              *invoiceStatus           `json:",omitempty"` // optional
	LineAmountTypes     *lineAmountType          `json:",omitempty"` // optional
	InvoiceNumber       *string                  `json:",omitempty"` // optional
	CurrencyCode        *currencies.CurrencyCode `json:",omitempty"` // optional
	CurrencyRate        *string                  `json:",omitempty"` // optional // max len = [18].[6] // default pulled from xe.com day rates
	BrandingThemeID     *string                  `json:",omitempty"` // optional
	Url                 *string                  `json:",omitempty"` // optional
	SentToContact       *bool                    `json:",omitempty"` // optional
	ExpectedPaymentDate *string                  `json:",omitempty"` // optional
	PlannedPaymentDate  *string                  `json:",omitempty"` // optional
}
