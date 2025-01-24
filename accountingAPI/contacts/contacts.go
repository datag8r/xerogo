package contacts

import (
	brandingthemes "github.com/datag8r/xerogo/accountingAPI/brandingThemes"
	contactgroups "github.com/datag8r/xerogo/accountingAPI/contactGroups"
	"github.com/datag8r/xerogo/accountingAPI/endpoints"
	trackingcategories "github.com/datag8r/xerogo/accountingAPI/trackingCategories"
	"github.com/datag8r/xerogo/filter"
	"github.com/datag8r/xerogo/helpers"
	"github.com/datag8r/xerogo/pagination"
	"github.com/datag8r/xerogo/utils"
)

type Contact struct {
	// Retreived On Multi Contact Get Requests
	ContactID                 string           `xero:"update,*id"`
	Name                      string           `xero:"create,*update,*id"`
	ContactNumber             string           `xero:"*create,*update"`
	AccountNumber             string           `xero:"*create,*update"`
	ContactStatus             contactStatus    `xero:"*create,*update"`
	FirstName                 string           `xero:"*create,*update"`
	LastName                  string           `xero:"*create,*update"`
	EmailAddress              string           `xero:"*create,*update"`
	BankAccountDetails        string           `xero:"*create,*update"`
	CompanyNumber             string           `xero:"*create,*update"` // max length 50 char
	TaxNumber                 string           `xero:"*create,*update"`
	AccountsReceivableTaxType string           `xero:"*create,*update"` // will be taxType
	AccountsPayableTaxType    string           `xero:"*create,*update"` // will be taxType
	Addresses                 []addressDetails `xero:"*create,*update"`
	Phones                    []phoneDetails   `xero:"*create,*update"`
	IsSupplier                bool             `xero:"*create,*update"`
	IsCustomer                bool             `xero:"*create,*update"`
	DefaultCurrency           string           `xero:"*create,*update"`
	UpdatedDateUTC            string           `xero:"*create,*update"`
	// Only Retreived with pagination or single contact request
	ContactPersons                 []ContactPerson                       `xero:"*create,*update"`
	XeroNetworkKey                 string                                `xero:"*create,*update"`
	SalesDefaultAccountCode        string                                `xero:"*create,*update"`
	PurchaseDefaultAccountCode     string                                `xero:"*create,*update"`
	SalesTrackingTCategories       []trackingcategories.TrackingCategory `xero:"*create,*update"` //
	PurchaseTrackingCategories     []trackingcategories.TrackingCategory `xero:"*create,*update"` //
	TrackingCategoryName           string                                `xero:"*create,*update"`
	TrackingOptionName             string                                `xero:"*create,*update"`
	PaymentTerms                   paymentTerms                          `xero:"*create,*update"`
	MergedToContactID              string
	SalesDefaultLineAmountType     lineAmountType
	PurchasesDefaultLineAmountType lineAmountType
	ContactGroups                  []contactgroups.ContactGroup
	Website                        string
	BrandingTheme                  brandingthemes.BrandingTheme
	// BatchPayments                  batchPaymentDetails // ??
	Discount       string
	Balances       balances
	HasAttachments bool
}

// includeArchived
func GetContacts(tenantId, accessToken string, page *uint, where *filter.Filter) (contacts []Contact, pageData *pagination.PaginationData, err error) {
	url := endpoints.EndpointContacts
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
		Pagination *pagination.PaginationData
		Contacts   []Contact
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	contacts = responseBody.Contacts
	pageData = responseBody.Pagination
	return
}

func GetContact(tenantId, accessToken, contactIdOrNumber string) (contact Contact, err error) {
	url := endpoints.EndpointContacts + "/" + contactIdOrNumber
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
		Contacts []Contact `json:"Contacts"`
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	if len(responseBody.Contacts) == 1 {
		contact = responseBody.Contacts[0]
	}
	return
}

func CreateContact(contactToCreate Contact, tenantId, accessToken string) (contact Contact, err error) {
	url := endpoints.EndpointContacts
	buf, err := helpers.MarshallJsonToBuffer(contactToCreate)
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
		Contacts []Contact
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	if err != nil {
		return
	}
	if len(responseBody.Contacts) == 1 {
		contact = responseBody.Contacts[0]
	}
	return
}

func CreateContacts(contactsToCreate []Contact, tenantId, accessToken string) (contacts []Contact, err error) {
	url := endpoints.EndpointContacts
	iter, err := utils.XeroCustomMarshal(contactsToCreate, "create")
	if err != nil {
		return
	}
	var requestBody = map[string]interface{}{"Contacts": iter}
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
		Contacts []Contact
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	contacts = responseBody.Contacts
	return
}

func UpdateContact(contact Contact, tenantId, accessToken string) (err error) {
	url := endpoints.EndpointContacts
	iter, err := utils.XeroCustomMarshal(contact, "update")
	if err != nil {
		return
	}
	buf, err := helpers.MarshallJsonToBuffer(iter)
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
		Contacts []Contact
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	return
}

func UpdateContacts(contacts []Contact, tenantId, accessToken string) (err error) {
	url := endpoints.EndpointContacts
	iter, err := utils.XeroCustomMarshal(contacts, "update")
	if err != nil {
		return
	}
	var requestBody = map[string]interface{}{"Contacts": iter}
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
		Contacts []Contact
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	return
}

func ArchiveContact(contactId, tenantId, accessToken string) (err error) {
	url := endpoints.EndpointContacts
	var requestBody struct {
		ContactID     string
		ContactStatus string
	}
	requestBody.ContactID = contactId
	requestBody.ContactStatus = string(ContactStatusArchived)
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

func ArchiveContacts(contactIds []string, tenantId, accessToken string) (err error) {
	url := endpoints.EndpointContacts
	var requestBody = map[string][]map[string]string{
		"Contacts": {},
	}
	for _, id := range contactIds {
		requestBody["Contacts"] = append(requestBody["Contacts"], map[string]string{
			"ContactID":     id,
			"ContactStatus": string(ContactStatusArchived),
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
