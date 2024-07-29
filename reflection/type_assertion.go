package reflection

import "reflect"

func fieldTypeIsSubtypeOf(fieldType reflect.Type, parentType reflect.Type) bool {
	if fieldType.Kind() != reflect.Struct {
		return false
	}
	for i := 0; i < fieldType.NumField(); i++ {
		field := fieldType.Field(i)
		if field.Anonymous && (field.Type == parentType || fieldTypeIsSubtypeOf(field.Type, parentType)) {
			return true
		}
	}
	return false
}

func IsTypeOf(child interface{}, parent interface{}) bool {
	childType := reflect.TypeOf(child)
	parentType := reflect.TypeOf(parent)
	return fieldTypeIsSubtypeOf(childType, parentType)
}
