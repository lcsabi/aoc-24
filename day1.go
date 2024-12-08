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
	rightFreqMap := make(map[int]int)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		left, _ := strconv.Atoi(scanner.Text())
		lefts = append(lefts, left)

		scanner.Scan()
		right, _ := strconv.Atoi(scanner.Text())
		rights = append(rights, right)
		rightFreqMap[right]++
	}
	sort.Ints(lefts)
	sort.Ints(rights)

	totalDistance := 0
	similarityScore := 0
	for i := range len(lefts) {
		totalDistance += absDiff(lefts[i], rights[i])
		similarityScore += lefts[i] * rightFreqMap[lefts[i]]
	}
	fmt.Println(totalDistance)   // Part One
	fmt.Println(similarityScore) // Part Two
}
