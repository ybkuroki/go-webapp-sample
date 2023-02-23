package util

import (
	"bufio"
	"embed"
	"strings"
)

const (
	// CommentChar represents that the line is the comment line.
	CommentChar = "#"
	// EqualsChar represents that the equals symbol.
	EqualsChar = "="
)

// ReadPropertiesFile reads a properties file and it returns a map has the keys and values in the file.
func ReadPropertiesFile(fs embed.FS, fileName string) map[string]string {
	config := make(map[string]string)

	file, err := fs.Open(fileName)
	if err != nil {
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !isCommentLine(line) && hasProperty(line) {
			setPorperty(line, config)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil
	}

	return config
}

// isCommentLine judge whether a given line is a comment line or not.
func isCommentLine(line string) bool {
	return strings.HasPrefix(line, CommentChar)
}

// hasProperty judge whether a given line has a property: a key and value or not.
func hasProperty(line string) bool {
	return strings.Contains(line, EqualsChar)
}

// setProperty sets the key and value in a properties file to the given map.
func setPorperty(line string, config map[string]string) {
	equal := strings.Index(line, EqualsChar)
	if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
		value := ""
		if len(line) > equal {
			value = strings.TrimSpace(line[equal+1:])
		}
		config[key] = value
	}
}
