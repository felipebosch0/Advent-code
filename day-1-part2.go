package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main1() {
	lines := ReadToLines2(os.Stdin)

	Part2(lines)
}

// this func read .txt
func ReadToLines2(rd io.Reader) []string {
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

func Part2(lines []string) {
	ejex := 0
	ejey := 0
	poss := 0

	for _, v := range lines {
		order := 0
		if diagnosis, _ := fmt.Sscanf(v, "forward %d", &order); diagnosis > 0 {
			ejex += order
			ejey += poss * order
		} else if diagnosis, _ := fmt.Sscanf(v, "down %d", &order); diagnosis > 0 {
			poss += order
		} else if diagnosis, _ := fmt.Sscanf(v, "up %d", &order); diagnosis > 0 {
			poss -= order
		}
	}

	fmt.Printf("Result: %d\n", ejex*ejey)
}
