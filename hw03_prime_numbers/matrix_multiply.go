package calculate

import (
	"errors"
	"math/big"
)

type Matrix2D struct {
	x11 big.Int
	x12 big.Int
	x21 big.Int
	x22 big.Int
}

// return 2D matrix
// x11 x12
// x21 x22
func NewMatrix2D(x11, x12, x21, x22 big.Int) *Matrix2D {
	return &Matrix2D{
		x11: x11,
		x12: x12,
		x21: x21,
		x22: x22,
	}
}

//  a11 a12   b11 b12   a11*b11 + a12*b21; a11*b12 + a12*b22
//          X         =
//  a21 a22   b21 b22   a21*b11 + a22*b21; a21*b12 + a22*b22
func (a *Matrix2D) Multiply(b *Matrix2D) {
	a11Xb11 := big.NewInt(0).Mul(&a.x11, &b.x11)
	a12Xb21 := big.NewInt(0).Mul(&a.x12, &b.x21)
	a11Xb12 := big.NewInt(0).Mul(&a.x11, &b.x12)
	a12Xb22 := big.NewInt(0).Mul(&a.x12, &b.x22)
	a21Xb11 := big.NewInt(0).Mul(&a.x21, &b.x11)
	a22Xb21 := big.NewInt(0).Mul(&a.x22, &b.x21)
	a21Xb12 := big.NewInt(0).Mul(&a.x21, &b.x12)
	a22Xb22 := big.NewInt(0).Mul(&a.x22, &b.x22)
	a.x11 = *big.NewInt(0).Add(a11Xb11, a12Xb21)
	a.x12 = *big.NewInt(0).Add(a11Xb12, a12Xb22)
	a.x21 = *big.NewInt(0).Add(a21Xb11, a22Xb21)
	a.x22 = *big.NewInt(0).Add(a21Xb12, a22Xb22)
}

func (a *Matrix2D) Power(n uint64) (err error) {
	if n == 0 {
		return errors.New("power must be 1 or greater uint64")
	}
	pow := NewMatrix2D(*big.NewInt(1), *big.NewInt(0), *big.NewInt(1), *big.NewInt(0))
	for ; n > 1; n /= 2 {
		if n%2 == 1 {
			pow.Multiply(a)
		}
		a.Multiply(a)
	}
	pow.Multiply(a)
	a.x11 = pow.x11
	a.x12 = pow.x12
	a.x21 = pow.x21
	a.x22 = pow.x22
	return nil
}

// func PowerViaPowOfTwo(x float64, n int) (pow float64) {
// 	pow = 1.0
// 	if n == 0 {
// 		return pow
// 	}
// 	for ; n > 1; n /= 2 {
// 		if n%2 == 1 {
// 			pow *= x
// 		}
// 		x *= x
// 	}
// 	pow *= x
// 	return pow
// }

func (a *Matrix2D) X11() *big.Int {
	return &a.x11
}

func (a *Matrix2D) X12() *big.Int {
	return &a.x12
}

func (a *Matrix2D) X21() *big.Int {
	return &a.x21
}

func (a *Matrix2D) X22() *big.Int {
	return &a.x22
}
