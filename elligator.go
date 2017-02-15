package ed448

// q is a odd prime power
// u : a non square in F_q
// A, B := elements of Fq such that AB(A^2 - 4B) != 0
// R := {r in F_q : 1 + ur^2 != 0, A^2ur^2 != B(1 + ur^2)^2

// v := -A/(1/ur^2)
// ep := x_(v^3+Av^2+Bv)
// x:= epv - (1 - ep)A/2
// y := -ep * sqrt(x^3_Ax^2_Bx)

// x_ :=  : F_q -> F_q by x_(a)^((q-1)/2)
//u = (1 - t)/(1 + t)
//v = u^5 + (r^2 - 2)u^3 + u
//X = x_(v)u
//Y = (x_(v)v^((q+1)/4) x_(v)*x_(u^2+1/c^2)
//x = (c - 1)sX(1 + X)/Y
//y = (rX -(1 + X)^2) = (rX + (1 + X)^2)

const (
	quadraticNonresidue = -1
)

// Inverse square root using addition chain.
// test this

func decafIsqrtChk(y, x *bigNumber, zero word) word {
	tmp0, tmp1 := &bigNumber{}, &bigNumber{}
	y.isr(x)
	tmp0.square(y)
	tmp1.mul(tmp0, x)
	tmp2 := tmp1.decafEq(bigZero)
	tmp3 := tmp1.decafEq(bigOne)
	return tmp3 | (zero & tmp2)
}

// 2-torque a point
func point2Torque(p, q *twExtendedPoint) {
	p.x.sub(bigZero, q.x)
	p.y.sub(bigZero, q.y)
	p.z = q.z.copy()
	p.t = q.t.copy()
}

//This function runs Elligator2 on the decaf Jacobi quartic model.  It then
// uses the isogeny to put the result in twisted Edwards form.  As a result,
// it is safe (cannot produce points of order 4), and would be compatible with
// hypothetical other implementations of Decaf using a Montgomery or untwisted
// Edwards model.
// gives out the data hashed to the curve
// buff bytes are less than 56
func decafNonUnifromHashToCurve(ser [56]byte) (*twExtendedPoint, word) {
	// XXD = (u*r^2 + 1) * (d - u*r^2) * (1 - u*d*r^2) / (d+1) // c=u*r^2 && s = r
	//  factor(XX / (1/XXD))
	//  (u*r^2 - d)^2
	//  factor((ey-1)/(ey+1)/(1/d * 1/XXD))
	//  (u*d*r^2 - 1)^2
	//  factor(XX2 / (u*r^2/XXD))
	//  (u*d*r^2 - 1)^2
	//  factor((ey2-1)/(ey2+1)/(1/d * u*r^2/XXD))
	//  (u*r^2 - d)^2
	r, a, b, c, dee, d2, n, rn, e := &bigNumber{}, &bigNumber{}, &bigNumber{}, &bigNumber{}, &bigNumber{}, &bigNumber{}, &bigNumber{}, &bigNumber{}, &bigNumber{}

	p := &twExtendedPoint{
		x: new(bigNumber),
		y: new(bigNumber),
		z: new(bigNumber),
		t: new(bigNumber),
	}

	// probable nonresidue

	r0, overT := deserializeReturnMask(ser) //is it neg r0?
	over := ^(overT)

	sgnR0 := highBit(r0)
	r0.strongReduce()
	a.square(r0) //r^2
	r.sub(bigZero, a)
	//r.decafMulW(a, QuadraticNonresidue) // urr = u*r^2
	dee.mulWSignedCurveConstant(bigOne, sdword(edwardsD)) // dee = 1*D
	c.mulWSignedCurveConstant(r, edwardsD)

	// Compute D2 := (dr+a-d)(dr-ar-d) with a=1 */ // from Decaf paper
	a.sub(c, dee)    // D - D
	a.add(a, bigOne) // D + 1
	specialIdentity := a.decafEq(bigZero)
	b.sub(c, r)
	b.sub(b, dee)
	d2.mul(a, b)

	// compute N := (r+1)(a-2d)
	a.add(r, bigOne)
	n.mulWSignedCurveConstant(a, 1-2*edwardsD)

	// e = +-1/sqrt(+-ND)
	rn.mul(r, n)
	a.mul(rn, d2)

	square := decafIsqrtChk(e, a, word(0))
	isZero := r.decafEq(bigZero)
	square |= isZero
	square |= specialIdentity

	// b <- t/s
	c.decafConstTimeSel(r0, r, square) // r? = sqr ? r : 1

	// In two steps to avoid overflow
	a.mulWSignedCurveConstant(c, 1-2*edwardsD)
	b.mulWSignedCurveConstant(a, 1-2*edwardsD)
	c.sub(r, bigOne)
	a.mul(b, c) /* = r? * (r-1) * (a-2d)^2 with a=1 */
	b.mul(a, e)
	b.decafCondNegate(^square)
	c.decafConstTimeSel(r0, bigOne, square)
	a.mul(e, c)
	c.mul(a, d2) // 1/s except for sign.
	b.sub(b, c)

	// a <- s = e * N * (sqr ? r : r0)
	// e^2 r N D = 1
	// 1/s =  1/(e * N * (sqr ? r : r0)) = e * D * (sqr ? 1 : r0)
	a.mul(n, r0)
	rn.decafConstTimeSel(a, rn, square)
	a.mul(rn, e)
	c.mul(a, b)

	// Normalize/negate
	negS := highBit(a) ^ (^square)
	a.decafCondNegate(negS)
	// ends up negative if !square

	sgnOverS := highBit(b) ^ negS
	sgnOverS &= ^(n.decafEq(bigZero))
	sgnOverS |= d2.decafEq(bigZero)

	// b <- t
	tmp := c.decafEq(bigZero)
	b.decafConstTimeSel(c, bigOne, tmp) // 0,0 -> 1,0

	// isogenize
	c.square(a) // s^2
	a.add(a, a) // 2s
	e.add(c, bigOne)
	p.t.mul(a, e) // 2s(1+s^2)
	p.x.mul(a, b) // 2st
	a.sub(bigOne, c)
	p.y.mul(e, a) // (1+s^2)(1-s^2)
	p.z.mul(a, b) // (1-s^2)t

	succ := (^square & 1) | (sgnOverS & 2) | (sgnR0 & 4) | (over & 8)
	return p, succ
}

// twice and adding
// not sure how to do this
func decafUniformFromHashToCurve(in []byte) (*twExtendedPoint, word) {

	p1 := &twExtendedPoint{
		x: new(bigNumber),
		y: new(bigNumber),
		z: new(bigNumber),
		t: new(bigNumber),
	}

	var in1 [56]byte
	copy(in1[:], in[:56])

	var in2 [56]byte
	copy(in2[:], in[56:])

	p1, ret1 := decafNonUnifromHashToCurve(in1)
	p2, ret2 := decafNonUnifromHashToCurve(in2)

	p1.add(p1, p2)

	succ := ret1 | (ret2 << 4)

	return p1, succ
}

// Inverse of elligator-like hash to curve.
func decafInvertNonUniformElligator(p *twExtendedPoint, hint word) ([]byte, word) {
	sgnS := ^(hint & 1)
	sgnOverS := ^(hint >> 1 & 1)
	sgnR0 := ^(hint >> 2 & 1)

	a, b, c, d := &bigNumber{}, &bigNumber{}, &bigNumber{}, &bigNumber{}

	p.deisogenize(sgnS, sgnOverS)

	//ok, s = a; c = -t/s

	b.mul(c, a)
	b.sub(bigOne, b) /* t+1 */
	c.square(a)

	// s^2
	// identity adjustments
	// in case of identity, currently c=0, t=0, b=1, will encode to 1
	// if hint is 0, -> 0
	// if hint is to neg t/s, then go to infinity, effectively set s to 1
	isIdentity := p.x.decafEq(bigZero)
	c.decafConstTimeSel(c, bigOne, isIdentity&sgnS)
	b.decafConstTimeSel(b, bigZero, isIdentity & ^(sgnOverS) & ^(sgnS)) // identity adjust // is it not?

	d.mulWSignedCurveConstant(c, (2*edwardsD - 1)) // $d = (2d-a)s^2
	a.add(b, d)                                    // num?
	d.sub(b, d)                                    // den?
	b.mul(a, d)                                    // n*d
	a.decafConstTimeSel(d, a, sgnS)
	succ := decafIsqrtChk(c, b, word(0xffffffff))
	b.mul(a, c)
	b.decafCondNegate(sgnR0 ^ highBit(b))

	succ &= ^(b.decafEq(bigZero) & sgnR0)

	var recovered [56]byte
	serialize(recovered[:], b)

	return recovered[:], succ
}

//func decafInvertUniformElligator(p *pointT, hint dword_t, ser serialized) ([]byte, dword_t) {
//
//	p2 := &pointT{
//		x: new(bigNumber),
//		y: new(bigNumber),
//		z: new(bigNumber),
//		t: new(bigNumber),
//	}
//
//	p1, _ := decafNonunifromHashToCurve(ser)
//	p2.decafPointSub(p, p1)
//	hash, succ := decafInvertNonUniformElligator(p2, hint)
//	return hash, succ
//}
