package brandingthemes

import (
	"github.com/datag8r/xerogo/accountingAPI/endpoints"
	"github.com/datag8r/xerogo/filter"
	"github.com/datag8r/xerogo/helpers"
)

type BrandingTheme struct {
	BrandingThemeID string `xero:"id"`
	Name            string
	LogoURL         string
	Type            string // document type
	SortOrder       int
	CreatedDateUTC  string
}

func GetBrandingThemes(tenantId, accessToken string, where *filter.Filter) (themes []BrandingTheme, err error) {
	url := endpoints.EndpointBrandingThemes
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
		BrandingThemes []BrandingTheme
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	themes = responseBody.BrandingThemes
	return
}

func GetBrandingTheme(tenantId, accessToken, brandingThemeId string) (theme BrandingTheme, err error) {
	url := endpoints.EndpointBrandingThemes + "/" + brandingThemeId
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
		BrandingThemes []BrandingTheme
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	if len(responseBody.BrandingThemes) == 1 {
		theme = responseBody.BrandingThemes[0]
	}
	return
}

// // for applying a payment service to a branding theme
// func UpdateBrandingTheme(tenantId, accessToken string, theme BrandingTheme) (err error) {
// 	// POST /BrandingThemes/{brandingThemeId}/PaymentServices
// 	return
// }

// not sure how best to implement the sub endpoint for payment services
