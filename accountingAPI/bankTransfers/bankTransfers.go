package banktransfers

import (
	"github.com/datag8r/xerogo/accountingAPI/endpoints"
	"github.com/datag8r/xerogo/filter"
	"github.com/datag8r/xerogo/helpers"
)

type BankTransfer struct {
	FromBankAccount       bankAccountDetails // required for creation
	ToBankAccount         bankAccountDetails // required for creation
	Amount                float64            `json:",string"`    // required for creation
	Date                  string             `json:",omitempty"` // optional for creation
	FromIsReconciled      *bool              `json:",omitempty"` // optional for creation // default false probs
	ToIsReconciled        *bool              `json:",omitempty"` // optional for creation // default false probs
	Reference             string             `json:",omitempty"` // optional for creation
	BankTransferID        string
	CurrencyRate          string // ?
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
	if transferId == "" {
		err = ErrInvalidBankTransferID
		return
	}
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
	if !bt.validForCreation() {
		err = ErrInvalidBankTransferForCreation
		return
	}
	buf, err := helpers.MarshallJsonToBuffer(bt.toCreate())
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
