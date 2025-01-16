package users

import (
	"github.com/datag8r/xerogo/accountingAPI/endpoints"
	"github.com/datag8r/xerogo/filter"
	"github.com/datag8r/xerogo/helpers"
)

type User struct {
	UserID           string
	EmailAddress     string
	FirstName        string
	LastName         string
	UpdatedDateUTC   string
	IsSubscriber     bool
	OrganisationRole userRole
}

func GetUsers(tenantId, accessToken string, where *filter.Filter) (users []User, err error) {
	url := endpoints.EndpointUsers
	request, err := helpers.BuildRequest("GET", url, nil, where, nil)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	body, err := helpers.DoRequest(request, 200)
	if err != nil {
		return
	}
	var responseBody struct {
		Users []User
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	users = responseBody.Users
	return
}

func GetUser(tenantId, accessToken, userId string) (user User, err error) {
	if len(userId) != len("297c2dc5-cc47-4afd-8ec8-74990b8761e9") { // figure out the number
		err = ErrInvalidUserID
		return
	}
	url := endpoints.EndpointUsers + "/" + userId
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
		Users []User
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	if len(responseBody.Users) == 1 {
		user = responseBody.Users[0]
	}
	return
}
