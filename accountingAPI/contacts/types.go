package contacts

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
