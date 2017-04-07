package ed448

import (
	"io"

	"golang.org/x/crypto/sha3"
)

type publicKey []byte

type privateKey []byte

func generateKeys(rand io.Reader) (publicKey, privateKey, error) {
	priv := make([]byte, 56)
	pub := make([]byte, 56)
	_, err := io.ReadFull(rand, priv[:56])
	if err != nil {
		return nil, nil, err
	}

	digest := sha3.Sum512(priv[:56])
	// XXX: change the numbers
	// XXX: test this
	digest[55] &= 255
	digest[0] &= 127
	digest[0] |= 64

	sc := &decafScalar{}
	barrettDeserializeAndReduce(sc[:], digest[:], &curvePrimeOrder)
	b := multiplyByBase(*sc)
	ok := b.OnCurve()
	if !ok {
		return nil, nil, err
	}

	bb := untwist(b).serializeExtensible()
	serializeBytes(pub, bb)
	// another validity here
	return priv, pub, nil
}

func generateKeys2(rand io.Reader) (publicKey, privateKey, error) {
	priv := make([]byte, 56)
	pub := make([]byte, 56)
	_, err := io.ReadFull(rand, priv[:56])
	if err != nil {
		return nil, nil, err
	}

	digest := sha3.Sum512(priv[:56])

	secretKey := decafScalar{}

	barrettDeserializeAndReduce(secretKey[:], digest[:], &curvePrimeOrder)
	secretKey.serialize(priv)

	publicKey := multiplyByBase(secretKey)
	ok := publicKey.OnCurve()
	if !ok {
		return nil, nil, err
	}

	serializedPublicKey := publicKey.untwistAndDoubleAndSerialize()
	serialize(pub, serializedPublicKey)

	return priv, pub, nil
}

func serializeBytes(dst []byte, n *bigNumber) {
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
