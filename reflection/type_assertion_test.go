package reflection

import (
	"reflect"
	"testing"
)

// Define some test structs
type Parent struct{}
type Child struct {
	Parent
}
type Unrelated struct{}

func TestFieldTypeIsSubtypeOf(t *testing.T) {
	tests := []struct {
		fieldType  reflect.Type
		parentType reflect.Type
		expected   bool
	}{
		{reflect.TypeOf(Child{}), reflect.TypeOf(Parent{}), true},
		{reflect.TypeOf(Unrelated{}), reflect.TypeOf(Parent{}), false},
		{reflect.TypeOf(Parent{}), reflect.TypeOf(Parent{}), false},
		{reflect.TypeOf(Child{}), reflect.TypeOf(Unrelated{}), false},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := fieldTypeIsSubtypeOf(test.fieldType, test.parentType)
			if result != test.expected {
				t.Errorf("fieldTypeIsSubtypeOf(%v, %v) = %v; want %v", test.fieldType, test.parentType, result, test.expected)
			}
		})
	}
}

func TestIsTypeOf(t *testing.T) {
	tests := []struct {
		child    interface{}
		parent   interface{}
		expected bool
	}{
		{Child{}, Parent{}, true},
		{Unrelated{}, Parent{}, false},
		{Parent{}, Parent{}, false},
		{Child{}, Unrelated{}, false},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := IsTypeOf(test.child, test.parent)
			if result != test.expected {
				t.Errorf("IsTypeOf(%v, %v) = %v; want %v", test.child, test.parent, result, test.expected)
			}
		})
	}
}
