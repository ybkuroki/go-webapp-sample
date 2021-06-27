package test

import "strings"

// RequestBuilder builds request URL.
type RequestBuilder struct {
	url           string
	pathParams    []string
	requestParams map[string]string
}

// NewRequestBuilder is constructor.
func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{url: "", pathParams: nil, requestParams: make(map[string]string)}
}

// URL sets request base url.
func (b *RequestBuilder) URL(url string) *RequestBuilder {
	b.url = url
	return b
}

// PathParams set a path parameter of the url.
func (b *RequestBuilder) PathParams(value string) *RequestBuilder {
	b.pathParams = append(b.pathParams, value)
	return b
}

// RequestParams set a request parameter of the reqest.
func (b *RequestBuilder) RequestParams(name string, value string) *RequestBuilder {
	b.requestParams[name] = value
	return b
}

// Build builds request url.
func (b *RequestBuilder) Build() *RequestURL {
	return &RequestURL{url: b.url, pathParams: b.pathParams, requestParams: b.requestParams}
}

// RequestURL is the element to compose request url.
type RequestURL struct {
	url           string
	pathParams    []string
	requestParams map[string]string
}

// GetRequestURL returns request url builded by request builder.
func (r *RequestURL) GetRequestURL() string {
	return r.url + r.getPathParams() + "?" + r.getRequestParams()
}

func (r *RequestURL) getPathParams() string {
	result := ""
	for i, value := range r.pathParams {
		if i == 0 && strings.HasSuffix(r.url, "/") {
			result += value
		} else {
			result += "/" + value
		}
	}
	return result
}

func (r *RequestURL) getRequestParams() string {
	count := 0
	result := ""
	for key, value := range r.requestParams {
		result += key + "=" + value
		if count != len(r.requestParams)-1 {
			result += "&"
			count++
		}
	}
	return result
}
