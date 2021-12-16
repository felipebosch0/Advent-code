package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// this func read .txt
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

func Part1(lines []string) {
	ejex := 0
	ejey := 0

	for _, v := range lines {
		order := 0
		if parsed, _ := fmt.Sscanf(v, "forward %d", &order); parsed > 0 {
			ejex += order
		} else if parsed, _ := fmt.Sscanf(v, "down %d", &order); parsed > 0 {
			ejey += order
		} else if parsed, _ := fmt.Sscanf(v, "up %d", &order); parsed > 0 {
			ejey -= order
		}
	}

	fmt.Printf("Result: %d\n", ejex*ejey)
}

func main() {
	lines := ReadToLines(os.Stdin)

	Part1(lines)
}
