package utils

import (
	"reflect"
	"slices"
	"strings"
)

func XeroCustomMarshal(value any, requestType string) (interface{}, error) {
	mp := map[string]interface{}{}
	var lst []map[string]interface{}
	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)
	if t.Kind() == reflect.Array || t.Kind() == reflect.Slice { // check if value is an array
		lst = []map[string]interface{}{}
		if v.Len() == 0 { // check if array is empty
			return nil, nil
		}
		for i := range v.Len() { // for each element in slice
			item := v.Index(i)                                  // grab element
			for i, f := range reflect.VisibleFields(t.Elem()) { // go through all struct fields
				if inc, opt, embeddedId := include(f, requestType); inc { // go through all field tags, to check if value is included and optional
					val := item.Field(i)
					if val.Kind() == reflect.Struct { // check if field is a struct
						recursiveRequestType := requestType
						if embeddedId { // check if field is an embeddedId
							recursiveRequestType = "id"
						}
						inter, _ := XeroCustomMarshal(val.Interface(), recursiveRequestType) // recursively call customMarshal
						mp[f.Name] = inter
						continue
					}
					if opt && isZero(val) { // check if its the default value and ignore if it is and the field is optional
						continue
					} else {
						mp[f.Name] = val.Interface() // add interface of value to map
					}
				}
			}
			if len(mp) > 0 {
				lst = append(lst, mp) // add map to list
				mp = map[string]interface{}{}
			}
		}
	} else { // if value is not an array
		for i, f := range reflect.VisibleFields(t) { // go through all struct fields
			if inc, opt, embeddedId := include(f, requestType); inc { // go through all field tags, to check if value is included and optional
				val := v.Field(i)
				if val.Kind() == reflect.Struct { // check if field is a struct
					recursiveRequestType := requestType
					if embeddedId { // check if field is an embeddedId
						recursiveRequestType = "id"
					}
					inter, _ := XeroCustomMarshal(val.Interface(), recursiveRequestType) // recursively call customMarshal
					mp[f.Name] = inter
					continue
				}
				if opt && isZero(val) { // check if its the default value and ignore if it is and the field is optional
					continue
				} else {
					mp[f.Name] = val.Interface() // add interface of value to map
				}
			}
		}
	}
	if lst != nil {
		return lst, nil
	} else {
		return mp, nil
	}
}

func include(f reflect.StructField, requestType string) (inc, optional, embeddedId bool) {
	xeroTags := strings.Split(f.Tag.Get("xero"), ",")                                          // split tags on comma
	embeddedId = slices.Contains(xeroTags, "embeddedId")                                       // check if embeddedId is in xeroTags
	inc = slices.Contains(xeroTags, requestType) || slices.Contains(xeroTags, "*"+requestType) // check if requestType// optional + requestType is in xeroTags
	optional = slices.Contains(xeroTags, "*"+requestType)                                      // check if optional + requestType is in xeroTags
	return
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Func, reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.Array:
		z := true
		for i := 0; i < v.Len(); i++ {
			z = z && isZero(v.Index(i))
		}
		return z
	case reflect.Struct:
		z := true
		for i := 0; i < v.NumField(); i++ {
			z = z && isZero(v.Field(i))
		}
		return z
	}
	// Compare other types directly:
	z := reflect.Zero(v.Type())
	return v.Interface() == z.Interface()
}
