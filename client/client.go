package client

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/datag8r/xerogo/auth"
)

type client struct {
	clientID      string
	clientSecret  string
	redirectURI   string
	scope         []string
	accessToken   string // Expiry: 30 minutes
	refreshToken  string // Expiry: 60 days
	identityToken string // Expiry: 5 minutes
	lastRefresh   time.Time
	Tenants       []tenant
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

// GetClientID returns the client ID
func (c *client) GetClientID() string {
	return c.clientID
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
	c.accessToken = accessToken
	c.refreshToken = refreshToken
	c.lastRefresh = time.Now()
	return
}

func (c *client) RefreshTokens() (err error) {
	if c.refreshToken == "" {
		return auth.ErrOfflineAccessNotEnabled
	}
	identityToken, accessToken, refreshToken, err := auth.RefreshToken(c.refreshToken, c.clientID, c.clientSecret)
	if err != nil {
		return
	}
	c.identityToken = identityToken
	c.accessToken = accessToken
	c.refreshToken = refreshToken
	c.lastRefresh = time.Now()
	return
}

func (c *client) Debug() string {
	return fmt.Sprintf("%+v", c)
}

func (c client) requiresRefresh() bool {
	return time.Since(c.lastRefresh) > 30*time.Minute
}

func (c client) GetTenants() (t []tenant, err error) {
	if c.requiresRefresh() {
		err = c.RefreshTokens()
		if err != nil {
			return
		}
	}
	t, err = getTenants(c.accessToken)
	if err == nil {
		c.Tenants = t
	}
	return
}
