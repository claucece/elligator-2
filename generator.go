package ed448

import (
	"fmt"

	"golang.org/x/crypto/sha3"
)

// this is not 512 as MH said. is it better to use elligator?
// std::string seed = [a fixed 512-bit constant which you chose at random];
// Point::from_hash(SHAKE<256>::Hash(std::string(Point::base()) + seed, Point::HASH_BYTES*2))
func findGenerator() *twExtendedPoint {

	var encodedBase [56]byte
	var encodedPoint [56]byte
	var seed [64]byte       //sure?
	var hashedBase [56]byte //sure?

	BasePoint.decafEncode(encodedBase[:])

	h1 := sha3.NewShake256()
	h1.Write(encodedBase[:])
	h1.Read(hashedBase[:])

	n := 0

	p := &twExtendedPoint{
		&bigNumber{},
		&bigNumber{},
		&bigNumber{},
		&bigNumber{},
	}

	valid := word(0x00)

	for valid == word(0x00) {
		h2 := sha3.NewShake256()
		for i := 0; i < n; i++ {
			h2.Write([]byte("g2")) //why?
		}
		h2.Read(seed[:])

		h3 := sha3.NewShake256()
		h3.Write([]byte("g2"))
		h3.Write(hashedBase[:]) //sure?
		h3.Write(seed[:])       //sure?
		h3.Read(encodedPoint[:])

		valid = decafDecode(p, encodedPoint, false) // not allowing identity?
		n++
		fmt.Printf("trying %d \n", n)
	}

	fmt.Printf("found %d", p)
	return p
}

func findGeneratorElligator() (*twExtendedPoint, word) {

	var encodedBase [56]byte
	var encodedPoint [56]byte
	var seed [64]byte       //sure?
	var hashedBase [56]byte //sure?

	BasePoint.decafEncode(encodedBase[:])

	h1 := sha3.NewShake256()
	h1.Write(encodedBase[:])
	h1.Read(hashedBase[:])

	h2 := sha3.NewShake256()
	h2.Write([]byte("decaf_448_g2")) //why?
	h2.Read(seed[:])

	h3 := sha3.NewShake256()
	h3.Write([]byte("decaf_448_g2"))
	h3.Write(hashedBase[:]) //sure?
	h3.Write(seed[:])       //sure?
	h3.Read(encodedPoint[:])

	p, hint := decafUniformFromHashToCurve(encodedPoint[:]) // not allowing identity?
	fmt.Printf("hint %#x \n", hint)

	fmt.Printf("found %d", p)
	return p, hint
}
