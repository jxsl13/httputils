package pickle

import "github.com/jxsl13/httputils/encoding"

func NewBinary() Pickle {
	return newPickle(
		ContentTypeBinary,
		encoding.NewEncoderFuncBinary(),
		encoding.NewDecoderFuncBinary(),
	)
}
