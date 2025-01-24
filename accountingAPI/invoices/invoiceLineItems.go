package invoices

import (
	"github.com/datag8r/xerogo/accountingAPI/items"
	trackingcategories "github.com/datag8r/xerogo/accountingAPI/trackingCategories"
	"github.com/datag8r/xerogo/accountingAPI/types"
)

type InvoiceLineItem struct {
	Description    string                        `xero:"*create,*update"`
	ItemCode       string                        `xero:"*create,*update"`
	AccountCode    string                        `xero:"*create,*update"`
	LineItemID     string                        `xero:"*create,*update"`
	TaxType        types.TaxType                 `xero:"*create,*update"`
	Tracking       []trackingcategories.Tracking `xero:"*create,*update"`
	Quantity       float64                       `json:",string" xero:"*create,*update"`
	UnitAmount     float64                       `json:",string" xero:"*create,*update"`
	TaxAmount      float64                       `json:",string" xero:"*create,*update"`
	LineAmount     float64                       `json:",string" xero:"*create,*update"`
	DiscountRate   float64                       `json:",string" xero:"*create,*update"` // percentage
	DiscountAmount float64                       `json:",string" xero:"*create,*update"` // amount
	Item           items.Item                    `xero:"*create,*update,embeddedId"`
}
