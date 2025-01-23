package client

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/datag8r/xerogo/auth"
	"github.com/datag8r/xerogo/utils"
)

type Client struct {
	clientID      string
	clientSecret  string
	redirectURI   string
	scope         []string
	AccessToken   string // Expiry: 30 minutes
	refreshToken  string // Expiry: 60 days
	identityToken string // Expiry: 5 minutes
	lastRefresh   time.Time
	Tenants       []*Tenant
	rateLimit     time.Duration `json:"-"`
	lastCall      time.Time     `json:"-"`
	rateMutex     sync.Mutex    `json:"-"`
}

// NewClient creates a new Client object
// redirectURI is the URL that the user will be redirected to after authenticating with Xero -- This must be an https address, however for testing you can use http://localhost/ Please note that http://127.0.0.1 cannot be used
// scope is a list of the permissions that your app will need to access -- https://developer.xero.com/documentation/guides/oauth2/scopes/
func NewClient(clientID string, clientSecret string, redirectURI string, scope []string) *Client {
	return &Client{
		clientID:     clientID,
		clientSecret: clientSecret,
		redirectURI:  redirectURI,
		scope:        scope,
		rateLimit:    time.Millisecond * 6, // Default Rate Limited For Client: 10000 / minute == 6ms minimum between requests
	}
}

type tokenData struct {
	IdentityToken   string    `json:"identity_token"`
	AccessToken     string    `json:"access_token"`
	RefreshToken    string    `json:"refresh_token"`
	TimeLastUpdated time.Time `json:"time_last_updated"`
}

func (c *Client) SaveTokenDataToJsonFile(filePath string) (err error) {
	path := utils.PathTo(filePath)
	var t tokenData = tokenData{
		IdentityToken:   c.identityToken,
		AccessToken:     c.AccessToken,
		RefreshToken:    c.refreshToken,
		TimeLastUpdated: c.lastRefresh,
	}
	b, err := json.Marshal(t)
	if err != nil {
		return
	}
	err = os.WriteFile(path, b, os.ModePerm)
	return
}

func (c *Client) LoadTokenDataFromJsonFile(filePath string) (err error) {
	path := utils.PathTo(filePath)
	b, err := os.ReadFile(path)
	if err != nil {
		return
	}
	var t tokenData
	err = json.Unmarshal(b, &t)
	if err != nil {
		return
	}
	c.identityToken = t.IdentityToken
	c.AccessToken = t.AccessToken
	c.refreshToken = t.RefreshToken
	c.lastRefresh = t.TimeLastUpdated
	return
}

func (c *Client) GetStandardAuthRedirectURL() (url string, state string) {
	state = fmt.Sprintf("%d", rand.Int31n(1000))
	url = auth.NewAuthRedirectUrl("code", c.clientID, c.redirectURI, c.scope, state)
	return
}

func (c *Client) VerifyStandardAuthRedirectCode(code, state, expectedState string) (err error) {
	c.Call()
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

func (c *Client) Refresh() error {
	if c.requiresRefresh() {
		c.Call()
		return c.refreshTokens()
	}
	return nil
}

func (c *Client) refreshTokens() (err error) {
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

func (c *Client) requiresRefresh() bool {
	return time.Since(c.lastRefresh) > 30*time.Minute
}

func (c *Client) GetTenants() (t []*Tenant, err error) {
	err = c.Refresh()
	if err != nil {
		return
	}
	return c.getTenants()
}

func (c *Client) GetTenant(tenantId string) (t *Tenant, err error) {
	err = c.Refresh()
	if err != nil {
		return
	}
	return c.getTenant(tenantId)
}

// impl get tenant (single tenant)
func (c *Client) getTenants() (t []*Tenant, err error) {
	c.Call()
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
			ten.Client = c
			ten.tenantEndpoints = NewTenantEndpoints(ten)
			ten.rateLimit = time.Second
			ten.lastCall = time.Now()
		}
		c.Tenants = t
	}
	return
}

func (c *Client) getTenant(tenantId string) (tenant *Tenant, err error) {
	c.Call()
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
	t := []*Tenant{}
	err = json.Unmarshal(b, &t)
	if err == nil {
		for _, ten := range t {
			ten.Client = c
			ten.tenantEndpoints = NewTenantEndpoints(ten)
			ten.rateLimit = time.Second
			ten.lastCall = time.Now()
			if ten.TenantID == tenantId {
				tenant = ten
			}
		}
		c.Tenants = t
	}
	return
}

// This Function is Used For Rate Limiting, To avoid being rate lmited you should call this before every request
// Built In (*Tenant) and (*Client) Methods Call This For you
func (c *Client) Call() {
	c.rateMutex.Lock()
	next := c.lastCall.Add(c.rateLimit) // Min Time Of Next Call
	current := time.Now()
	toWait := next.Sub(current)
	<-time.After(toWait)
	c.lastCall = time.Now()
	c.rateMutex.Unlock()
}
