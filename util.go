package fmp_client

import (
	"fmt"
	"reflect"
	"time"
)

// StructToMap converts a struct to a map[string]string using field names as keys.
// If a field is a pointer and nil, the property is omitted from the map.
func StructToMap(params interface{}) map[string]string {
	result := make(map[string]string)
	v := reflect.ValueOf(params)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		if field.PkgPath != "" { // skip unexported fields
			continue
		}
		key := field.Name
		value := v.Field(i)
		if value.Kind() == reflect.Ptr {
			if value.IsNil() {
				continue // omit nil pointer fields
			}
			value = value.Elem()
		}
		
		// Handle time.Time specially
		if value.Type() == reflect.TypeOf(time.Time{}) {
			timeVal := value.Interface().(time.Time)
			result[key] = timeVal.Format("2006-01-02")
		} else {
			result[key] = fmt.Sprintf("%v", value.Interface())
		}
	}
	return result
}
