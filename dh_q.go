package ed448

import (
	"fmt"
	"math/big"
)

var (
	p         *big.Int // prime field, defined in RFC3526 as Diffie-Hellman Group 5
	pMinusTwo *big.Int
	q         *big.Int // prime order
	g1        *big.Int // group generator
)

func sub(l, r *big.Int) *big.Int {
	return new(big.Int).Sub(l, r)
}

func div(l, r *big.Int) *big.Int {
	return new(big.Int).Div(l, r)
}

func getQ(p *big.Int) *big.Int {
	pMinusOne := sub(p, big.NewInt(1))
	q := div(pMinusOne, big.NewInt(2))

	fmt.Println("the q")
	fmt.Printf("%#x", q)
	return q
}

func lte(l, r *big.Int) bool {
	return l.Cmp(r) != 1
}

func gte(l, r *big.Int) bool {
	return l.Cmp(r) != -1
}

func isGroupElement(p, n *big.Int) bool {
	pMinusTwo = sub(p, big.NewInt(2))
	g1 = big.NewInt(2)
	return gte(n, g1) && lte(n, pMinusTwo)
}
