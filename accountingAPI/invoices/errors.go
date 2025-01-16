package invoices

import "errors"

var (
	ErrInvalidInvoiceID          = errors.New("invalid invoice id")
	ErrInvalidInvoiceForCreation = errors.New("invalid invoice for creation, check: line items, contact and type fields")
	ErrInvalidInvoiceForUpdating = errors.New("invalid invoice for updating, check: invoice id")
)
