package pickle

import "github.com/jxsl13/httputils/encoding"

func NewText() Pickle {
	return newPickle(
		ContentTypeText,
		encoding.NewEncoderFuncText(),
		encoding.NewDecoderFuncText(),
	)
}
