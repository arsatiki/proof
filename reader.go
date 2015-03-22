package proof

import (
	"hash"
	"io"
)

type Reader32 struct {
	r   io.Reader
	h   hash.Hash32
	sum uint32
}

func NewReader32(r io.reader, h hash.Hash32, sum uint32) {
	// WIP
}

func (r *Reader32) Read(p []byte) (int, error) {
	// WIP
}
