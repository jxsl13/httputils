package pickle

import (
	"io"

	"github.com/jxsl13/httputils/encoding"
)

type Jar = map[string]Pickle

func NewJar(pickle ...Pickle) Jar {
	m := make(map[string]Pickle, len(pickle))
	for _, v := range pickle {
		m[v.ContentType()] = v
	}
	return m
}

// Pickle is borrowed from Python and is a serializer and decerializer but for a given content-type.
type Pickle interface {
	Encoder(w io.Writer) encoding.Encoder
	Decoder(r io.Reader) encoding.Decoder
	ContentType() string
}

func newPickle(contentType string, enc encoding.EncoderFunc, dec encoding.DecoderFunc) *pickle {
	return &pickle{
		contentType: contentType,
		enc:         enc,
		dec:         dec,
	}
}

type pickle struct {
	contentType string
	enc         encoding.EncoderFunc
	dec         encoding.DecoderFunc
}

func (p *pickle) Encoder(w io.Writer) encoding.Encoder {
	return p.enc(w)
}

func (p *pickle) Decoder(r io.Reader) encoding.Decoder {
	return p.dec(r)
}

func (p *pickle) ContentType() string {
	return p.contentType
}
