package items

import (

	// 	"github.com/datag8r/xerogo/accountingAPI/endpoints"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/datag8r/xerogo/accountingAPI/endpoints"
	"github.com/datag8r/xerogo/filter"
	"github.com/datag8r/xerogo/utils"
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

type ItemHistory struct {
	Changes       string
	DateUTCString string
	DateUTC       string
	User          string
	Details       string
}

var (
	ErrInvalidItemForCreation = errors.New("some fields are invalid to create an item")
	ErrInvalidItemForUpdating = errors.New("some fields are invalid to update an item" +
		" - usually this is the code field")
	ErrInvalidItemID = errors.New("invalid item id for request")
)

func GetItems(tenantID, accessToken string, where *filter.Filter) (items []Item, err error) {
	url := endpoints.EndpointItems
	var request *http.Request
	if where != nil {
		request, err = where.BuildRequest("GET", url, nil)
	} else {
		request, err = http.NewRequest("GET", url, nil)
	}
	if err != nil {
		return
	}
	utils.AddXeroHeaders(request, accessToken, tenantID)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	var bod struct {
		Items []Item
	}
	defer response.Body.Close()
	b, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	if response.StatusCode != 200 {
		err = errors.New(string(b))
		return
	}
	err = json.Unmarshal(b, &bod)
	items = bod.Items
	return
}

func GetItem(itemIdOrCode, tenantID, accessToken string) (item Item, err error) {
	url := endpoints.EndpointItems + "/" + itemIdOrCode
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	utils.AddXeroHeaders(request, accessToken, tenantID)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	b, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	if response.StatusCode != 200 {
		err = errors.New(string(b))
		return
	}
	var body struct {
		Items []Item
	}
	err = json.Unmarshal(b, &body)
	if len(body.Items) == 1 {
		item = body.Items[0]
	}
	return
}

func GetItemHistory(itemIdOrCode, tenantID, accessToken string) (history []ItemHistory, err error) {
	if itemIdOrCode == "" {
		err = ErrInvalidItemID
		return
	}
	url := endpoints.EndpointItems + "/" + itemIdOrCode + "/History"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	utils.AddXeroHeaders(request, accessToken, tenantID)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	b, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	if response.StatusCode != 200 {
		err = errors.New(string(b))
		return
	}
	var responseBody struct {
		HistoryRecords []ItemHistory
	}
	err = json.Unmarshal(b, &responseBody)
	history = responseBody.HistoryRecords
	return
}

func CreateItem(item Item, tenantID, accessToken string) (itm Item, err error) {
	url := endpoints.EndpointItems
	if !item.validForCreation() {
		err = ErrInvalidItemForCreation
		return
	}
	b, err := json.Marshal(item.toCreate())
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(b)
	request, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return
	}
	utils.AddXeroHeaders(request, accessToken, tenantID)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	b, err = io.ReadAll(response.Body)
	if err != nil {
		return
	}
	if response.StatusCode != 200 {
		err = errors.New(string(b))
		return
	}
	var responseBody struct {
		Items []Item
	}
	err = json.Unmarshal(b, &responseBody)
	if err != nil {
		return
	}
	if len(responseBody.Items) == 1 {
		itm = responseBody.Items[0]
	}
	return
}

func UpdateItem(item Item, tenantID, accessToken string) (err error) {
	url := endpoints.EndpointItems
	if !item.validForUpdating() {
		err = ErrInvalidItemForUpdating
		return
	}
	b, err := json.Marshal(item.toUpdate())
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(b)
	request, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return
	}
	utils.AddXeroHeaders(request, accessToken, tenantID)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	if response.StatusCode != 200 {
		defer response.Body.Close()
		b, err = io.ReadAll(response.Body)
		if err != nil {
			return
		}
		err = errors.New(string(b))
		return
	}
	return
}

func DeleteItem(itemID, tenantID, accessToken string) (err error) {
	if itemID == "" {
		return ErrInvalidItemID
	}
	url := endpoints.EndpointItems + "/" + itemID
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return
	}
	utils.AddXeroHeaders(request, accessToken, tenantID)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	if response.StatusCode != 204 {
		defer response.Body.Close()
		b, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		err = errors.New(string(b))
		return err
	}
	return nil
}

func AddNoteToItem(itemID, note, tenantID, accessToken string) (err error) {
	if itemID == "" {
		return ErrInvalidItemID
	}
	url := endpoints.EndpointItems + "/" + itemID + "/History"
	var requestBody struct {
		HistoryRecords []struct {
			Details string
		}
	}
	requestBody.HistoryRecords = []struct{ Details string }{{note}}
	b, err := json.Marshal(requestBody)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(b)
	request, err := http.NewRequest("PUT", url, buf)
	if err != nil {
		return
	}
	utils.AddXeroHeaders(request, accessToken, tenantID)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	if response.StatusCode != 200 {
		defer response.Body.Close()
		b, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		err = errors.New(string(b))
		return err
	}
	return nil
}
