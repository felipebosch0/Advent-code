package main

import "fmt"

func d2p1(lines []string) {
	xpos := 0
	ypos := 0
	total := len(lines) - 1
	for int, pos := range lines {
		act := 0
		if pass, _ := fmt.Sscanf(pos, "forward %d", act); pass > 0 {

			xpos += act

		} else if pass, _ := fmt.Sscanf(pos, "up &d", act); pass > 0 {

			ypos = ypos + act

		} else if pass, _ := fmt.Sscanf(pos, "down %d", act); pass > 0 {

			ypos = ypos - act
		}
		if total == int {
			break
		}

	}

	fmt.Printf("The position is :%d ", xpos*ypos)
}
