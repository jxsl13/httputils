package encoding

import (
	"encoding/xml"
	"io"
)

func NewEncoderFuncXML() EncoderFunc {
	return func(w io.Writer) Encoder {
		return xml.NewEncoder(w)
	}
}

func NewDecoderFuncXML() DecoderFunc {
	return func(r io.Reader) Decoder {
		return xml.NewDecoder(r)
	}
}
