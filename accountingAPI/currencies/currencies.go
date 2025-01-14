package currencies

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/datag8r/xerogo/accountingAPI/endpoints"
	"github.com/datag8r/xerogo/filter"
	"github.com/datag8r/xerogo/utils"
)

type Currency struct {
	Code        CurrencyCode
	Description string
}

func GetCurrencies(tenantId, accessToken string, where *filter.Filter) (currencies []Currency, err error) {
	url := endpoints.EndpointCurrencies
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
		Currencies []Currency
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
	currencies = responseBody.Currencies
	return
}

func AddCurrency(tenantId, accessToken, currencyCode string) (err error) {
	url := endpoints.EndpointCurrencies
	var request *http.Request
	var requestBody struct {
		Code string
	}
	requestBody.Code = currencyCode
	b, err := json.Marshal(requestBody)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(b)
	request, err = http.NewRequest("PUT", url, buf)
	if err != nil {
		return
	}
	utils.AddXeroHeaders(request, accessToken, tenantId)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	b, err = io.ReadAll(response.Body)
	if err != nil {
		return
	}
	if response.StatusCode != 200 {
		err = errors.New(string(b))
		return
	}
	return
}
