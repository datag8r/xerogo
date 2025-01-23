package banktransfers

import (
	"github.com/datag8r/xerogo/accountingAPI/accounts"
	"github.com/datag8r/xerogo/accountingAPI/endpoints"
	"github.com/datag8r/xerogo/filter"
	"github.com/datag8r/xerogo/helpers"
	"github.com/datag8r/xerogo/utils"
)

type BankTransfer struct {
	FromBankAccount       accounts.Account `xero:"create,embeddedId"`     // required for creation
	ToBankAccount         accounts.Account `xero:"create,embeddedId"`     // required for creation
	Amount                float64          `xero:"create" json:",string"` // required for creation
	Date                  string           `xero:"*create"`               // optional for creation
	FromIsReconciled      bool             `xero:"*create"`               // optional for creation
	ToIsReconciled        bool             `xero:"*create"`               // optional for creation
	Reference             string           `xero:"*create"`               // optional for creation
	BankTransferID        string
	CurrencyRate          string
	FromBankTransactionID string
	ToBankTransactionID   string
	HasAttachments        bool
	CreatedDateUTC        string
}

func GetBankTransfers(tenantId, accessToken string, where *filter.Filter) (transfers []BankTransfer, err error) {
	url := endpoints.EndpointBankTransfers
	request, err := helpers.BuildRequest("GET", url, nil, where, nil)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	b, err := helpers.DoRequest(request, 200)
	if err != nil {
		return
	}
	var responseBody struct {
		BankTransfers []BankTransfer
	}
	err = helpers.UnmarshalJson(b, &responseBody)
	transfers = responseBody.BankTransfers
	return
}

func GetBankTransfer(tenantId, accessToken, transferId string) (transfer BankTransfer, err error) {
	url := endpoints.EndpointBankTransfers + "/" + transferId
	request, err := helpers.BuildRequest("GET", url, nil, nil, nil)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	b, err := helpers.DoRequest(request, 200)
	if err != nil {
		return
	}
	var responseBody struct {
		BankTransfers []BankTransfer
	}
	err = helpers.UnmarshalJson(b, &responseBody)
	if len(responseBody.BankTransfers) == 1 {
		transfer = responseBody.BankTransfers[0]
	}
	return
}

func CreateBankTransfer(tenantId, accessToken string, bt BankTransfer) (transfer BankTransfer, err error) {
	url := endpoints.EndpointBankTransfers
	inter, err := utils.XeroCustomMarshal(bt, "create")
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
		BankTransfers []BankTransfer
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	if err != nil {
		return
	}
	if len(responseBody.BankTransfers) == 1 {
		transfer = responseBody.BankTransfers[0]
	}
	return
}

func CreateBankTransfers(tenantId, accessToken string, bts []BankTransfer) (transfers []BankTransfer, err error) {
	url := endpoints.EndpointBankTransfers
	inter, err := utils.XeroCustomMarshal(bts, "create")
	if err != nil {
		return
	}
	var requestBody = map[string]interface{}{"BankTransfers": inter}
	buf, err := helpers.MarshallJsonToBuffer(requestBody)
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
		BankTransfers []BankTransfer
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	transfers = responseBody.BankTransfers
	return
}
