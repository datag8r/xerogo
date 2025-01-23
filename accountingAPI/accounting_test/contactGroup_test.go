package accounting_test

import (
	"testing"

	contactgroups "github.com/datag8r/xerogo/accountingAPI/contactGroups"
	config "github.com/datag8r/xerogo/testing"
)

func TestGetContactGroups(t *testing.T) {
	conf, token, err := config.Setup(2)
	if err != nil {
		t.Fatal(err)
	}
	contactgroupList, err := contactgroups.GetContactGroups(conf.TenantID, token.AccessToken, nil)
	if err != nil {
		t.Fatal(err)
	}
	if l := len(contactgroupList); l == 0 {
		t.Fatal("len of contact groups is 0")
	}
}
func TestGetContactGroup(t *testing.T) {
	conf, token, err := config.Setup(2)
	if err != nil {
		t.Fatal(err)
	}
	conGroupId := "91dbdc3f-86c5-4bfe-b227-5d1735945cea"
	contactgroup, err := contactgroups.GetContactGroup(conf.TenantID, token.AccessToken, conGroupId)
	if err != nil {
		t.Fatal(err)
	}
	if contactgroup.Name == "" {
		t.Fatal("empty contact group name returned")
	}
}
