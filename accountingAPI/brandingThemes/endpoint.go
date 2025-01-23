package brandingthemes

import (
	"github.com/datag8r/xerogo/errors"
	"github.com/datag8r/xerogo/filter"
)

type brandingThemeEndpoint struct {
	tenantId          string
	accessToken       string
	rateLimitCallback func()
}

func NewBrandingThemeEndpoint(tenantId, accessToken string, rateLimitCallback func()) *brandingThemeEndpoint {
	return &brandingThemeEndpoint{
		tenantId:          tenantId,
		accessToken:       accessToken,
		rateLimitCallback: rateLimitCallback,
	}
}

func (e *brandingThemeEndpoint) RateLimitCallback() {
	if e.rateLimitCallback != nil {
		e.rateLimitCallback()
	}
}

func (e *brandingThemeEndpoint) GetOne(id string) (theme BrandingTheme, err error) {
	e.RateLimitCallback()
	return GetBrandingTheme(e.tenantId, e.accessToken, id)
}

func (e *brandingThemeEndpoint) GetMulti(where *filter.Filter) (themes []BrandingTheme, err error) {
	e.RateLimitCallback()
	return GetBrandingThemes(e.tenantId, e.accessToken, where)
}

// Not Supported
func (e *brandingThemeEndpoint) CreateOne(account BrandingTheme) (theme BrandingTheme, err error) {
	return BrandingTheme{}, errors.ErrEndpointCallNotSupported
}

// Not Supported
func (e *brandingThemeEndpoint) CreateMulti(accounts []BrandingTheme) (themes []BrandingTheme, err error) {
	return nil, errors.ErrEndpointCallNotSupported
}

// Not Supported
func (e *brandingThemeEndpoint) UpdateOne(account BrandingTheme) (err error) {
	return errors.ErrEndpointCallNotSupported
}

// Not Supported
func (e *brandingThemeEndpoint) UpdateMulti(accounts []BrandingTheme) (err error) {
	return errors.ErrEndpointCallNotSupported
}

// Not Supported
func (e *brandingThemeEndpoint) ArchiveOne(id string) (err error) {
	return errors.ErrEndpointCallNotSupported
}

// Not Supported
func (e *brandingThemeEndpoint) ArchiveMulti(ids []string) (err error) {
	return errors.ErrEndpointCallNotSupported
}

// Not Supported
func (e *brandingThemeEndpoint) DeleteOne(id string) (err error) {
	return errors.ErrEndpointCallNotSupported
}

// Not Supported
func (e *brandingThemeEndpoint) DeleteMulti(ids []string) (err error) {
	return errors.ErrEndpointCallNotSupported
}
