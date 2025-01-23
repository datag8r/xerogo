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
	err = currencies.AddCurrency(conf.TenantID, token.AccessToken, currencyCode)
	if err != nil {
		t.Fatal(err)
	}
}
