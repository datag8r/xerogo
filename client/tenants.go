package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type tenant struct {
	TenantID       string `json:"tenantId"`
	Name           string `json:"tenantName"`
	CreatedDateUTC string `json:"createdDateUtc"`
	UpdatedDateUTC string `json:"updatedDateUtc"`
}

func getTenants(accessToken string) (t []tenant, err error) {
	url := "https://api.xero.com/connections"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	request.Header.Add("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	b, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &t)
	return
}
