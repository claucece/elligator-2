package ed448

func multiplyByBase(scalar decafScalar) *twExtensible {
	out := &twExtensible{
		new(bigNumber),
		new(bigNumber),
		new(bigNumber),
		new(bigNumber),
		new(bigNumber),
	}

	n := combNumber
	t := combTeeth
	s := combSpacing

	schedule := make([]word, scalarWords)
	scheduleScalarForCombs(schedule, scalar)

	var ni *twNiels

	for i := uint(0); i < s; i++ {
		if i != 0 {
			out.double()
		}

		for j := uint(0); j < n; j++ {
			tab := word(0)

			for k := uint(0); k < t; k++ {
				bit := (s - 1 - i) + k*s + j*(s*t)
				if bit < scalarWords*wordBits {
					tab |= (schedule[bit/wordBits] >> (bit % wordBits) & 1) << k
				}
			}

			invert := word(tab>>(t-1)) - 1
			tab ^= invert
			tab &= (1 << (t - 1)) - 1

			ni = baseTable.lookup(j, t, uint(tab))
			ni.conditionalNegate(invert)

			if i != 0 || j != 0 {
				out.addTwNiels(ni)
			} else {
				convertTwNielsToTwExtensible(out, ni)
			}
		}
	}

	return out
}
