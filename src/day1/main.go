package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, err := ioutil.ReadFile("src/day1/input.txt")

	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(file), "\n")
	nums := make([]int, 0, len(rows))

	for _, row := range rows {
		value, _ := strconv.Atoi(strings.TrimSpace(row))
		nums = append(nums, value)
	}

	for i := 0; i < len(nums); i++ {
		var multiplicands [3]int
		var found bool

		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]

				if sum == 2020 {
					found = true
					multiplicands[0] = nums[i]
					multiplicands[1] = nums[j]
					multiplicands[2] = nums[k]
					break
				}
			}
		}

		if found {
			fmt.Println(multiplicands[0] * multiplicands[1] * multiplicands[2])
			break
		}
	}
}
