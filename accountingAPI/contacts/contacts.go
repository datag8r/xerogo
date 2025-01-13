package contacts

import (
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
	ContactID                 string
	ContactNumber             string
	AccountNumber             string
	ContactState              contactStatus
	Name                      string
	FirstName                 string
	LastName                  string
	EmailAddress              string
	BankAccountDetails        string
	CompanyNumber             string // max length 50 char
	TaxNumber                 string
	AccountsReceivableTaxType string // will be taxType
	AccountsPayableTaxType    string // will be taxType
	Addresses                 []addressDetails
	Phones                    []phoneDetails
	IsSupplier                bool
	IsCustomer                bool
	DefaultCurrency           string
	UpdatedDateUTC            string
	// Only Retreived with pagination or single contact request
	ContactPersons                 []ContactPerson
	XeroNetworkKey                 string
	MergedToContactID              string
	SalesDefaultAccountCode        string
	PurchaseDefaultAccountCode     string
	SalesTrackingTCategories       []trackingcategories.TrackingCategory
	PurchaseTrackingCategories     []trackingcategories.TrackingCategory
	SalesDefaultLineAmountType     lineAmountType
	PurchasesDefaultLineAmountType lineAmountType
	TrackingCategoryName           string
	TrackingOptionName             string
	PaymentTerms                   paymentTerms
	// ContactGroups                  ContactGroups
	Website       string
	BrandingTheme brandingthemes.BrandingTheme
	// BatchPayments                  batchPaymentDetails // ??
	Discount       string
	Balances       balances // idek
	HasAttachments bool
}

type balances struct {
	AccountsPayable    *balance `json:"AccountsPayable,omitempty"`
	AccountsReceivable *balance `json:"AccountsReceivable,omitempty"`
}

type balance struct {
	Outstanding float64
	Overdue     float64
}

type paymentTerms string

// payment terms? not sure where number goes
// DAYSAFTERBILLDATE	day(s) after bill date
// DAYSAFTERBILLMONTH	day(s) after bill month
// OFCURRENTMONTH	of the current month
// OFFOLLOWINGMONTH

// This will probably move
type lineAmountType string

var (
	LineAmountTypeInclusive lineAmountType = "INCLUSIVE"
	LineAmountTypeExclusive lineAmountType = "EXCLUSIVE"
	LineAmountTypeNone      lineAmountType = "NONE"
)

type CISSettings struct {
	CISEnabled bool
	Rate       int
}

type contactStatus string

var (
	ContactStatusActive      contactStatus = "ACTIVE"
	ContactStatusArchived    contactStatus = "ARCHIVED"
	ContactStatusGDPRRequest contactStatus = "GDPRREQUEST"
)

type phoneDetails struct {
	Phonetype        phoneType
	PhoneNumber      string // max length 50 chars
	PhoneAreaCode    string // max length 10 chars
	PhoneCountryCode string // max length 20 chars
}

type phoneType string

var (
	PhoneTypeDefault phoneType = "DEFAULT"
	PhoneTypeDDI     phoneType = "DDI"
	PhoneTypeMobile  phoneType = "MOBILE"
	PhoneTypeFax     phoneType = "FAX"
)

type addressDetails struct {
	AddressType  addressType
	AddressLine1 string // max length = 500
	AddressLine2 string // max length = 500
	AddressLine3 string // max length = 500
	City         string // max length = 500
	Region       string // max length = 500
	PostalCode   string // max length = 500
	Country      string // max length = 50 | [A-Z], [a-z] only
	AttentionTo  string // max length = 255
}

type addressType string

var (
	AddressTypePOBox    addressType = "POBOX"
	AddressTypeStreet   addressType = "STREET"
	AddressTypeDelivery addressType = "DELIVERY" // Not Valid For Contacts
)

type ContactPerson struct {
	FirstName       string
	LastName        string
	EmailAddress    string
	IncludeInEmails bool
}

var (
	ErrInvalidContactForCreation = errors.New("one or more required fields were invalid to create this contact")
	ErrInvalidContactForUpdating = errors.New("a valid contact id is required to update an contact")
	ErrInvalidContactID          = errors.New("invalid contact id for request")
)

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

func GetContact(tenantId, accessToken, contactId string) (contact Contact, err error) {
	if len(contactId) != len("297c2dc5-cc47-4afd-8ec8-74990b8761e9") { // figure out the number
		err = ErrInvalidContactID
		return
	}
	url := endpoints.EndpointContacts + "/" + contactId
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

func CreateContact(contact Contact, tenantId, accessToken string) (err error)  { return }
func UpdateContact(contact Contact, tenantId, accessToken string) (err error)  { return }
func ArchiveContact(contact Contact, tenantId, accessToken string) (err error) { return }
func DeleteContact(contact Contact, tenantId, accessToken string) (err error)  { return }
