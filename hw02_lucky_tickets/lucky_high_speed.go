package calculate

func LuckyNumbersHighSpeed(compareNDigits int) (result int) {
	counter := NewLuckyCountWithArrays()
	return counter.Run(compareNDigits)
}

type LuckyCountWithArrays struct {
	count int
}

func NewLuckyCountWithArrays() *LuckyCountWithArrays {
	return &LuckyCountWithArrays{
		count: 0,
	}
}

func (l *LuckyCountWithArrays) Run(N int) int {
	l.count = 0
	arr := []int{1} // первый массив состоит из одного элемента
	for N != 0 {    // нужное количество раз
		arr = getNextArr(arr) // строим следующий массив из предыдущего
		N--
	}
	for _, v := range arr {
		l.count += v * v // сводим квадраты значений в получившемся массиве
	}
	return l.count
}

// функция для построения следующего массива из предыдущего
// Например если на вход поступает массив из 10 единиц
// Создаются промежуточные массивы а затем результат складывается по каждому из индексов
// 	1	1	1	1	1	1	1	1	1	1
// 		1	1	1	1	1	1	1	1	1	1
// 			1	1	1	1	1	1	1	1	1	1
// 				1	1	1	1	1	1	1	1	1	1
// 					1	1	1	1	1	1	1	1	1	1
// 						1	1	1	1	1	1	1	1	1	1
// 							1	1	1	1	1	1	1	1	1	1
// 								1	1	1	1	1	1	1	1	1	1
// 									1	1	1	1	1	1	1	1	1	1
// 										1	1	1	1	1	1	1	1	1	1
// Результирующий массив:
// 1	2	3	4	5	6	7	8	9	10	9	8	7	6	5	4	3	2	1
// возвращает функция.
func getNextArr(prevArr []int) []int {
	// длина следующего массива будет больше на 9
	newLen := len(prevArr) + 9
	arr := make([]int, newLen) // заготовка результата
	for i := 0; i < 10; i++ {
		tmpArr := make([]int, newLen)
		copy(tmpArr[i:], prevArr)
		for j := 0; j < newLen; j++ {
			arr[j] += tmpArr[j]
		}
	}
	return arr
}
