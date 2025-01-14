package users

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/datag8r/xerogo/accountingAPI/endpoints"
	"github.com/datag8r/xerogo/filter"
	"github.com/datag8r/xerogo/utils"
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
		Users []User
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
	users = responseBody.Users
	return
}

func GetUser(tenantId, accessToken, userId string) (user User, err error) {
	if len(userId) != len("297c2dc5-cc47-4afd-8ec8-74990b8761e9") { // figure out the number
		err = ErrInvalidUserID
		return
	}
	url := endpoints.EndpointUsers + "/" + userId
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
	if response.StatusCode != 200 {
		err = errors.New(string(b))
		return
	}
	var responseBody struct {
		Users []User
	}
	err = json.Unmarshal(b, &responseBody)
	if len(responseBody.Users) == 1 {
		user = responseBody.Users[0]
	}
	return
}
