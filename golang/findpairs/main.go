package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseArray(input string) ([]int, error) {
	input = strings.TrimSpace(input)

	if len(input) < 2 || input[0] != '[' || input[len(input)-1] != ']' {
		return nil, fmt.Errorf("invalid input")
	}

	content := strings.TrimSpace(input[1 : len(input)-1])

	if content == "" {
		return []int{}, nil
	}

	parts := strings.Split(content, ",")
	result := make([]int, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)

		if part == "" {
			return nil, fmt.Errorf("invalid input")
		}

		n, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("invalid number: %s", part)
		}

		result = append(result, n)
	}

	return result, nil
}

func findPairs(arr []int, target int) [][]int {
	pairs := [][]int{}

	used := make([]bool, len(arr))

	for i := 0; i < len(arr); i++ {
		if used[i] {
			continue
		}

		for j := i + 1; j < len(arr); j++ {
			if used[j] {
				continue
			}

			if arr[i]+arr[j] == target {
				pairs = append(pairs, []int{i, j})
				used[i] = true
				used[j] = true
				break
			}
		}
	}

	return pairs
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid input.")
		return
	}

	arr, err := parseArray(os.Args[1])
	if err != nil {
		if err.Error() == "invalid input" {
			fmt.Println("Invalid input.")
		} else {
			fmt.Println(err.Error())
		}
		return
	}

	targetString := strings.TrimSpace(os.Args[2])

	target, err := strconv.Atoi(targetString)
	if err != nil || strings.ContainsAny(targetString, " \t\n\r") {
		fmt.Println("Invalid target sum.")
		return
	}

	pairs := findPairs(arr, target)

	if len(pairs) == 0 {
		fmt.Println("No pairs found.")
		return
	}

	fmt.Printf("Pairs with sum %d: %v\n", target, pairs)
}
