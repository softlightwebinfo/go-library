package helpers

import (
	"path/filepath"
	"strconv"
)

//GenerateUniqueId Generate uniqueid random timestamp
func GenerateUniqueId() string {
	return strconv.FormatInt(MakeTimestamp(), 10)
}

//GetExtensionName Get extension name from filename
func GetExtensionName(fileName string) string {
	return filepath.Ext(fileName)
}

//GenerateFileName get generate filName
func GenerateFileName(fileName string) string {
	name := GenerateUniqueId()
	return name + GetExtensionName(fileName)
}
