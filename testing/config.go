package testing

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/datag8r/xerogo/auth"
	"github.com/datag8r/xerogo/utils"
)

type TestConfig struct {
	ClientID     string   `json:"clientID"`
	ClientSecret string   `json:"clientSecret"`
	Scopes       []string `json:"scopes"`
	RedirectURI  string   `json:"redirectURI"`
	TenantID     string   `json:"tenantID"`
}

type TokenData struct {
	IdentityToken   string    `json:"identity_token"`
	AccessToken     string    `json:"access_token"`
	RefreshToken    string    `json:"refresh_token"`
	TimeLastUpdated time.Time `json:"time_last_updated"`
}

func loadConfig(filedepth int) (conf TestConfig, err error) {
	path := utils.PathToMinus("test_config.json", filedepth)
	b, err := os.ReadFile(path)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &conf)
	return
}

func saveTokenData(t TokenData, filedepth int) error {
	path := utils.PathToMinus("tokens.json", filedepth)
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

func loadTokenData(filedepth int) (t TokenData, err error) {
	path := utils.PathToMinus("tokens.json", filedepth)
	b, err := os.ReadFile(path)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &t)
	return
}

func Setup(filedepth int) (conf TestConfig, token TokenData, err error) {
	conf, err = loadConfig(filedepth)
	if err != nil {
		return
	}
	token, err = loadTokenData(filedepth)
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
	err = saveTokenData(token, filedepth)
	return
}
