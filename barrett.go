package ed448

type barrettPrime struct {
	wordsInP word
	pShift   word
	lowWords []word
}

var lowWords = []word{
	0x54a7bb0d,
	0xdc873d6d,
	0x723a70aa,
	0xde933d8d,
	0x5129c96f,
	0x3bb124b6,
	0x8335dc16,
}

func convertToSignedWindowForm(out []word, scalar []word, preparedData []word) {
	mask := word(dword(-(scalar[0] & 1)) & lmask)

	carry := addExtPacked(out, scalar, preparedData[:scalarWords], word(^mask))
	carry += addExtPacked(out, out, preparedData[scalarWords:], word(mask))

	for i := 0; i < scalarWords-1; i++ {
		out[i] >>= 1
		out[i] |= out[i+1] << (wordBits - 1)
	}

	out[scalarWords-1] >>= 1
	out[scalarWords-1] |= carry << (wordBits - 1)
}

func addExtPacked(dst, x, y []word, mask word) word {
	carry := sdword(0)
	for i := 0; i < len(y); i++ {
		carry += sdword(x[i]) + sdword(y[i]&mask)
		dst[i] = word(carry)
		carry >>= wordBits
	}

	for i := len(y); i < len(x); i++ {
		carry += sdword(x[i])
		dst[i] = word(carry)
		carry >>= wordBits
	}

	return word(carry)
}

func subExtPacked(dst, x, y []word, mask word) word {
	carry := sdword(0x00)
	for i := 0; i < len(y); i++ {
		carry += sdword(x[i]) - (sdword(y[i]) & sdword(mask))
		dst[i] = word(carry)
		carry >>= wordBits
	}

	for i := len(y); i < len(x); i++ {
		carry += sdword(x[i])
		dst[i] = word(carry)
		carry >>= wordBits
	}

	return word(carry)
}

func barrettDeserializeAndReduce(dst []word, serial []byte, p *barrettPrime) {
	wordLen := wordBits / 8
	size := (len(serial) + wordLen - 1) / wordLen
	if size < int(p.wordsInP) {
		size = int(p.wordsInP)
	}

	tmp := make([]word, size)
	bytesToWords(tmp[:], serial[:])
	barrettReduce(tmp[:], 0, p)

	for i := word(0); i < p.wordsInP; i++ {
		dst[i] = tmp[i]
	}
}

func barrettReduce(dst []word, carry word, p *barrettPrime) {
	for wordsLeft := word(len(dst)); wordsLeft >= p.wordsInP; wordsLeft-- {
		//XXX PERF unroll
		for repeat := 0; repeat < 2; repeat++ {
			mand := dst[wordsLeft-1] >> p.pShift
			dst[wordsLeft-1] &= (word(1) << p.pShift) - 1

			if p.pShift != 0 && repeat == 0 {
				if wordsLeft < word(len(dst)) {
					mand |= dst[wordsLeft] << (wordBits - p.pShift)
					dst[wordsLeft] = 0
				} else {
					mand |= carry << (wordBits - p.pShift)
				}
			}

			carry = widemac(
				dst[wordsLeft-p.wordsInP:wordsLeft],
				p.lowWords, mand, 0)
		}
	}

	cout := addExtPacked(dst, dst[:p.wordsInP], p.lowWords, lmask)

	if p.pShift != 0 {
		cout = (cout << (wordBits - p.pShift)) + (dst[p.wordsInP-1] >> p.pShift)
		dst[p.wordsInP-1] &= word(1)<<p.pShift - 1
	}

	// mask = carry-1: if no carry then do sub, otherwise don't
	subExtPacked(dst, dst[:p.wordsInP], p.lowWords, cout-1)
}

//XXX Is this the same as mulAddVWW_g() ?
func widemac(accum []word, mier []word, mand, carry word) word {
	for i := 0; i < len(mier); i++ {
		product := dword(mand) * dword(mier[i])
		product += dword(accum[i])
		product += dword(carry)

		accum[i] = word(product)
		carry = word(product >> wordBits)
	}

	for i := len(mier); i < len(accum); i++ {
		sum := dword(carry) + dword(accum[i])
		accum[i] = word(sum)
		carry = word(sum >> wordBits)
	}

	return carry
}

// Deserializes an array of bytes (little-endian) into an array of words
func bytesToWords(dst []word, src []byte) {
	wordBytes := uint(wordBits / 8)
	srcLen := uint(len(src))

	dstLen := uint((srcLen + wordBytes - 1) / wordBytes)
	if dstLen < uint(len(dst)) {
		panic("wrong dst size")
	}

	for i := uint(0); i*wordBytes < srcLen; i++ {
		out := word(0)
		for j := uint(0); j < wordBytes && wordBytes*i+j < srcLen; j++ {
			out |= word(src[wordBytes*i+j]) << (8 * j)
		}

		dst[i] = out
	}
}

func scheduleScalarForCombs(schedule []word, scalar decafScalar) {
	table := baseTable
	tmp := make([]word, len(schedule))
	copy(tmp, scalar[:])

	tmp[len(tmp)-1] &= (word(1) << (446 % wordBits)) - 1

	convertToSignedWindowForm(schedule, tmp, table.adjustments[:])
}
