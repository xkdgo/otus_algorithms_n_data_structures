package calculate

import (
	"math"
)

func PowerStandard(x float64, n int) (pow float64) {
	return math.Pow(x, float64(n))
}

func PowerViaMultiply(x float64, n int) (pow float64) {
	if n == 0 {
		return 1
	}
	xpow1 := x
	for j := 2; j < n; j++ {
		x *= xpow1
	}
	x *= xpow1
	return x
}

func PowerViaPowOfTwo(x float64, n int) (pow float64) {
	pow = 1.0
	if n == 0 {
		return pow
	}
	for ; n > 1; n /= 2 {
		if n%2 == 1 {
			pow *= x
		}
		x *= x
	}
	pow *= x
	return pow
}
