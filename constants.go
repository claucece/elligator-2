package ed448

const (
	symKeyBytes  = 32
	pubKeyBytes  = fieldBytes
	privKeyBytes = 2*fieldBytes + symKeyBytes // 144

	signatureBytes = 2 * fieldBytes // 112

)
