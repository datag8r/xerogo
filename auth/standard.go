package auth

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	netUrl "net/url"
	"strings"

	"github.com/datag8r/xerogo/auth/endpoints"
)

// Redirect User To:
// https://login.xero.com/identity/connect/authorize?response_type={code}&client_id={YOURCLIENTID}&redirect_uri={YOURREDIRECTURI}&scope={openid profile email accounting.transactions}&state={123}
// User is Returned to your redirect_uri with a code and a state
// Exchange Code For Access Token (optional: refresh token, id token)

// This function is used to generate the URL that the user will be redirected to in order to authenticate with Xero.
func NewAuthRedirectUrl(response_type string, client_id string, redirect_uri string, scope []string, state string) string {
	return "https://login.xero.com/identity/connect/authorize?response_type=" + response_type + "&client_id=" + client_id + "&redirect_uri=" + redirect_uri + "&scope=" + strings.Join(scope, "%20") + "&state=" + state
}

var (
	ErrMisMatchedState         = errors.New("state mismatch")
	ErrInvalidInput            = errors.New("one or more inputs are invalid")
	ErrOfflineAccessNotEnabled = errors.New("in order to refresh tokens, offline access scope must be provided at setup")
)

func ExchangeCode(code, clientID, clientSecret, redirectURI string) (identityToken, accessToken, refreshToken string, err error) {
	if code == "" || clientID == "" || clientSecret == "" {
		err = ErrInvalidInput
		return
	}
	url := endpoints.EndpointToken
	authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(clientID+":"+clientSecret))
	data := netUrl.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)

	request, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		return
	}
	request.Header.Add("Authorization", authHeader)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
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
	type responseBody struct {
		AccessToken   string `json:"access_token"`
		RefreshToken  string `json:"refresh_token"`
		IdentityToken string `json:"id_token"`
	}
	var resp = &responseBody{}
	err = json.Unmarshal(b, resp)
	if err != nil {
		return
	}
	accessToken = resp.AccessToken
	refreshToken = resp.RefreshToken
	identityToken = resp.IdentityToken
	return
}

func RefreshToken(clientId, clientSecret, RefreshToken string) (identityToken, accessToken, refreshToken string, err error) {
	url := endpoints.EndpointToken
	authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(clientId+":"+clientSecret))

	data := netUrl.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", RefreshToken)

	request, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		return
	}
	request.Header.Add("Authorization", authHeader)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
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
	type responseBody struct {
		AccessToken   string `json:"access_token"`
		RefreshToken  string `json:"refresh_token"`
		IdentityToken string `json:"id_token"`
	}
	var resp = &responseBody{}
	err = json.Unmarshal(b, resp)
	if err != nil {
		return
	}
	accessToken = resp.AccessToken
	refreshToken = resp.RefreshToken
	identityToken = resp.IdentityToken
	return
}
