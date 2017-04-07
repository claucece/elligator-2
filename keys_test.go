package ed448

import (
	"crypto/rand"

	. "gopkg.in/check.v1"
)

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
