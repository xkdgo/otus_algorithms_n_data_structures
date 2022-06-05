package calculate

import "math"

func CountPrimesLinear(limit int) int {
	// массив в который будет записываться наименьший делитель
	lp := make([]int, limit+1)
	// массив простых чисел
	primes := make([]int, 0, (limit+1)/2)
	for i := 2; i <= limit; i++ {
		if lp[i] == 0 {
			primes = append(primes, i)
			lp[i] = i
		}
		for _, p := range primes {
			if i*p < len(lp) && p <= lp[i] {
				lp[i*p] = p
			}
		}
	}
	return len(primes)
}

func EratosphenClassic(limit int) int {
	divs := make([]bool, limit+1)
	count := 0
	for p := 2; p <= limit; p++ {
		if !divs[p] {
			count++
			if p > int(math.Ceil(math.Sqrt(float64(limit)))) {
				continue
			}
			for j := p * p; j <= limit; j += p {
				// if true then not prime
				divs[j] = true
			}
		}
	}
	return count
}
