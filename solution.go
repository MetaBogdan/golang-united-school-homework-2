package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import (
	"math"
)

type SidesType int

const (
	SidesTriangle = 3
	SidesSquare   = 4
	SidesCircle   = 0
)

func CalcSquare(sideLen float64, sidesNum SidesType) float64 {
	var square float64

	switch sidesNum {
	case SidesCircle:
		square = math.Pi * sideLen * sideLen
	case SidesTriangle:
		p := (3 * sideLen) / 2
		square = math.Sqrt(p * ((p - sideLen) * (p - sideLen) * (p - sideLen)))
	case SidesSquare:
		square = sideLen * sideLen
	default:
		square = 0
	}

	return float64(square)
}
