package employees

import (
	"github.com/datag8r/xerogo/accountingAPI/endpoints"
	"github.com/datag8r/xerogo/filter"
	"github.com/datag8r/xerogo/helpers"
	"github.com/datag8r/xerogo/utils"
)

type Employee struct {
	EmployeeID   string         `xero:"update,*id"`
	FirstName    string         `xero:"create,*update"`  // Required For Creation
	LastName     string         `xero:"create,*update"`  // Required For Creation
	Status       employeeStatus `xero:"*create,*update"` // Required For Creation
	ExternalLink Link           `xero:"*create,*update"`
}

func GetEmployees(tenantId, accessToken string, where *filter.Filter) (employees []Employee, err error) {
	url := endpoints.EndpointEmployees
	request, err := helpers.BuildRequest("GET", url, nil, where, nil)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	body, err := helpers.DoRequest(request, 200)
	if err != nil {
		return
	}
	var responseBody struct {
		Employees []Employee
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	employees = responseBody.Employees
	return
}

func GetEmployee(tenantId, accessToken, employeeID string) (employee Employee, err error) {
	url := endpoints.EndpointEmployees + "/" + employeeID
	request, err := helpers.BuildRequest("GET", url, nil, nil, nil)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	b, err := helpers.DoRequest(request, 200)
	if err != nil {
		return
	}
	var responseBody struct {
		Employees []Employee
	}
	err = helpers.UnmarshalJson(b, &responseBody)
	if len(responseBody.Employees) == 1 {
		employee = responseBody.Employees[0]
	}
	return
}

func CreateEmployee(tenantId, accessToken string, employee Employee) (emp Employee, err error) {
	url := endpoints.EndpointEmployees
	inter, err := utils.XeroCustomMarshal(employee, "create")
	if err != nil {
		return
	}
	mp := map[string]interface{}{"Employees": []interface{}{inter}}
	buf, err := helpers.MarshallJsonToBuffer(mp)
	if err != nil {
		return
	}
	request, err := helpers.BuildRequest("POST", url, nil, nil, buf)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	body, err := helpers.DoRequest(request, 200)
	if err != nil {
		return
	}
	var responseBody struct {
		Employees []Employee
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	if len(responseBody.Employees) == 1 {
		emp = responseBody.Employees[0]
	}
	return
}

func CreateEmployees(tenantId, accessToken string, employees []Employee) (emps []Employee, err error) {
	url := endpoints.EndpointEmployees
	inter, err := utils.XeroCustomMarshal(employees, "create")
	if err != nil {
		return
	}
	mp := map[string]interface{}{"Employees": inter}
	buf, err := helpers.MarshallJsonToBuffer(mp)
	if err != nil {
		return
	}
	request, err := helpers.BuildRequest("POST", url, nil, nil, buf)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	body, err := helpers.DoRequest(request, 200)
	if err != nil {
		return
	}
	var responseBody struct {
		Employees []Employee
	}
	err = helpers.UnmarshalJson(body, &responseBody)
	emps = responseBody.Employees
	return
}

func UpdateEmployee(tenantId, accessToken string, employee Employee) (err error) {
	url := endpoints.EndpointEmployees
	inter, err := utils.XeroCustomMarshal(employee, "update")
	if err != nil {
		return
	}
	mp := map[string]interface{}{"Employees": []interface{}{inter}}
	buf, err := helpers.MarshallJsonToBuffer(mp)
	if err != nil {
		return
	}
	request, err := helpers.BuildRequest("POST", url, nil, nil, buf)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	_, err = helpers.DoRequest(request, 200)
	return
}

func UpdateEmployees(tenantId, accessToken string, employees []Employee) (err error) {
	url := endpoints.EndpointEmployees
	inter, err := utils.XeroCustomMarshal(employees, "update")
	if err != nil {
		return
	}
	mp := map[string]interface{}{"Employees": inter}
	buf, err := helpers.MarshallJsonToBuffer(mp)
	if err != nil {
		return
	}
	request, err := helpers.BuildRequest("POST", url, nil, nil, buf)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	_, err = helpers.DoRequest(request, 200)
	return
}

func ArchiveEmployee(tenantId, accessToken, employeeId string) (err error) {
	url := endpoints.EndpointEmployees
	var requestBody = map[string][]map[string]string{
		"Employees": {
			0: {
				"EmployeeID": employeeId,
				"Status":     EmployeeStatusArchived,
			},
		},
	}
	buf, err := helpers.MarshallJsonToBuffer(requestBody)
	if err != nil {
		return
	}
	request, err := helpers.BuildRequest("POST", url, nil, nil, buf)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	_, err = helpers.DoRequest(request, 200)
	if err != nil {
		return
	}
	return
}

func ArchiveEmployees(tenantId, accessToken string, contactIds []string) (err error) {
	url := endpoints.EndpointEmployees
	var requestBody = map[string][]map[string]string{
		"Employees": {},
	}
	for _, id := range contactIds {
		requestBody["Employees"] = append(requestBody["Employees"], map[string]string{
			"EmployeeID": id,
			"Status":     EmployeeStatusArchived,
		})
	}
	buf, err := helpers.MarshallJsonToBuffer(requestBody)
	if err != nil {
		return
	}
	request, err := helpers.BuildRequest("POST", url, nil, nil, buf)
	if err != nil {
		return
	}
	helpers.AddXeroHeaders(request, accessToken, tenantId)
	_, err = helpers.DoRequest(request, 200)
	if err != nil {
		return
	}
	return
}
