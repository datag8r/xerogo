package client

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/datag8r/xerogo/auth"
)

type client struct {
	clientID      string
	clientSecret  string
	redirectURI   string
	scope         []string
	AccessToken   string // Expiry: 30 minutes
	refreshToken  string // Expiry: 60 days
	identityToken string // Expiry: 5 minutes
	lastRefresh   time.Time
	Tenants       []*Tenant
}

// NewClient creates a new client object
// redirectURI is the URL that the user will be redirected to after authenticating with Xero -- This must be an https address, however for testing you can use http://localhost/ Please note that http://127.0.0.1 cannot be used
// scope is a list of the permissions that your app will need to access -- https://developer.xero.com/documentation/guides/oauth2/scopes/
func NewClient(clientID string, clientSecret string, redirectURI string, scope []string) *client {
	return &client{
		clientID:     clientID,
		clientSecret: clientSecret,
		redirectURI:  redirectURI,
		scope:        scope,
	}
}

func (c *client) GetStandardAuthRedirectURL() (url string, state string) {
	state = fmt.Sprintf("%d", rand.Int31n(1000))
	url = auth.NewAuthRedirectUrl("code", c.clientID, c.redirectURI, c.scope, state)
	return
}

func (c *client) VerifyStandardAuthRedirectCode(code, state, expectedState string) (err error) {
	if state != expectedState {
		return auth.ErrMisMatchedState
	}
	if code == "" {
		return auth.ErrInvalidInput
	}
	identityToken, accessToken, refreshToken, err := auth.ExchangeCode(code, c.clientID, c.clientSecret, c.redirectURI)
	if err != nil {
		return
	}
	c.identityToken = identityToken
	c.AccessToken = accessToken
	c.refreshToken = refreshToken
	c.lastRefresh = time.Now()
	return
}

func (c *client) Refresh() error {
	if c.requiresRefresh() {
		return c.refreshTokens()
	}
	return nil
}

func (c *client) refreshTokens() (err error) {
	if c.refreshToken == "" {
		return auth.ErrOfflineAccessNotEnabled
	}
	identityToken, accessToken, refreshToken, err := auth.RefreshToken(c.refreshToken, c.clientID, c.clientSecret)
	if err != nil {
		return
	}
	c.identityToken = identityToken
	c.AccessToken = accessToken
	c.refreshToken = refreshToken
	c.lastRefresh = time.Now()
	return
}

func (c client) requiresRefresh() bool {
	return time.Since(c.lastRefresh) > 30*time.Minute
}

func (c client) GetTenants() (t []*Tenant, err error) {
	err = c.Refresh()
	if err != nil {
		return
	}
	return c.getTenants()
}

func (c *client) getTenants() (t []*Tenant, err error) {
	url := "https://api.xero.com/connections"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
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
	if err == nil {
		for _, ten := range t {
			ten.c = c
		}
		c.Tenants = t
	}
	return
}
