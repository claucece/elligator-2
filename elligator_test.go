package ed448

import . "gopkg.in/check.v1"

func (s *Ed448Suite) Test_2Torsion(c *C) {

	p := &twExtendedPoint{
		&bigNumber{0x00},
		&bigNumber{0x01},
		&bigNumber{0x01},
		&bigNumber{0x00},
	}

	q := &twExtendedPoint{
		new(bigNumber),
		new(bigNumber),
		new(bigNumber),
		new(bigNumber),
	}

	e := &twExtendedPoint{
		// gives you p
		&bigNumber{0xfffffff, 0xfffffff, 0xfffffff, 0xfffffff, 0xfffffff, 0xfffffff, 0xfffffff, 0xfffffff, 0xffffffe, 0xfffffff, 0xfffffff, 0xfffffff, 0xfffffff, 0xfffffff, 0xfffffff, 0xfffffff},
		&bigNumber{0xffffffe, 0xfffffff, 0xfffffff, 0xfffffff, 0xfffffff, 0xfffffff, 0xfffffff, 0xfffffff, 0xffffffe, 0xfffffff, 0xfffffff, 0xfffffff, 0xfffffff, 0xfffffff, 0xfffffff, 0xfffffff},
		&bigNumber{0x01},
		&bigNumber{0x00},
	}

	point2Torque(q, p)

	c.Assert(q.x, DeepEquals, e.x)
	c.Assert(q.y, DeepEquals, e.y)
	c.Assert(q.z, DeepEquals, e.z)
	c.Assert(q.t, DeepEquals, e.t)
}

//The empty or all-zero string maps to the identity, as does the string "\x01". If the buffer is shorter than 2*HASH_BYTES, well, it won't be as uniform, but the buffer will be zero-padded on the right.
//less than 56 is non uniform

func (s *Ed448Suite) Test_SetFromNonUniformHash(c *C) {
	//identity
	b := serialized{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	exp := &twExtendedPoint{
		&bigNumber{0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0ffffffe, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
		},
		&bigNumber{0x00000000, 0x10000000,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0ffffffe, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
		},
		&bigNumber{0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
		},
		&bigNumber{0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0ffffffe, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
		},
	}

	p, hint := decafNonUnifromHashToCurve(b)
	c.Assert(p, DeepEquals, exp)
	c.Assert(hint, DeepEquals, word(0))

	// if the size is 56 or less. If less, pad with zeros
	b2 := serialized{
		0xba, 0x70, 0x84, 0x14, 0x7f, 0x7d, 0x48, 0x31,
		0x99, 0x89, 0x54, 0x10, 0x35, 0x5b, 0xf7, 0xd0,
		0x1c, 0x5d, 0x4e, 0x0c, 0x0e, 0x10, 0xc7, 0xa0,
		0x60, 0x82, 0x51, 0x00, 0x29, 0x23, 0xbb, 0xeb,
		0xb9, 0x13, 0x92, 0x2c, 0xb5, 0xc7, 0xc8, 0x01,
		0x80, 0xd4, 0x30, 0x38, 0xb0, 0x3d, 0xb7, 0x21,
		0x27, 0xbf, 0x7a, 0xc9, 0x3b, 0x80, 0xe2, 0x87,
	}

	exp2 := &twExtendedPoint{
		&bigNumber{
			0x093e25dd, 0x0607e543,
			0x0307034f, 0x0a8d41c6,
			0x0a06b83f, 0x03e0ae7f,
			0x0b1a0d8c, 0x0ebd989e,
			0x0380386c, 0x08dca9d2,
			0x0f9d2ab7, 0x0416e2b2,
			0x0980ce3c, 0x0b4b6da1,
			0x0c3cbfe2, 0x0e532bcc,
		},
		&bigNumber{
			0x0caad1a0, 0x02dd8093,
			0x02e2f4e7, 0x03e0841b,
			0x09f9bcfb, 0x0e8af421,
			0x03b7d734, 0x04e2e957,
			0x0f43817c, 0x0f3fd154,
			0x0551ee34, 0x01375c4b,
			0x0032d670, 0x0f1702df,
			0x07836cd6, 0x0ff97486,
		},
		&bigNumber{
			0x0519f530, 0x0ecd1c56,
			0x0b302878, 0x0ec241bc,
			0x03f73adc, 0x00c051ff,
			0x0aba9eda, 0x02596634,
			0x0867640f, 0x0867d056,
			0x0c6dc48f, 0x0b70c707,
			0x004fe2d5, 0x05f455ac,
			0x0d314047, 0x075dd11d,
		},
		&bigNumber{
			0x0dd0d9ef, 0x01b7bc0a,
			0x08bc7dd4, 0x024fae37,
			0x0fb145b8, 0x0e4a498e,
			0x056e62f3, 0x01a28112,
			0x0c6d3c3b, 0x0439a868,
			0x038f9913, 0x0dbfcb19,
			0x0f5b60ca, 0x0ebd3101,
			0x046d19ce, 0x01082d37,
		},
	}

	q, hint2 := decafNonUnifromHashToCurve(b2)
	c.Assert(q, DeepEquals, exp2)
	c.Assert(hint2, DeepEquals, word(0x04))

}

func (s *Ed448Suite) Test_SetFromUniformHash(c *C) {
	c.Skip("not sure")
	b := []byte{
		0xba, 0x70, 0x84, 0x14, 0x7f, 0x7d, 0x48, 0x31,
		0x99, 0x89, 0x54, 0x10, 0x35, 0x5b, 0xf7, 0xd0,
		0x1c, 0x5d, 0x4e, 0x0c, 0x0e, 0x10, 0xc7, 0xa0,
		0x60, 0x82, 0x51, 0x00, 0x29, 0x23, 0xbb, 0xeb,
		0xb9, 0x13, 0x92, 0x2c, 0xb5, 0xc7, 0xc8, 0x01,
		0x80, 0xd4, 0x30, 0x38, 0xb0, 0x3d, 0xb7, 0x21,
		0x27, 0xbf, 0x7a, 0xc9, 0x3b, 0x80, 0xe2, 0x87,
		0xb2, 0xa7, 0x8b, 0x53, 0xce, 0xf6, 0x95, 0x47,
		0x47, 0x20, 0xdc, 0x44, 0xa1, 0x57, 0x12, 0x48,
		0xc9, 0x1c, 0xa3, 0x88, 0xad, 0x4f, 0x5f, 0x1a,
		0xcc, 0x76, 0x61, 0x12, 0x01, 0x6e, 0x27, 0x6c,
		0x16, 0xe9, 0xb9, 0x77, 0xcf, 0x94, 0x07, 0x8d,
		0x8e, 0x95, 0xc3, 0x74, 0xe5, 0xd4, 0xe2, 0x22,
		0xbc, 0xc1, 0x8c, 0x1e, 0x2f, 0xfa, 0x48, 0xb5,
	}

	exp := &twExtendedPoint{
		&bigNumber{
			0x0ede1343, 0x0d4c78f1,
			0x04668513, 0x02ee10cf,
			0x0a4fb950, 0x0d5c14c0,
			0x08ed74a9, 0x04f5e108,
			0x083a08cb, 0x07d7e67f,
			0x0d239385, 0x09dcf285,
			0x010e73bc, 0x00ce3785,
			0x072b7060, 0x04cd05cb,
		},
		&bigNumber{
			0x0d5995ee, 0x0b551f0d,
			0x033eb345, 0x05ef736c,
			0x01a7683c, 0x094f9002,
			0x026f9b74, 0x03f71838,
			0x0283376c, 0x054d1567,
			0x0ce66af6, 0x02b8cdce,
			0x01e68477, 0x04ef089a,
			0x0def56ea, 0x014ee78e,
		},
		&bigNumber{
			0x0a61edec, 0x07162e9c,
			0x0b6232c6, 0x04684b17,
			0x013026db, 0x079c761e,
			0x0d70c52c, 0x020d43f1,
			0x03a4bc3f, 0x07c456af,
			0x0727cba3, 0x00fdb784,
			0x06fab20f, 0x0197bf04,
			0x01297231, 0x074dd768,
		},
		&bigNumber{
			0x0611791e, 0x0e46c648,
			0x0bbdd730, 0x03f70a7f,
			0x06d13ffe, 0x0dbbf727,
			0x0a030c18, 0x0ebf07bf,
			0x06c26a04, 0x0f3a6f0d,
			0x0f668abd, 0x0a604e72,
			0x00c13a2d, 0x0f16970a,
			0x0c741b2c, 0x08a43f74,
		},
	}

	p, hint := decafUniformFromHashToCurve(b)
	c.Assert(p.x, DeepEquals, exp.x)
	c.Assert(p.y, DeepEquals, exp.y)
	c.Assert(p.z, DeepEquals, exp.z)
	c.Assert(p.t, DeepEquals, exp.t)
	c.Assert(hint, DeepEquals, word(0x44))
}

// Not sure
func (s *Ed448Suite) Test_InvertNonUniformElligator(c *C) {
	p := &twExtendedPoint{
		&bigNumber{
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
		},
		&bigNumber{
			0x00000001, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
		},
		&bigNumber{
			0x00000001, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
		},
		&bigNumber{
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
		},
	}

	hint := word(0xfffffff)

	exp := []byte{
		0xfe, 0xff, 0xff, 0xef, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xfe, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

	b, succ := decafInvertNonUniformElligator(p, hint)

	c.Assert(b, DeepEquals, exp)
	c.Assert(succ, Equals, word(0xffffffff))
}
