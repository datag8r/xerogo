package items

type PurchaseDetails struct {
	UnitPrice       float64
	AccountCode     string
	TaxType         string  // Will be taxType when I make them
	COGSAccountCode *string // TODO
}

type SalesDetails struct {
	UnitPrice   float64
	AccountCode string
	TaxType     string // Will be taxType when I make them
}

type itemForCreate struct {
	Code                      string
	InventoryAssetAccountCode *string `json:",omitempty"`
	Name                      string  `json:",omitempty"`
	IsSold                    bool
	IsPurchased               bool
	Description               string           `json:",omitempty"`
	PurchaseDescription       string           `json:",omitempty"`
	PurchaseDetails           *PurchaseDetails `json:",omitempty"`
	SalesDetails              *SalesDetails    `json:",omitempty"`
}

type itemForUpdate struct {
	ItemID                    string `json:",omitempty"`
	Code                      string
	InventoryAssetAccountCode *string `json:",omitempty"`
	Name                      string  `json:",omitempty"`
	IsSold                    bool
	IsPurchased               bool
	Description               string           `json:",omitempty"`
	PurchaseDescription       string           `json:",omitempty"`
	PurchaseDetails           *PurchaseDetails `json:",omitempty"`
	SalesDetails              *SalesDetails    `json:",omitempty"`
}
