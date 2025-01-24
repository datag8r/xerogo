package invoices

import (
	"github.com/datag8r/xerogo/accountingAPI/contacts"
	creditnotes "github.com/datag8r/xerogo/accountingAPI/creditNotes"
	"github.com/datag8r/xerogo/accountingAPI/currencies"
	"github.com/datag8r/xerogo/accountingAPI/endpoints"
	"github.com/datag8r/xerogo/accountingAPI/overpayments"
	"github.com/datag8r/xerogo/accountingAPI/payments"
	"github.com/datag8r/xerogo/accountingAPI/prepayments"
	"github.com/datag8r/xerogo/accountingAPI/types"
	"github.com/datag8r/xerogo/filter"
	"github.com/datag8r/xerogo/helpers"
)

type Invoice struct {
	InvoiceID       string                  `xero:"update,id"`
	InvoiceNumber   string                  `xero:"*create,*update,id"`
	Reference       string                  `xero:"*create,*update"` // ACCREC only
	Type            invoiceType             `xero:"create"`
	Contact         contacts.Contact        `xero:"create,embeddedId"` // only fills id and name without pagination or single resource request
	DateString      string                  `xero:"*create,*update"`   // YYYY-MM-DD
	DueDateString   string                  `xero:"*create,*update"`   // YYYY-MM-DD
	Status          invoiceStatus           `xero:"*create,*update"`
	LineAmountTypes types.LineAmountType    `xero:"*create,*update"`
	CurrencyCode    currencies.CurrencyCode `xero:"*create,*update"`
	SubTotal        float64                 `json:",string"`
	TotalTax        float64                 `json:",string"`
	Total           float64                 `json:",string"`
	TotalDiscount   float64                 `json:",string"`
	AmountDue       float64                 `json:",string"`
	AmountCredited  float64                 `json:",string"`
	AmountPaid      float64                 `json:",string"`
	Date            string
	DueDate         string
	UpdatedDateUTC  string
	// the following are only filled using pagination or single resource request
	LineItems                    []InvoiceLineItem `xero:"create"`
	CurrencyRate                 string            `xero:"*create,*update"`
	BrandingThemeID              string            `xero:"*create,*update"`
	Url                          string            `xero:"*create,*update"`
	SentToContact                bool              `xero:"*create,*update"`
	ExpectedPaymentDate          string            `xero:"*create,*update"`
	PlannedPaymentDate           string            `xero:"*create,*update"`
	HasAttachments               bool
	RepeatingInvoiceID           string
	Payments                     []payments.Payment
	CreditNotes                  []creditnotes.CreditNote
	Prepayments                  []prepayments.Prepayment
	Overpayments                 []overpayments.Overpayment
	CISDeduction                 float64
	FullyPaidOnDate              string
	SalesTaxCalclulationTypeCode taxCalculationtype
}

func GetInvoices(tenantId, accessToken string, page *uint, where *filter.Filter) (invoices []Invoice, err error) {
	url := endpoints.EndpointInvoices
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
		Invoices []Invoice
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	invoices = responseBody.Invoices
	return
}

func GetInvoice(invoiceId string, tenantId, accessToken string) (invoice Invoice, err error) {
	if invoiceId == "" {
		err = ErrInvalidInvoiceID
		return
	}
	url := endpoints.EndpointInvoices + "/" + invoiceId
	request, err := helpers.BuildRequest("GET", url, nil, nil, nil)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	body, err := helpers.DoRequest(request, 200)
	if err != nil {
		return
	}
	err = helpers.UnmarshalJson(body, &invoice)
	return
}

func CreateInvoice(invoiceToCreate Invoice, tenantId string, accessToken string) (invoice Invoice, err error) {
	url := endpoints.EndpointInvoices

	buf, err := helpers.MarshallJsonToBuffer(invoiceToCreate)
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
		Invoices []Invoice
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	if err != nil {
		return
	}
	if len(responseBody.Invoices) == 1 {
		invoice = responseBody.Invoices[0]
	}
	return
}

func UpdateInvoice(invoice Invoice, tenantId string, accessToken string) (err error) {
	url := endpoints.EndpointInvoices
	buf, err := helpers.MarshallJsonToBuffer(invoice)
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

func DeleteInvoice(invoiceId string, tenantId, accessToken string) (err error) {
	if invoiceId == "" {
		err = ErrInvalidInvoiceID
		return
	}
	url := endpoints.EndpointInvoices + "/" + invoiceId
	var requestBody struct {
		InvoiceID string
		Status    string
	}
	requestBody.InvoiceID = invoiceId
	requestBody.Status = string(InvoiceStatusDeleted)
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
	return
}

func VoidInvoice(invoiceId string, tenantId, accessToken string) (err error) {
	if invoiceId == "" {
		err = ErrInvalidInvoiceID
		return
	}
	url := endpoints.EndpointInvoices + "/" + invoiceId
	var requestBody struct {
		InvoiceID string
		Status    string
	}
	requestBody.InvoiceID = invoiceId
	requestBody.Status = string(InvoiceStatusVoided)
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
	return
}
