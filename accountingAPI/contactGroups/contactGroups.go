package contactgroups

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/datag8r/xerogo/accountingAPI/endpoints"
	"github.com/datag8r/xerogo/filter"
	"github.com/datag8r/xerogo/utils"
)

type ContactGroup struct {
	ContactGroupID string
	Name           string
	Status         contactGroupStatus
	Contacts       []contactIdentifier // Only filled on Get by group id
}

func GetContactGroups(tenantId, accessToken string, where *filter.Filter) (contactGroups []ContactGroup, err error) {
	url := endpoints.EndpointContactGroups
	var request *http.Request
	if where != nil {
		request, err = where.BuildRequest("GET", url, nil)
	} else {
		request, err = http.NewRequest("GET", url, nil)
	}
	if err != nil {
		return
	}
	utils.AddXeroHeaders(request, accessToken, tenantId)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	var responseBody struct {
		ContactGroups []ContactGroup
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
	contactGroups = responseBody.ContactGroups
	return
}

func GetContactGroup(tenantId, accessToken, contactGroupId string) (contactGroup ContactGroup, err error) {
	url := endpoints.EndpointContactGroups + "/" + contactGroupId
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
	var responseBody struct {
		ContactGroups []ContactGroup
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
	if len(responseBody.ContactGroups) == 1 {
		contactGroup = responseBody.ContactGroups[0]
	}
	return
}
