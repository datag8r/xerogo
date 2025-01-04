package auth

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

// Redirect User To:
// https://login.xero.com/identity/connect/authorize?response_type={code}&client_id={YOURCLIENTID}&redirect_uri={YOURREDIRECTURI}&scope={openid profile email accounting.transactions}&state={123}
// User is Returned to your redirect_uri with a code and a state
// Exchange Code For Access Token (optional: refresh token, id token)

// This function is used to generate the URL that the user will be redirected to in order to authenticate with Xero.
func NewAuthRedirectUrl(response_type string, client_id string, redirect_uri string, scope []string, state string) string {
	return "https://login.xero.com/identity/connect/authorize?response_type=" + response_type + "&client_id=" + client_id + "&redirect_uri=" + redirect_uri + "&scope=" + strings.Join(scope, " ") + "&state=" + state
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
	url := "https://identity.xero.com/connect/token"
	authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(clientID+":"+clientSecret))
	type tokenRequest struct {
		Grant_type   string `json:"grant_type"`
		Code         string `json:"code"`
		Redirect_uri string `json:"redirect_uri"`
	}
	var reqBody = tokenRequest{
		Grant_type:   "authorization_code",
		Code:         code,
		Redirect_uri: redirectURI,
	}
	b, err := json.Marshal(reqBody)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(b)
	request, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return
	}
	request.Header.Add("Authorization", authHeader)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, err := authClient.Do(request)
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
	url := "https://identity.xero.com/connect/token"
	authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(clientId+":"+clientSecret))
	type tokenRequest struct {
		GrantType    string `json:"grant_type"`
		RefreshToken string `json:"refresh_token"`
	}
	var reqBody = tokenRequest{
		GrantType:    "refresh_token",
		RefreshToken: RefreshToken,
	}
	b, err := json.Marshal(reqBody)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(b)
	request, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return
	}
	request.Header.Add("Authorization", authHeader)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, err := authClient.Do(request)
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
