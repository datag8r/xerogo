package contactgroups

import (
	"github.com/datag8r/xerogo/accountingAPI/endpoints"
	"github.com/datag8r/xerogo/filter"
	"github.com/datag8r/xerogo/helpers"
)

type ContactGroup struct {
	ContactGroupID string
	Name           string
	Status         contactGroupStatus
	Contacts       []contactIdentifier // Only filled on Get by group id
}

func GetContactGroups(tenantId, accessToken string, where *filter.Filter) (contactGroups []ContactGroup, err error) {
	url := endpoints.EndpointContactGroups
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
		ContactGroups []ContactGroup
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	contactGroups = responseBody.ContactGroups
	return
}

func GetContactGroup(tenantId, accessToken, contactGroupId string) (contactGroup ContactGroup, err error) {
	url := endpoints.EndpointContactGroups + "/" + contactGroupId
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
		ContactGroups []ContactGroup
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	if len(responseBody.ContactGroups) == 1 {
		contactGroup = responseBody.ContactGroups[0]
	}
	return
}
