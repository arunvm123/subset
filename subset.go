package subset

import (
	"reflect"
)

// Subset checks if structA is a subset of structB.
// structA is said to be a subset of structB when all
// the fields of structA are present in structB.
func Subset(structA interface{}, structB interface{}) bool {
	structAFields := getFieldsMap(structA)
	structBFields := getFieldsMap(structB)

	for key := range structBFields {
		delete(structAFields, key)
	}

	if len(structAFields) != 0 {
		return false
	}

	return true
}

func getFieldsMap(s interface{}) map[string]struct{} {
	res := make(map[string]struct{})
	getFieldsMapHelper(res, "", s)
	return res
}

func getFieldsMapHelper(res map[string]struct{}, parentName string, s interface{}) {
	typ := reflect.TypeOf(s)
	n := typ.NumField()

	for i := 0; i < n; i++ {
		fieldName := typ.Field(i).Name
		fieldType := typ.Field(i).Type
		if fieldType.Kind().String() == "struct" {
			getFieldsMapHelper(res, parentName+"."+fieldName, reflect.ValueOf(s).Field(i).Interface())
			continue
		}

		res[parentName+"."+fieldName] = struct{}{}
	}
}
