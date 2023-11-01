package main

import (
	"bufio"
	"fmt"
	"os"
)

func part_1() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewScanner(file)

	sum := 0

	for reader.Scan() {
		seen := map[byte]bool{}
		line := reader.Text()
		halfllen := len(line) / 2
		var shared int
		left, right := line[:halfllen], line[halfllen:]

		for i := 0; i < halfllen; i++ {
			seen[left[i]] = true
		}
		for i := 0; i < halfllen; i++ {
			if seen[right[i]] {
				shared = int(right[i])
				break
			}
		}
		sum += priority(shared)
	}

	fmt.Printf("Answer to part 1: %d\n", sum)

}

func part_2() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewScanner(file)

	sum := 0
	const a byte = 'a'
	const a_subtr int = int(a) - 1
	const A byte = 'A'
	const A_subtr int = int(A) - 27
	var seen [3]map[byte]bool
	for i := range seen {
		seen[i] = make(map[byte]bool)
	}

	line_num := 0
	for reader.Scan() {
		line := reader.Text()
		for i := 0; i < len(line); i += 1 {
			seen[line_num][line[i]] = true
		}

		if line_num == 2 {
			// Intersect
			inter := intersect(intersect(seen[0], seen[1]), seen[2])
			if len(inter) != 1 {
				panic(fmt.Sprintf("Size of intersection is not 1 but %d", len(inter)))
			}

			// Get single item in intersection
			var shared byte
			for k := range inter {
				shared = k
			}

			sum += priority(int(shared))

			// Clear seen items
			for i := range seen {
				clear(seen[i])
			}
		}
		line_num = (line_num + 1) % 3
	}

	fmt.Printf("Answer to part 2: %d\n", sum)
}

const a byte = 'a'
const a_subtr int = int(a) - 1
const A byte = 'A'
const A_subtr int = int(A) - 27

func priority(item int) int {
	if 'a' <= item && item <= 'z' {
		return item - a_subtr
	} else {
		return item - A_subtr
	}
}

func intersect(a map[byte]bool, b map[byte]bool) map[byte]bool {
	out := map[byte]bool{}
	for ai := range a {
		if b[ai] {
			out[ai] = true
		}
	}
	return out
}

func main() {
	part_1()
	part_2()
}
