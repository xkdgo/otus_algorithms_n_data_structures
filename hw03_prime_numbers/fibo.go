package calculate

import "math/big"

func FiboRecursion(N uint64) (res *big.Int) {
	res = big.NewInt(0)
	if N == 0 {
		return big.NewInt(0)
	}
	if N == 1 || N == 2 {
		return big.NewInt(1)
	}

	return res.Add(FiboRecursion(N-1), FiboRecursion(N-2))
}

func FiboCicle(N uint64) (res *big.Int) {
	if N == 0 {
		return big.NewInt(0)
	}
	if N == 1 {
		return big.NewInt(1)
	}
	f1 := big.NewInt(1)
	f2 := big.NewInt(1)
	for j := uint64(2); j < N; j++ {
		f3 := big.NewInt(0)
		f3 = f3.Add(f1, f2)
		f1, f2 = f2, f3
	}
	return f2
}
