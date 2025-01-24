package currencies

import (
	"github.com/datag8r/xerogo/accountingAPI/endpoints"
	"github.com/datag8r/xerogo/filter"
	"github.com/datag8r/xerogo/helpers"
	"github.com/datag8r/xerogo/utils"
)

type Currency struct {
	Code        CurrencyCode `xero:"create"`
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

func CreateCurrency(tenantId, accessToken string, currencyToCreate Currency) (currency Currency, err error) {
	url := endpoints.EndpointCurrencies
	inter, err := utils.XeroCustomMarshal(currencyToCreate, "create")
	if err != nil {
		return
	}
	mp := map[string]interface{}{"Currencies": []interface{}{inter}}
	buf, err := helpers.MarshallJsonToBuffer(mp)
	if err != nil {
		return
	}
	request, err := helpers.BuildRequest("PUT", url, nil, nil, buf)
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
	if len(responseBody.Currencies) == 1 {
		currency = responseBody.Currencies[0]
	}
	return
}

func CreateCurrencies(tenantId, accessToken string, currenciesToCreate []Currency) (currencies []Currency, err error) {
	url := endpoints.EndpointCurrencies
	inter, err := utils.XeroCustomMarshal(currenciesToCreate, "create")
	if err != nil {
		return
	}
	mp := map[string]interface{}{"Currencies": inter}
	buf, err := helpers.MarshallJsonToBuffer(mp)
	if err != nil {
		return
	}
	request, err := helpers.BuildRequest("PUT", url, nil, nil, buf)
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
