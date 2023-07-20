package httputils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/jxsl13/httputils/pickle"
)

// New creates an extended http tester object that is able to encode and decode
// bodys
func New(handler http.Handler, options ...TesterOption) *HTTPTester {

	t := &HTTPTester{
		handler: handler,
		pickle: pickle.NewJar(
			pickle.NewJSON(),
			pickle.NewXML(),
			pickle.NewText(),
			pickle.NewBinary()),

		fallbackPickle: pickle.NewJSON(),
	}

	for _, opt := range options {
		opt(t)
	}
	return t
}

type HTTPTester struct {
	handler    http.Handler
	reqOptions []RequestOption
	pickle     pickle.Jar

	fallbackPickle pickle.Pickle
}

type TesterOption func(*HTTPTester)

func WithRequestOptions(options ...RequestOption) TesterOption {
	return func(h *HTTPTester) {
		h.reqOptions = append(h.reqOptions, options...)
	}
}

func WithPickle(pickle pickle.Pickle) TesterOption {
	return func(h *HTTPTester) {
		h.pickle[pickle.ContentType()] = pickle

	}
}

func WithPickleJar(pickle map[string]pickle.Pickle) TesterOption {
	return func(h *HTTPTester) {
		for k, v := range pickle {
			h.pickle[k] = v
		}
	}
}

func WithFallbackPickle(pickle pickle.Pickle) TesterOption {
	return func(h *HTTPTester) {
		h.fallbackPickle = pickle
	}
}

func (t *HTTPTester) decode(resp *httptest.ResponseRecorder, result any) error {
	if result == nil {
		return nil
	}

	if resp.Code/100 != 2 {
		return fmt.Errorf("unexpected response code: %d", resp.Code)
	}

	contentType := resp.Header().Get("Content-Type")
	if contentType == "" {
		return t.fallbackPickle.Decoder(resp.Body).Decode(result)
	}
	p, ok := t.pickle[contentType]
	if !ok {
		return fmt.Errorf("no decoder found for content type %s", contentType)
	}
	decoder := p.Decoder(resp.Body)
	return decoder.Decode(result)
}

func (t *HTTPTester) encode(contentType string, body any) (io.Reader, error) {
	if body == nil {
		return &bytes.Buffer{}, nil
	}
	buf := bytes.NewBuffer(nil)

	if contentType == "" {
		err := t.fallbackPickle.Encoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
		return buf, nil
	}

	p, ok := t.pickle[contentType]
	if !ok {
		return nil, fmt.Errorf("no encoder found for content type %s", contentType)
	}
	encoder := p.Encoder(buf)
	err := encoder.Encode(body)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (t *HTTPTester) Get(pathQuery string, result any, requestOptions ...RequestOption) error {
	resp := GET(t.handler, pathQuery, append(t.reqOptions, requestOptions...)...)
	return t.decode(resp, result)
}

func (t *HTTPTester) Post(pathQuery string, body, result any, contentType string, requestOptions ...RequestOption) error {
	r, err := t.encode(contentType, body)
	if err != nil {
		return err
	}
	resp := POST(t.handler, pathQuery, r, append(t.reqOptions, requestOptions...)...)
	return t.decode(resp, result)
}

func (t *HTTPTester) Put(pathQuery string, body, result any, contentType string, requestOptions ...RequestOption) error {
	r, err := t.encode(contentType, body)
	if err != nil {
		return err
	}
	resp := PUT(t.handler, pathQuery, r, append(t.reqOptions, requestOptions...)...)
	return t.decode(resp, result)
}

func (t *HTTPTester) Patch(pathQuery string, body, result any, contentType string, requestOptions ...RequestOption) error {
	r, err := t.encode(contentType, body)
	if err != nil {
		return err
	}
	resp := PATCH(t.handler, pathQuery, r, append(t.reqOptions, requestOptions...)...)
	return t.decode(resp, result)
}

func (t *HTTPTester) Delete(pathQuery string, result any, requestOptions ...RequestOption) error {
	resp := DELETE(t.handler, pathQuery, append(t.reqOptions, requestOptions...)...)
	return t.decode(resp, result)
}

func (t *HTTPTester) Head(pathQuery string, requestOptions ...RequestOption) (http.Header, error) {
	resp := HEAD(t.handler, pathQuery, append(t.reqOptions, requestOptions...)...)
	return resp.Header(), nil
}

func (t *HTTPTester) Connect(pathQuery string, requestOptions ...RequestOption) (http.Header, error) {
	resp := CONNECT(t.handler, pathQuery, append(t.reqOptions, requestOptions...)...)
	return resp.Header(), nil
}

func (t *HTTPTester) Options(pathQuery string, requestOptions ...RequestOption) (http.Header, error) {
	resp := CONNECT(t.handler, pathQuery, append(t.reqOptions, requestOptions...)...)
	return resp.Header(), nil
}

func (t *HTTPTester) Trace(pathQuery string, requestOptions ...RequestOption) (*httptest.ResponseRecorder, error) {
	resp := CONNECT(t.handler, pathQuery, append(t.reqOptions, requestOptions...)...)
	return resp, nil
}
