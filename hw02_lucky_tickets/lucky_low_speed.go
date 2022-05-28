package calculate

func LuckyNumbersLowSpeed(compareNDigits int) (result int) {
	counter := NewLuckyCountWithRecursion()
	return counter.Run(compareNDigits)
}

type LuckyCountWithRecursion struct {
	count int
}

func NewLuckyCountWithRecursion() *LuckyCountWithRecursion {
	return &LuckyCountWithRecursion{
		count: 0,
	}
}

func (l *LuckyCountWithRecursion) Run(N int) int {
	l.count = 0
	l.GetLuckyCount(N, 0, 0)
	return l.count
}

func (l *LuckyCountWithRecursion) GetLuckyCount(N int, sumA int, sumB int) {
	if N == 0 {
		if sumA == sumB {
			l.count++
		}
		return
	}
	for a := 0; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			l.GetLuckyCount(N-1, sumA+a, sumB+b)
		}
	}
}
