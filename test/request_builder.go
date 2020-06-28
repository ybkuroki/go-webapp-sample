package test

// RequestBuilder builds request URL.
type RequestBuilder struct {
	url    string
	params map[string]string
}

// NewRequestBuilder is constructor.
func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{url: "", params: make(map[string]string)}
}

// URL sets request base url.
func (b *RequestBuilder) URL(url string) *RequestBuilder {
	b.url = url
	return b
}

// Params set a request parameter of the reqest.
func (b *RequestBuilder) Params(name string, value string) *RequestBuilder {
	b.params[name] = value
	return b
}

// Build builds request url.
func (b *RequestBuilder) Build() *RequestURL {
	return &RequestURL{url: b.url, params: b.params}
}

// RequestURL is the element to compose request url.
type RequestURL struct {
	url    string
	params map[string]string
}

// GetRequestURL returns request url builded by request builder.
func (r *RequestURL) GetRequestURL() string {
	return r.url + "?" + r.getParams()
}

func (r *RequestURL) getParams() string {
	count := 0
	result := ""
	for key, value := range r.params {
		result += key + "=" + value
		if count != len(r.params)-1 {
			result += "&"
			count++
		}
	}
	return result
}
