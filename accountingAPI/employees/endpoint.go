package employees

import (
	"github.com/datag8r/xerogo/errors"
	"github.com/datag8r/xerogo/filter"
)

type employeesEndpoint struct {
	tenantId          string
	accessToken       string
	rateLimitCallback func()
}

func NewEmployeesEndpoint(tenantId, accessToken string, rateLimitCallback func()) *employeesEndpoint {
	return &employeesEndpoint{
		tenantId:          tenantId,
		accessToken:       accessToken,
		rateLimitCallback: rateLimitCallback,
	}
}

func (e *employeesEndpoint) RateLimitCallback() {
	if e.rateLimitCallback != nil {
		e.rateLimitCallback()
	}
}

func (e *employeesEndpoint) GetOne(id string) (emp Employee, err error) {
	e.RateLimitCallback()
	return GetEmployee(e.tenantId, e.accessToken, id)
}

func (e *employeesEndpoint) GetMulti(where *filter.Filter) (accs []Employee, err error) {
	e.RateLimitCallback()
	return GetEmployees(e.tenantId, e.accessToken, where)
}

func (e *employeesEndpoint) CreateOne(employee Employee) (emp Employee, err error) {
	e.RateLimitCallback()
	return CreateEmployee(e.tenantId, e.accessToken, employee)
}

func (e *employeesEndpoint) CreateMulti(employees []Employee) (emps []Employee, err error) {
	e.RateLimitCallback()
	return CreateEmployees(e.tenantId, e.accessToken, employees)
}

func (e *employeesEndpoint) UpdateOne(employee Employee) (err error) {
	e.RateLimitCallback()
	return UpdateEmployee(e.tenantId, e.accessToken, employee)
}

func (e *employeesEndpoint) UpdateMulti(employees []Employee) (err error) {
	e.RateLimitCallback()
	return UpdateEmployees(e.tenantId, e.accessToken, employees)
}

func (e *employeesEndpoint) ArchiveOne(id string) (err error) {
	e.RateLimitCallback()
	return ArchiveEmployee(id, e.tenantId, e.accessToken)
}

func (e *employeesEndpoint) ArchiveMulti(ids []string) (err error) {
	e.RateLimitCallback()
	return ArchiveEmployees(e.tenantId, e.accessToken, ids)
}

// Not Supported
func (e *employeesEndpoint) DeleteOne(id string) (err error) {
	return errors.ErrEndpointCallNotSupported
}

// Not Supported
func (e *employeesEndpoint) DeleteMulti(ids []string) (err error) {
	return errors.ErrEndpointCallNotSupported
}
