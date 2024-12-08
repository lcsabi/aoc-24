package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isSafe(levels []int) bool {
	isIncreasing := levels[1] > levels[0]
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		if diff == 0 || diff > 3 || diff < -3 {
			return false
		}
		if (isIncreasing && diff < 0) || (!isIncreasing && diff > 0) {
			return false
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	safeCount := 0

	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		levels := make([]int, len(parts))
		for i, part := range parts {
			levels[i], _ = strconv.Atoi(part)
		}
		if isSafe(levels) {
			safeCount++
		}
	}

	fmt.Println(safeCount) // Part One
}
