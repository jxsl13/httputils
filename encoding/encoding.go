package encoding

import "io"

type Encoder interface {
	Encode(v any) error
}

type Decoder interface {
	Decode(v any) error
}

type DecoderFunc func(io.Reader) Decoder
type EncoderFunc func(io.Writer) Encoder

func writeAll(w io.Writer, data []byte) error {
	var (
		err     error
		written = 0
		n       = 0
		length  = len(data)
	)

	for written < length {
		n, err = w.Write(data)
		if err != nil {
			return err
		}
		written += n
	}

	return nil
}
