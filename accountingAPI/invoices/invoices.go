package invoices

import (
	"github.com/datag8r/xerogo/accountingAPI/contacts"
	"github.com/datag8r/xerogo/accountingAPI/currencies"
	"github.com/datag8r/xerogo/accountingAPI/endpoints"
	"github.com/datag8r/xerogo/filter"
	"github.com/datag8r/xerogo/helpers"
)

type Invoice struct {
	Reference       string // ACCREC only
	Type            invoiceType
	Contact         contacts.Contact // only fills id and name without pagination or single resource request
	Date            string
	DateString      string
	DueDate         string
	DueDateString   string
	Status          invoiceStatus
	LineAmountTypes lineAmountType
	SubTotal        float64 `json:",string"`
	TotalTax        float64 `json:",string"`
	Total           float64 `json:",string"`
	TotalDiscount   float64 `json:",string"`
	UpdatedDateUTC  string
	InvoiceID       string
	InvoiceNumber   string
	CurrencyCode    currencies.CurrencyCode
	AmountCredited  float64 `json:",string"`
	AmountDue       float64 `json:",string"`
	AmountPaid      float64 `json:",string"`
	// the following are only filled using pagination or single resource request
	LineItems           []InvoiceLineItem
	CurrencyRate        string //
	BrandingThemeID     string `json:",omitempty"`
	Url                 string
	SentToContact       bool
	ExpectedPaymentDate string
	PlannedPaymentDate  string
	HasAttachments      bool
	RepeatingInvoiceID  string
	// Payments []payment
	// CreditNotes []creditNotes
	// Prepayments []prepayment
	// Overpayments []overpayment
	CISDeduction                 float64 `json:",string"`
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
	if !invoiceToCreate.validForCreation() {
		err = ErrInvalidInvoiceForCreation
		return
	}
	buf, err := helpers.MarshallJsonToBuffer(invoiceToCreate.toCreate())
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
	if !invoice.validForUpdating() {
		err = ErrInvalidInvoiceForUpdating
		return
	}
	buf, err := helpers.MarshallJsonToBuffer(invoice.toUpdate())
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
