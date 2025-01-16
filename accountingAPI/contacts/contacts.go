package contacts

import (
	brandingthemes "github.com/datag8r/xerogo/accountingAPI/brandingThemes"
	"github.com/datag8r/xerogo/accountingAPI/endpoints"
	"github.com/datag8r/xerogo/accountingAPI/pagination"
	trackingcategories "github.com/datag8r/xerogo/accountingAPI/trackingCategories"
	"github.com/datag8r/xerogo/filter"
	"github.com/datag8r/xerogo/helpers"
)

type Contact struct {
	// Retreived On Multi Contact Get Requests
	ContactID                 string           `json:",omitempty"`
	ContactNumber             string           `json:",omitempty"`
	AccountNumber             string           `json:",omitempty"`
	ContactStatus             contactStatus    `json:",omitempty"`
	Name                      string           `json:",omitempty"`
	FirstName                 string           `json:",omitempty"`
	LastName                  string           `json:",omitempty"`
	EmailAddress              string           `json:",omitempty"`
	BankAccountDetails        string           `json:",omitempty"`
	CompanyNumber             string           `json:",omitempty"` // max length 50 char
	TaxNumber                 string           `json:",omitempty"`
	AccountsReceivableTaxType string           `json:",omitempty"` // will be taxType
	AccountsPayableTaxType    string           `json:",omitempty"` // will be taxType
	Addresses                 []addressDetails `json:",omitempty"`
	Phones                    []phoneDetails   `json:",omitempty"`
	IsSupplier                bool             `json:",omitempty"`
	IsCustomer                bool             `json:",omitempty"`
	DefaultCurrency           string           `json:",omitempty"`
	UpdatedDateUTC            string           `json:",omitempty"`
	// Only Retreived with pagination or single contact request
	ContactPersons                 []ContactPerson                       `json:",omitempty"`
	XeroNetworkKey                 string                                `json:",omitempty"`
	MergedToContactID              string                                `json:",omitempty"`
	SalesDefaultAccountCode        string                                `json:",omitempty"`
	PurchaseDefaultAccountCode     string                                `json:",omitempty"`
	SalesTrackingTCategories       []trackingcategories.TrackingCategory `json:",omitempty"`
	PurchaseTrackingCategories     []trackingcategories.TrackingCategory `json:",omitempty"`
	SalesDefaultLineAmountType     lineAmountType                        `json:",omitempty"`
	PurchasesDefaultLineAmountType lineAmountType                        `json:",omitempty"`
	TrackingCategoryName           string                                `json:",omitempty"`
	TrackingOptionName             string                                `json:",omitempty"`
	PaymentTerms                   paymentTerms                          `json:",omitempty"`
	// ContactGroups                  ContactGroups
	Website       string                       `json:",omitempty"`
	BrandingTheme brandingthemes.BrandingTheme `json:",omitzero"`
	// BatchPayments                  batchPaymentDetails // ??
	Discount       string   `json:",omitempty"`
	Balances       balances // idek
	HasAttachments bool
}

func (c Contact) IsZero() bool {
	return c.ContactID == ""
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
	if contactIdOrNumber == "" {
		err = ErrInvalidContactID
		return
	}
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
	if contactToCreate.Name == "" {
		err = ErrInvalidContactForCreation
		return
	}
	url := endpoints.EndpointContacts
	buf, err := helpers.MarshallJsonToBuffer(contactToCreate)
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

func UpdateContact(contact Contact, tenantId, accessToken string) (err error) {
	if contact.ContactID == "" {
		return ErrInvalidContactForUpdating
	}
	url := endpoints.EndpointContacts
	buf, err := helpers.MarshallJsonToBuffer(contact)
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
		Contacts []Contact
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	return
}

func ArchiveContact(contactId, tenantId, accessToken string) (err error) {
	if contactId == "" {
		return ErrInvalidContactID
	}
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

// summaryOnly
// searchTerm
