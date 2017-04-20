package ed448

var (
	modulus = &bigNumber{
		0xffffff, 0xffffffff, 0xffffff, 0xffffffff,
		0xffffff, 0xffffffff, 0xffffff, 0xffffffff,
		0xfffffe, 0xffffffff, 0xffffff, 0xffffffff,
		0xffffff, 0xffffffff, 0xffffff, 0xffffffff,
	}
)

// the u coorof the base point is 5
func desisogenizeFromMontgomery(src []byte) (*twExtendedPoint, word) {
	u, t1, t2, t3, t4, v, y, o := &bigNumber{}, &bigNumber{}, &bigNumber{}, &bigNumber{}, &bigNumber{}, &bigNumber{}, &bigNumber{}, &bigNumber{}

	p := &twExtendedPoint{
		&bigNumber{},
		&bigNumber{},
		&bigNumber{},
		&bigNumber{},
	}

	succ := decafTrue
	succ &= dsaLikeDeserialize(u, src[:])

	// given u recover edwards v
	// v^2 = u^3 + A*u^2 + u
	// this is u(u^2 + Au + 1)
	bigA := dword(156326)
	t1.square(u)
	t2.mulW(u, bigA) // this prob can be a normal mul
	t3.add(t1, t2)
	t3.add(t3, bigOne) //check
	t4.mul(u, t3)
	v.mul(t4, t4)
	// a square of v

	//given u and v, recover y
	// -(u^5 - 2*u^3 - 4*u*v^2 + u)/
	// (u^5 - 2*u^2*v^2 - 2*u^3 - 2*v^2 + u))

	// o = (u^2)^2
	// -(u(o - 2*u^2 - 4*v^2 + 1)
	t1.square(u)
	o.square(t1)
	t2.square(v)
	t3.mul(t1, &bigNumber{2})
	t4.mul(t2, &bigNumber{4})
	t1.sub(o, t3)
	t2.sub(t1, t4)
	t3.add(t4, bigOne)
	t1.mul(u, t4)
	t3.decafCondNegate(lmask) // num

	// (u^5 - 2*u^2*v^2 - 2*u^3 - 2*v^2 + u))

	// o = (u^2)^2
	// (u(o - 2*u*v^2 - 2*v^2 - 2*v + 1)
	t4.mul(u, t2)
	t1.mul(t4, &bigNumber{2})
	t5, t6 := &bigNumber{}, &bigNumber{}
	t5.mul(t2, &bigNumber{2})
	t6.mul(v, &bigNumber{2})
	t2.sub(o, t1)
	t1.sub(t2, t5)
	t5.sub(t1, t6)
	t6.add(t5, bigOne)
	t5.mul(t6, u)

	t7 := invert(t3)
	y.mul(t7, t5)

	// recover the x with the y
	p.x.square(y)
	p.z.sub(bigOne, p.x)                       // num = 1 - (y^2)
	p.t.mulWSignedCurveConstant(p.x, edwardsD) // d * (y^2)
	p.t.sub(bigOne, p.t)                       // denom = 1 - d * (y^2)
	p.x.mul(p.z, p.t)
	p.t.isr(p.x)      // 1/sqrt(num * denom) // implement it with check
	p.x.mul(p.t, p.z) // sqrt(num / denom)
	p.x.decafCondNegate(^lowBit(p.x) ^ word(1))
	p.z = bigOne.copy()

	// 4-isogeny 2xy/(y^2-ax^2), (y^2+ax^2)/(2-y^2-ax^2)
	a, b, c, d := &bigNumber{}, &bigNumber{}, &bigNumber{}, &bigNumber{}
	c.square(p.x)
	a.square(p.y)
	d.add(c, a)
	p.t.add(p.y, p.x)
	b.square(p.t)
	b.sub(b, d)
	p.t.sub(a, c)
	p.x.square(p.z)
	p.z.add(p.x, p.x)
	a.sub(p.z, d)
	p.x.mul(a, b)
	p.z.mul(p.t, a)
	p.y.mul(p.t, d)
	p.t.mul(b, d)

	// wipe out
	a.set(bigZero)
	b.set(bigZero)
	c.set(bigZero)
	d.set(bigZero)
	src = make([]byte, 57)

	ok := p.isOnCurve()
	if !ok {
		return nil, decafFalse
	}

	return p, succ
}

func dsaLikeDeserialize(n *bigNumber, in []byte) word {
	j, fill := uint(0), uint(0)
	buffer := dword(0x00)
	scarry := sdword(0x00)

	// XXX: unroll my power!!
	for i := uint(0); i < 16; i++ {
		for fill < radix && j < fieldBytes {
			buffer |= dword(in[j]) << fill
			fill += 8
			j++
		}

		if !(i < 16-1) {

			n[i] = word(buffer)
		}
		n[i] = word(buffer & ((dword(1 << radix)) - 1))

		fill -= radix
		buffer >>= radix
		scarry = sdword((word(scarry) + n[i] - modulus[i]) >> 8 * 4)
	}

	// XXX: check me, and add case when hibit is one
	var high word = 0x01
	succ := -(high)
	succ &= isZeroMask(word(buffer))
	succ &= ^(isZeroMask(word(scarry)))

	return succ
}
