package validutf8

import (
	"bufio"
	"io"
	"unicode"
	"unicode/utf8"
)

// Reader implements a Reader which reads only bytes that constitute valid UTF-8
type Reader struct {
	buffer *bufio.Reader
}

// Function Read reads bytes in the byte array b. n is the number of bytes read.
func (rd Reader) Read(b []byte) (n int, err error) {
	for {
		var r rune
		var size int
		r, size, err = rd.buffer.ReadRune()
		if err != nil {
			return
		}
		if r == unicode.ReplacementChar && size == 1 {
			continue
		} else if n+size < len(b) {
			utf8.EncodeRune(b[n:], r)
			n += size
		} else {
			rd.buffer.UnreadRune()
			break
		}
	}
	return
}

// NewReader constructs a new Reader that wraps an existing io.Reader
func NewReader(rd io.Reader) Reader {
	return Reader{bufio.NewReader(rd)}
}
