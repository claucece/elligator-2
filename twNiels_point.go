package ed448

import "errors"

type twNiels struct {
	a, b, c *bigNumber
}

func newNielsPoint(a, b, c [fieldBytes]byte) *twNiels {
	return &twNiels{
		a: mustDeserialize(serialized(a)),
		b: mustDeserialize(serialized(b)),
		c: mustDeserialize(serialized(c)),
	}
}

func (p *twNiels) copy() *twNiels {
	return &twNiels{
		a: p.a.copy(),
		b: p.b.copy(),
		c: p.c.copy(),
	}
}

func (p *twExtensible) addTwNiels(p2 *twNiels) *twExtensible {
	x := p.x
	y := p.y
	z := p.z
	t := p.t
	u := p.u

	l0 := new(bigNumber)
	l1 := new(bigNumber)

	l1 = l1.sub(y, x)
	l0.mul(p2.a, l1)
	l1 = l1.addRaw(x, y)
	y.mul(p2.b, l1)
	l1.mul(u, t)
	x.mul(p2.c, l1)

	u = u.addRaw(l0, y)
	// This is equivalent do subx_nr in 32 bits. Change if using 64-bits
	t = t.sub(y, l0)

	// This is equivalent do subx_nr in 32 bits. Change if using 64-bits
	y = y.sub(z, x)
	l0 = l0.addRaw(x, z)

	z.mul(l0, y)
	x.mul(y, t)
	y.mul(l0, u)

	return p
}

func convertTwNielsToTwExtensible(dst *twExtensible, src *twNiels) {
	dst.y = dst.y.add(src.b, src.a)
	dst.x = dst.x.sub(src.b, src.a)
	dst.z = dst.z.setUI(1)
	dst.t = dst.x.copy()
	dst.u = dst.y.copy()
}

func (p *twNiels) conditionalNegate(neg word) {
	p.a.conditionalSwap(p.b, neg)
	p.c = p.c.conditionalNegate(neg)
}

func (s *decafScalar) serialize(dst []byte) error {
	wordBytes := wordBits / 8
	if len(dst) < fieldBytes {
		return errors.New("dst length smaller than fieldBytes")
	}

	for i := 0; i*wordBytes < fieldBytes; i++ {
		for j := 0; j < wordBytes; j++ {
			b := s[i] >> uint(8*j)
			dst[wordBytes*i+j] = byte(b)
		}
	}
	return nil
}
