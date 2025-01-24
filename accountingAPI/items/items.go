package items

import (
	"github.com/datag8r/xerogo/accountingAPI/endpoints"
	"github.com/datag8r/xerogo/filter"
	"github.com/datag8r/xerogo/helpers"
	"github.com/datag8r/xerogo/utils"
)

type Item struct {
	ItemID                    string `xero:"*id"`
	Code                      string `xero:"*id"`
	Name                      string `xero:"*id"`
	IsSold                    bool
	IsPurchased               bool
	Description               string
	PurchaseDescription       string
	PurchaseDetails           *PurchaseDetails
	SalesDetails              *SalesDetails
	InventoryAssetAccountCode *string // Account must be of type INVENTORY
	IsTrackedAsInventory      bool    // True for items that are tracked as inventory. An item will be tracked as inventory if the InventoryAssetAccountCode and COGSAccountCode are set.
	TotalCostPool             float64
	QuantityOnHand            float64
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
	var responseBody struct {
		Items []Item
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	items = responseBody.Items
	return
}

func GetItem(tenantId, accessToken, itemIdOrCode string) (item Item, err error) {
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
	var responseBody struct {
		Items []Item
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	if len(responseBody.Items) == 1 {
		item = responseBody.Items[0]
	}
	return
}

func CreateItem(tenantId, accessToken string, itemToCreate Item) (item Item, err error) {
	url := endpoints.EndpointItems
	inter, err := utils.XeroCustomMarshal(itemToCreate, "create")
	if err != nil {
		return
	}
	buf, err := helpers.MarshallJsonToBuffer(inter)
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

func CreateItems(tenantId, accessToken string, itemsToCreate []Item) (items []Item, err error) {
	url := endpoints.EndpointItems
	inter, err := utils.XeroCustomMarshal(itemsToCreate, "create")
	if err != nil {
		return
	}
	mp := map[string]interface{}{"Items": inter}
	buf, err := helpers.MarshallJsonToBuffer(mp)
	if err != nil {
		return
	}
	request, err := helpers.BuildRequest("POST", url, nil, nil, buf)
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
	items = responseBody.Items
	return
}

func UpdateItem(tenantId, accessToken string, item Item) (err error) {
	url := endpoints.EndpointItems
	inter, err := utils.XeroCustomMarshal(item, "update")
	if err != nil {
		return
	}
	buf, err := helpers.MarshallJsonToBuffer(inter)
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

func UpdateItems(tenantId, accessToken string, items []Item) (err error) {
	url := endpoints.EndpointItems
	inter, err := utils.XeroCustomMarshal(items, "update")
	if err != nil {
		return
	}
	mp := map[string]interface{}{"Items": inter}
	buf, err := helpers.MarshallJsonToBuffer(mp)
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

func DeleteItem(tenantId, accessToken, itemId string) (err error) {
	url := endpoints.EndpointItems + "/" + itemId
	request, err := helpers.BuildRequest("DELETE", url, nil, nil, nil)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	_, err = helpers.DoRequest(request, 200)
	return
}
