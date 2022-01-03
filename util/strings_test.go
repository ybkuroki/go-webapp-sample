package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNumeric_True(t *testing.T) {
	result := IsNumeric("123")
	assert.True(t, result)
}

func TestIsNumeric_False(t *testing.T) {
	result := IsNumeric("abcde")
	assert.False(t, result)
}

func TestConvertToInt_Number(t *testing.T) {
	result := ConvertToInt("123")
	assert.Exactly(t, int(123), result)
}

func TestConvertToInt_NotNumber(t *testing.T) {
	result := ConvertToInt("abcde/:12@s#$%'()")
	assert.Equal(t, 0, result)
}

func TestConvertToUint_Number(t *testing.T) {
	result := ConvertToUint("123")
	assert.Exactly(t, uint(123), result)
}
