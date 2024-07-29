package _map

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var testIntArray = []int{1, 2, 3, 4, 5}
var testMap = map[int]string{
	2: "2",
	5: "5",
	3: "3",
	1: "1",
	4: "4",
}

func TestSortedKeys(t *testing.T) {
	assert := assert.New(t)
	assert.True(reflect.DeepEqual(testIntArray, GetSortedKeys(testMap)),
		"The result list does not match the expected list.")
}
