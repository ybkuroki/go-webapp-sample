package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRequestURL_PathParam(t *testing.T) {
	result := NewRequestBuilder().URL("https://www.test/").PathParams("1").PathParams("test").Build().GetRequestURL()
	assert.Equal(t, "https://www.test/1/test", result)
}

func TestGetRequestURL_RequestParam(t *testing.T) {
	result := NewRequestBuilder().URL("https://www.test/").RequestParams("hoge", "123").RequestParams("huga", "abcd").Build().GetRequestURL()
	assert.Equal(t, "https://www.test/?hoge=123&huga=abcd", result)
}
