package errors

import (
	"errors"
	"testing"
)

func TestJoinErrors(t *testing.T) {
	tests := []struct {
		name      string
		errorList []error
		expected  string
	}{
		{
			name:      "No errors",
			errorList: []error{},
			expected:  "",
		},
		{
			name: "One error",
			errorList: []error{
				errors.New("error 1"),
			},
			expected: "#1: error 1\n",
		},
		{
			name: "Multiple errors",
			errorList: []error{
				errors.New("error 1"),
				errors.New("error 2"),
				errors.New("error 3"),
			},
			expected: "#1: error 1\n#2: error 2\n#3: error 3\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := JoinErrors(test.errorList)
			if result == nil && test.expected != "" {
				t.Errorf("JoinErrors() = nil; want %v", test.expected)
			} else if result != nil && result.Error() != test.expected {
				t.Errorf("JoinErrors() = %v; want %v", result.Error(), test.expected)
			}
		})
	}
}
