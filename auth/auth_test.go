package auth_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/datag8r/xerogo/auth"
	"github.com/datag8r/xerogo/utils"
)

func ATestCodeExchange(t *testing.T) {
	conf, err := loadConfig()
	if err != nil {
		t.Fatal(err)
	}
	var code string
	m := http.NewServeMux()
	s := &http.Server{Addr: fmt.Sprintf("0.0.0.0:%d", 80), Handler: m}
	var codeHandler = func(w http.ResponseWriter, r *http.Request) {
		code = r.URL.Query().Get("code")
		s.Close()
	}
	m.HandleFunc("GET /", codeHandler)
	_ = s.ListenAndServe()
	identityToken, accessToken, refreshToken, err := auth.ExchangeCode(code, conf.ClientID, conf.ClientSecret, conf.RedirectURI)
	if err != nil {
		t.Fatal(err)
	}
	if len(accessToken) == 0 {
		t.FailNow()
	}
	if conf.Scopes[0] == "offline_access" && len(refreshToken) == 0 {
		t.FailNow()
	}
	token := tokenData{
		IdentityToken:   identityToken,
		AccessToken:     accessToken,
		RefreshToken:    refreshToken,
		TimeLastUpdated: time.Now(),
	}
	err = saveTokenData(token)
	if err != nil {
		t.Fatal(err)
	}
}

func ATestTokenRefresh(t *testing.T) {
	conf, err := loadConfig()
	if err != nil {
		t.Fatal(err)
	}
	token, err := loadTokenData()
	if err != nil {
		t.Fatal(err)
	}
	if token.AccessToken == "" {
		t.Fatal("Empty Access Token")
	}
	if token.RefreshToken == "" {
		t.Fatal("Empty Refresh Token")
	}
	if token.TimeLastUpdated.IsZero() {
		t.Fatal("TimeLastUpdate IsZero")
	}
	if time.Since(token.TimeLastUpdated) > (time.Hour*24)*60 {
		t.Fatal("Expired Refresh Token")
	}
	id, access, refresh, err := auth.RefreshToken(conf.ClientID, conf.ClientSecret, token.RefreshToken)
	if err != nil {
		t.Fatal(err)
	}
	newTokenData := tokenData{
		IdentityToken:   id,
		AccessToken:     access,
		RefreshToken:    refresh,
		TimeLastUpdated: time.Now(),
	}
	err = saveTokenData(newTokenData)
	if err != nil {
		t.Fatal(err)
	}
}

type testConfig struct {
	ClientID     string   `json:"clientID"`
	ClientSecret string   `json:"clientSecret"`
	Scopes       []string `json:"scopes"`
	RedirectURI  string   `json:"redirectURI"`
}

type tokenData struct {
	IdentityToken   string    `json:"identity_token"`
	AccessToken     string    `json:"access_token"`
	RefreshToken    string    `json:"refresh_token"`
	TimeLastUpdated time.Time `json:"time_last_updated"`
}

func loadConfig() (conf testConfig, err error) {
	path := utils.PathToMinus("test_config.json", 1)
	b, err := os.ReadFile(path)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &conf)
	return
}

func saveTokenData(t tokenData) error {
	path := utils.PathToMinus("tokens.json", 1)
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, b, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func loadTokenData() (t tokenData, err error) {
	path := utils.PathToMinus("tokens.json", 1)
	b, err := os.ReadFile(path)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &t)
	return
}
