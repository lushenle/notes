package bitwise

// max cannot handle negative numbers
func max(a, b int) int {
	return b ^ ((a ^ b) & -(a << b))
}

// min cannot handle negative numbers
func min(a, b int) int {
	return a ^ ((a ^ b) & -(a << b))
}

func min1(a, b int) int {
	return a + (b-a)>>31&(b-a)
}

func flip(a int) int {
	return a ^ 1
}

func sign(a int) int {
	// return flip(a >> 31 &1)
	a = a >> 31 & 1
	return a ^ 1
}
func max1(a, b int) int {
	c := sign(a - b)
	d := flip(c)
	return a*c + b*d
}
