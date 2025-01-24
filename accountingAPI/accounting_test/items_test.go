package accounting_test

import (
	"testing"

	"github.com/datag8r/xerogo/accountingAPI/endpoints"
	"github.com/datag8r/xerogo/accountingAPI/history"
	"github.com/datag8r/xerogo/accountingAPI/items"
	config "github.com/datag8r/xerogo/testing"
)

func TestGetItems(t *testing.T) {
	conf, token, err := config.Setup(2)
	if err != nil {
		t.Fatal(err)
	}
	allItems, err := items.GetItems(conf.TenantID, token.AccessToken, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(allItems) == 0 {
		t.Fatal("No Items Returned")
	}
}

func TestGetItem(t *testing.T) {
	conf, token, err := config.Setup(2)
	if err != nil {
		t.Fatal(err)
	}
	item, err := items.GetItem("TESTItemCode", conf.TenantID, token.AccessToken)
	if err != nil {
		t.Fatal(err)
	}
	if item.ItemID == "" {
		t.Fatal("Invalid Item ID")
	}
}

func TestGetItemHistory(t *testing.T) {
	conf, token, err := config.Setup(2)
	if err != nil {
		t.Fatal(err)
	}
	history, err := history.GetResourceHistory(endpoints.EndpointItems, "14df66b4-463b-4b1e-b302-b39cc2865304", conf.TenantID, token.AccessToken)
	if err != nil {
		t.Fatal(err)
	}
	if len(history) == 0 {
		t.Fatal("No History Returned")
	}
}

func TestCreateItem(t *testing.T) {
	i := items.Item{
		Code: "TESTItemCode",
	}
	conf, token, err := config.Setup(2)
	if err != nil {
		t.Fatal(err)
	}
	i, err = items.CreateItem(conf.TenantID, token.AccessToken, i)
	if err != nil {
		t.Fatal(err)
	}
	if i.ItemID == "" {
		t.Fatal("Item ID Field Empty after creation")
	}
}

func TestUpdateItem(t *testing.T) {
	i := items.Item{
		Code: "TESTItemCode",
		Name: "Updated name TEST",
	}
	conf, token, err := config.Setup(2)
	if err != nil {
		t.Fatal(err)
	}
	err = items.UpdateItem(conf.TenantID, token.AccessToken, i)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteItem(t *testing.T) {
	conf, token, err := config.Setup(2)
	if err != nil {
		t.Fatal(err)
	}
	err = items.DeleteItem("591f6fc1-17f9-4b03-9788-d844b70656f4", conf.TenantID, token.AccessToken)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddNoteToItem(t *testing.T) {
	conf, token, err := config.Setup(2)
	if err != nil {
		t.Fatal(err)
	}
	err = history.AddNoteToResource(endpoints.EndpointItems, "14df66b4-463b-4b1e-b302-b39cc2865304", "test note auto", conf.TenantID, token.AccessToken)
	if err != nil {
		t.Fatal(err)
	}
}
