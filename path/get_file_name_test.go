package path

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileNameWithoutExt(t *testing.T) {
	myAssert := assert.New(t)

	myAssert.True(GetFileNameFromPathWithoutExt("a/b/c/d.ext") == "d")
	myAssert.True(GetFileNameFromPathWithoutExt("a/b/c/d") == "d")
	myAssert.True(GetFileNameFromPathWithoutExt("/a/b/c/d.ext") == "d")
	myAssert.True(GetFileNameFromPathWithoutExt("d.ext") == "d")
	myAssert.True(GetFileNameFromPathWithoutExt("d") == "d")
	myAssert.True(GetFileNameFromPathWithoutExt("/d.ext") == "d")
	myAssert.True(GetFileNameFromPathWithoutExt("/d") == "d")
	myAssert.True(GetFileNameFromPathWithoutExt("/b/c/d.ext1.ext2") == "d.ext1")
}
