package calculate

import (
	"fmt"
	"math"
)

func PowerStandard(x, y float64) (pow float64) {
	return math.Pow(x, y)
}

func PowerViaMultiply(x, y float64) (pow float64) {
	if int(y) == 0 {
		return 1
	}
	xpow1 := x
	j := 2
	for ; j <= int(y/2); j *= 2 {
		x *= x
	}
	fmt.Println("j=", j)
	if j > int(y/2) {
		j = int(y / 2)
	}
	for ; j <= int(y); j++ {
		x *= xpow1
	}
	return x
}
