package errors

import (
	"errors"
	"fmt"
)

func JoinErrors(errorList []error) error {
	var result string
	for ix, e := range errorList {
		result += fmt.Sprintf("#%d: %s\n", ix+1, e.Error())
	}
	return errors.New(result)
}
