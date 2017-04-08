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
	digest[55] &= 0
	digest[0] &= 0
	digest[0] |= 1

	sc := &decafScalar{}
	barrettDeserializeAndReduce(sc[:], digest[:], &curvePrimeOrder)
	p := multiplyByBase(*sc)
	ok := p.OnCurve()
	if !ok {
		return nil, nil, err
	}

	pubBytes := untwist(p).serializeExtensible()
	serializeBytes(pub, pubBytes)
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

// On C

// XXX: change the types
//
//static void inline
//random_bytes_long_term(void * const buf, const size_t size) {
//#ifndef FAST_RANDOM
//  gcry_randomize(buf, size, GCRY_VERY_STRONG_RANDOM);
//#else
//  gcry_randomize(buf, size, GCRY_STRONG_RANDOM);
//#endif
//}
//
//static void inline
//random_bytes_strong(void * const buf, const size_t size) {
//  gcry_randomize(buf, size, GCRY_STRONG_RANDOM);
//}
//
//typedef unsigned char symmetric_key_t[32];
//
//int
//derive_private_key (
//    struct goldilocks_private_key_t *privkey,
//    const unsigned char proto[GOLDI_SYMKEY_BYTES]
//) {
//
//    const char *magic = "derive_private_key";
//    uint8_t encoded_scalar[DECAF_448_SCALAR_OVERKILL_BYTES];
//    decaf_448_point_t pub;
//
//    symmetric_key_t proto;
//    random_bytes_strong(proto,sizeof(proto));
//
//    // take this from decaf brach or similar
//    keccak_sponge_t sponge;
//    shake256_init(sponge);
//    shake256_update(sponge, proto, sizeof(symmetric_key_t));
//
//    // not needed part
//    shake256_update(sponge, (const unsigned char *)magic, strlen(magic));
//
//    shake256_final(sponge, encoded_scalar, sizeof(encoded_scalar));
//    shake256_destroy(sponge);
//
//    // is this how is done in C?
//    encoded_scalar[55] &= 0;
//    encoded_scalar[0] |= 1;
//    encoded_scalar[0] &= 0;
//
//    barrett_deserialize_and_reduce(sk, encoded_scalar, sizeof(encoded_scalar), &curve_prime_order);
//    barrett_serialize(privkey->opaque, sk, GOLDI_FIELD_BYTES);
//
//    tw_extensible_a_t exta;
//    scalarmul_fixed_base(exta, sk, 446, &goldilocks_global.fixed_base);
//
//    // do a check if not, return nil
//    mask_t ok;
//    ok = validate_tw_extensible(exta)
//
//    tw_extensible_a_t tmp;
//    twist_and_double(tmp, exta);
//
//    field_a_t pk;
//    serialize_extensible(pk, tmp);
//
//    field_serialize(&privkey->opaque[446], pk);
//
//    return GOLDI_EOK;
//}
//
//
//static void inline
//random_bytes_long_term(void * const buf, const size_t size) {
//#ifndef FAST_RANDOM
//  gcry_randomize(buf, size, GCRY_VERY_STRONG_RANDOM);
//#else
//  gcry_randomize(buf, size, GCRY_STRONG_RANDOM);
//#endif
//}
//
//static void inline
//random_bytes_strong(void * const buf, const size_t size) {
//  gcry_randomize(buf, size, GCRY_STRONG_RANDOM);
//}
//
//
//struct key_pair_t {
//    uint8_t pub[56];
//    uint8_t priv[56];
//};
//
//typedef unsigned char random_data_t[DECAF_448_SYMMETRIC_KEY_BYTES];
//static const int CORRUPT = 44801;
//static const int OK = 0;
//static const mask_t MASK_FAILURE = 0, MASK_SUCCESS = -(mask_t)1
//
//int
//derive_private_key (
//    struct key_pair_t *pair,
//) {
//
//    random_data_t proto;
//    random_bytes_strong(proto,sizeof(random_data_t));
//
//    unsigned char skb[64];
//    word_t sk[14];
//    const char *magic = "derive_private_key";
//
//    sha512_ctx_a_t ctx;
//    tw_extensible_a_t exta;
//    field_a_t pk;
//    mask_t ok;
//
//    sha512_init(ctx);
//    sha512_update(ctx, (const unsigned char *)magic, strlen(magic));
//    sha512_update(ctx, proto, sizeof(random_data_t));
//    field_hash_final(ctx, (unsigned char *)skb);
//
//    barrett_deserialize_and_reduce(sk, skb, sizeof(skb), &curve_prime_order);
//    barrett_serialize(pair->priv, sk, 56);
//
//    scalarmul_fixed_base(exta, sk, GOLDI_SCALAR_BITS, &goldilocks_global.fixed_base);
//    succ = validate_tw_extensible(exta);
//    if (succ != MASK_SUCCESS) { // how in C?
//        return CORRUPT;
//    }
//
//    untwist_and_double_and_serialize(pk, exta);
//    field_serialize(&pair->pub[56], pk);
//
//    return OK;
//}
//
