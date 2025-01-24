package contacts

import (
	"github.com/datag8r/xerogo/errors"
	"github.com/datag8r/xerogo/filter"
)

type contactsEndpoint struct {
	tenantId          string
	accessToken       string
	rateLimitCallback func()
}

func NewContactsEndpoint(tenantId, accessToken string, rateLimitCallback func()) *contactsEndpoint {
	return &contactsEndpoint{
		tenantId:          tenantId,
		accessToken:       accessToken,
		rateLimitCallback: rateLimitCallback,
	}
}

func (e *contactsEndpoint) RateLimitCallback() {
	if e.rateLimitCallback != nil {
		e.rateLimitCallback()
	}
}

func (e *contactsEndpoint) GetOne(id string) (acc Contact, err error) {
	e.RateLimitCallback()
	return GetContact(id, e.tenantId, e.accessToken)
}

func (e *contactsEndpoint) GetMulti(where *filter.Filter) (contacts []Contact, err error) {
	var currentPage uint = 1
	for {
		// rate limit handling
		e.RateLimitCallback()
		// actual fetching
		c, pData, err := GetContacts(e.tenantId, e.accessToken, &currentPage, where)
		if err != nil {
			return nil, err
		}
		if pData == nil {
			return nil, errors.ErrNoPageDataReturned
		}
		contacts = append(contacts, c...)
		if pData.PageCount == currentPage {
			break
		}
		currentPage++
	}
	return
}

func (e *contactsEndpoint) CreateOne(con Contact) (contact Contact, err error) {
	e.RateLimitCallback()
	return CreateContact(con, e.tenantId, e.accessToken)
}

func (e *contactsEndpoint) CreateMulti(cons []Contact) (contacts []Contact, err error) {
	e.RateLimitCallback()
	return CreateContacts(cons, e.tenantId, e.accessToken)
}

func (e *contactsEndpoint) UpdateOne(account Contact) (err error) {
	e.RateLimitCallback()
	return UpdateContact(account, e.tenantId, e.accessToken)
}

func (e *contactsEndpoint) UpdateMulti(contacts []Contact) (err error) {
	e.RateLimitCallback()
	return UpdateContacts(contacts, e.tenantId, e.accessToken)
}

func (e *contactsEndpoint) ArchiveOne(id string) (err error) {
	e.RateLimitCallback()
	return ArchiveContact(id, e.tenantId, e.accessToken)
}

func (e *contactsEndpoint) ArchiveMulti(ids []string) (err error) {
	e.RateLimitCallback()
	return ArchiveContacts(ids, e.tenantId, e.accessToken)
}

// Not Supported
func (e *contactsEndpoint) DeleteOne(id string) (err error) {
	return errors.ErrEndpointCallNotSupported
}

// Not Supported
func (e *contactsEndpoint) DeleteMulti(ids []string) (err error) {
	return errors.ErrEndpointCallNotSupported
}
