package _map

import "testing"

func TestConvertInterfaceToMap(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected map[string]int
	}{
		{
			name: "Valid conversion",
			input: map[string]int{
				"one": 1,
				"two": 2,
			},
			expected: map[string]int{
				"one": 1,
				"two": 2,
			},
		},
		{
			name:     "Invalid conversion",
			input:    "not a map",
			expected: nil,
		},
		{
			name:     "Empty map",
			input:    map[string]int{},
			expected: map[string]int{},
		},
		{
			name:     "Nil input",
			input:    nil,
			expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ConvertInterfaceToMap[string, int](test.input)
			if len(result) != len(test.expected) {
				t.Errorf("ConvertInterfaceToMap() = %v; want %v", result, test.expected)
			}
			for k, v := range test.expected {
				if result[k] != v {
					t.Errorf("ConvertInterfaceToMap() = %v; want %v", result, test.expected)
				}
			}
		})
	}
}
