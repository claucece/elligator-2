package ed448

import (
	"crypto/rand"
	"io"

	. "gopkg.in/check.v1"
)

type fixedRandReader struct {
	data []byte
	at   int
}

func fixedRand(data []byte) io.Reader {
	return &fixedRandReader{data, 0}
}

func (r *fixedRandReader) Read(p []byte) (n int, err error) {
	if r.at < len(r.data) {
		n = copy(p, r.data[r.at:])
		r.at += 56
		return
	}
	return 0, io.ErrUnexpectedEOF
}

func (s *Ed448Suite) Test_GenerateKeys(c *C) {
	pub, priv, err := generateKeys(rand.Reader)

	c.Assert(err, IsNil)
	c.Assert(pub, Not(IsNil))
	c.Assert(priv, Not(IsNil))
}

func (s *Ed448Suite) Test_GenerateKeys2(c *C) {
	pub, priv, err := generateKeys2(rand.Reader)

	c.Assert(err, IsNil)
	c.Assert(pub, Not(IsNil))
	c.Assert(priv, Not(IsNil))
}
