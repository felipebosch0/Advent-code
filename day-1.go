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
		if num > nbrs[i] {
			up++
		}
	}
	fmt.Printf("d1p2 %d\n", up)
	return up
}

func d1pXX(nbrs []int, offset int) int {
	up := 0

	for i, num := range nbrs[offset:] {
		if num > nbrs[i] {
			up++
		}
	}
	fmt.Printf("d1pXX, peek %d %d\n", offset, up)
	return up
}

func d1pX(lines []string, peekAhead int) int {
	result := 0
	total := len(lines) - peekAhead
	for order, line := range lines {

		current, _ := strconv.Atoi(line)
		if total == order {
			break
		}
		next, _ := strconv.Atoi(lines[order+peekAhead])
		if next > current {
			result++
		}

	}
	fmt.Printf("d1pX Peek %d:%d\n", peekAhead, result)
	return result
}
