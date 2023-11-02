package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part_1() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	n_overl := 0

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			continue
		}

		inters := strings.Split(line, ",")

		inter_1 := str_to_interv(inters[0])
		inter_2 := str_to_interv(inters[1])

		if (inter_1[0] <= inter_2[0] && inter_2[1] <= inter_1[1]) || (inter_2[0] <= inter_1[0] && inter_1[1] <= inter_2[1]) {
			n_overl++
		}
	}

	fmt.Println("Answer for part 1:", n_overl)
}

func part_2() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	n_overl := 0

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			continue
		}

		inters := strings.Split(line, ",")

		inter_1 := str_to_interv(inters[0])
		inter_2 := str_to_interv(inters[1])

		/* non-overlapping cases:
		   [--a--][---b---]
		   [---b---]   [--a--]
		*/
		if !(inter_1[1] < inter_2[0] || inter_2[1] < inter_1[0]) {
			n_overl++
		}
	}

	fmt.Println("Answer for part 2:", n_overl)
}

func str_to_interv(interv string) [2]int {
	inter_str := strings.Split(interv, "-")
	start, err := strconv.Atoi(inter_str[0])
	if err != nil {
		panic(err)
	}
	end, err := strconv.Atoi(inter_str[1])
	if err != nil {
		panic(err)
	}
	return [2]int{start, end}
}

func main() {
	part_1()
	part_2()
}
