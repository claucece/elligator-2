package ed448

// Curve is the interface that wraps the basic curve methods.
//XXX It would be better with the use of privateKey and publicKey types.
type Curve interface {
}

type curveT struct{}

var (
	curve = &curveT{}
)

// NewCurve returns a Curve.
func NewCurve() Curve {
	return curve
}
