package util

import (
	"bufio"
	"embed"
	"strings"
)

const (
	CommentChar = "#"
	EqualsChar  = "="
)

func ReadPropertiesFile(fs embed.FS) map[string]string {
	config := make(map[string]string)

	file, err := fs.Open("messages.properties")
	if err != nil {
		return nil
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !isCommentLine(line) && hasProperty(line) {
			if key, value := getProperty(line); key != "" {
				config[key] = value
			}
		}
	}
	return config
}

func isCommentLine(line string) bool {
	return strings.HasPrefix(line, CommentChar)
}

func hasProperty(line string) bool {
	return strings.Contains(line, EqualsChar)
}

func getProperty(line string) (string, string) {
	equal := strings.Index(line, EqualsChar)
	if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
		value := ""
		if len(line) > equal {
			value = strings.TrimSpace(line[equal+1:])
		}
		return key, value
	}
	return "", ""
}
