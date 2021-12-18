package main

import (
	"fmt"
	"strconv"
)

func d1p1(lines []string) {
	result := 0
	total := len(lines) - 1
	for order, line := range lines {
		fmt.Printf("%d:%s\n", order, line)
		current, _ := strconv.Atoi(line)
		if total == order {
			break
		}
		next, _ := strconv.Atoi(lines[order+1])
		if next > current {
			result++
		}

	}
	fmt.Printf("Part 1:%d\n", result)
}

func d1p2(nbrs []int) int {
	up := 0
	for i, num := range nbrs[3:] {
		fmt.Printf("%d: %d, %d\n", i, num, nbrs[i])
		if num < nbrs[i] {
			up++
		}
	}
	return up
}
