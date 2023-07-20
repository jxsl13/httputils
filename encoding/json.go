package encoding

import (
	"encoding/json"
	"io"
)

func NewEncoderFuncJSON() EncoderFunc {
	return func(w io.Writer) Encoder {
		return json.NewEncoder(w)
	}
}

func NewDecoderFuncJSON() DecoderFunc {
	return func(r io.Reader) Decoder {
		return json.NewDecoder(r)
	}
}
