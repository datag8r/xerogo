package accounting_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/datag8r/xerogo/accountingAPI/accounts"
	"github.com/datag8r/xerogo/auth"
	"github.com/datag8r/xerogo/utils"
)

func TestGetAccounts(t *testing.T) {
	conf, token, err := setup()
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
	conf, token, err := setup()
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

func TestCreateAccount(t *testing.T) {
	conf, token, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	a := accounts.Account{Code: "TESTNEW4", Name: "Test New Account4", Type: accounts.AccountTypeExpense}
	a, err = accounts.CreateAccount(a, conf.TenantID, token.AccessToken)
	if err != nil {
		t.Fatal(err)
	}
	if a.AccountID == "" {
		t.Fatal("Account ID Field Empty after creation")
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

func setup() (conf testConfig, token tokenData, err error) {
	conf, err = loadConfig()
	if err != nil {
		return
	}
	token, err = loadTokenData()
	if err != nil {
		return
	}
	if token.AccessToken == "" {
		err = errors.New("Empty Access Token")
		return
	}
	if token.RefreshToken == "" {
		err = errors.New("Empty Refresh Token")
		return
	}
	if token.TimeLastUpdated.IsZero() {
		err = errors.New("TimeLastUpdate IsZero")
		return
	}
	if time.Since(token.TimeLastUpdated) > (time.Hour*24)*60 {
		err = errors.New("Expired Refresh Token")
		return
	}
	id, access, refresh, err := auth.RefreshToken(conf.ClientID, conf.ClientSecret, token.RefreshToken)
	if err != nil {
		return
	}
	token.IdentityToken = id
	token.AccessToken = access
	token.RefreshToken = refresh
	token.TimeLastUpdated = time.Now()
	err = saveTokenData(token)
	return
}
