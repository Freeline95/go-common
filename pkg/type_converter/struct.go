package type_converter

import (
	"fmt"
	"reflect"
)

func structToMap(obj interface{}) map[string]string {
	result := make(map[string]string)

	typ := reflect.TypeOf(obj)
	val := reflect.ValueOf(obj)
	for i := 0; i < typ.NumField(); i++ {
		fieldName := typ.Field(i).Name
		fieldValue := fmt.Sprintf("%v", val.Field(i).Interface())
		result[fieldName] = fieldValue
	}

	return result
}
