package path

import (
	"path"
	"path/filepath"
)

func GetFileNameFromPathWithoutExt(filePath string) string {
	base := path.Base(filePath)
	return base[:len(base)-len(filepath.Ext(filePath))]
}
