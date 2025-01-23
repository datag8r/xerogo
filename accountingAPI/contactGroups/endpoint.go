package contactgroups

import (
	"github.com/datag8r/xerogo/errors"
	"github.com/datag8r/xerogo/filter"
)

type contactGroupsEndpoint struct {
	tenantId          string
	accessToken       string
	rateLimitCallback func()
}

func NewContactGroupsEndpoint(tenantId, accessToken string, rateLimitCallback func()) *contactGroupsEndpoint {
	return &contactGroupsEndpoint{
		tenantId:          tenantId,
		accessToken:       accessToken,
		rateLimitCallback: rateLimitCallback,
	}
}

func (e *contactGroupsEndpoint) RateLimitCallback() {
	if e.rateLimitCallback != nil {
		e.rateLimitCallback()
	}
}

func (e *contactGroupsEndpoint) GetOne(id string) (group ContactGroup, err error) {
	e.RateLimitCallback()
	return GetContactGroup(id, e.tenantId, e.accessToken)
}

func (e *contactGroupsEndpoint) GetMulti(where *filter.Filter) (groups []ContactGroup, err error) {
	e.RateLimitCallback()
	return GetContactGroups(e.tenantId, e.accessToken, where)
}

func (e *contactGroupsEndpoint) CreateOne(contactGroup ContactGroup) (group ContactGroup, err error) {
	e.RateLimitCallback()
	return CreateContactGroup(contactGroup, e.tenantId, e.accessToken)
}

func (e *contactGroupsEndpoint) CreateMulti(contactGroups []ContactGroup) (groups []ContactGroup, err error) {
	e.RateLimitCallback()
	return CreateContactGroups(contactGroups, e.tenantId, e.accessToken)
}

func (e *contactGroupsEndpoint) UpdateOne(contactGroup ContactGroup) (err error) {
	e.RateLimitCallback()
	return UpdateContactGroup(contactGroup, e.tenantId, e.accessToken)
}

// Not Supported
func (e *contactGroupsEndpoint) UpdateMulti(contactGroups []ContactGroup) (err error) {
	e.RateLimitCallback()
	return UpdateContactGroups(contactGroups, e.tenantId, e.accessToken)
}

// Not Supported
func (e *contactGroupsEndpoint) ArchiveOne(id string) (err error) {
	return errors.ErrEndpointCallNotSupported
}

// Not Supported
func (e *contactGroupsEndpoint) ArchiveMulti(ids []string) (err error) {
	return errors.ErrEndpointCallNotSupported
}

func (e *contactGroupsEndpoint) DeleteOne(id string) (err error) {
	e.RateLimitCallback()
	return DeleteContactGroup(id, e.tenantId, e.accessToken)
}

// Not Supported
func (e *contactGroupsEndpoint) DeleteMulti(ids []string) (err error) {
	return errors.ErrEndpointCallNotSupported
}
