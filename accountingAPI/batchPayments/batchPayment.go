package batchpayments

import (
	"github.com/datag8r/xerogo/accountingAPI/accounts"
	"github.com/datag8r/xerogo/accountingAPI/endpoints"
	"github.com/datag8r/xerogo/accountingAPI/payments"
	"github.com/datag8r/xerogo/filter"
	"github.com/datag8r/xerogo/helpers"
	"github.com/datag8r/xerogo/utils"
)

type BatchPayment struct {
	BatchPaymentID string             `xero:"id"`
	Account        accounts.Account   `xero:"create,embeddedId"`
	Particulars    string             `xero:"*create"` // NZ Only
	Code           string             `xero:"*create"` // NZ Only
	Reference      string             `xero:"*create"` // NZ Only
	Details        string             `xero:"*create"` // Non NZ Only
	Narrative      string             `xero:"*create"` // UK Only
	Date           string             `xero:"create"`  // YYYY-MM-DD
	Payments       []payments.Payment `xero:"create"`
	Type           batchPaymentType
	Status         batchPaymentStatus
	TotalAmount    float64 `json:",string"`
	IsReconciled   bool
	UpdatedDateUTC string
}

func GetBatchPayments(tenantId, accessToken string, where *filter.Filter) (batchPayments []BatchPayment, err error) {
	url := endpoints.EndpointBatchPayments
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
		BatchPayments []BatchPayment
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	batchPayments = responseBody.BatchPayments
	return
}

func GetBatchPayment(tenantId, accessToken, batchPaymentId string) (batchPayment BatchPayment, err error) {
	url := endpoints.EndpointBatchPayments + "/" + batchPaymentId
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
		BatchPayments []BatchPayment
	}
	err = helpers.UnmarshalJson(b, &responseBody)
	if len(responseBody.BatchPayments) == 1 {
		batchPayment = responseBody.BatchPayments[0]
	}
	return
}

func CreateBatchPayment(tenantId string, accessToken string, bp BatchPayment) (batchPayment BatchPayment, err error) {
	url := endpoints.EndpointBatchPayments
	inter, err := utils.XeroCustomMarshal(bp, "create")
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
		BatchPayments []BatchPayment
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	if err != nil {
		return
	}
	if len(responseBody.BatchPayments) == 1 {
		batchPayment = responseBody.BatchPayments[0]
	}
	return
}

func CreateBatchPayments(tenantId, accessToken string, bp []BatchPayment) (batchPayments []BatchPayment, err error) {
	url := endpoints.EndpointBatchPayments
	iter, err := utils.XeroCustomMarshal(bp, "create")
	if err != nil {
		return
	}
	var requestBody = map[string]interface{}{"BatchPayments": iter}
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
		BatchPayments []BatchPayment
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	batchPayments = responseBody.BatchPayments
	return
}

func DeleteBatchPayment(tenantId, accessToken string, batchPaymentId string) (err error) {
	url := endpoints.EndpointBatchPayments + "/" + batchPaymentId
	var requestBody = map[string]string{"Status": string(BatchPaymentStatusDeleted)}
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

func DeleteBatchPayments(tenantId, accessToken string, batchPaymentIds []string) (err error) {
	url := endpoints.EndpointBatchPayments
	var requestBody = map[string][]map[string]string{
		"BatchPayments": {},
	}
	for _, id := range batchPaymentIds {
		requestBody["BatchPayments"] = append(requestBody["BatchPayments"], map[string]string{
			"BatchPaymentID": id,
			"Status":         string(BatchPaymentStatusDeleted),
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
