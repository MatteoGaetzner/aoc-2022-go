package main

import (
	"bufio"
	"fmt"
	"os"
)

var choices2points = [3][3]uint8{
	{3, 6, 0},
	{0, 3, 6},
	{6, 0, 3},
}

var choice2outcome2choice = [3][3]byte{
	{'X', 'Z', 'Y'},
	{'Z', 'Y', 'X'},
	{'Y', 'X', 'Z'},
}

func symbol_points(self byte) uint8 {
	return self - 'X' + 1
}

func win_points(other byte, self byte) uint8 {
	return choices2points[other-'A'][self-'X']
}

func part_1() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var total uint32 = 0

	for scanner.Scan() {
		line := scanner.Text()
		other, self := line[0], line[2]
		wp := win_points(other, self)
		sp := symbol_points(self)
		total += uint32(wp + sp)
	}
	fmt.Println("Part 1 answer:", total)
}

func part_2() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var total uint32 = 0

	for scanner.Scan() {
		line := scanner.Text()
		other := line[0]
		outcome := line[2]
		println(string(other), string(outcome))
		println('C'-other, 'Z'-outcome, string(choice2outcome2choice['C'-other]['Z'-outcome]))
		println()
		self := choice2outcome2choice['C'-other]['Z'-outcome]
		// fmt.Printf("other=%s, outcome=%s, self=%s\n", string(other), string(outcome), string(self))
		wp := win_points(other, self)
		sp := symbol_points(self)
		total += uint32(wp + sp)
	}
	fmt.Println("Part 2 answer:", total)
}

func main() {
	part_1()
	part_2()
}
