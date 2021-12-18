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
	nbrs := StringsToInts(os.Args)

	d1p1(lines)
	d1p2(nbrs)
}
