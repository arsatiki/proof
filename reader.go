package proof

import (
	"errors"
	"hash"
	"io"
)

var EHASHFAIL error = errors.New("checksum mismatch")

type Reader32 struct {
	r   io.Reader
	h   hash.Hash32
	sum uint32
}

func NewReader32(r io.Reader, h hash.Hash32, sum uint32) *Reader32 {
	if r == nil || h == nil {
		return nil
	}
	r32 := &Reader32{r, h, sum}
	return r32
}

func (r *Reader32) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	r.h.Write(p[:n])

	if err == io.EOF && r.h.Sum32() != r.sum {
		err = EHASHFAIL
	}

	return n, err
}
