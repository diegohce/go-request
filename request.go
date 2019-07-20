package request

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
)

// RequestBuilder struct for request builder pattern.
type RequestBuilder struct {
	method   string
	url      url.URL
	headers  http.Header
	queryStr *url.Values
	payload  io.Reader
}

// Scheme sets the url scheme on RequestBuilder. Defaults to 'http'.
func (rb *RequestBuilder) Scheme(s string) *RequestBuilder {
	rb.url.Scheme = s
	return rb
}

// Host sets the url host[:port] on RequestBuilder.
func (rb *RequestBuilder) Host(host string) *RequestBuilder {
	rb.url.Host = host
	return rb
}

// Method sets the request method on RequestBuilder.
// Defaults to GET if no payload was set, otherwise defaults to POST.
func (rb *RequestBuilder) Method(method string) *RequestBuilder {
	rb.method = method
	return rb
}

// URL sets the url part on RequestBuilder. It is supposed to start with a slash (/).
func (rb *RequestBuilder) URL(url string) *RequestBuilder {
	rb.url.Path = url
	return rb
}

// UserPassword set the user authentication part of the url on RequestBuilder.
func (rb *RequestBuilder) UserPassword(user, password string) *RequestBuilder {
	rb.url.User = url.UserPassword(user, password)
	return rb
}

// Payload sets the request body.
func (rb *RequestBuilder) Payload(payload []byte) *RequestBuilder {
	rb.payload = bytes.NewReader(payload)
	return rb
}

// AddValue adds a value to an existing key.
// It will duplicate the key on the resulting querystring.
func (rb *RequestBuilder) AddValue(key, value string) *RequestBuilder {
	if rb.queryStr == nil {
		rb.queryStr = &url.Values{}
	}
	rb.queryStr.Add(key, value)
	return rb
}

// SetValue sets value to a new key. If key exists, it is replaced.
func (rb *RequestBuilder) SetValue(key, value string) *RequestBuilder {
	if rb.queryStr == nil {
		rb.queryStr = &url.Values{}
	}
	rb.queryStr.Set(key, value)
	return rb
}

// DelValue removes key and all its values from querystring.
func (rb *RequestBuilder) DelValue(key string) *RequestBuilder {
	if rb.queryStr != nil {
		rb.queryStr.Del(key)
	}
	return rb
}

// Values returns the encoded querystring that will be used to form the request.
func (rb *RequestBuilder) Values() string {
	return rb.queryStr.Encode()
}

// Build creates a *http.Request object using the collected settings from RequestBuilder.
func (rb *RequestBuilder) Build() (*http.Request, error) {
	if rb.url.Scheme == "" {
		rb.url.Scheme = "http"
	}
	if rb.method == "" {
		rb.method = "GET"
	}
	rb.url.RawQuery = rb.queryStr.Encode()
	return http.NewRequest(rb.method, rb.url.String(), rb.payload)
}
