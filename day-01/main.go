package main

import (
	"bufio"
	"fmt" // for printing
	"os"
	"strconv"
)

func part_one() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var max_sum int64 = 0
	var cur_sum int64 = 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			cur_sum = 0
			continue
		}

		cur_val, err := strconv.ParseInt(line, 10, 32)
		if err != nil {
			panic(err)
		}
		cur_sum += cur_val
		max_sum = max(cur_sum, max_sum)
	}
	fmt.Println("Answer to part one: ", max_sum)
}

func part_two() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	max_sums := [3]int64{0, 0, 0}
	var cur_sum int64 = 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			max_sums = maybe_update(max_sums, cur_sum)
			cur_sum = 0
			continue
		}

		cur_val, err := strconv.ParseInt(line, 10, 32)
		if err != nil {
			panic(err)
		}
		cur_sum += cur_val
	}

	if cur_sum != 0 {
		max_sums = maybe_update(max_sums, cur_sum)
	}

	fmt.Printf("Top 3 total calories: %v\n", max_sums)
	fmt.Printf("Answer to part two: %d\n", max_sums[0]+max_sums[1]+max_sums[2])
}

func maybe_update(max_sums [3]int64, cur_sum int64) [3]int64 {
	if max_sums[0] >= cur_sum {
		return max_sums
	}
	a, b, c := cur_sum, max_sums[1], max_sums[2]

	if a <= b {
		max_sums[0] = a
	} else if c <= a {
		max_sums[0], max_sums[1], max_sums[2] = b, c, a
	} else {
		max_sums[0], max_sums[1] = b, a
	}

	return max_sums
}

func main() {
	part_one()
	fmt.Println()
	part_two()
}
