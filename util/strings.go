package util

import "strconv"

// IsNumeric judges whether given string is numeric or not.
func IsNumeric(number string) bool {
	_, err := strconv.Atoi(number)
	return err == nil
}

// ConvertToInt converts given string to int.
func ConvertToInt(number string) int {
	value, err := strconv.Atoi(number)
	if err != nil {
		return 0
	}
	return value
}

// ConvertToUint converts given string to uint.
func ConvertToUint(number string) uint {
	return uint(ConvertToInt(number))
}
