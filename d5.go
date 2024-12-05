package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func ReverseIndexOf(haystack []int, needle int, startPosition int) int {
	for i := startPosition; i >= 0; i-- {
		if haystack[i] == needle {
			return i
		}
	}
	return -1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	defer file.Close()

	rules := make(map[int][]int)
	var updates [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			if len(parts) == 2 {
				key, err := strconv.Atoi(parts[0])
				if err != nil {
					log.Fatalf("Error parsing key: %v", err)
					return
				}

				value, err := strconv.Atoi(parts[1])
				if err != nil {
					log.Fatalf("1 Error parsing value: %v", err)
					return
				}

				rules[key] = append(rules[key], value)
			}

		} else {
			parts := strings.Split(line, ",")
			var values []int
			for _, part := range parts {
				value, err := strconv.Atoi(part)
				if err != nil {
					log.Fatalf("2 Error parsing value: %v", err)
				}
				values = append(values, value)
			}
			updates = append(updates, values)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning file: %v", err)
	}

	sumA, incorrectUpdates := PartOne(rules, updates)
	sumB := PartTwo(rules, incorrectUpdates)
	fmt.Printf("Part 1 sum: %d\n", sumA)
	fmt.Printf("Part 2 sum: %d\n", sumB)
}

func PartOne(rules map[int][]int, updates [][]int) (int, [][]int) {
	sum := 0
	var incorrectUpdates [][]int
	for i, update := range updates {
		isCorrect := true
		for j, number := range updates[i] {
			if !isCorrect {
				break
			}
			currRuleNumbers := rules[number]
			for _, ruleNumber := range currRuleNumbers {
				index := ReverseIndexOf(update, ruleNumber, j)
				if index != -1 && index < j {
					isCorrect = false
					fmt.Printf("Update |%d| is INCORRECT, FIRST RULE BROKEN: %d|%d\n", update, number, ruleNumber)
					incorrectUpdates = append(incorrectUpdates, update)
					break
				}
			}
		}
		if isCorrect {
			fmt.Printf("Update |%d| is CORRECT\n", update)
			sum += update[len(update)/2]
		}
	}
	return sum, incorrectUpdates
}

func PartTwo(rules map[int][]int, updates [][]int) int {
	sum := 0
	for _, update := range updates {
		slices.SortFunc(update, func(a, b int) int {
			x := rules[a]
			if slices.Contains(x, b) {
				return 1
			} else {
				return -1
			}
		})
		sum += update[len(update)/2]
	}
	return sum
}
