package ed448

//void
//find_generator()
//{
//    decaf_448_point_t p;
//    unsigned char seed[decaf_448_ser_bytes+8];
//    const char *magic = "g2";
//    unsigned char encoded_point[decaf_448_ser_bytes],
//                encoded_base[decaf_448_ser_bytes],
//                hashed_base[decaf_448_ser_bytes];
//
//    decaf_448_point_encode(encoded_base, decaf_448_point_base);
//    keccak_sponge_t sponge;
//    shake256_init(sponge);
//    shake256_update(sponge, (const unsigned char *)encoded_base, sizeof(encoded_base));
//    shake256_final(sponge, hashed_base, sizeof(hashed_base));
//    decaf_bool_t valid = decaf_false;
//    int n=0;
//    do {
//        keccak_sponge_t magic_sponge;
//        shake256_init(magic_sponge);
//        int i;
//        for (i=0; i<n; i++){
//            shake256_update(magic_sponge, (const unsigned char *)magic, strlen(magic));
//        }
//        shake256_final(magic_sponge, seed, sizeof(seed));
//        shake256_destroy(magic_sponge);
//        /* random_bytes(seed,sizeof(seed)); */
//        shake256_init(sponge);
//        shake256_update(sponge, (const unsigned char *)magic, strlen(magic));
//        shake256_update(sponge, (const unsigned char *)hashed_base, sizeof(hashed_base));
//        shake256_update(sponge, (const unsigned char *)seed, sizeof(seed));
//        shake256_final(sponge, encoded_point, sizeof(encoded_point));
//        valid = decaf_448_point_decode(p, encoded_point, decaf_false);
//        n++;
//        printf("trial n:%d\n",n);
//    } while (!valid);
//    shake256_destroy(sponge);
//    print_point(p,"g2 found");
//}
