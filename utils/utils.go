package utils

func AbsDiff(a int, b int) int {
	diff := a - b
	if diff < 0 {
		return -diff
	}
	return diff
}
