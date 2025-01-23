package utils_test

import (
	"encoding/json"
	"testing"

	"github.com/datag8r/xerogo/utils"
)

func TestCustomMarshall(t *testing.T) {
	type testStruct struct {
		Name string `xero:"update"`
		Age  int    `xero:"update"`
	}
	test := testStruct{
		Name: "test",
		Age:  1,
	}
	result, err := utils.XeroCustomMarshal(test, "update")
	if err != nil {
		t.Error(err)
	}
	if result == nil {
		t.Error("Expected result to not be nil")
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		t.Error(err)
	}
	expected := `{"Age":1,"Name":"test"}`
	if string(jsonResult) != expected {
		t.Errorf("Expected %s, got %s", expected, string(jsonResult))
	}
}

func TestCustomMarshallWithEmbeddedID(t *testing.T) {
	type testEmbed struct {
		ID string `xero:"id"`
	}
	type testStruct struct {
		Name  string    `xero:"update"`
		Age   int       `xero:"update"`
		Embed testEmbed `xero:"update,embeddedId"`
	}
	test := testStruct{
		Name: "test",
		Age:  1,
		Embed: testEmbed{
			ID: "test",
		},
	}
	result, err := utils.XeroCustomMarshal(test, "update")
	if err != nil {
		t.Error(err)
	}
	if result == nil {
		t.Error("Expected result to not be nil")
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		t.Error(err)
	}
	expected := `{"Age":1,"Embed":{"ID":"test"},"Name":"test"}`
	if string(jsonResult) != expected {
		t.Errorf("Expected %s, got %s", expected, string(jsonResult))
	}
}

func TestCustomMarshallWithOptional(t *testing.T) {
	type testStruct struct {
		Name string `xero:"update"`
		Age  int    `xero:"*update"`
	}
	test := testStruct{
		Name: "test",
	}
	result, err := utils.XeroCustomMarshal(test, "update")
	if err != nil {
		t.Error(err)
	}
	if result == nil {
		t.Error("Expected result to not be nil")
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		t.Error(err)
	}
	expected := `{"Name":"test"}`
	if string(jsonResult) != expected {
		t.Errorf("Expected %s, got %s", expected, string(jsonResult))
	}
}
