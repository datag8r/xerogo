package banktransactions

import (
	"github.com/datag8r/xerogo/accountingAPI/accounts"
	batchpayments "github.com/datag8r/xerogo/accountingAPI/batchPayments"
	"github.com/datag8r/xerogo/accountingAPI/contacts"
	"github.com/datag8r/xerogo/accountingAPI/endpoints"
	trackingcategories "github.com/datag8r/xerogo/accountingAPI/trackingCategories"
	"github.com/datag8r/xerogo/accountingAPI/types"
	"github.com/datag8r/xerogo/filter"
	"github.com/datag8r/xerogo/helpers"
	"github.com/datag8r/xerogo/pagination"
	"github.com/datag8r/xerogo/utils"
)

type BankTransaction struct {
	BankTransactionID string                `xero:"update,id"`
	Type              bankTransactionType   `xero:"create,*update"`
	Contact           contacts.Contact      `xero:"create,*update,embeddedId"`
	LineItems         []LineItem            `xero:"create,*update"`
	BankAccount       accounts.Account      `xero:"create,*update,embeddedId"`
	IsReconciled      bool                  `xero:"*create,*update"`
	Reference         string                `xero:"*create,*update"`
	Date              string                `xero:"*create,*update"`
	CurrencyCode      string                `xero:"*create,*update"`
	CurrencyRate      float64               `xero:"*create,*update" json:",string"`
	Url               string                `xero:"*create,*update"`
	Status            bankTransactionStatus `xero:"*create,*update"`
	LineAmountTypes   types.LineAmountType  `xero:"*create,*update"`
	BatchPayment      batchpayments.BatchPayment
	SubTotal          float64 `json:",string"`
	TotalTax          float64 `json:",string"`
	Total             float64 `json:",string"`
	PrepaymentID      string
	OverpaymentID     string
	HasAttachments    bool
	UpdatedDateUTC    string
}

type LineItem struct {
	LineItemID  string                        `xero:"update,id"`
	Description string                        `xero:"create,*update"`
	Quantity    float64                       `xero:"create,*update" json:",string"`
	UnitAmount  float64                       `xero:"create,*update" json:",string"`
	ItemCode    string                        `xero:"create,*update"`
	AccountCode string                        `xero:"create,*update"`
	TaxType     types.TaxType                 `xero:"create,*update"`
	TaxAmount   float64                       `xero:"create,*update" json:",string"`
	LineAmount  float64                       `xero:"create,*update" json:",string"`
	Tracking    []trackingcategories.Tracking `xero:"*create,*update"`
}

func GetBankTransactions(tenantId, accessToken string, page *uint, where *filter.Filter) (transactions []BankTransaction, pageData *pagination.PaginationData, err error) {
	url := endpoints.EndpointBankTransactions
	request, err := helpers.BuildRequest("GET", url, page, where, nil)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	body, err := helpers.DoRequest(request, 200)
	if err != nil {
		return
	}
	var responseBody struct {
		Pagination       *pagination.PaginationData
		BankTransactions []BankTransaction
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	transactions = responseBody.BankTransactions
	pageData = responseBody.Pagination
	return
}

func GetBankTransaction(tenantId, accessToken, transactionId string) (transaction BankTransaction, err error) {
	url := endpoints.EndpointBankTransactions + "/" + transactionId
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
		BankTransactions []BankTransaction
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	if len(responseBody.BankTransactions) == 1 {
		transaction = responseBody.BankTransactions[0]
	}
	return
}

func CreateBankTransaction(tenantId, accessToken string, transactionToCreate BankTransaction) (transaction BankTransaction, err error) {
	url := endpoints.EndpointBankTransactions
	inter, err := utils.XeroCustomMarshal(transactionToCreate, "create")
	if err != nil {
		return
	}
	var requestBody = map[string]interface{}{"BankTransactions": []interface{}{inter}}
	buf, err := helpers.MarshallJsonToBuffer(requestBody)
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
		BankTransactions []BankTransaction
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	if err != nil {
		return
	}
	if len(responseBody.BankTransactions) == 1 {
		transaction = responseBody.BankTransactions[0]
	}
	return
}

func CreateBankTransactions(tenantId, accessToken string, transactionsToCreate []BankTransaction) (transactions []BankTransaction, err error) {
	url := endpoints.EndpointBankTransactions
	iter, err := utils.XeroCustomMarshal(transactionsToCreate, "create")
	if err != nil {
		return
	}
	var requestBody = map[string]interface{}{"BankTransactions": iter}
	buf, err := helpers.MarshallJsonToBuffer(requestBody)
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
		BankTransactions []BankTransaction
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	transactions = responseBody.BankTransactions
	return
}

func UpdateBankTransaction(tenantId, accessToken string, transaction BankTransaction) (err error) {
	url := endpoints.EndpointBankTransactions
	inter, err := utils.XeroCustomMarshal(transaction, "update")
	if err != nil {
		return
	}
	var requestBody = map[string]interface{}{"BankTransactions": []interface{}{inter}}
	buf, err := helpers.MarshallJsonToBuffer(requestBody)
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
		BankTransactions []BankTransaction
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	return
}

func UpdateBankTransactions(tenantId, accessToken string, transactions []BankTransaction) (err error) {
	url := endpoints.EndpointBankTransactions
	inter, err := utils.XeroCustomMarshal(transactions, "update")
	if err != nil {
		return
	}
	var requestBody = map[string]interface{}{"BankTransactions": inter}
	buf, err := helpers.MarshallJsonToBuffer(requestBody)
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
		BankTransactions []BankTransaction
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	return
}

func DeleteBankTransaction(tenantId, accessToken, transactionId string) (err error) {
	url := endpoints.EndpointBankTransactions
	var requestBody struct {
		BankTransactionID string
		TransactionStatus string
	}
	requestBody.BankTransactionID = transactionId
	requestBody.TransactionStatus = string(BankTransactionStatusDeleted)
	buf, err := helpers.MarshallJsonToBuffer(requestBody)
	if err != nil {
		return
	}
	request, err := helpers.BuildRequest("PUT", url, nil, nil, buf)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	_, err = helpers.DoRequest(request, 200)
	if err != nil {
		return
	}
	return
}

func DeleteBankTransactions(tenantId, accessToken string, transactionIds []string) (err error) {
	url := endpoints.EndpointBankTransactions
	var requestBody = map[string][]map[string]string{
		"BankTransactions": {},
	}
	for _, id := range transactionIds {
		requestBody["BankTransactions"] = append(requestBody["BankTransactions"], map[string]string{
			"BankTransactionID": id,
			"TransactionStatus": string(BankTransactionStatusDeleted),
		})
	}
	buf, err := helpers.MarshallJsonToBuffer(requestBody)
	if err != nil {
		return
	}
	request, err := helpers.BuildRequest("POST", url, nil, nil, buf)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	_, err = helpers.DoRequest(request, 200)
	if err != nil {
		return
	}
	return
}
