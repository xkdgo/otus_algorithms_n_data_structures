package bitboard

// This function count bits with 1 in uint64.
func PopcntCicle(mask uint64) (count int) {
	for mask > 0 {
		// 0101
		// &
		// 0001
		//    1
		if mask&1 == 1 {
			count++
		}
		// 0101 >> 1
		// =
		// 010
		mask >>= 1
	}
	return count
}

// This function count bits with 1 in uint64.
func PopcntSparseFast(mask uint64) (count int) {
	for mask > 0 {
		count++
		// 1010000001
		// &
		// 1010000000
		// =
		// 1010000000
		// &
		// 1001111111
		// =
		// 1000000000
		// &
		// 0111111111
		// =
		// 0
		mask &= mask - 1

	}
	return count
}

// Hacker’s Delight version
// divide and conquer
// good way to count the number of 1-bits is to first set
// each 2-bit field equal to the sum of the two single bits that were originally in the field,
// and then sum adjacent 2-bit fields,
// putting the results in each 4-bit field, and so on.
func Popcnt(mask uint64) (count int) {
	m0 := uint64(0x5555555555555555) // 01010101 ...
	m1 := uint64(0x3333333333333333) // 00110011 ...
	m2 := uint64(0x0f0f0f0f0f0f0f0f) // 00001111 ...
	m := uint64(1<<64 - 1)
	mask = mask>>1&(m0&m) + mask&(m0&m)
	mask = mask>>2&(m1&m) + mask&(m1&m)
	mask = (mask>>4 + mask) & (m2 & m)
	mask += mask >> 8
	mask += mask >> 16
	mask += mask >> 32
	return int(mask) & (1<<7 - 1)
}
