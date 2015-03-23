package proof

import (
	"bytes"
	"hash/crc32"
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
