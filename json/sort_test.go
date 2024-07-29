package json

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var testMap = map[int]string{
	2: "2",
	5: "5",
	3: "3",
	1: "1",
	4: "4",
}
var testSortedMap = map[int]string{
	1: "1",
	2: "2",
	3: "3",
	4: "4",
	5: "5",
}

func TestSortJsonWithValidData(t *testing.T) {
	assert := assert.New(t)

	jsonStr, err := json.Marshal(testMap)
	assert.True(err == nil, "Marshaling cannot be performed for test data.")

	result, err := SortJson(jsonStr)
	assert.True(err == nil, "An error has occurred in tested method.")

	var resultMap map[int]string
	err = json.Unmarshal(result, &resultMap)
	assert.True(err == nil && reflect.DeepEqual(testSortedMap, resultMap),
		"The result map does not match the expected map.")
}
