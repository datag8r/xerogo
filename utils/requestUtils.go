package utils

import "net/http"

func AddXeroHeaders(req *http.Request, accessToken, tenantID string) {
	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("xero-tenant-id", tenantID)
}
