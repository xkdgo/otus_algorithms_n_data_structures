package bitboard

// https://gekomad.github.io/Cinnamon/BitboardCalculator/
const (
	noGHMask              = uint64(0x3f3f3f3f3f3f3f3f)
	noABMask              = uint64(0xfcfcfcfcfcfcfcfc)
	noHMask               = uint64(0x7f7f7f7f7f7f7f7f)
	noAMask               = uint64(0xfefefefefefefefe)
	mainSlashDiagonal     = uint64(0x8040201008040201)
	aboveSlashMask        = uint64(0x7f3f1f0f07030100)
	downSlashMask         = uint64(0x0080c0e0f0f8fcfe)
	mainBackSlashDiagonal = uint64(0x0102040810204080)
	aboveBackslashMask    = uint64(0xfefcf8f0e0c08000)
	downBackslashMask     = uint64(0x000103070f1f3f7f)
)

func GetKingBitboardMoves(pos int) (mask uint64) {
	if pos < 0 || pos > 63 {
		return 0
	}
	currentBitboardPosition := uint64(1) << pos
	currA := currentBitboardPosition & noAMask
	currH := currentBitboardPosition & noHMask
	mask =
		(currA << 7) | (currentBitboardPosition << 8) | (currH << 9) |
			(currA >> 1) | (currH << 1) |
			(currA >> 9) | (currentBitboardPosition >> 8) | (currH >> 7)
	return mask
}

func GetKnightBitboardMoves(pos int) (mask uint64) {
	if pos < 0 || pos > 63 {
		return 0
	}
	currentBitboardPosition := uint64(1) << pos
	currA := currentBitboardPosition & noAMask
	currH := currentBitboardPosition & noHMask
	currAB := currentBitboardPosition & noABMask
	currGH := currentBitboardPosition & noGHMask
	mask = currA << 15
	mask = mask | currAB<<6
	mask = mask | currAB>>10
	mask = mask | currA>>17
	mask = mask | currH<<17
	mask = mask | currGH<<10
	mask = mask | currGH>>6
	mask = mask | currH>>15
	return mask
}

func GetRookBitboardMoves(pos int) (mask uint64) {
	if pos < 0 || pos > 63 {
		return 0
	}
	curPos := uint64(1) << pos
	mask = curPos ^ getRookMask(pos)
	return mask
}

func GetBishopBitboardMoves(pos int) (mask uint64) {
	if pos < 0 || pos > 63 {
		return 0
	}
	curPos := uint64(1) << pos
	mask = curPos ^ getBishopMask(pos)
	return mask
}

func GetQueenBitboardMoves(pos int) (mask uint64) {
	if pos < 0 || pos > 63 {
		return 0
	}
	curPos := uint64(1) << pos
	rookMask := curPos ^ getRookMask(pos)
	bishopMask := curPos ^ getBishopMask(pos)
	return rookMask | bishopMask
}

func getGorizontalMaskPosition(pos int) int {
	switch {
	case pos >= 0 && pos <= 7:
		return 0
	case pos <= 15:
		return 1
	case pos <= 23:
		return 2
	case pos <= 31:
		return 3
	case pos <= 39:
		return 4
	case pos <= 47:
		return 5
	case pos <= 55:
		return 6
	case pos <= 63:
		return 7
	}
	return 0
}

func getVerticalMaskPosition(pos int) int {
	switch {
	case pos >= 0 && pos <= 7:
		return pos
	case pos <= 15:
		return pos % 8
	case pos <= 23:
		return pos % 16
	case pos <= 31:
		return pos % 24
	case pos <= 39:
		return pos % 32
	case pos <= 47:
		return pos % 40
	case pos <= 55:
		return pos % 48
	case pos <= 63:
		return pos % 56
	}
	return 0
}

func getRookMask(pos int) uint64 {
	gorizontalInitMask := uint64(0xff)
	verticalInitMask := uint64(0x101010101010101)
	posGorizontal := gorizontalInitMask << (getGorizontalMaskPosition(pos) * 8)
	posVertical := verticalInitMask << getVerticalMaskPosition(pos)
	return posGorizontal | posVertical
}

// check position is down a8\h1
func checkDownBackslashDiagonal(pos int) bool {
	return ((uint64(1) << pos) & downBackslashMask) > 0
}

// check position is above a1/h8
func checkAboveSlashDiagonal(pos int) bool {
	return ((uint64(1) << pos) & aboveSlashMask) > 0
}

func getBishopMask(pos int) uint64 {
	var slashDiagonal uint64
	var backslashDaigonal uint64
	curPos := uint64(1) << pos
	if mainSlashDiagonal&curPos > 0 {
		slashDiagonal = mainSlashDiagonal
	} else {
		slashDiagonal = getSlashDiagonal(pos)
	}
	if mainBackSlashDiagonal&curPos > 0 {
		backslashDaigonal = mainBackSlashDiagonal
	} else {
		backslashDaigonal = getBackslashDiagonal(pos)
	}
	return slashDiagonal | backslashDaigonal
}

func getSlashDiagonal(pos int) uint64 {
	if pos == 0 || pos == 63 {
		return 0
	}

	if checkAboveSlashDiagonal(pos) {
		return (mainSlashDiagonal >> getSlashDiagonalAboveMoveBits(pos)) & aboveSlashMask
	}
	return (mainSlashDiagonal << getSlashDiagonalDownMoveBits(pos)) & downSlashMask
}

func getBackslashDiagonal(pos int) uint64 {
	if pos == 56 || pos == 7 {
		return 0
	}
	if checkDownBackslashDiagonal(pos) {
		return (mainBackSlashDiagonal >> getBackslashDiagonalDownMoveBits(pos)) & downBackslashMask
	}
	return (mainBackSlashDiagonal << getBackslashDiagonalAboveMoveBits(pos)) & aboveBackslashMask
}

//	56	57	58	59	60	61	62	63
//	48	49	50	51	52	53	54
//	40	41	42	43	44	45
//	32	33	34	35	36
//	24	25	26	27
//	16	17	18
//	08	09
//	00
func getSlashDiagonalAboveMoveBits(pos int) int {
	mainSlashDiag := []int{0, 9, 18, 27, 36, 45, 54, 63}
	for _, mainSlashPos := range mainSlashDiag {
		if pos < mainSlashPos {
			return mainSlashPos % pos
		}
	}
	return 0
}

//								63
//							54	55
//						45	46	47
//					36	37	38	39
//				27	28	29	30	31
//			18	19	20	21	22	23
//		09	10	11	12	13	14	15
//	00	01	02	03	04	05	06	07
func getSlashDiagonalDownMoveBits(pos int) int {
	mainSlashDiagBackSorted := []int{63, 54, 45, 36, 27, 18, 9, 0}
	for _, mainSlashPos := range mainSlashDiagBackSorted {
		if mainSlashPos == 0 {
			return pos
		}
		if pos > mainSlashPos {
			return pos % mainSlashPos
		}
	}
	return 0
}

//	56
//	48	49
//	40	41	42
//	32	33	34	35
//	24	25	26	27	28
//	16	17	18	19	20	21
//	08	09	10	11	12	13	14
//	00	01	02	03	04	05	06	07
func getBackslashDiagonalDownMoveBits(pos int) int {
	mainBackSlashDiag := []int{07, 14, 21, 28, 35, 42, 49, 56}
	if pos >= 0 && pos < 7 {
		return 7 - pos
	}
	for _, mainBackSlashPos := range mainBackSlashDiag {
		if pos < mainBackSlashPos {
			return mainBackSlashPos % pos
		}
	}
	return 0
}

//	56	57	58	59	60	61	62	63
//		49	50	51	52	53	54	55
//			42	43	44	45	46	47
//				35	36	37	38	39
//					28	29	30	31
//						21	22	23
//							14	15
//								07
func getBackslashDiagonalAboveMoveBits(pos int) int {
	mainBackSlashDiag := []int{56, 49, 42, 35, 28, 21, 14, 07}
	for _, mainBackSlashPos := range mainBackSlashDiag {
		if pos > mainBackSlashPos {
			return pos % mainBackSlashPos
		}
	}
	return 0
}
