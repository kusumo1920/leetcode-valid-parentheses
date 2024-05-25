package main

import "fmt"

func main() {
	input := "({[]})"
	output := isValidSolution1(input)
	fmt.Println(output)
}

func isValidSolution1(s string) bool {
	var stack []string
	mapper := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}

	for _, c := range s {
		if _, ok := mapper[string(c)]; ok {
			stack = append(stack, string(c))
		} else {
			if len(stack) == 0 {
				return false
			}
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if mapper[pop] != string(c) {
				return false
			}
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}
