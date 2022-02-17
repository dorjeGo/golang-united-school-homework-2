package square

import "math"

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
		result = math.Sqrt(3) * sideLen * sideLen / 4
	case SidesSquare:
		result = sideLen * sideLen
	case SidesCircle:
		result = math.Pi * sideLen * sideLen
	}

	return result
}
