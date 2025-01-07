package tenant

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	acc "github.com/datag8r/xerogo/accountingAPI/accounts"
)

type Tenant struct {
	TenantID       string `json:"tenantId"`
	Name           string `json:"tenantName"`
	CreatedDateUTC string `json:"createdDateUtc"`
	UpdatedDateUTC string `json:"updatedDateUtc"`
}

func GetTenants(accessToken string) (t []Tenant, err error) {
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

func (t *Tenant) GetAccounts(accessToken string) (accounts []acc.Account, err error) {
	return acc.GetAccounts(t.TenantID, accessToken)
}
