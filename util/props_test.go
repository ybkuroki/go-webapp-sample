package util

import (
	"embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed test.properties
var testPropsFile embed.FS

func TestReadPropertiesFile_FileExists(t *testing.T) {
	messages := ReadPropertiesFile(testPropsFile, "test.properties")

	assert.Equal(t, "testtest", messages["props1"])
	assert.Equal(t, "test test", messages["props2"])
	assert.Equal(t, "", messages["props3"])

	_, ok := messages["props4"]
	assert.Equal(t, false, ok)

	assert.Equal(t, "test #test", messages["props5"])

	_, ok = messages["# comment line"]
	assert.Equal(t, false, ok)
}

func TestReadPropertiesFile_FileIsNil(t *testing.T) {
	messages := ReadPropertiesFile(testPropsFile, "testnil.properties")
	assert.Equal(t, map[string]string(nil), messages)
}
