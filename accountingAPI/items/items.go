package items

import (
	"github.com/datag8r/xerogo/accountingAPI/endpoints"
	"github.com/datag8r/xerogo/filter"
	"github.com/datag8r/xerogo/helpers"
)

type Item struct {
	ItemID                    string
	Code                      string
	Name                      string `json:",omitempty"`
	IsSold                    bool
	IsPurchased               bool
	Description               string
	PurchaseDescription       string
	PurchaseDetails           *PurchaseDetails `json:",omitempty"`
	SalesDetails              *SalesDetails    `json:",omitempty"`
	InventoryAssetAccountCode *string          `json:",omitempty"` // Account must be of type INVENTORY
	IsTrackedAsInventory      bool             // True for items that are tracked as inventory. An item will be tracked as inventory if the InventoryAssetAccountCode and COGSAccountCode are set.
	UpdatedDateUTC            string
}

func GetItems(tenantId, accessToken string, where *filter.Filter) (items []Item, err error) {
	url := endpoints.EndpointItems
	request, err := helpers.BuildRequest("GET", url, nil, where, nil)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	body, err := helpers.DoRequest(request, 200)
	if err != nil {
		return
	}
	var bod struct {
		Items []Item
	}
	err = helpers.UnmarshalJson(body, &bod)
	items = bod.Items
	return
}

func GetItem(itemIdOrCode, tenantId, accessToken string) (item Item, err error) {
	url := endpoints.EndpointItems + "/" + itemIdOrCode
	request, err := helpers.BuildRequest("GET", url, nil, nil, nil)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	body, err := helpers.DoRequest(request, 200)
	if err != nil {
		return
	}
	var bod struct {
		Items []Item
	}
	err = helpers.UnmarshalJson(body, &bod)
	if len(bod.Items) == 1 {
		item = bod.Items[0]
	}
	return
}

func CreateItem(itemToCreate Item, tenantId, accessToken string) (item Item, err error) {
	url := endpoints.EndpointItems
	if !itemToCreate.validForCreation() {
		err = ErrInvalidItemForCreation
		return
	}
	buf, err := helpers.MarshallJsonToBuffer(itemToCreate.toCreate())
	if err != nil {
		return
	}
	request, err := helpers.BuildRequest("PUT", url, nil, nil, buf)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	body, err := helpers.DoRequest(request, 200)
	if err != nil {
		return
	}
	var responseBody struct {
		Items []Item
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	if err != nil {
		return
	}
	if len(responseBody.Items) == 1 {
		item = responseBody.Items[0]
	}
	return
}

func UpdateItem(item Item, tenantId, accessToken string) (err error) {
	url := endpoints.EndpointItems
	if !item.validForUpdating() {
		err = ErrInvalidItemForUpdating
		return
	}
	buf, err := helpers.MarshallJsonToBuffer(item.toUpdate())
	if err != nil {
		return
	}
	request, err := helpers.BuildRequest("POST", url, nil, nil, buf)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	_, err = helpers.DoRequest(request, 200)
	return
}

func DeleteItem(itemID, tenantId, accessToken string) (err error) {
	if itemID == "" {
		return ErrInvalidItemID
	}
	url := endpoints.EndpointItems + "/" + itemID
	request, err := helpers.BuildRequest("DELETE", url, nil, nil, nil)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	_, err = helpers.DoRequest(request, 200)
	return
}
