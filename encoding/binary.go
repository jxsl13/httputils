package encoding

import (
	"encoding"
	"fmt"
	"io"
)

func NewEncoderFuncBinary() EncoderFunc {
	return func(w io.Writer) Encoder {
		return NewBinaryEncoder(w)
	}
}

func NewDecoderFuncBinary() DecoderFunc {
	return func(r io.Reader) Decoder {
		return NewBinaryDecoder(r)
	}
}

func NewBinaryEncoder(w io.Writer) *BinaryEncoder {
	return &BinaryEncoder{
		w: w,
	}
}

type BinaryEncoder struct {
	w io.Writer
}

func (enc *BinaryEncoder) Encode(v any) error {
	var (
		data []byte
		err  error
	)
	switch x := v.(type) {
	case encoding.BinaryMarshaler:
		data, err = x.MarshalBinary()
	case io.Reader:
		data, err = io.ReadAll(x)
	case string:
		data = []byte(x)
	case []byte:
		data = x
	default:
		return fmt.Errorf("unsupported binary encoding type: %#T", v)
	}
	if err != nil {
		return err
	}

	return writeAll(enc.w, data)
}

func NewBinaryDecoder(r io.Reader) *BinaryDecoder {
	return &BinaryDecoder{
		r: r,
	}
}

type BinaryDecoder struct {
	r io.Reader
}

func (dec *BinaryDecoder) Decode(v any) error {
	if v == nil {
		return fmt.Errorf("decode out value is nil")
	}
	data, err := io.ReadAll(dec.r)
	if err != nil {
		return err
	}

	switch x := v.(type) {
	case encoding.BinaryUnmarshaler:
		return x.UnmarshalBinary(data)
	case *string:
		*x = string(data)
	case *[]byte:
		*x = data
	case io.Writer:
		return writeAll(x, data)
	default:
		return fmt.Errorf("unsupported binary decoding type: %#T", v)
	}
	return nil
}
