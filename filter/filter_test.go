package filter_test

import (
	"strings"
	"testing"
	"time"

	"github.com/datag8r/xerogo/filter"
)

func TestWhereFieldEquals(t *testing.T) {
	f := filter.NewFilter(nil, nil, filter.WhereFieldEquals("TEST_FIELD_NAME", "TEST_FIELD_VALUE"))
	req, err := f.BuildRequest("GET", "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}
	actualUrl := req.URL.String()
	expectedUrl := "http://localhost?where=TEST_FIELD_NAME%3D%3D%22TEST_FIELD_VALUE%22"
	if actualUrl != expectedUrl {
		t.Fatal("Req URL Did Not Match Expected:\n\texpected:\t" + expectedUrl + "\n\tgot:\t\t" + actualUrl)
	}
}

func TestWhereFieldContains(t *testing.T) {
	vals := []string{
		"TEST_FIELD_VALUE_1",
		"TEST_FIELD_VALUE_2",
		"TEST_FIELD_VALUE_3",
	}
	f := filter.NewFilter(nil, nil, filter.WhereFieldContains("TEST_FIELD_NAME", vals))
	req, err := f.BuildRequest("GET", "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}
	actualUrl := req.URL.String()
	expectedUrl := "http://localhost?TEST_FIELD_NAME=" + strings.Join(vals, ",")
	if actualUrl != expectedUrl {
		t.Fatal("Req URL Did Not Match Expected:\n\texpected:\t" + expectedUrl + "\n\tgot:\t\t" + actualUrl)
	}
}

func TestOrderByAcscending(t *testing.T) {
	f := filter.NewFilter(nil, filter.OrderBy("TEST_FIELD_NAME", true), nil)
	req, err := f.BuildRequest("GET", "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}
	actualUrl := req.URL.String()
	expectedUrl := "http://localhost?order=TEST_FIELD_NAME"
	if actualUrl != expectedUrl {
		t.Fatal("Req URL Did Not Match Expected:\n\texpected:\t" + expectedUrl + "\n\tgot:\t\t" + actualUrl)
	}
}
func TestOrderByDescending(t *testing.T) {
	f := filter.NewFilter(nil, filter.OrderBy("TEST_FIELD_NAME", false), nil)
	req, err := f.BuildRequest("GET", "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}
	actualUrl := req.URL.String()
	expectedUrl := "http://localhost?order=TEST_FIELD_NAME%20DESC"
	if actualUrl != expectedUrl {
		t.Fatal("Req URL Did Not Match Expected:\n\texpected:\t" + expectedUrl + "\n\tgot:\t\t" + actualUrl)
	}
}

func TestModifiedAfter(t *testing.T) {
	testTime := &time.Time{}
	f := filter.NewFilter(testTime, nil, nil)
	req, err := f.BuildRequest("GET", "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}
	actualHeader := req.Header.Get("If-Modified-Since")
	expectedHeader := testTime.UTC().String() // This Probably needs a custom format to match xero
	if actualHeader != expectedHeader {
		t.Fatal("Req Header Value Did Not Match Expected:\n\texpected:\t" + expectedHeader + "\n\tgot:\t\t" + actualHeader)
	}
}

func TestFilterMulti(t *testing.T) {
	testTime := &time.Time{}
	orderBy := filter.OrderBy("TEST_FIELD_NAME_1", false)
	vals := []string{
		"TEST_FIELD_VALUE_1",
		"TEST_FIELD_VALUE_2",
		"TEST_FIELD_VALUE_3",
	}
	opt1 := filter.WhereFieldContains("TEST_FIELD_NAME_2", vals)
	opt2 := filter.WhereFieldEquals("TEST_FIELD_NAME_3", "TEST_FIELD_VALUE_4")
	opt3 := filter.WhereFieldEquals("TEST_FIELD_NAME_4", "TEST_FIELD_VALUE_5")

	f := filter.NewFilter(testTime, orderBy, opt1, opt2, opt3)
	req, err := f.BuildRequest("GET", "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}
	actualUrl := req.URL.String()
	expectedUrl := "http://localhost?order=TEST_FIELD_NAME_1%20DESC&TEST_FIELD_NAME_2=" +
		strings.Join(vals, ",") +
		"&where=TEST_FIELD_NAME_3%3D%3D%22TEST_FIELD_VALUE_4%22" +
		"%26%26TEST_FIELD_NAME_4%3D%3D%22TEST_FIELD_VALUE_5%22"
	actualHeader := req.Header.Get("If-Modified-Since")
	expectedHeader := testTime.UTC().String() // This Probably needs a custom format to match xero
	if actualUrl != expectedUrl {
		t.Fatal("Req URL Did Not Match Expected:\n\texpected:\t" + expectedUrl + "\n\tgot:\t\t" + actualUrl)
	}
	if actualHeader != expectedHeader {
		t.Fatal("Req Header Value Did Not Match Expected:\n\texpected:\t" + expectedHeader + "\n\tgot:\t\t" + actualHeader)
	}
}
