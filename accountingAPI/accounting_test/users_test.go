package accounting_test

import (
	"testing"

	"github.com/datag8r/xerogo/accountingAPI/users"
	config "github.com/datag8r/xerogo/testing"
)

func TestGetUsers(t *testing.T) {
	conf, token, err := config.Setup(2)
	if err != nil {
		t.Fatal(err)
	}
	allUsers, err := users.GetUsers(conf.TenantID, token.AccessToken, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(allUsers) == 0 {
		t.Fatal("No Users Returned")
	}
}

func TestGetUser(t *testing.T) {
	conf, token, err := config.Setup(2)
	if err != nil {
		t.Fatal(err)
	}
	userId := "73f0cf2f-ede4-49ea-b93e-6723d9dfca32"
	user, err := users.GetUser(conf.TenantID, token.AccessToken, userId)
	if err != nil {
		t.Fatal(err)
	}
	if user.FirstName == "" {
		t.Fatal("empty name field")
	}
}
