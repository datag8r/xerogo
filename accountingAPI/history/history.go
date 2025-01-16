package history

import "github.com/datag8r/xerogo/helpers"

type History struct {
	Changes       string
	DateUTCString string
	DateUTC       string
	User          string
	Details       string
}

func AddNoteToResource(endpoint, resourceID, note, tenantId, accessToken string) (err error) {
	if resourceID == "" {
		return ErrInvalidResourceID
	}
	url := endpoint + "/" + resourceID + "/History"
	var requestBody struct {
		HistoryRecords []struct {
			Details string
		}
	}
	requestBody.HistoryRecords = []struct{ Details string }{{note}}
	buf, err := helpers.MarshallJsonToBuffer(requestBody)
	if err != nil {
		return
	}
	request, err := helpers.BuildRequest("PUT", url, nil, nil, buf)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	_, err = helpers.DoRequest(request, 200)
	return
}

func GetResourceHistory(endpoint, resourceID, tenantId, accessToken string) (history []History, err error) {
	if resourceID == "" {
		err = ErrInvalidResourceID
		return
	}
	url := endpoint + "/" + resourceID + "/History"
	request, err := helpers.BuildRequest("GET", url, nil, nil, nil)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	body, err := helpers.DoRequest(request, 200)
	if err != nil {
		return
	}
	var responseBody struct {
		HistoryRecords []History
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	history = responseBody.HistoryRecords
	return
}
