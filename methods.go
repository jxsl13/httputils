package httputils

import (
	"io"
	"net/http"
	"net/http/httptest"
)

// GET simulates a Get request
func GET(handler http.Handler, pathQuery string, requestOptions ...RequestOption) *httptest.ResponseRecorder {
	return Exec(handler, http.MethodGet, pathQuery, nil, requestOptions...)
}

// Post simulates a POST request
func POST(handler http.Handler, pathQuery string, body io.Reader, requestOptions ...RequestOption) *httptest.ResponseRecorder {
	return Exec(handler, http.MethodPost, pathQuery, body, requestOptions...)
}

func PUT(handler http.Handler, pathQuery string, body io.Reader, requestOptions ...RequestOption) *httptest.ResponseRecorder {
	return Exec(handler, http.MethodPut, pathQuery, body, requestOptions...)
}

func DELETE(handler http.Handler, pathQuery string, requestOptions ...RequestOption) *httptest.ResponseRecorder {
	return Exec(handler, http.MethodPut, pathQuery, nil, requestOptions...)
}

func HEAD(handler http.Handler, pathQuery string, requestOptions ...RequestOption) *httptest.ResponseRecorder {
	return Exec(handler, http.MethodHead, pathQuery, nil, requestOptions...)
}

func PATCH(handler http.Handler, pathQuery string, body io.Reader, requestOptions ...RequestOption) *httptest.ResponseRecorder {
	return Exec(handler, http.MethodPatch, pathQuery, body, requestOptions...)
}

func CONNECT(handler http.Handler, pathQuery string, requestOptions ...RequestOption) *httptest.ResponseRecorder {
	return Exec(handler, http.MethodConnect, pathQuery, nil, requestOptions...)
}

func OPTIONS(handler http.Handler, pathQuery string, requestOptions ...RequestOption) *httptest.ResponseRecorder {
	return Exec(handler, http.MethodOptions, pathQuery, nil, requestOptions...)
}

func TRACE(handler http.Handler, pathQuery string, requestOptions ...RequestOption) *httptest.ResponseRecorder {
	return Exec(handler, http.MethodTrace, pathQuery, nil, requestOptions...)
}

func Exec(handler http.Handler, method, pathQuery string, body io.Reader, requestOptions ...RequestOption) *httptest.ResponseRecorder {
	resp := httptest.NewRecorder()
	req := httptest.NewRequest(method, pathQuery, body)

	for _, opt := range requestOptions {
		opt(req)
	}

	handler.ServeHTTP(resp, req)
	return resp
}
