package pickle

import "github.com/jxsl13/httputils/encoding"

func NewXML() Pickle {
	return newPickle(
		ContentTypeXML,
		encoding.NewEncoderFuncXML(),
		encoding.NewDecoderFuncXML(),
	)
}
