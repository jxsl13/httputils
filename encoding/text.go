package encoding

import (
	"encoding"
	"fmt"
	"io"
)

func NewEncoderFuncText() EncoderFunc {
	return func(w io.Writer) Encoder {
		return NewTextEncoder(w)
	}
}

func NewDecoderFuncText() DecoderFunc {
	return func(r io.Reader) Decoder {
		return NewTextDecoder(r)
	}
}

func NewTextEncoder(w io.Writer) *TextEncoder {
	return &TextEncoder{
		w: w,
	}
}

type TextEncoder struct {
	w io.Writer
}

func (enc *TextEncoder) Encode(v any) error {
	var (
		data []byte
		err  error
	)
	switch x := v.(type) {
	case encoding.TextMarshaler:
		data, err = x.MarshalText()
	case io.Reader:
		data, err = io.ReadAll(x)
	case string:
		data = []byte(x)
	case []byte:
		data = x
	default:
		return fmt.Errorf("unsupported Text encoding type: %#T", v)
	}
	if err != nil {
		return err
	}

	return writeAll(enc.w, data)
}

func NewTextDecoder(r io.Reader) *TextDecoder {
	return &TextDecoder{
		r: r,
	}
}

type TextDecoder struct {
	r io.Reader
}

func (dec *TextDecoder) Decode(v any) error {
	if v == nil {
		return fmt.Errorf("decode out value is nil")
	}
	data, err := io.ReadAll(dec.r)
	if err != nil {
		return err
	}

	switch x := v.(type) {
	case encoding.TextUnmarshaler:
		return x.UnmarshalText(data)
	case *string:
		*x = string(data)
	case *[]byte:
		*x = data
	case io.Writer:
		return writeAll(x, data)
	default:
		return fmt.Errorf("unsupported Text decoding type: %#T", v)
	}
	return nil
}
