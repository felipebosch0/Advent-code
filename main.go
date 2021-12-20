package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// this func read .txt (strings)
func ReadToLines(rd io.Reader) []string {
	var result []string
	reader := bufio.NewReader(rd)

	for {
		text, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		result = append(result, strings.TrimSuffix(text, "\n"))
	}
	return result
}

// this func read .txt (int)
func StringsToInts(input []string) []int {
	var result []int

	for _, current := range input {
		convertedInt, err := strconv.Atoi(current)
		if err != nil {
			fmt.Printf("Error parsing string: %s\n", err)
		} else {
			result = append(result, convertedInt)
		}
	}
	return result
}

func main() {
	lines := ReadToLines(os.Stdin)

	nbrs := StringsToInts(lines)

	fmt.Printf("%d\n", len(nbrs))

	d1p2(nbrs)
	d1pX(lines, 1)
	d1pXX(nbrs, 1)
	d1pX(lines, 3)
	d1pXX(nbrs, 3)
	d2p1(lines)

}
