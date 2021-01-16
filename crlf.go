package crlf

import (
	"io"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"

	"golang.org/x/text/transform"
)

// Transformer implements a io.Writer object
type Crlf struct {
}

func (Crlf) Reset() {}

// Transform only converts CRLF. This must be chained with Windows1252 encoder to encode to 1252
func (Crlf) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	for nDst < len(dst) && nSrc < len(src) {
		if c := src[nSrc]; c == '\n' {
			if nDst+1 == len(dst) {
				break
			}
			// break if the newline seen has \r before it
			if nSrc != 0 && src[nSrc-1] == '\r' {
				dst[nDst] = c
				nSrc++
				nDst++
				continue
			}
			dst[nDst] = '\r'
			dst[nDst+1] = '\n'
			nSrc++
			nDst += 2
		} else {
			dst[nDst] = c
			nSrc++
			nDst++
		}
	}
	if nSrc < len(src) {
		err = transform.ErrShortDst
	}
	return
}

// NewWriter returns an io.Writer that converts LF line endings to CRLF.
func NewWriter(w io.Writer) io.Writer {
	return transform.NewWriter(w, Windows1252Crlf())
}

// Windows1252Crlf returns a transformer that converts to Crlf and and windows 1252
func Windows1252Crlf() transform.Transformer {
	return transform.Chain(Crlf{}, encoding.ReplaceUnsupported(charmap.Windows1252.NewEncoder()))
}
