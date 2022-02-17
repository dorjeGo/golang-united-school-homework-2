package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type intCustomType int

const (
	SidesTriangle intCustomType = 3
	SidesSquare   intCustomType = 4
	SidesCircle   intCustomType = 0
)

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	var result float64 = 0

	switch sidesNum {
	case SidesTriangle:
		result = math.Sqrt(3) * 2 * sideLen / 4
	case SidesSquare:
		result = sideLen * sideLen
	case SidesCircle:
		result = math.Pi * 2 * sideLen
	}
	return result
}
