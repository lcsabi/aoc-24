package main

import (
	"aoc-24/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	safeCount := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), " ")
		fst, _ := strconv.Atoi(strs[0])
		snd, _ := strconv.Atoi(strs[1])
		if fst == snd {
			continue
		}
		if utils.AbsDiff(fst, snd) > 3 {
			continue
		}
		idx := 1
		if fst > snd { // Decreasing
			for idx < len(strs)-1 {
				left, _ := strconv.Atoi(strs[idx])
				right, _ := strconv.Atoi(strs[idx+1])
				if left <= right || left-right > 3 {
					break
				}
				idx++
			}
		} else { // Increasing
			for idx < len(strs)-1 {
				left, _ := strconv.Atoi(strs[idx])
				right, _ := strconv.Atoi(strs[idx+1])
				if left >= right || left-right < -3 {
					break
				}
				idx++
			}
		}
		if idx == len(strs)-1 {
			safeCount++
		}
	}

	fmt.Println(safeCount) // Part One
}
