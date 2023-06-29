package utils

import "math"

func TrunFloat(f float32, prec int) float32 {
	x := math.Pow10(prec)
	return float32(math.Trunc(float64(f)*x) / x)
}
