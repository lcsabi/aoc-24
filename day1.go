package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func absDiff(a int, b int) int {
	diff := a - b
	if diff < 0 {
		return -diff
	}
	return diff
}

func main() {
	lefts := make([]int, 0)
	rights := make([]int, 0)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		left, _ := strconv.Atoi(scanner.Text())
		lefts = append(lefts, left)

		scanner.Scan()
		right, _ := strconv.Atoi(scanner.Text())
		rights = append(rights, right)
	}
	sort.Ints(lefts)
	sort.Ints(rights)

	diff := 0
	for i := range len(lefts) {
		diff += absDiff(lefts[i], rights[i])
	}
	fmt.Println(diff)
}
