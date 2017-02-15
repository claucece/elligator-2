package ed448

// will look something like this in C

//void
//find_generator_elligator()
//{
//	decaf_448_point_t p;
//	unsigned char seed[DECAF_448_SER_BYTES+8];
//	const char *magic = "decaf_448_g2";
//	unsigned char encoded_point[DECAF_448_SER_BYTES],
//	encoded_base[DECAF_448_SER_BYTES],
//	hashed_base[DECAF_448_SER_BYTES];
//
//	decaf_448_point_encode(encoded_base, decaf_448_point_base);
//
//	//hash the base point
//	keccak_sponge_t sponge;
//	shake256_init(sponge);
//	shake256_update(sponge, (const unsigned char *)encoded_base, sizeof(encoded_base));
//	shake256_final(sponge, hashed_base, sizeof(hashed_base));
//
//	unsigned char hint;
//
//	//the seed
//	keccak_sponge_t magic_sponge;
//	shake256_init(magic_sponge);
//	shake256_update(magic_sponge, (const unsigned char *)magic, strlen(magic));
//	shake256_final(magic_sponge, seed, sizeof(seed));
//	shake256_destroy(magic_sponge);
//
//	// random_bytes(seed,sizeof(seed))
//	shake256_init(sponge);
//	shake256_update(sponge, (const unsigned char *)magic, strlen(magic));
//	shake256_update(sponge, (const unsigned char *)hashed_base, sizeof(hashed_base));
//	shake256_update(sponge, (const unsigned char *)seed, sizeof(seed));
//	shake256_final(sponge, encoded_point, sizeof(encoded_point));
//	shake256_destroy(sponge);
//
//	//elligator
//	hint = decaf_448_point_from_hash_uniform(p, encoded_point);
//	printf("hint n:%#x\n", hint);
//
//	print_point(p,"g2 found");
//}

// will give

//hint n:0x24

//g2 found :
//0x00ac97f43cf14237,
//0x00dc98db8a9543bc,
//0x007874a17bcca6a6,
//0x00fffa76321af78f,
//0x0074f2a89cf2ac0b,
//0x00356a31ef89f88d,
//0x000c010839f61e5a,
//0x000bf3b5cc84b7a5,
//
//0x006b775bc0c9a64c,
//0x00ee0c3e126148bb,
//0x004fad09b303aa98,
//0x003008555efaf59d,
//0x0023bc0fa72a0bf6,
//0x00f0f61f9c52ee5b,
//0x00b8b7f385cf8d7f,
//0x006a9849a18a4398,
//
//0x00198c24c14e2fce,
//0x00080f748b74b290,
//0x008ab2f53fb60b6e,
//0x0069791886c32b60,
//0x0087ecac7e87a66d,
//0x0035faebff354ebd,
//0x00c96f513e30d07f,
//0x00da28e58fab82ed,
//
//0x005c67537702239a,
//0x00fae388ece76a54,
//0x006b5fe3d34bcae9,
//0x009cac77dd3c37ae,
//0x00a02246f761a657,
//0x009448b046490757,
//0x00e0bd3d45281bbe,
//0x007c655f9abc5ecb,

var (
	g2 = &twExtendedPoint{
		&bigNumber{
			0x0cf14237, 0x0ac97f43,
			0x0a9543bc, 0x0dc98db8,
			0x0bcca6a6, 0x07874a17,
			0x021af78f, 0x0fffa763,
			0x0cf2ac0b, 0x074f2a89,
			0x0f89f88d, 0x0356a31e,
			0x09f61e5a, 0x00c01083,
			0x0c84b7a5, 0x00bf3b5c,
		},

		&bigNumber{
			0x00c9a64c, 0x06b775bc,
			0x026148bb, 0x0ee0c3e1,
			0x0303aa98, 0x04fad09b,
			0x0efaf59d, 0x03008555,
			0x072a0bf6, 0x023bc0fa,
			0x0c52ee5b, 0x0f0f61f9,
			0x05cf8d7f, 0x0b8b7f38,
			0x018a4398, 0x06a9849a,
		},

		&bigNumber{
			0x014e2fce, 0x0198c24c,
			0x0b74b290, 0x0080f748,
			0x0fb60b6e, 0x08ab2f53,
			0x06c32b60, 0x06979188,
			0x0e87a66d, 0x087ecac7,
			0x0f354ebd, 0x035faebf,
			0x0e30d07f, 0x0c96f513,
			0x0fab82ed, 0x0da28e58,
		},

		&bigNumber{
			0x0702239a, 0x05c67537,
			0x0ce76a54, 0x0fae388e,
			0x034bcae9, 0x06b5fe3d,
			0x0d3c37ae, 0x09cac77d,
			0x0761a657, 0x0a02246f,
			0x06490757, 0x09448b04,
			0x05281bbe, 0x0e0bd3d4,
			0x0abc5ecb, 0x07c655f9,
		},
	}
)
