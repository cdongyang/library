package utils

func Lowbit(x int) int {
	return x & (-x)
}
