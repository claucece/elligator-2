package ed448

import "fmt"

type affinePoint struct {
	x, y *bigNumber
}

// for debuggin
func printPoint(q *affinePoint) {
	fmt.Println()
	fmt.Println("the x")
	for i := 0; i < 16; i++ {
		fmt.Printf("0x%08x, \n", q.x[i])
	}

	fmt.Println()
	fmt.Println("the y")
	for i := 0; i < 16; i++ {
		fmt.Printf("0x%08x, \n", q.y[i])
	}
}

func printBig(n *bigNumber) {
	for i := 0; i < 16; i++ {
		fmt.Printf("0x%08x, \n", n[i])
	}
}

func (p *twExtendedPoint) toAffine() *affinePoint {
	identity := &twExtendedPoint{
		&bigNumber{0x00},
		&bigNumber{0x01},
		&bigNumber{0x01},
		&bigNumber{0x00},
	}

	out := &affinePoint{
		&bigNumber{},
		&bigNumber{},
	}

	if (p.equals(identity) == word(0xfffffff) || p.z.equals(&bigNumber{0x00})) {
		return out
	}

	tmp1, tmp2 := &bigNumber{}, &bigNumber{}
	tmp0 := invert(p.z)
	tmp1.square(tmp0)

	out.x.mul(p.x, tmp1).strongReduce()
	tmp2.mul(tmp1, tmp0)
	out.y.mul(p.y, tmp2).strongReduce()

	return out
}

func serializeAffine(dst []byte, n *bigNumber) {
	src := n.copy()
	src.strongReduce()

	for i := 0; i < 8; i++ {
		l := dword(src[2*i]) + dword(src[2*i+1])<<28
		for j := 0; j < 7; j++ {
			dst[7*i+j] = byte(l)
			l >>= 8
		}
	}
}
