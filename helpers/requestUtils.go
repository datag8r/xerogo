package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/datag8r/xerogo/accountingAPI/pagination"
	"github.com/datag8r/xerogo/filter"
)

func BuildRequest(method, url string, page *uint, where *filter.Filter, body io.Reader) (request *http.Request, err error) {
	if where != nil {
		if page != nil {
			where.AddPagination(*page, pagination.CustomPageSize)
		}
		request, err = where.BuildRequest(method, url, body)
	} else {
		if page != nil {
			url += "?page=" + fmt.Sprint(*page)
			if !pagination.IsDefaultPageSize() {
				url += "&pageSize=" + fmt.Sprint(pagination.CustomPageSize)
			}
		}
		request, err = http.NewRequest(method, url, body)
	}
	return
}

func AddXeroHeaders(req *http.Request, accessToken, tenantID string) {
	req.Header.Add("xero-tenant-id", tenantID)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+accessToken)
}

func DoRequest(request *http.Request, expectedStatus int) (body []byte, err error) {
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	body, err = io.ReadAll(response.Body)
	if err != nil {
		return
	}
	if response.StatusCode != expectedStatus {
		err = errors.New(string(body))
		return
	}
	return
}

func MarshallJsonToBuffer(v any) (buf *bytes.Buffer, err error) {
	b, err := json.Marshal(v)
	buf = bytes.NewBuffer(b)
	return
}

func UnmarshalJson(b []byte, v any) error {
	return json.Unmarshal(b, v)
}
