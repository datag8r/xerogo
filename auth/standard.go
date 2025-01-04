package auth

import "strings"

// Redirect User To:
// https://login.xero.com/identity/connect/authorize?response_type={code}&client_id={YOURCLIENTID}&redirect_uri={YOURREDIRECTURI}&scope={openid profile email accounting.transactions}&state={123}
// User is Returned to your redirect_uri with a code and a state

// This function is used to generate the URL that the user will be redirected to in order to authenticate with Xero.
func NewAuthRedirectUrl(response_type string, client_id string, redirect_uri string, scope []string, state string) string {
	return "https://login.xero.com/identity/connect/authorize?response_type=" + response_type + "&client_id=" + client_id + "&redirect_uri=" + redirect_uri + "&scope=" + strings.Join(scope, " ") + "&state=" + state
}
