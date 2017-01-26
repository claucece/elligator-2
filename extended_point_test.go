package ed448

import (
	. "gopkg.in/check.v1"
)

func (s *Ed448Suite) Test_PointDouble(c *C) {
	q := &twExtendedPoint{
		&bigNumber{0x08354b7a, 0x0895b3e8,
			0x06ae5175, 0x0644b394,
			0x0b7faf9e, 0x0c5237db,
			0x013a0c90, 0x08f5bce0,
			0x09a3d79b, 0x00f17559,
			0x0de8f041, 0x073e222f,
			0x0dc2b7ee, 0x005ac354,
			0x0766db38, 0x065631fe,
		},
		&bigNumber{0x00398885, 0x055c9bed,
			0x0ae443ca, 0x0fd70ea4,
			0x09e2a7d2, 0x04ac2e9d,
			0x00678287, 0x0294768e,
			0x0b604cea, 0x07b49317,
			0x0dc2a6d9, 0x0e44a6fb,
			0x09db3965, 0x049d3bf5,
			0x03e655fe, 0x003a9c02,
		},
		&bigNumber{0x0fd57162, 0x0a39f768,
			0x03009756, 0x065d735f,
			0x0d1da282, 0x0589ecd7,
			0x003196b1, 0x0c001dfe,
			0x019f1050, 0x0152e8d2,
			0x0c14ff38, 0x00f7a446,
			0x028053f6, 0x0f8a91e9,
			0x05a8d694, 0x09d5ae86,
		},
		&bigNumber{0x04198f2e, 0x0d82440f,
			0x0fce100e, 0x0af4829d,
			0x0d5c3516, 0x0094a0da,
			0x078cdb39, 0x0e738836,
			0x01ec536d, 0x06dfd1e9,
			0x0ee16173, 0x0addc8c0,
			0x0797fb1d, 0x059741a3,
			0x0a7f9c34, 0x088fe0a6,
		},
	}

	p := &twExtendedPoint{
		&bigNumber{0}, &bigNumber{0},
		&bigNumber{0}, &bigNumber{0},
	}

	p.double(q, false)

	expected := &twExtendedPoint{
		&bigNumber{0x00d8f04c, 0x03e54689,
			0x0eb4db2b, 0x0887ba34,
			0x0a5b4ebc, 0x0f6c0261,
			0x03bfa803, 0x0408ff02,
			0x03b4ef26, 0x0465c028,
			0x0cd47378, 0x064c55b4,
			0x08245850, 0x01912682,
			0x0dcbf92c, 0x07a7fa30,
		},
		&bigNumber{0x0d94d1a6, 0x0f7306e8,
			0x0278b336, 0x04362b7b,
			0x0faf02b9, 0x06b01d18,
			0x07a597da, 0x0bd6add0,
			0x047afa98, 0x0e64e897,
			0x0bbf88e6, 0x01d0a534,
			0x04a52b9d, 0x0af374e0,
			0x05091d54, 0x00fcf1a5,
		},
		&bigNumber{0x042318ce, 0x04aecdae,
			0x0e8f196b, 0x0019d2e3,
			0x045d147c, 0x060b153e,
			0x0adf2c37, 0x0419cdd8,
			0x06d19046, 0x00d18821,
			0x06c7b9c2, 0x0c0ffd68,
			0x0b7e4ca2, 0x06da0d56,
			0x0952b40f, 0x03008395,
		},
		&bigNumber{0x04643593, 0x000e0fdd,
			0x013f29f3, 0x0bb8992d,
			0x0a30d344, 0x09151eec,
			0x0d12bb82, 0x05c7a054,
			0x0103c2c6, 0x08a61fe2,
			0x0aced4bf, 0x0f76d481,
			0x0db774be, 0x065ef8a8,
			0x0ff47a71, 0x0f49f73e,
		}}

	c.Assert(p, DeepEquals, expected)

	resetPoint(p)

	w := &twExtendedPoint{
		&bigNumber{1},
		&bigNumber{2},
		&bigNumber{3},
		&bigNumber{4},
	}

	p.double(w, true)

	expected2 := &twExtendedPoint{
		&bigNumber{0x0000003b, 0x10000000,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0ffffffe, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
		},
		&bigNumber{0x0000000e, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
		},
		&bigNumber{0x0000002c, 0x10000000,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0ffffffe, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
		},
		&bigNumber{0x00000002, 0x10000000,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0ffffffe, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
		}}

	c.Assert(p, DeepEquals, expected2)
}

func (s *Ed448Suite) Test_PointDoubleAgain(c *C) {
	c.Skip(":|")
	p := &twExtendedPoint{
		&bigNumber{0x0e0fbf9e, 0xba1bcd70, 0x01cc6d39, 0x53b56e80, 0x0635d142, 0x383307a0, 0x0f8a159b, 0x97fd2cf0, 0x0fa310f6, 0x5522bde0, 0x0b981703, 0xb095b1e0, 0x042d4780, 0x5ae11df0, 0x0934fe80, 0xdc6474d0},
		&bigNumber{0x02c1149c, 0xe72febf0, 0x05259893, 0x723e1840, 0x0f7232ff, 0x19a56000, 0x05581d2c, 0x73314440, 0x04e0124a, 0x9c3c5e50, 0x0945536e, 0xb786a200, 0x0f75623f, 0x0ba30e80, 0x0cc589a3, 0x4a2eea80},
		&bigNumber{0x02406c71, 0xb2fdb670, 0x02591aa2, 0x85fc24e0, 0x0dc50d09, 0x8692c5b0, 0x0ba917d7, 0xaefea740, 0x037d0084, 0x4d5defa0, 0x08bbe7ad, 0x50da9770, 0x08adf827, 0x5425cdd0, 0x037d816d, 0xd59cd0a0},
		&bigNumber{0x0baf8c30, 0x6686ad30, 0x0c149bac, 0xf57f68d0, 0x05cd321a, 0x2ff8d600, 0x09dcc4bd, 0xf731ec20, 0x0cd7ea75, 0xbe970e40, 0x043d30e0, 0xdd64b9b0, 0x04f78bf1, 0xd1fde200, 0x05c88e97, 0x26ce3140},
	}

	expected := &twExtendedPoint{
		&bigNumber{0x0cbc0465, 0x0ae92226, 0x091cf2c2, 0x06088c39, 0x06adcba9, 0x02b78bc6, 0x06b0b2aa, 0x0d1b1697, 0x0cf760b0, 0x09cc4ae8, 0x06ff4c95, 0x06f9eaff, 0x03f88cb2, 0x0dfe1f47, 0x015828d5, 0x003f26ff},
		&bigNumber{0x0f73319b, 0x007fa279, 0x0c6ab58a, 0x0224b18c, 0x0bbc2c00, 0x012c8b76, 0x07428968, 0x06047979, 0x0a19d606, 0x0cf6c1c8, 0x06a83f0d, 0x0725b63c, 0x0bc33839, 0x0d9114a0, 0x07eec286, 0x0908447d},
		&bigNumber{0x01e2013d, 0x0287266b, 0x0f434216, 0x08bac041, 0x03321096, 0x096004ef, 0x0c9f384c, 0x0ae584aa, 0x02a4d456, 0x02d20148, 0x0df9a8dd, 0x051a0acf, 0x0b3886b1, 0x00ad3c7f, 0x06f0b032, 0x0e09bf7e},
		&bigNumber{0x0419dfec, 0x0723fb8c, 0x00f6c661, 0x045419d4, 0x00449ef9, 0x09f3003d, 0x0d77bd3d, 0x03aabfcb, 0x05f1ad37, 0x018339b6, 0x0e4963f2, 0x0060cbfc, 0x078b53f3, 0x023457ac, 0x048d759c, 0x02a6c760},
	}

	p.double(p, false)

	c.Assert(p, DeepEquals, expected)
}

func resetPoint(p *twExtendedPoint) {
	p = &twExtendedPoint{
		&bigNumber{0},
		&bigNumber{0},
		&bigNumber{0},
		&bigNumber{0},
	}
}

func (s *Ed448Suite) Test_AddNielsToExtended_BeforeDouble(c *C) {
	extdPoint := &twExtendedPoint{
		&bigNumber{},
		&bigNumber{1},
		&bigNumber{1},
		&bigNumber{},
	}
	n := &twNiels{
		&bigNumber{0x068d5b74},
		&bigNumber{0x068d5b74},
		&bigNumber{0x068d5b74},
	}

	expected := &twExtendedPoint{
		&bigNumber{0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0ffffffe, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
		},
		&bigNumber{0x0d1ab6e7, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x00000000, 0x00000000,
			0x0fffffff, 0x0fffffff,
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
		}, &bigNumber{},
	}

	extdPoint.addNielsToExtended(n, true)

	c.Assert(extdPoint.x, DeepEquals, expected.x)
	c.Assert(extdPoint.y, DeepEquals, expected.y)
	c.Assert(extdPoint.z, DeepEquals, expected.z)
	c.Assert(extdPoint.t, DeepEquals, expected.t)
}

func (s *Ed448Suite) Test_AddNielsToProjective(c *C) {
	n := &twNiels{&bigNumber{0x08fcb20f, 0x04611087,
		0x01cc6f32, 0x0df43db2,
		0x04516644, 0x0ffdde9f,
		0x091686b9, 0x05199177,
		0x0fd34473, 0x0b72b441,
		0x0cb1c72b, 0x08d45684,
		0x00fc17a5, 0x01518137,
		0x007f74d3, 0x0a456d13},
		&bigNumber{0x09b607dc, 0x01430f14,
			0x016715fc, 0x0e992ccd,
			0x00a32a09, 0x0a62209b,
			0x0c26b8e4, 0x0b889ced,
			0x0ac109cf, 0x059bf9a3,
			0x0b7feac2, 0x06871bb3,
			0x0d9a0e6b, 0x0f4a4d5f,
			0x00cd69a5, 0x0b95db46},
		&bigNumber{0x08bda702, 0x03630441,
			0x01561558, 0x07bc5686,
			0x0e30416f, 0x0f344bc8,
			0x080f59d7, 0x0a645370,
			0x07d00ace, 0x0b4c2007,
			0x0b26f8cc, 0x0ee79620,
			0x00b5403d, 0x0a6a558e,
			0x066f3d19, 0x08f1d2c7},
	}

	extdPoint := twExtendedPoint{
		&bigNumber{0x00d8f04c, 0x03e54689,
			0x0eb4db2b, 0x0887ba34,
			0x0a5b4ebc, 0x0f6c0261,
			0x03bfa803, 0x0408ff02,
			0x03b4ef26, 0x0465c028,
			0x0cd47378, 0x064c55b4,
			0x08245850, 0x01912682,
			0x0dcbf92c, 0x07a7fa30,
		},
		&bigNumber{0x0d94d1a6, 0x0f7306e8,
			0x0278b336, 0x04362b7b,
			0x0faf02b9, 0x06b01d18,
			0x07a597da, 0x0bd6add0,
			0x047afa98, 0x0e64e897,
			0x0bbf88e6, 0x01d0a534,
			0x04a52b9d, 0x0af374e0,
			0x05091d54, 0x00fcf1a5,
		},
		&bigNumber{0x042318ce, 0x04aecdae,
			0x0e8f196b, 0x0019d2e3,
			0x045d147c, 0x060b153e,
			0x0adf2c37, 0x0419cdd8,
			0x06d19046, 0x00d18821,
			0x06c7b9c2, 0x0c0ffd68,
			0x0b7e4ca2, 0x06da0d56,
			0x0952b40f, 0x03008395,
		},
		&bigNumber{0x04643593, 0x000e0fdd,
			0x013f29f3, 0x0bb8992d,
			0x0a30d344, 0x09151eec,
			0x0d12bb82, 0x05c7a054,
			0x0103c2c6, 0x08a61fe2,
			0x0aced4bf, 0x0f76d481,
			0x0db774be, 0x065ef8a8,
			0x0ff47a71, 0x0f49f73e,
		},
	}
	expected := &twExtendedPoint{
		&bigNumber{0x0662c9a5, 0x0e2bc383,
			0x09b2fc38, 0x0042d545,
			0x0431bbe8, 0x09e2a364,
			0x03b8e92e, 0x0df6d043,
			0x07136f20, 0x00bde4fe,
			0x0ca79859, 0x0c484320,
			0x099507c4, 0x0ef683e6,
			0x09f8221d, 0x0b1fdcb8,
		},
		&bigNumber{0x0aaf871f, 0x08fcadaf,
			0x0974aaea, 0x07d73c92,
			0x0bdaba0c, 0x069d1bf6,
			0x0906e75c, 0x0020e493,
			0x07a2e1ec, 0x06e27878,
			0x00e9c9d2, 0x08e429f5,
			0x026f7c86, 0x0420e6c5,
			0x0304fccb, 0x0599fe0e,
		},
		&bigNumber{0x01b26129, 0x071c89cf,
			0x0b012391, 0x0074b87c,
			0x0331b5fb, 0x0a2cbc8d,
			0x0d1a4729, 0x0ab451d3,
			0x0308cad6, 0x0e086c2b,
			0x03bd396c, 0x0cd2bd87,
			0x0910f41c, 0x090be75a,
			0x0a8d7a0e, 0x07ec7ea8,
		},
		&bigNumber{0x08b7d023, 0x05bc6276,
			0x03e2082d, 0x09d3eba3,
			0x0ecc2af3, 0x07a4c7be,
			0x08ca49b8, 0x0ebe1040,
			0x0cf6ddeb, 0x015ec1ff,
			0x010eed61, 0x0882e84d,
			0x07fefb78, 0x0d97e204,
			0x02e940a1, 0x0537d7c0,
		},
	}

	extdPoint.addNielsToExtended(n, false)

	c.Assert(extdPoint.x, DeepEquals, expected.x)
	c.Assert(extdPoint.y, DeepEquals, expected.y)
	c.Assert(extdPoint.z, DeepEquals, expected.z)
	c.Assert(extdPoint.t, DeepEquals, expected.t)
}

func (s *Ed448Suite) Test_ConvertNielsToExtended(c *C) {
	p := &twExtendedPoint{
		&bigNumber{},
		&bigNumber{0x01},
		&bigNumber{0x01},
		&bigNumber{},
	}
	niels := &twNiels{
		&bigNumber{0x068d5b74},
		&bigNumber{0x068d5b74},
		&bigNumber{0x068d5b74},
	}

	expected := &twExtendedPoint{
		&bigNumber{0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0ffffffe, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
		},
		&bigNumber{0x0d1ab6e8},
		&bigNumber{0x00000001},
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

	p.nielsToExtended(niels)

	c.Assert(p.x, DeepEquals, expected.x)
	c.Assert(p.y, DeepEquals, expected.y)
	c.Assert(p.z, DeepEquals, expected.z)
	c.Assert(p.t, DeepEquals, expected.t)
}

func (s *Ed448Suite) Test_PrecomputedScalarMultiplication(c *C) {
	scalar := Scalar{0}

	p := curve.precomputedScalarMul(scalar)

	expP := &twExtendedPoint{
		&bigNumber{0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0ffffffe, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
			0x0fffffff, 0x0fffffff,
		},
		&bigNumber{0x0b1ff82e, 0x05e98b74,
			0x000cecf1, 0x0277711a,
			0x0f9b17c5, 0x0c98aadc,
			0x05b06211, 0x0bc17782,
			0x0809fef2, 0x08bb648f,
			0x0323239f, 0x0d37d81d,
			0x0389402c, 0x0cbabc81,
			0x087aaae9, 0x01b50b05,
		},
		&bigNumber{0x04e007d1, 0x0a16748b,
			0x0ff3130e, 0x0d888ee5,
			0x0064e83a, 0x03675523,
			0x0a4f9dee, 0x043e887d,
			0x07f6010c, 0x07449b70,
			0x0cdcdc60, 0x02c827e2,
			0x0c76bfd3, 0x0345437e,
			0x07855516, 0x0e4af4fa,
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

	c.Assert(p.x, DeepEquals, expP.x)
	c.Assert(p.y, DeepEquals, expP.y)
	c.Assert(p.z, DeepEquals, expP.z)
	c.Assert(p.t, DeepEquals, expP.t)
}

func (s *Ed448Suite) Test_PrepareFixedWindow(c *C) {
	c.Skip("waiting for subfuctions to be implemented")
	p := &twExtendedPoint{
		&bigNumber{0x0e0fbf9e, 0xba1bcd70, 0x01cc6d39, 0x53b56e80, 0x0635d142, 0x383307a0, 0x0f8a159b, 0x97fd2cf0, 0x0fa310f6, 0x5522bde0, 0x0b981703, 0xb095b1e0, 0x042d4780, 0x5ae11df0, 0x0934fe80, 0xdc6474d0},
		&bigNumber{0x02c1149c, 0xe72febf0, 0x05259893, 0x723e1840, 0x0f7232ff, 0x19a56000, 0x05581d2c, 0x73314440, 0x04e0124a, 0x9c3c5e50, 0x0945536e, 0xb786a200, 0x0f75623f, 0x0ba30e80, 0x0cc589a3, 0x4a2eea80},
		&bigNumber{0x02406c71, 0xb2fdb670, 0x02591aa2, 0x85fc24e0, 0x0dc50d09, 0x8692c5b0, 0x0ba917d7, 0xaefea740, 0x037d0084, 0x4d5defa0, 0x08bbe7ad, 0x50da9770, 0x08adf827, 0x5425cdd0, 0x037d816d, 0xd59cd0a0},
		&bigNumber{0x0baf8c30, 0x6686ad30, 0x0c149bac, 0xf57f68d0, 0x05cd321a, 0x2ff8d600, 0x09dcc4bd, 0xf731ec20, 0x0cd7ea75, 0xbe970e40, 0x043d30e0, 0xdd64b9b0, 0x04f78bf1, 0xd1fde200, 0x05c88e97, 0x26ce3140},
	}
	ntable := 16

	w := p.prepareFixedWindow(ntable)

	c.Assert(len(w), Equals, ntable)

	expected := []*twPNiels{
		&twPNiels{
			&twNiels{
				&bigNumber{0x04b15504, 0x0d141e7f, 0x03592b5c, 0x0e88a9c0, 0x093c61be, 0x01725860, 0x05ce078f, 0x0b34174f, 0x053d0167, 0x0719a06f, 0x0dad3c6f, 0x06f0f01f, 0x0b481abf, 0x00c1f090, 0x03908b2e, 0x0dca75b0},
				&bigNumber{0x00d0d43c, 0x014bb961, 0x06f205d6, 0x05f386c0, 0x05a8044d, 0x01d867a1, 0x04e232cc, 0x0b2e7131, 0x04832342, 0x015f1c31, 0x04dd6a80, 0x081c53e1, 0x03a2a9c5, 0x06842c71, 0x05fa8829, 0x06935f51},
				&bigNumber{0x045debb9, 0x07781d40, 0x0bbd5b0e, 0x0650e537, 0x0ffb5de2, 0x0b635849, 0x0d26b188, 0x094ba94b, 0x0acf701b, 0x0e741ce8, 0x04b5f6b7, 0x09b8105b, 0x010e8851, 0x0e6f7837, 0x0852bd6d, 0x0baa27a2},
			},
			&bigNumber{0x0480d8ec, 0x05fb6ce0, 0x04b2354a, 0x0bf849c0, 0x0b8a1a12, 0x0d258b61, 0x07522fae, 0x0dfd4e81, 0x06fa0117, 0x0abbdf40, 0x0177cf63, 0x01b52ee1, 0x015bf058, 0x084b9ba1, 0x06fb02e4, 0x0b39a140},
		},
		&twPNiels{
			&twNiels{
				&bigNumber{0x0bb9b611, 0x0e04e576, 0x00e7b53c, 0x03219258, 0x08627cd9, 0x081aaaac, 0x0ad5b7cc, 0x0efb2e81, 0x04f77f74, 0x001b42ef, 0x0c53bd25, 0x08c13078, 0x014474a6, 0x00e4ea6c, 0x0095ef47, 0x0685fc8e},
				&bigNumber{0x0efecaa8, 0x04a0861e, 0x0002531f, 0x0b48047d, 0x0d0e3086, 0x0a217912, 0x0181947d, 0x05b3c0fa, 0x0213ffd3, 0x017a8c28, 0x0dc38782, 0x058dcf53, 0x03deb16c, 0x09ba8e34, 0x00e3ea4f, 0x0da9b6ec},
				&bigNumber{0x04f60122, 0x0456af55, 0x049cd12f, 0x0fdff0b1, 0x0ec17b3e, 0x0c6bdfaf, 0x03ecbd31, 0x0ef04d76, 0x0a700d34, 0x0b9fbe06, 0x03ffa657, 0x0f640301, 0x00e26a71, 0x0fc3a319, 0x0438e0a7, 0x0c7f5f07},
			},
			&bigNumber{0x01f7cc5f, 0x08911bd2, 0x0b741049, 0x02ba701f, 0x0717b73a, 0x08df835f, 0x0eea2bd4, 0x04c57a6b, 0x087d22ef, 0x0e665f02, 0x0794d80a, 0x0b88810d, 0x05139e07, 0x0c0285e5, 0x06672eb3, 0x07c07b97},
		},
		&twPNiels{
			&twNiels{
				&bigNumber{0x0921e451, 0x0f6addc9, 0x0ae67810, 0x04bfefcf, 0x0b577222, 0x0b78187a, 0x008aac0e, 0x0b1880d0, 0x06fe51aa, 0x0b80c957, 0x01210039, 0x0064725d, 0x0eff49cf, 0x01c80809, 0x0c3a75fa, 0x043c018a},
				&bigNumber{0x02494fff, 0x00e44212, 0x0077e261, 0x05568783, 0x06863572, 0x0111ec07, 0x06898c13, 0x0dc9c65d, 0x0d62b6b5, 0x04f9774d, 0x050fab1b, 0x0427ea63, 0x0e2a1abd, 0x0a61c851, 0x04cb4f15, 0x02500201},
				&bigNumber{0x07b5f04a, 0x03ff4348, 0x0dddc73b, 0x076c1dee, 0x050057b9, 0x03e993f4, 0x0a5f81b6, 0x01d6baba, 0x0c4067c9, 0x08a960f2, 0x0e87c68b, 0x0f8defe3, 0x08fc1146, 0x0eccecfc, 0x095659f9, 0x0d12f5bc},
			},
			&bigNumber{0x0940d17c, 0x087e6ac0, 0x0b919fe4, 0x07e01de4, 0x0687afe8, 0x0455dc41, 0x07bdfd1f, 0x0dd27627, 0x0569692e, 0x01950411, 0x0e05ded0, 0x030dbf5d, 0x0fbdc7d7, 0x0e1600c0, 0x077beea2, 0x0d9bcb73},
		},
		&twPNiels{
			&twNiels{
				&bigNumber{0x09567019, 0x0a6ffe58, 0x08964ffd, 0x0985f405, 0x0a95e25d, 0x0f325c38, 0x05dd535c, 0x0899a763, 0x05c1c3aa, 0x0d79a4e9, 0x064e48e5, 0x02636d06, 0x08c41b49, 0x0130c022, 0x01a58247, 0x06e59231},
				&bigNumber{0x0c2db17f, 0x005966c7, 0x0af4e20a, 0x02893dfa, 0x07516c50, 0x0100744c, 0x09ab8407, 0x0af874fd, 0x0a79b80e, 0x05c24031, 0x0834a48b, 0x0192373f, 0x07f083a5, 0x0bcef16e, 0x0186e0ad, 0x0b6007e6},
				&bigNumber{0x06a2435b, 0x0f8de8f4, 0x02621c8a, 0x07700e13, 0x084d7864, 0x0ed1da31, 0x05a42140, 0x0f467df1, 0x0616494d, 0x0aba4047, 0x0bddf48f, 0x0fd37695, 0x0f970fd3, 0x080e59d8, 0x061a1b1f, 0x07701da1},
			},
			&bigNumber{0x06277d9a, 0x08564a71, 0x06a8551c, 0x085301e6, 0x0252a48c, 0x0c37ff7b, 0x0b8153a3, 0x0c4051bd, 0x0062d6ec, 0x09a69c0b, 0x0a865464, 0x0811c212, 0x050f4e97, 0x049a7480, 0x09d39f94, 0x0dca0b51},
		},
		&twPNiels{
			&twNiels{
				&bigNumber{0x0401829b, 0x03b32714, 0x09e8c592, 0x05834dd9, 0x0a513a0c, 0x0270a68e, 0x08c98781, 0x04af8368, 0x040d9ea9, 0x0d5d537d, 0x04381df2, 0x0b7a127b, 0x0a7e7f64, 0x094299c3, 0x079f82fd, 0x0a87d3b9},
				&bigNumber{0x09871b11, 0x012626f1, 0x0dbb9075, 0x066a8818, 0x03e5b8ee, 0x0148e221, 0x0e15ce5d, 0x0de23c37, 0x02f45e1e, 0x0a271cbf, 0x0acf74b9, 0x03ed4b51, 0x0d4c228e, 0x0ff08fc7, 0x0d007a9e, 0x0332791c},
				&bigNumber{0x0e704566, 0x001aa18f, 0x09e1aa0a, 0x05b9da5b, 0x02a18843, 0x06b7ba87, 0x0b62565a, 0x084daeb3, 0x0e49deb9, 0x032ca56e, 0x01334ba3, 0x0ce45ea3, 0x0a108917, 0x0152423b, 0x0597d39b, 0x06fe6fff},
			},
			&bigNumber{0x0e0e4387, 0x0629db63, 0x03635f6a, 0x0ff93274, 0x0f9fc920, 0x029fc483, 0x00af6edc, 0x0cba3b79, 0x0fffe466, 0x00028043, 0x02b4e6bf, 0x0d9d8ffa, 0x0a5b655b, 0x0a278d16, 0x0f25b8f5, 0x08889190},
		},
		&twPNiels{
			&twNiels{
				&bigNumber{0x07560076, 0x05823e24, 0x0025a478, 0x0dbc0186, 0x0c8fca8a, 0x0b5d4904, 0x0fcfc8e1, 0x079cb39b, 0x0259ebad, 0x043381c7, 0x015bdf27, 0x09c79dc7, 0x0d42589a, 0x0829a798, 0x03f85ee0, 0x02738d2d},
				&bigNumber{0x0c89a2ee, 0x0c08f51d, 0x0fa19cce, 0x005cb7ab, 0x0ef9e267, 0x06dec3dc, 0x07c0c5fa, 0x06c93687, 0x03ceb72a, 0x091594de, 0x04a21940, 0x0d421064, 0x06af5d2a, 0x06cabee9, 0x0ad0421d, 0x0f212b4c},
				&bigNumber{0x01069832, 0x03e14eef, 0x004211c0, 0x0bda09c4, 0x0736b10e, 0x048bf9c2, 0x0f5023f2, 0x0fa9d0c4, 0x07be00b5, 0x0f86c62d, 0x0d9c5c6d, 0x0215b5d4, 0x0b19870f, 0x00a94c1d, 0x0b17311b, 0x02d9239f},
			},
			&bigNumber{0x0e5ad33b, 0x019bb763, 0x06a83d49, 0x07451a5f, 0x045c81d4, 0x06686178, 0x02fae4fd, 0x0266cfcf, 0x02fd2ca4, 0x0c8120d0, 0x0e0673a0, 0x06db3227, 0x03cff2d4, 0x06206075, 0x01674689, 0x058e498a},
		},
		&twPNiels{
			&twNiels{
				&bigNumber{0x05c26c7f, 0x09bae1cb, 0x06333392, 0x01e75549, 0x078ce964, 0x03e7a4b7, 0x0215fc98, 0x02e850b3, 0x0b80325b, 0x0a7cb594, 0x04b8745b, 0x0301ef60, 0x0eeea155, 0x0c21b85c, 0x0650a7c1, 0x0e497b58},
				&bigNumber{0x0e1ea322, 0x0eb69c03, 0x06671ead, 0x09600754, 0x0fdac69c, 0x07a76baa, 0x0d07154d, 0x0b33c274, 0x02e11d88, 0x0238fab8, 0x099121a5, 0x0209727f, 0x0eea9985, 0x06a90391, 0x0d745360, 0x0312b0d4},
				&bigNumber{0x09c51c81, 0x0d94c3b7, 0x0afa26b8, 0x00a4f4e0, 0x05691370, 0x0f9ec0a6, 0x0197cbea, 0x000155df, 0x0534e1ff, 0x03cc434a, 0x04b60706, 0x0e1535c1, 0x08024788, 0x0796d3bb, 0x058f1367, 0x0d4a4b28},
			},
			&bigNumber{0x00a34eac, 0x0974e9fa, 0x029b9e0c, 0x029172eb, 0x0fc09f68, 0x03e978d2, 0x000c650f, 0x07876c82, 0x0a8702b8, 0x088e25b0, 0x087ea436, 0x0ef88b56, 0x090b593c, 0x047dc683, 0x035c1df5, 0x07c95327},
		},
		&twPNiels{
			&twNiels{
				&bigNumber{0x0c331b06, 0x0da3956b, 0x08b109f0, 0x0aefd801, 0x07aca32a, 0x0c3be476, 0x0f7408ca, 0x0432f993, 0x0075f10c, 0x04138e38, 0x08c7f415, 0x00984c20, 0x0fc56ee2, 0x04e31a65, 0x064cd457, 0x0d16aab2},
				&bigNumber{0x0f3dffed, 0x07ae5b30, 0x0c3d67ed, 0x0a281d2b, 0x05f81516, 0x03221b45, 0x0c660c70, 0x0c0abbe8, 0x0f1a175b, 0x018fc6bc, 0x010013fb, 0x04d474ae, 0x0eceebbb, 0x0f5d5e7a, 0x00040f4a, 0x07d2164b},
				&bigNumber{0x06903269, 0x0612c1b9, 0x01bf9aed, 0x0cbf0e06, 0x0d78c4f2, 0x044b8023, 0x0e7c45f7, 0x07357b84, 0x08870f86, 0x0cc658ed, 0x030f7f41, 0x02e7a5ae, 0x0f31178b, 0x0ccf41df, 0x06886379, 0x09e558c0},
			},
			&bigNumber{0x059ebc12, 0x002c1977, 0x0564c445, 0x0be7579b, 0x074b53c8, 0x01ae6f07, 0x0be22bc8, 0x06aa6576, 0x08ac90be, 0x08ffdcd4, 0x0e0503d6, 0x0549de6f, 0x04546647, 0x02267b40, 0x0470149a, 0x04d13a4f},
		},
		&twPNiels{
			&twNiels{
				&bigNumber{0x048a8013, 0x09ac5cd3, 0x0f6e0524, 0x07ad4570, 0x02c3808c, 0x05c20659, 0x0cebb3ae, 0x0d0cef59, 0x05e6a6e5, 0x0e0b1a86, 0x00a0c890, 0x0cfea425, 0x0934e246, 0x0ce83eef, 0x05b3943a, 0x0c7b1843},
				&bigNumber{0x046df945, 0x09b535ca, 0x0c3277fa, 0x060869d2, 0x048ea073, 0x0fa0a4a7, 0x0d2a8ae2, 0x07ed577c, 0x0eb554d3, 0x01d5a467, 0x08e6ecab, 0x04d36ea5, 0x0fd8a04a, 0x0ce771b8, 0x0c15d02d, 0x06287c31},
				&bigNumber{0x0b102830, 0x0500cdf7, 0x0ce151cc, 0x0f53280c, 0x0e2fe85b, 0x0cb91351, 0x045cf048, 0x0c96c4b8, 0x0a785c8c, 0x016ee144, 0x05c8f9d9, 0x0f821375, 0x016d01bf, 0x03f8b538, 0x08e8f714, 0x020ec408},
			},
			&bigNumber{0x0d1da8f1, 0x01578367, 0x08267195, 0x01cefa6b, 0x07c6b5bf, 0x046257ca, 0x0687677a, 0x0143f792, 0x03e23d99, 0x075a29af, 0x07f628e6, 0x08b79ac5, 0x00bb2996, 0x046d5282, 0x0fb6ac94, 0x0a33438a},
		},
		&twPNiels{
			&twNiels{
				&bigNumber{0x018edca2, 0x07debea1, 0x07a2ddca, 0x078c80e6, 0x079d2fb1, 0x049697c0, 0x0b1188b8, 0x0c5d3747, 0x00407854, 0x0534bf0f, 0x0f54f2d7, 0x0cfb7cf7, 0x0bf910ac, 0x030e053a, 0x0098d055, 0x0e558632},
				&bigNumber{0x0c53cf97, 0x0e974b5e, 0x086e98a6, 0x0a7819f4, 0x0585880c, 0x0f3f732e, 0x0811c0fe, 0x0e93ea22, 0x0e207d56, 0x0c03304c, 0x06b8cd79, 0x02d80a84, 0x0b67d408, 0x04b3a48d, 0x0a8785a1, 0x08005d97},
				&bigNumber{0x0c554038, 0x0df94fd5, 0x051a6923, 0x0c2f7058, 0x0805e017, 0x08d350e4, 0x01d6541a, 0x06d4bc92, 0x0259ed7b, 0x0e66e444, 0x0170e624, 0x073eb86c, 0x00322408, 0x0d2ca7ee, 0x0cd50d02, 0x0e653b93},
			},
			&bigNumber{0x0a88db76, 0x06322356, 0x0613a3db, 0x04748332, 0x07219fb7, 0x04bc2c9f, 0x0ffd0653, 0x0e8331c3, 0x01dd9837, 0x0896229e, 0x0b3451e8, 0x0812bd87, 0x00e68050, 0x0af7ae7b, 0x00994348, 0x0eb628be},
		},
		&twPNiels{
			&twNiels{
				&bigNumber{0x00d6b3c0, 0x00081170, 0x00ce21ac, 0x07daa590, 0x07e1cc39, 0x0655c1bf, 0x0fcf5149, 0x03ea006e, 0x030548cc, 0x0cca0c37, 0x00323efd, 0x058f4399, 0x0e9c4564, 0x015f4a70, 0x0d2b801a, 0x00486cb0},
				&bigNumber{0x04ed553e, 0x0586ceb2, 0x06507711, 0x04ef585c, 0x0f2b1778, 0x06240c83, 0x01c4560a, 0x00a27867, 0x065ee476, 0x01c59746, 0x0378b326, 0x0ffdf89a, 0x0fcb3ed4, 0x07542e08, 0x0941ddd9, 0x07e2960a},
				&bigNumber{0x04bf8afa, 0x0506e0a4, 0x077ff487, 0x0603692d, 0x053489ea, 0x0686566a, 0x0ebf2146, 0x0b852e09, 0x0edde792, 0x06cc7210, 0x020e9c1d, 0x02618f3d, 0x08621ed7, 0x079b375a, 0x09a5dbb9, 0x0b166e68},
			},
			&bigNumber{0x04d36152, 0x0bc727e4, 0x0f87de30, 0x07989f99, 0x087b5c7d, 0x00e1e8d2, 0x0d87fc58, 0x0951eb09, 0x03bbdd2d, 0x026c017d, 0x091154b9, 0x0cd2b3f7, 0x0227e8e2, 0x0d03ba6c, 0x08a69793, 0x0b9a3cda},
		},
		&twPNiels{
			&twNiels{
				&bigNumber{0x07868a95, 0x047e0efd, 0x026655ba, 0x046218bb, 0x057ca504, 0x0b9d515a, 0x066c93b3, 0x015770b1, 0x0bb4584c, 0x03d887bd, 0x0a2483ba, 0x08c35546, 0x03c4e826, 0x072e6b5d, 0x0d3d371b, 0x04d9664f},
				&bigNumber{0x04db7356, 0x08c76780, 0x00ef6c05, 0x0aab7cd0, 0x0481e042, 0x03606745, 0x097a2a3e, 0x08852cd5, 0x04a9f685, 0x033f39bb, 0x047353d3, 0x0d5a1c69, 0x03516276, 0x0c254b8f, 0x0a3ad6bd, 0x072decc9},
				&bigNumber{0x04e55f32, 0x0d94cc9a, 0x0af4d143, 0x09528685, 0x0f843ba3, 0x01915ac1, 0x0118eaf8, 0x0519cc24, 0x0fe649c6, 0x09b304b6, 0x0ead9af2, 0x0ff0e903, 0x04a0e812, 0x0694d37c, 0x05ca344b, 0x074fd317},
			},
			&bigNumber{0x021e237f, 0x0e98985b, 0x04dff5df, 0x023ac590, 0x03d852b8, 0x0e056411, 0x0dabc62e, 0x0b8a2280, 0x0d8c6cc9, 0x0c77aa99, 0x0e6c9856, 0x0ac1427e, 0x08707de1, 0x0eb076c5, 0x0e88cec8, 0x03f453a2},
		},
		&twPNiels{
			&twNiels{
				&bigNumber{0x08c37a32, 0x000dd866, 0x0fa71837, 0x0fd325df, 0x02d0d5c8, 0x0e9838e0, 0x0fa8c4d8, 0x05457c16, 0x0d465658, 0x0a174204, 0x0b49176a, 0x0c6926c0, 0x0123e28d, 0x0f569fbf, 0x072bbd4c, 0x04737cc7},
				&bigNumber{0x02f587fd, 0x0ee1a153, 0x0761264a, 0x07427832, 0x0ee5eb88, 0x05c1dccd, 0x07e9615c, 0x06f68ba9, 0x0e41c85f, 0x0ec4c48e, 0x0fa3f8fc, 0x06140e69, 0x0dfd97f2, 0x0a70e8f6, 0x0d2f2da9, 0x0bbd70bf},
				&bigNumber{0x0ec444a5, 0x04096dbc, 0x0a0cea49, 0x02861544, 0x014bfdd7, 0x03957aff, 0x03152338, 0x0e48271f, 0x0a78f948, 0x03b919ad, 0x057bc81e, 0x09e39095, 0x0dd011ba, 0x0f4ecfbe, 0x074d3595, 0x03c31d75},
			},
			&bigNumber{0x05dc0e17, 0x0375ef70, 0x0a8e113a, 0x0c733121, 0x0a3f6947, 0x055545ac, 0x04dae5c2, 0x0b5d48d1, 0x0139488c, 0x0f6a155c, 0x0ff9ba0b, 0x0dc1cdb8, 0x0c8dd21e, 0x085940d8, 0x0d8dda2b, 0x08d3c8e4},
		},
		&twPNiels{
			&twNiels{
				&bigNumber{0x0b28864b, 0x02dbc5e4, 0x017ec0d9, 0x02d228c5, 0x0837c212, 0x076af9f0, 0x0cc2e39a, 0x0a03a657, 0x01715670, 0x06397d13, 0x0108eeda, 0x0d7feab6, 0x01708bf6, 0x0efd7842, 0x0a83a289, 0x0382af8b},
				&bigNumber{0x0b58d2c9, 0x03d4140c, 0x027742bc, 0x094f30ed, 0x039b7954, 0x05268b29, 0x065bba5a, 0x008a9dc5, 0x0e19778b, 0x073e1c2c, 0x01ce3b2c, 0x063f43ba, 0x0d837535, 0x0aa8449c, 0x0b26446c, 0x07eac90f},
				&bigNumber{0x08a2eb6f, 0x034bcc4b, 0x0711b90b, 0x03419b77, 0x0e397292, 0x0e74e501, 0x030e4414, 0x0466a47d, 0x066a5db7, 0x08867270, 0x0340ca5f, 0x02eee9ec, 0x02a4d046, 0x0dba1ad8, 0x0ada2bb1, 0x003c5334},
			},
			&bigNumber{0x0c9d9a60, 0x083cedf7, 0x016f3589, 0x0205ba68, 0x0f626120, 0x0ad30b54, 0x0d48673f, 0x03d897ed, 0x0d421303, 0x0ad1ab33, 0x026b41e0, 0x00918b4d, 0x0d7508dc, 0x0d773c93, 0x001b1570, 0x047de932},
		},
		&twPNiels{
			&twNiels{
				&bigNumber{0x0f37f42c, 0x0c06920a, 0x06a5cc35, 0x0d6bb9fd, 0x0fc4c69f, 0x0bb25ee6, 0x06283092, 0x08912436, 0x0f9391a7, 0x0e041639, 0x0c126f0b, 0x05d87e5e, 0x0c2a51d1, 0x0a3c431a, 0x09823a9e, 0x0577881a},
				&bigNumber{0x0aa15c1f, 0x02281a8a, 0x01e7367d, 0x0b2e134b, 0x0f263908, 0x0fde343c, 0x06eb3cb0, 0x093a40fe, 0x0306b564, 0x008d054f, 0x023730bc, 0x004d4e3f, 0x06c6cfbc, 0x0ab25f6c, 0x0b8ac3f2, 0x022abad4},
				&bigNumber{0x0d5d6b9b, 0x0de60da1, 0x0c3e6171, 0x0c200be8, 0x0af5a396, 0x02188f51, 0x05d830f3, 0x082a55c3, 0x0907de57, 0x0d98f543, 0x06036ad7, 0x04369340, 0x056d2589, 0x003bab9a, 0x045c468c, 0x0fa883aa},
			},
			&bigNumber{0x062cefdd, 0x0653a383, 0x0a408897, 0x0b9c7d56, 0x0564e61c, 0x02968b2a, 0x01cfab58, 0x023e8beb, 0x0819135d, 0x0fcd4ffd, 0x03eb296a, 0x0682a42d, 0x02383e09, 0x0147b62f, 0x0ac692e6, 0x0139b627},
		},
		&twPNiels{
			&twNiels{
				&bigNumber{0x07199fda, 0x00583fe1, 0x04b731bd, 0x080cfd23, 0x0201e480, 0x035b45de, 0x0a627096, 0x048ecd2a, 0x04a24e0c, 0x08c5d452, 0x069b10e3, 0x022d1bd7, 0x06df42f0, 0x087b69f2, 0x057a9fa5, 0x03ca0b07},
				&bigNumber{0x0c2264e2, 0x0a216d66, 0x0b4aa2fa, 0x0a703af9, 0x03cf5b52, 0x0b9de086, 0x053e812a, 0x01b82aec, 0x00ee6056, 0x0414a98f, 0x0d80f676, 0x05b1800b, 0x0d6a7d02, 0x01709f47, 0x02074faa, 0x0bb9ad13},
				&bigNumber{0x05773c5e, 0x0ae1e0bb, 0x058a7cee, 0x053077f5, 0x02743be3, 0x0d2e3a05, 0x0278db5a, 0x011bdbdf, 0x04ca784e, 0x0e9dca6b, 0x0be9bd21, 0x0476254f, 0x0757b68d, 0x0505570d, 0x0d4f201f, 0x0975eafb},
			},
			&bigNumber{0x085a5086, 0x0b61dc0e, 0x0abf7943, 0x0f406ff5, 0x057e301b, 0x0c1a460e, 0x0983063d, 0x01658a83, 0x0baf0680, 0x00bc0ba4, 0x01ff544a, 0x05709b13, 0x0897f07f, 0x0df91c65, 0x097645c1, 0x0a01708f},
		},
	}
	c.Assert(w, DeepEquals, expected)
}

func (s *Ed448Suite) Test_ExtendedToPNiels(c *C) {
	c.Skip("waiting for subfuctions to be implemented")
	p := &twExtendedPoint{
		&bigNumber{0x0cbc0465, 0x0ae92226, 0x091cf2c2, 0x06088c39, 0x06adcba9, 0x02b78bc6, 0x06b0b2aa, 0x0d1b1697, 0x0cf760b0, 0x09cc4ae8, 0x06ff4c95, 0x06f9eaff, 0x03f88cb2, 0x0dfe1f47, 0x015828d5, 0x003f26ff},
		&bigNumber{0x0f73319b, 0x007fa279, 0x0c6ab58a, 0x0224b18c, 0x0bbc2c00, 0x012c8b76, 0x07428968, 0x06047979, 0x0a19d606, 0x0cf6c1c8, 0x06a83f0d, 0x0725b63c, 0x0bc33839, 0x0d9114a0, 0x07eec286, 0x0908447d},
		&bigNumber{0x01e2013d, 0x0287266b, 0x0f434216, 0x08bac041, 0x03321096, 0x096004ef, 0x0c9f384c, 0x0ae584aa, 0x02a4d456, 0x02d20148, 0x0df9a8dd, 0x051a0acf, 0x0b3886b1, 0x00ad3c7f, 0x06f0b032, 0x0e09bf7e},
		&bigNumber{0x0419dfec, 0x0723fb8c, 0x00f6c661, 0x045419d4, 0x00449ef9, 0x09f3003d, 0x0d77bd3d, 0x03aabfcb, 0x05f1ad37, 0x018339b6, 0x0e4963f2, 0x0060cbfc, 0x078b53f3, 0x023457ac, 0x048d759c, 0x02a6c760},
	}

	expected := &twPNiels{
		&twNiels{
			&bigNumber{0x02b72d36, 0x05968053, 0x034dc2c7, 0x0c1c2553, 0x050e6056, 0x0e74ffb0, 0x0091d6bd, 0x08e962e2, 0x0d227555, 0x032a76df, 0x0fa8f278, 0x002bcb3c, 0x07caab87, 0x0f92f559, 0x069699b0, 0x08c91d7e},
			&bigNumber{0x0c2f3600, 0x0b68c4a0, 0x0587a84c, 0x082d3dc6, 0x0269f7a9, 0x03e4173d, 0x0df33c12, 0x031f9010, 0x071136b7, 0x06c30cb1, 0x0da78ba3, 0x0e1fa13b, 0x0fbbc4eb, 0x0b8f33e7, 0x0946eb5c, 0x09476b7c},
			&bigNumber{0x03ba27f6, 0x097f5bcd, 0x0caccee8, 0x01a1e00a, 0x0814dab2, 0x08fb39de, 0x0c3f7d1f, 0x0d773562, 0x0d542a5f, 0x092ec8d9, 0x0fd7c5ba, 0x0d34c4ad, 0x0743ea0c, 0x0a5ebf98, 0x087662be, 0x0dbcdd9f},
		},
		&bigNumber{0x03c4027b, 0x050e4cd6, 0x0e86842c, 0x01758083, 0x0664212d, 0x02c009de, 0x093e7099, 0x05cb0955, 0x0549a8ae, 0x05a40290, 0x0bf351ba, 0x0a34159f, 0x06710d62, 0x015a78ff, 0x0de16064, 0x0c137efc},
	}

	twpn := p.twPNiels()

	c.Assert(twpn.n.a, DeepEquals, expected.n.a)
	c.Assert(twpn.n.b, DeepEquals, expected.n.b)
	c.Assert(twpn.n.c, DeepEquals, expected.n.c)
	c.Assert(twpn.z, DeepEquals, expected.z)
}

func (s *Ed448Suite) Test_AddPNielsToExtended(c *C) {
	p := &twExtendedPoint{
		&bigNumber{
			0x065fe00e, 0x070a5e7e, 0x0bfd74eb, 0x08eee5a2,
			0x01eb1d18, 0x06978a30, 0x0e05687d, 0x0332d83d,
			0x029bfaa5, 0x01d55025, 0x01a04e4d, 0x03387c04,
			0x0bfe9dbe, 0x0052df1b, 0x0adf7ef5, 0x0f7738e9,
		},
		&bigNumber{
			0x0e896a5f, 0x0e551bba, 0x0363107c, 0x0ad5faf3,
			0x00d4bdd0, 0x0025431b, 0x0a96726f, 0x06450cae,
			0x06e879ac, 0x076c4a06, 0x04816d63, 0x072fda21,
			0x02fc227b, 0x0548ed7b, 0x062e39e8, 0x07635ec5,
		},
		&bigNumber{
			0x0a9cb59b, 0x0e8e9c67, 0x08ab4927, 0x039fd088,
			0x05bafdda, 0x0010c3bf, 0x0621a18e, 0x0288d330,
			0x0e39f617, 0x0cf20409, 0x050aa964, 0x037696d0,
			0x0197ecc8, 0x0098a718, 0x039ab7ce, 0x0ccac870,
		},
		&bigNumber{
			0x003a81bc, 0x05aec25f, 0x0352215e, 0x03ef1e1e,
			0x0c216e7b, 0x0ff5ee46, 0x0c5784fd, 0x0f377ed9,
			0x0cd2678a, 0x0583af4e, 0x04da9308, 0x00eeaf60,
			0x043e4dc8, 0x0e77b786, 0x06aab96c, 0x0fe963f3,
		},
	}

	pn := &twPNiels{
		&twNiels{
			&bigNumber{
				0x0dffcb31, 0x0426f81c, 0x00faf45c, 0x06c50175,
				0x0b4d7ba0, 0x04bac22f, 0x0b0d28b6, 0x0b542495,
				0x04e4c5b6, 0x08405b7e, 0x0e2cc773, 0x0ef5cf36,
				0x0015a008, 0x056cdf03, 0x0f0b952e, 0x09ac1df4,
			},
			&bigNumber{
				0x0d598e52, 0x0ea2ac0e, 0x02b8d2d9, 0x068f2906,
				0x005ebdfc, 0x01f67a03, 0x052cc96f, 0x065635dd,
				0x044ca0b5, 0x087e7c65, 0x02b8e2cd, 0x05f826b9,
				0x0bdb6d31, 0x07119d4b, 0x07909767, 0x02c7692f,
			},
			&bigNumber{
				0x0c789450, 0x09efcd8a, 0x0917d143, 0x0eddc9f9,
				0x0c135cd5, 0x034d7893, 0x022ea365, 0x0999b447,
				0x06951972, 0x0925c008, 0x0fe1cd6c, 0x0f55a3fd,
				0x0081cfcf, 0x000a9eb6, 0x02bacafd, 0x0b8a3daf,
			},
		},
		&bigNumber{
			0x0154741b, 0x084a4939, 0x071479d1, 0x01c5a7d6,
			0x0e77458b, 0x0b209c54, 0x05f90c8d, 0x01e562a7,
			0x07f0fe64, 0x071b3d33, 0x042b13a2, 0x023bca1b,
			0x0f4a5e74, 0x0b19b036, 0x03b187a0, 0x008b6799,
		},
	}

	p.add(pn, false)

	expected := &twExtendedPoint{
		&bigNumber{
			0x0229a19b, 0x07640779, 0x0cd5c825, 0x00a542fe,
			0x0bb2362b, 0x0261d1af, 0x015db410, 0x0ca63130,
			0x00e9a7e4, 0x00a06584, 0x01029e89, 0x025cdd0e,
			0x0a11b601, 0x0c592ed0, 0x0b1854d2, 0x0c374936,
		},
		&bigNumber{
			0x045e9fe6, 0x09017d4c, 0x077f92d8, 0x079e77d5,
			0x081ea883, 0x0ad757f3, 0x088b186c, 0x03d4ae18,
			0x01f422c7, 0x048c7fc1, 0x052d50d6, 0x0b310d2d,
			0x0f98c559, 0x0ce135a0, 0x02cdff0e, 0x06c117fd,
		},
		&bigNumber{
			0x0ac69c60, 0x0fa58fee, 0x0e0b93bd, 0x0c5887dc,
			0x0dc231b4, 0x08d5bd5a, 0x0cbaa909, 0x06cda8c7,
			0x064c5ba3, 0x0b7a44bc, 0x0e076348, 0x008976ff,
			0x04ae0ed6, 0x03735252, 0x04c09fca, 0x075120c2,
		},
		&bigNumber{
			0x0f5f78ac, 0x0f2f972b, 0x0c7b74b5, 0x084ae55c,
			0x03fe0c33, 0x02441814, 0x0b335997, 0x052cd9d7,
			0x00af1474, 0x0d285ac9, 0x0ea87209, 0x087cd8a3,
			0x03d9718e, 0x0b8f00f3, 0x003d8c42, 0x0345b1a8,
		},
	}

	c.Assert(p.x, DeepEquals, expected.x)
	c.Assert(p.y, DeepEquals, expected.y)
	c.Assert(p.z, DeepEquals, expected.z)
	c.Assert(p.t, DeepEquals, expected.t)
}

func (s *Ed448Suite) Test_PointDoubleScalarmul_WithLargeNumbers(c *C) {
	c.Skip("waiting for subfuctions to be implemented")
	p1 := &twExtendedPoint{
		&bigNumber{
			0x0e0fbf9e, 0x0ba1bcd7, 0x01cc6d39, 0x053b56e8,
			0x0635d142, 0x0383307a, 0x0f8a159b, 0x097fd2cf,
			0x0fa310f6, 0x05522bde, 0x0b981703, 0x0b095b1e,
			0x042d4780, 0x05ae11df, 0x0934fe80, 0x0dc6474d,
		},
		&bigNumber{
			0x02c1149c, 0x0e72febf, 0x05259893, 0x0723e184,
			0x0f7232ff, 0x019a5600, 0x05581d2c, 0x07331444,
			0x04e0124a, 0x09c3c5e5, 0x0945536e, 0x0b786a20,
			0x0f75623f, 0x00ba30e8, 0x0cc589a3, 0x04a2eea8,
		},
		&bigNumber{
			0x02406c71, 0x0b2fdb67, 0x02591aa2, 0x085fc24e,
			0x0dc50d09, 0x08692c5b, 0x0ba917d7, 0x0aefea74,
			0x037d0084, 0x04d5defa, 0x08bbe7ad, 0x050da977,
			0x08adf827, 0x05425cdd, 0x037d816d, 0x0d59cd0a,
		},
		&bigNumber{
			0x0baf8c30, 0x06686ad3, 0x0c149bac, 0x0f57f68d,
			0x05cd321a, 0x02ff8d60, 0x09dcc4bd, 0x0f731ec2,
			0x0cd7ea75, 0x0be970e4, 0x043d30e0, 0x0dd64b9b,
			0x04f78bf1, 0x0d1fde20, 0x05c88e97, 0x026ce314,
		},
	}

	s1 := Scalar{
		0x9a1044c6, 0x92f78393, 0x68cea2bc, 0x5f23f942,
		0xd4384e9e, 0x76969060, 0x4d82f8cc, 0xb8016c73,
		0x1db9b587, 0x061aca05, 0x9f0333f5, 0x5a2a7f4a,
		0x216a1e70, 0x1d22f534,
	}

	p2 := &twExtendedPoint{
		&bigNumber{
			0x06172a44, 0x0731d576, 0x0da247e0, 0x0d9fd318,
			0x072d1c77, 0x073e77aa, 0x09a004b5, 0x012507b9,
			0x09a684c3, 0x08b559f8, 0x0d445c85, 0x07941c89,
			0x0c942cd4, 0x02bcfe3e, 0x022ccaaa, 0x0be3a6b3,
		},
		&bigNumber{
			0x03294fb1, 0x0e4336b5, 0x0fe125d6, 0x08c09f34,
			0x0f04e3ce, 0x0eac940d, 0x09c38a23, 0x0a2ec035,
			0x06545488, 0x0355e18f, 0x0522a0ec, 0x0ce0fd60,
			0x0bd3a6ce, 0x03fe9d85, 0x06e5c4f3, 0x018cf1e5,
		},
		&bigNumber{
			0x0e957107, 0x0f672aa2, 0x049b0276, 0x07a7ecf2,
			0x0e9a1c69, 0x04067d01, 0x03f2ddee, 0x0ffebccb,
			0x0d58b6cf, 0x0d95fb9c, 0x077d5935, 0x078ddbc3,
			0x085093f2, 0x03015d2f, 0x019d8e0a, 0x0388a2ac,
		},
		&bigNumber{
			0x0bf26ccb, 0x0b930dcd, 0x0e207a77, 0x0d8fdde5,
			0x04e2452b, 0x099e9922, 0x0ec0b62c, 0x04f9d73b,
			0x03811a2a, 0x0871aefb, 0x00f5e028, 0x0b6aa04c,
			0x0226cb55, 0x0b6e4ee0, 0x0f3eba42, 0x04409402,
		},
	}

	s2 := Scalar{
		0x3aad8a3d, 0x7cbae122, 0xed340da1, 0x1e37d7eb,
		0x2a2e914d, 0xcae48b24, 0x9e50875c, 0xc5b5e48b,
		0x89d9f0e4, 0xdf9d2321, 0x8775f116, 0xd1868de2,
		0x139f9896, 0x0ddda899,
	}

	expected := &twExtendedPoint{
		&bigNumber{
			0x0fd40c86, 0x09c9d284, 0x0b3c54a8, 0x01d164e6,
			0x0fb54a59, 0x05c4eddf, 0x089f31e2, 0x069f1376,
			0x08986bdf, 0x0a6960c4, 0x04394fb6, 0x02461954,
			0x0ebc7778, 0x03893eb9, 0x0dfe52f7, 0x0cb5ac4c,
		},
		&bigNumber{
			0x08e7c39b, 0x0cd0c757, 0x09afb3e2, 0x03ae4eb4,
			0x0a24045a, 0x0b710cf9, 0x06b2b8d0, 0x0ed5287b,
			0x0ba20243, 0x03154883, 0x0fff59ad, 0x05c300b6,
			0x0614dfab, 0x0b22e212, 0x0b0e2874, 0x07e5eb59,
		},
		&bigNumber{
			0x08fe1ce7, 0x0d6422e3, 0x0ee2fd89, 0x0bee7e8c,
			0x0c5140b3, 0x07fc0df8, 0x0110af8b, 0x04713185,
			0x0d410281, 0x054ba615, 0x05ac365c, 0x0e68cd22,
			0x093f4cce, 0x08fe59ac, 0x04ce12b5, 0x0ea3228d,
		},
		&bigNumber{
			0x0a6753d7, 0x033ccb22, 0x0a5ce2fc, 0x0cb4632c,
			0x08926d77, 0x057b0090, 0x0f360581, 0x03037a76,
			0x0bcbbecd, 0x06009fb0, 0x00a93751, 0x06598ba6,
			0x07e59b3b, 0x0e41fa30, 0x0d6178a6, 0x02d5bcc8,
		},
	}

	out := pointDoubleScalarMul(p1, s1, p2, s2)
	c.Assert(out, DeepEquals, expected)
}
