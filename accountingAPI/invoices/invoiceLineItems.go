package invoices

import (
	"github.com/datag8r/xerogo/accountingAPI/types"
)

type InvoiceLineItem struct {
	Description    string
	Quantity       float64 `json:",string"`
	UnitAmount     float64 `json:",string"`
	ItemCode       string
	AccountCode    string
	Item           Item
	LineItemID     string
	TaxType        types.TaxType
	TaxAmount      float64 `json:",string"`
	LineAmount     float64 `json:",string"`
	DiscountRate   float64 `json:",string"` // percentage
	DiscountAmount float64 `json:",string"` // amount
	Tracking       []Tracking
}
