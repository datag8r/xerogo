package currencies

import (
	"github.com/datag8r/xerogo/accountingAPI/endpoints"
	"github.com/datag8r/xerogo/filter"
	"github.com/datag8r/xerogo/helpers"
)

type Currency struct {
	Code        CurrencyCode
	Description string
}

func GetCurrencies(tenantId, accessToken string, where *filter.Filter) (currencies []Currency, err error) {
	url := endpoints.EndpointCurrencies
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
		Currencies []Currency
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	currencies = responseBody.Currencies
	return
}

func AddCurrency(tenantId, accessToken, currencyCode string) (err error) {
	url := endpoints.EndpointCurrencies
	var requestBody struct {
		Code string
	}
	requestBody.Code = currencyCode
	buf, err := helpers.MarshallJsonToBuffer(requestBody)
	if err != nil {
		return
	}
	request, err := helpers.BuildRequest("PUT", url, nil, nil, buf)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	_, err = helpers.DoRequest(request, 200)
	return
}
