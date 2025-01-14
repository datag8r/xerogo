package contacts

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	brandingthemes "github.com/datag8r/xerogo/accountingAPI/brandingThemes"
	"github.com/datag8r/xerogo/accountingAPI/endpoints"
	"github.com/datag8r/xerogo/accountingAPI/pagination"
	trackingcategories "github.com/datag8r/xerogo/accountingAPI/trackingCategories"
	"github.com/datag8r/xerogo/filter"
	"github.com/datag8r/xerogo/utils"
)

type Contact struct {
	// Retreived On Multi Contact Get Requests
	ContactID                 string        `json:",omitempty"`
	ContactNumber             string        `json:",omitempty"`
	AccountNumber             string        `json:",omitempty"`
	ContactState              contactStatus `json:",omitempty"`
	Name                      string
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
	Website       string
	BrandingTheme brandingthemes.BrandingTheme `json:",omitzero"`
	// BatchPayments                  batchPaymentDetails // ??
	Discount       string
	Balances       balances // idek
	HasAttachments bool
}

// includeArchived
func GetContacts(tenantID, accessToken string, page *uint, where *filter.Filter) (contacts []Contact, pageData *pagination.PaginationData, err error) {
	url := endpoints.EndpointContacts
	if page != nil { // make this a func later
		url += "?page=" + fmt.Sprint(*page)
		if !pagination.IsDefaultPageSize() {
			url += "?pageSize=" + fmt.Sprint(pagination.CustomPageSize)
		}
	}
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
	var responseBody struct {
		Pagination *pagination.PaginationData
		Contacts   []Contact
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
	err = json.Unmarshal(b, &responseBody)
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
	var request *http.Request
	request, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	utils.AddXeroHeaders(request, accessToken, tenantId)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	b, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	var responseBody struct {
		Contacts []Contact `json:"Contacts"`
	}
	if response.StatusCode != 200 {
		err = errors.New(string(b))
		return
	}
	err = json.Unmarshal(b, &responseBody)
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
	b, err := json.Marshal(contactToCreate)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(b)
	var request *http.Request
	request, err = http.NewRequest("POST", url, buf)
	if err != nil {
		return
	}
	utils.AddXeroHeaders(request, accessToken, tenantId)
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
		Contacts []Contact
	}
	err = json.Unmarshal(b, &responseBody)
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
	b, err := json.Marshal(contact)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(b)
	var request *http.Request
	request, err = http.NewRequest("POST", url, buf)
	if err != nil {
		return
	}
	utils.AddXeroHeaders(request, accessToken, tenantId)
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
		Contacts []Contact
	}
	err = json.Unmarshal(b, &responseBody)
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
	b, err := json.Marshal(requestBody)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(b)
	request, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return
	}
	utils.AddXeroHeaders(request, accessToken, tenantId)
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
		return errors.New(string(b))
	}
	return
}

// summaryOnly
// searchTerm
