package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
const (
	SidesTriangle = 3
	SidesSquare   = 4
	SidesCircle   = 0
)

type Sides int8

func CalcSquare(sideLen float64, sidesNum Sides) float64 {
	switch sidesNum {
	case SidesCircle:
		return math.Pow(sideLen, 2) * math.Pi
	case SidesSquare:
		return sideLen * sideLen
	case SidesTriangle:
		return (math.Pow(sideLen, 2) * math.Sqrt(3.0)) / 4.0
	}
	return 0
}
