package accounting_test

import (
	"fmt"
	"testing"

	"github.com/datag8r/xerogo/accountingAPI/currencies"
	config "github.com/datag8r/xerogo/testing"
)

func TestGetCurrencies(t *testing.T) {
	conf, token, err := config.Setup(2)
	if err != nil {
		t.Fatal(err)
	}
	currencyList, err := currencies.GetCurrencies(conf.TenantID, token.AccessToken, nil)
	if err != nil {
		t.Fatal(err)
	}
	if l := len(currencyList); l == 0 {
		t.Fatal("len of accs: " + fmt.Sprint(l))
	}
}

func TestAddCurrency(t *testing.T) {
	conf, token, err := config.Setup(2)
	if err != nil {
		t.Fatal(err)
	}
	currencyCode := "SGD"
	cur := currencies.Currency{Code: currencyCode}
	cur, err = currencies.CreateCurrency(conf.TenantID, token.AccessToken, cur)
	if err != nil {
		t.Fatal(err)
	}
	if cur.Code != currencyCode {
		t.Fatal("currency code mismatch")
	}
}
