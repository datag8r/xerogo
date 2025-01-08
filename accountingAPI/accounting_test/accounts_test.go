package accountingtest

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/datag8r/xerogo/accountingAPI/accounts"
	"github.com/datag8r/xerogo/auth"
	"github.com/datag8r/xerogo/utils"
)

func TestGetAccounts(t *testing.T) {
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
	accs, err := accounts.GetAccounts(conf.TenantID, token.AccessToken, nil)
	if err != nil {
		t.Fatal(err)
	}
	if l := len(accs); l == 0 {
		t.Fatal("len of accs: " + fmt.Sprint(l))
	}
}

func TestUpdateAccount(t *testing.T) {
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
	accs, err := accounts.GetAccounts(conf.TenantID, token.AccessToken, nil)
	if err != nil {
		t.Fatal(err)
	}
	if l := len(accs); l == 0 {
		t.Fatal("len of accs: " + fmt.Sprint(l))
	}

	a := accs[0]
	var before = a.Name

	a.Name = before + "TEST"

	err = accounts.UpdateAccount(a, conf.TenantID, token.AccessToken)
	if err != nil {
		t.Fatal(err)
	}
	accs, err = accounts.GetAccounts(conf.TenantID, token.AccessToken, nil)
	if err != nil {
		t.Fatal(err)
	}
	if l := len(accs); l == 0 {
		t.Fatal("len of accs: " + fmt.Sprint(l))
	}

	a = accs[0]
	var after = a.Name
	if before == after {
		t.Fatal("name didnt change:\n\tbefore:\t" + before + "\n\tafter:" + after)
	}
}

type testConfig struct {
	ClientID     string   `json:"clientID"`
	ClientSecret string   `json:"clientSecret"`
	Scopes       []string `json:"scopes"`
	RedirectURI  string   `json:"redirectURI"`
	TenantID     string   `json:"tenantID"`
}

type tokenData struct {
	IdentityToken   string    `json:"identity_token"`
	AccessToken     string    `json:"access_token"`
	RefreshToken    string    `json:"refresh_token"`
	TimeLastUpdated time.Time `json:"time_last_updated"`
}

func loadConfig() (conf testConfig, err error) {
	path := utils.PathToMinus("test_config.json", 2)
	b, err := os.ReadFile(path)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &conf)
	return
}

func saveTokenData(t tokenData) error {
	path := utils.PathToMinus("tokens.json", 2)
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
	path := utils.PathToMinus("tokens.json", 2)
	b, err := os.ReadFile(path)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &t)
	return
}
