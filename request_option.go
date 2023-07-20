package httputils

import "net/http"

type RequestOption func(*http.Request)

// WithAddHeader adds a header without overwriting an existing header
func WithContentType(contentType string) RequestOption {
	return func(r *http.Request) {
		r.Header.Set("Content-Type", contentType)
	}
}

// WithAddHeader adds a header without overwriting an existing header
func WithAddHeader(key, value string) RequestOption {
	return func(r *http.Request) {
		r.Header.Add(key, value)
	}
}

// WithAddHeader sets a header overwriting an existing header
func WithSetHeader(key, value string) RequestOption {
	return func(r *http.Request) {
		r.Header.Set(key, value)
	}
}

// WithHeaders overwrites existing headers with the key value pars from the map
func WithHeaders(headers map[string]string) RequestOption {
	return func(r *http.Request) {
		for k, v := range headers {
			r.Header.Set(k, v)
		}
	}
}

// WithMultiHeaders allows to set multi value headers
func WithMultiHeaders(headers http.Header) RequestOption {
	return func(r *http.Request) {
		r.Header = headers
	}
}
