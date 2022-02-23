package util

import "strings"

// RequestBuilder builds request URL.
type RequestBuilder interface {
	URL(url string) RequestBuilder
	PathParams(value string) RequestBuilder
	RequestParams(name string, value string) RequestBuilder
	Build() RequestURL
}

type requestBuilder struct {
	url           string
	pathParams    []string
	requestParams map[string]string
}

// NewRequestBuilder is constructor.
func NewRequestBuilder() RequestBuilder {
	return &requestBuilder{url: "", pathParams: nil, requestParams: make(map[string]string)}
}

// URL sets request base url.
func (b *requestBuilder) URL(url string) RequestBuilder {
	b.url = url
	return b
}

// PathParams set a path parameter of the url.
func (b *requestBuilder) PathParams(value string) RequestBuilder {
	b.pathParams = append(b.pathParams, value)
	return b
}

// RequestParams set a request parameter of the reqest.
func (b *requestBuilder) RequestParams(name string, value string) RequestBuilder {
	b.requestParams[name] = value
	return b
}

// Build builds request url.
func (b *requestBuilder) Build() RequestURL {
	return &requestURL{url: b.url, pathParams: b.pathParams, requestParams: b.requestParams}
}

// RequestURL is the element to compose request url.
type RequestURL interface {
	GetRequestURL() string
}

type requestURL struct {
	url           string
	pathParams    []string
	requestParams map[string]string
}

// GetRequestURL returns request url builded by request builder.
func (r *requestURL) GetRequestURL() string {
	result := r.url + r.getPathParams()
	if r.getRequestParams() != "" {
		result = result + "?" + r.getRequestParams()
	}
	return result
}

func (r *requestURL) getPathParams() string {
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

func (r *requestURL) getRequestParams() string {
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
