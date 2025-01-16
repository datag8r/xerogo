package filter

import (
	"fmt"
	"io"
	"net/http"
	netUrl "net/url"
	"strings"
	"time"
)

type Filter struct {
	modifiedAfter *time.Time
	where         []whereOpt
	orderBy       *orderBy
}

func NewFilter(modifiedAfter *time.Time, orderBy *orderBy, opts ...whereOpt) *Filter {
	f := &Filter{modifiedAfter: modifiedAfter, orderBy: orderBy, where: opts}
	return f
}

func (f *Filter) AddPagination(page uint, pageSize uint) {
	opts := []whereOpt{
		WhereFieldContains("page", []string{fmt.Sprint(page)}),
		WhereFieldContains("pageSize", []string{fmt.Sprint(pageSize)}),
	}
	f.where = append(f.where, opts...)
}

func (f *Filter) BuildRequest(method, url string, body io.Reader) (req *http.Request, err error) {
	req, err = http.NewRequest(method, f.buildUrl(url), body)
	if err != nil {
		return
	}
	if f.modifiedAfter != nil {
		req.Header.Add("If-Modified-Since", f.modifiedAfter.UTC().String()) // This might need fiddling
	}
	return
}

func (f *Filter) buildUrl(oldUrl string) (url string) {
	var orderString string
	var whereString string
	if len(f.where) == 1 && f.where[0] == nil {
		f.where = []whereOpt{}
	}
	if f.orderBy == nil && (len(f.where) == 0) {
		return oldUrl
	}
	url = oldUrl + "?"
	if f.orderBy != nil {
		orderString = f.orderBy.build()
	}
	url += orderString
	if len(orderString) > 0 && len(f.where) > 0 {
		url += "&"
	}
	if len(f.where) > 0 {
		var lists []string
		var fields []string
		for _, w := range f.where {
			if w.isList() {
				lists = append(lists, w.build())
			} else {
				fields = append(fields, w.build())
			}
		}

		for i, l := range lists {
			var str string
			if i != len(lists)-1 {
				str = l + "&"
			} else {
				str = l
			}
			whereString += str
		}

		if len(fields) > 0 {
			if len(lists) > 0 {
				whereString += "&"
			}
			whereString += "where="
		}
		for i, f := range fields {
			var str string
			if i != len(fields)-1 {
				str = f + "&&"
			} else {
				str = f
			}
			whereString += netUrl.QueryEscape(str)
		}
	}
	url += whereString
	return
}

type orderBy struct {
	field     string
	ascending bool
}

func OrderBy(field string, ascending bool) *orderBy {
	return &orderBy{
		field:     field,
		ascending: ascending,
	}
}

func (o orderBy) build() string {
	var directionString string
	if !o.ascending {
		directionString = "%20DESC"
	}
	return fmt.Sprintf("order=%s%s", o.field, directionString)
}

type whereOpt interface {
	isList() bool
	build() string
}

type whereField struct {
	field string
	logic string
	value string
}

func WhereFieldEquals(field, value string) whereOpt {
	return whereField{
		field: field,
		value: value,
		logic: "==",
	}
}

func WhereFieldNotEqual(field, value string) whereOpt {
	return whereField{
		field: field,
		value: value,
		logic: "!=",
	}
}

func (whereField) isList() bool { return false }

func (w whereField) build() string {
	return fmt.Sprintf("%s%s\"%s\"", w.field, w.logic, w.value)
}

type whereFieldList struct {
	field      string
	listValues []string
}

func WhereFieldContains(field string, values []string) whereOpt {
	return whereFieldList{
		field:      field,
		listValues: values,
	}
}

func (whereFieldList) isList() bool { return true }

func (w whereFieldList) build() string {
	if len(w.listValues) == 0 {
		return ""
	}
	return fmt.Sprintf("%s=%s", w.field, strings.Join(w.listValues, ","))
}
