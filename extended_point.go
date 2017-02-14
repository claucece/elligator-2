package ed448

type twExtendedPoint struct {
	x, y, z, t *bigNumber
}

func (p *twExtendedPoint) isValidPoint() bool {
	a, b, c := &bigNumber{}, &bigNumber{}, &bigNumber{}
	a.mul(p.x, p.y)
	b.mul(p.z, p.t)
	valid := a.decafEq(b)
	a.square(p.x)
	b.square(p.y)
	a.sub(b, a)
	b.square(p.t)
	c.mulW(b, 1-edwardsD)
	b.square(p.z)
	b.sub(b, c)
	valid &= a.decafEq(b)
	valid &= ^(p.z.decafEq(bigZero))

	return valid == decafTrue
}

func (p *twExtendedPoint) copy() *twExtendedPoint {
	n := &twExtendedPoint{}
	n.x = p.x.copy()
	n.y = p.y.copy()
	n.z = p.z.copy()
	n.t = p.t.copy()
	return n
}

func (p *twExtendedPoint) setIdentity() {
	p.x.setUI(0)
	p.y.setUI(1)
	p.z.setUI(1)
	p.t.setUI(0)
}

func (p *twExtendedPoint) equals(q *twExtendedPoint) word {
	a, b := &bigNumber{}, &bigNumber{}
	a.mul(p.y, q.x)
	b.mul(q.y, p.x)
	return a.decafEq(b)
}

func (p *twExtendedPoint) add(q, r *twExtendedPoint) {
	a, b, c, d := &bigNumber{}, &bigNumber{}, &bigNumber{}, &bigNumber{}
	b.sub(q.y, q.x)
	c.sub(r.y, r.x)
	d.addRaw(r.y, r.x)
	a.mul(c, b)
	b.addRaw(q.y, q.x)
	p.y.mul(d, b)
	b.mul(r.t, q.t)
	p.x.mulW(b, 2-2*edwardsD)
	b.addRaw(a, p.y)
	c.sub(p.y, a)
	a.mul(q.z, r.z)
	a.addRaw(a, a)
	p.y.addRaw(a, p.x)
	a.sub(a, p.x)
	p.z.mul(a, p.y)
	p.x.mul(p.y, c)
	p.y.mul(a, b)
	p.t.mul(b, c)
}

func (p *twExtendedPoint) sub(q *twExtendedPoint, r *twExtendedPoint) {
	a, b, c, d := &bigNumber{}, &bigNumber{}, &bigNumber{}, &bigNumber{}
	b.sub(q.y, q.x)
	d.sub(r.y, r.x)
	c.addRaw(r.y, r.x)
	a.mul(c, b)
	b.addRaw(q.y, q.x)
	p.y.mul(d, b)
	b.mul(r.t, q.t)
	p.x.mulW(b, 2-2*edwardsD)
	b.addRaw(a, p.y)
	c.sub(p.y, a)
	a.mul(q.z, r.z)
	a.addRaw(a, a)
	p.y.sub(a, p.x)
	a.addRaw(a, p.x)
	p.z.mul(a, p.y)
	p.x.mul(p.y, c)
	p.y.mul(a, b)
	p.t.mul(b, c)
}

// Based on Hisil's formula 5.1.3: Doubling in E^e
func (p *twExtendedPoint) double(beforeDouble bool) *twExtendedPoint {
	a, b, c, d := &bigNumber{}, &bigNumber{}, &bigNumber{}, &bigNumber{}
	c.square(p.x)
	a.square(p.y)
	d.addRaw(c, a)
	p.t.addRaw(p.y, p.x)
	b.square(p.t)
	exponentBias := word(3)
	b.subXBias(b, d, exponentBias)
	p.t.sub(a, c)
	p.x.square(p.z)
	p.z.addRaw(p.x, p.x)
	exponentBias = word(4)
	a.subXBias(p.z, p.t, exponentBias)
	p.x.mul(a, b)
	p.z.mul(p.t, a)
	p.y.mul(p.t, d)
	if !beforeDouble {
		p.t.mul(b, d)
	}
	return p
}

func (p *twExtendedPoint) decafEncode(dst []byte) {
	if len(dst) != fieldBytes {
		panic("Attempted an encode with a destination that is not 56 bytes")
	}
	t := word(0x00)
	overT := word(0x00)
	serialize(dst, p.deisogenize(t, overT))
}

func (p *twExtendedPoint) deisogenize(t, overT word) *bigNumber {
	a, b, c, d, s := &bigNumber{}, &bigNumber{}, &bigNumber{}, &bigNumber{}, &bigNumber{}
	a.mulWSignedCurveConstant(p.y, 1-(edwardsD))
	c.mul(a, p.t)
	a.mul(p.x, p.z)
	d.sub(c, a)
	a.add(p.z, p.y)
	b.sub(p.z, p.y)
	c.mul(b, a)
	b.mulWSignedCurveConstant(c, (-(edwardsD)))
	a.isr(b)
	b.mulWSignedCurveConstant(a, (-(edwardsD)))
	c.mul(b, a)
	a.mul(c, d)
	d.add(b, b)
	c.mul(d, p.z)
	b.decafCondNegate(overT ^ (^(highBit(c))))
	c.decafCondNegate(overT ^ (^(highBit(c))))
	d.mul(b, p.y)
	s.add(a, d)
	s.decafCondNegate(overT ^ highBit(s))

	return s
}

func decafDecode(dst *twExtendedPoint, src serialized, useIdentity bool) word {
	a, b, c, d, e := &bigNumber{}, &bigNumber{}, &bigNumber{}, &bigNumber{}, &bigNumber{}

	n, succ := deserializeReturnMask(src)
	zero := n.decafEq(bigZero)
	if useIdentity {
		succ &= decafTrue | ^zero
	} else {
		succ &= decafFalse | ^zero
	}
	succ &= ^highBit(n)
	a.square(n)
	dst.z.sub(bigOne, a)
	b.square(dst.z)
	c.mulWSignedCurveConstant(a, 4-4*(edwardsD))
	c.add(c, b)
	b.mul(c, a)
	d.isr(b)
	e.square(d)
	a.mul(e, b)
	a.add(a, bigOne)
	succ &= ^(a.decafEq(bigZero))
	b.mul(c, d)
	d.decafCondNegate(highBit(b))
	dst.x.add(n, n)
	c.mul(d, n)
	b.sub(bigTwo, dst.z)
	a.mul(b, c)
	dst.y.mul(a, dst.z)
	dst.t.mul(dst.x, a)
	dst.y[0] -= zero

	return succ
}
