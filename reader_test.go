package proof

import (
	"bytes"
	"hash/crc32"
	"io"
	"io/ioutil"
	"testing"
)

func TestCRC(t *testing.T) {
	data := []byte("Hello, World!")
	cases := []struct {
		csum uint32
		err  error
	}{
		{0xec4ac3d0, nil},
		{0xdeadbeef, EHASHFAIL},
	}

	for _, c := range cases {
		buf := bytes.NewBuffer(data)
		h := crc32.NewIEEE()
		r := NewReader32(buf, h, c.csum)

		_, err := ioutil.ReadAll(r)
		if err != c.err {
			t.Fatalf("expected error: %v, got: %v", c.err, err)
		}
	}
}

func TestTeeReader(t *testing.T) {
	data := []byte("Hello, World!")
	sum := uint32(0xec4ac3d0)

	buf := bytes.NewBuffer(data)
	h := crc32.NewIEEE()
	r := io.TeeReader(buf, h)

	_, err := ioutil.ReadAll(r)
	if err != nil || h.Sum32() != sum {
		t.Fatalf("Either error %v or %d != %d", err, h.Sum32(), sum)
	}
}
