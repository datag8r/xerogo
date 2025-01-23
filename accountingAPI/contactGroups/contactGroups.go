package contactgroups

import (
	"github.com/datag8r/xerogo/accountingAPI/endpoints"
	"github.com/datag8r/xerogo/filter"
	"github.com/datag8r/xerogo/helpers"
	"github.com/datag8r/xerogo/utils"
)

type ContactGroup struct {
	ContactGroupID string              `xero:"update"` // Only used on update
	Name           string              `xero:"*create,*update"`
	Status         contactGroupStatus  `xero:"*create,*update"`
	Contacts       []contactIdentifier `xero:"*create,*update"` // Only filled on Get by group id
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

func CreateContactGroup(contactGroup ContactGroup, tenantId, accessToken string) (createdContactGroup ContactGroup, err error) {
	url := endpoints.EndpointContactGroups
	inter, err := utils.XeroCustomMarshal(contactGroup, "create")
	if err != nil {
		return
	}
	buf, err := helpers.MarshallJsonToBuffer(inter)
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
		ContactGroups []ContactGroup
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	if len(responseBody.ContactGroups) == 1 {
		createdContactGroup = responseBody.ContactGroups[0]
	}
	return
}

func CreateContactGroups(contactGroups []ContactGroup, tenantId, accessToken string) (createdContactGroups []ContactGroup, err error) {
	url := endpoints.EndpointContactGroups
	inter, err := utils.XeroCustomMarshal(contactGroups, "create")
	if err != nil {
		return
	}
	buf, err := helpers.MarshallJsonToBuffer(inter)
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
		ContactGroups []ContactGroup
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	createdContactGroups = responseBody.ContactGroups
	return
}

func UpdateContactGroup(contactGroup ContactGroup, tenantId, accessToken string) (err error) {
	url := endpoints.EndpointContactGroups
	inter, err := utils.XeroCustomMarshal(contactGroup, "update")
	if err != nil {
		return
	}
	buf, err := helpers.MarshallJsonToBuffer(inter)
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

func UpdateContactGroups(contactGroups []ContactGroup, tenantId, accessToken string) (err error) {
	url := endpoints.EndpointContactGroups
	inter, err := utils.XeroCustomMarshal(contactGroups, "update")
	if err != nil {
		return
	}
	buf, err := helpers.MarshallJsonToBuffer(inter)
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

func DeleteContactGroup(tenantId, accessToken, contactGroupId string) (err error) {
	url := endpoints.EndpointContactGroups + "/" + contactGroupId
	mp := map[string]string{"ContactGroupID": contactGroupId, "Status": "DELETED"}
	buf, err := helpers.MarshallJsonToBuffer(mp)
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
