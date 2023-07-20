package pickle

import "github.com/jxsl13/httputils/encoding"

func NewJSON() Pickle {
	return newPickle(
		ContentTypeJSON,
		encoding.NewEncoderFuncJSON(),
		encoding.NewDecoderFuncJSON(),
	)
}
