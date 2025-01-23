package accounting_test

import (
	"testing"

	"github.com/datag8r/xerogo/accountingAPI/accounts"
	config "github.com/datag8r/xerogo/testing"
)

func TestGetAccounts(t *testing.T) {
	conf, token, err := config.Setup(2)
	if err != nil {
		t.Fatal(err)
	}
	accs, err := accounts.GetAccounts(conf.TenantID, token.AccessToken, nil)
	if err != nil {
		t.Fatal(err)
	}
	if l := len(accs); l == 0 {
		t.Fatal("len of accs is 0")
	}
}
func TestGetAccount(t *testing.T) {
	conf, token, err := config.Setup(2)
	if err != nil {
		t.Fatal(err)
	}
	acc, err := accounts.GetAccount("bd9e85e0-0478-433d-ae9f-0b3c4f04bfe4", conf.TenantID, token.AccessToken)
	if err != nil {
		t.Fatal(err)
	}
	if acc.AccountID == "" {
		t.Fatal("Account ID Field Empty")
	}
}

func TestUpdateAccount(t *testing.T) {
	conf, token, err := config.Setup(2)
	if err != nil {
		t.Fatal(err)
	}
	accs, err := accounts.GetAccounts(conf.TenantID, token.AccessToken, nil)
	if err != nil {
		t.Fatal(err)
	}
	if l := len(accs); l == 0 {
		t.Fatal("len of accs is 0")
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
		t.Fatal("len of accs is 0")
	}

	a = accs[0]
	var after = a.Name
	if before == after {
		t.Fatal("name didnt change:\n\tbefore:\t" + before + "\n\tafter:" + after)
	}
}

func TestCreateAccount(t *testing.T) {
	conf, token, err := config.Setup(2)
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
