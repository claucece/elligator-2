package ed448

import (
	"encoding/hex"
	"fmt"

	. "gopkg.in/check.v1"
)

func (s *Ed448Suite) Test_IsValidAffine(c *C) {
	exp := &affinePoint{
		&bigNumber{
			0x0a6862a1, 0x0b9509e3,
			0x0f633a09, 0x01bbe8fd,
			0x0055bfe7, 0x04b7a267,
			0x098fec7a, 0x02b43bdb,
			0x038728f3, 0x0e50a54c,
			0x06da2f47, 0x0b1844b2,
			0x03e1ddfe, 0x03f84a5f,
			0x0517a1cc, 0x0bc8e0e4,
		},

		&bigNumber{
			0x0b002bae, 0x046b63b0,
			0x00e0f577, 0x093028d2,
			0x04ae3673, 0x0cb031e4,
			0x0a1b1455, 0x0ef16821,
			0x0505815a, 0x0c83cd8d,
			0x0753d9cc, 0x06691155,
			0x0cfa1242, 0x0bcee146,
			0x03335bee, 0x0dfe21fd,
		},
	}

	affg2 := g2.toAffine()
	c.Assert(affg2, DeepEquals, exp)

	var dst [56]byte
	var dst1 [56]byte

	serializeAffine(dst[:], affg2.x)
	serializeAffine(dst1[:], affg2.y)

	dst2 := make([]byte, hex.EncodedLen(len(dst)))
	hex.Encode(dst2, dst[:])

	fmt.Println(" \n")
	fmt.Printf("0x%s\n", dst2)

	dst3 := make([]byte, hex.EncodedLen(len(dst1)))
	hex.Encode(dst3, dst1[:])

	fmt.Printf("0x%s\n", dst3)
}
