package items

type PurchaseDetails struct {
	UnitPrice       float64 `xero:"create,update"`
	AccountCode     string  `xero:"create,update"`
	TaxType         string  `xero:"*create,*update"` // Will be taxType when I make them
	COGSAccountCode string  `xero:"*create,*update"`
}

type SalesDetails struct {
	UnitPrice   float64 `xero:"create,update"`
	AccountCode string  `xero:"create,update"`
	TaxType     string  `xero:"*create,*update"` // Will be taxType when I make them
}
