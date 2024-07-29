package slice_and_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStringArray = []string{"a", "b", "c", "d", "e"}
var testIntArray = []int{1, 2, 3, 4, 5}

type first struct{}
type second struct{ first }
type third struct{ second }
type fourth struct{ first int }

func TestContainsWithStringElement(t *testing.T) {
	assert := assert.New(t)

	assert.True(contains(testStringArray, "b"), "The item should be in the list.")
	assert.False(contains(testStringArray, "f"), "The element should be not in the list.")
}

func TestContainsWithIntElement(t *testing.T) {
	assert := assert.New(t)

	assert.True(contains(testIntArray, 3), "The item should be in the list.")
	assert.False(contains(testIntArray, 6), "The element should not in the list.")
}
