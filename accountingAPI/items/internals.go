package items

func (i Item) validForCreation() bool {
	if i.Code == "" {
		return false
	}
	if i.IsTrackedAsInventory &&
		(i.InventoryAssetAccountCode == nil || *i.InventoryAssetAccountCode == "") {
		return false
	}
	return true
}

func (i Item) validForUpdating() bool {
	if i.Code == "" && i.ItemID == "" {
		return false
	}
	if i.IsTrackedAsInventory &&
		(i.InventoryAssetAccountCode == nil || *i.InventoryAssetAccountCode == "") {
		return false
	}
	return true
}

func (i Item) toCreate() itemForCreate {
	return itemForCreate{
		Code:                      i.Code,
		InventoryAssetAccountCode: i.InventoryAssetAccountCode,
		Name:                      i.Name,
		IsSold:                    i.IsSold,
		IsPurchased:               i.IsPurchased,
		Description:               i.Description,
		PurchaseDescription:       i.PurchaseDescription,
		PurchaseDetails:           i.PurchaseDetails,
		SalesDetails:              i.SalesDetails,
	}
}

func (i Item) toUpdate() itemForUpdate {
	return itemForUpdate{
		ItemID:                    i.ItemID,
		Code:                      i.Code,
		InventoryAssetAccountCode: i.InventoryAssetAccountCode,
		Name:                      i.Name,
		IsSold:                    i.IsSold,
		IsPurchased:               i.IsPurchased,
		Description:               i.Description,
		PurchaseDescription:       i.PurchaseDescription,
		PurchaseDetails:           i.PurchaseDetails,
		SalesDetails:              i.SalesDetails,
	}
}
