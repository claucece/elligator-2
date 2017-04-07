package ed448

type twExtensible struct {
	x, y, z, t, u *bigNumber
}

func (p *twExtensible) double() *twExtensible {
	x := p.x
	y := p.y
	z := p.z
	t := p.t
	u := p.u

	l0 := new(bigNumber)
	l1 := new(bigNumber)
	l2 := new(bigNumber)

	l2.square(x)
	l0.square(y)
	u = u.addRaw(l2, l0)
	t = t.addRaw(y, x)
	l1.square(t)
	t = t.subRaw(l1, u)
	t.bias(3)
	t.weakReduce()
	// This is equivalent do subx_nr in 32 bits. Change if using 64-bits
	l1 = l1.sub(l0, l2)
	x.square(z)
	x.bias(1)
	z = z.addRaw(x, x)
	l0 = l0.subRaw(z, l1)
	l0.weakReduce()
	z.mul(l1, l0)
	x.mul(l0, t)
	y.mul(l1, u)

	return p
}

func untwist(p *twExtensible) *twExtensible {
	out := &twExtensible{
		&bigNumber{},
		&bigNumber{},
		&bigNumber{},
		&bigNumber{},
		&bigNumber{},
	}

	out.y.square(p.z)
	out.z.square(p.x)
	out.u.sub(p.y, p.z)
	out.z.sub(p.z, p.x)
	out.y.mul(out.z, p.y)
	out.z.sub(p.z, p.y)
	out.x.mul(out.z, out.y)
	out.t.mul(out.x, out.u)
	out.y.mul(out.x, out.t)
	out.t.isr(out.y)
	out.u.mul(out.x, out.t)
	out.x.square(out.t)
	out.t.mul(out.y, out.x)
	out.x.mul(p.x, out.u)
	out.y.mul(p.y, out.u)
	out.y.addW(^(out.z.zeroMask())) //check
	out.z.setUI(dword(1))
	out.x = out.t.copy()
	out.y = out.u.copy()

	return out
}

func (p *twExtensible) serializeExtensible() *bigNumber {
	b, l0, l1, l2 := &bigNumber{}, &bigNumber{}, &bigNumber{}, &bigNumber{}
	l0.sub(p.y, p.z)
	b.add(p.z, p.y)
	l1.mul(p.z, p.x)
	l2.mul(l0, l1)
	l1.mul(l2, l0)
	l0.mul(l2, b)
	l2.mul(l1, l0)
	l0.isr(l2)
	b.mul(l1, l0)
	l1.square(l0)
	l0.mul(l2, l1)
	return b //sure?
}

func (p *twExtensible) OnCurve() bool {
	l0 := new(bigNumber)
	l1 := new(bigNumber)
	l2 := new(bigNumber)
	l3 := new(bigNumber)

	// Check invariant:
	// 0 = -x*y + z*t*u
	l1 = l1.mul(p.t, p.u)
	l2 = l2.mul(p.z, l1)
	l0 = l0.mul(p.x, p.y)
	l1 = l1.neg(l0)
	l0 = l0.add(l1, l2)
	l5 := l0.zeroMask()

	// Check invariant:
	// 0 = d*t^2*u^2 + x^2 - y^2 + z^2 - t^2*u^2

	l2 = l2.square(p.y)
	l1 = l1.neg(l2)
	l0 = l0.square(p.x)
	l2 = l2.add(l0, l1)
	l3 = l3.square(p.u)
	l0 = l0.square(p.t)
	l1 = l1.mul(l0, l3)
	l3 = l3.mulWSignedCurveConstant(l1, edwardsD)
	l0 = l0.add(l3, l2)
	l3 = l3.neg(l1)
	l2 = l2.add(l3, l0)
	l1 = l1.square(p.z)
	l0 = l0.add(l1, l2)
	l4 := l0.zeroMask()

	ret := l4 & l5 & (^p.z.zeroMask())
	return ret == decafTrue
}

func (p *twExtensible) untwistAndDoubleAndSerialize() *bigNumber {
	l0 := new(bigNumber)
	l1 := new(bigNumber)
	l2 := new(bigNumber)
	l3 := new(bigNumber)
	b := new(bigNumber)

	l3.mul(p.y, p.x)
	b.add(p.y, p.x)
	l1.square(b)
	l2.add(l3, l3)
	b.sub(l1, l2)
	l2.square(p.z)
	l1.square(l2)
	b.add(b, b)
	l2.mulWSignedCurveConstant(b, edwardsD-1)
	b.mulWSignedCurveConstant(l2, edwardsD-1)
	l0.mul(l2, l1)
	l2.mul(b, l0)
	l0.isr(l2)
	l1.mul(b, l0)

	return b.mul(l1, l3)
}
