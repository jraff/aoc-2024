package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

type Rule []int
type Update []int

func readInput() (*[]Rule, *[]Update) {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal("Could not open input")
	}
	defer file.Close()

	rules := make([]Rule, 0)
	updates := make([]Update, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()

		// Found rule
		if bytes.Contains(line, []byte("|")) {
			s := bytes.Split(line, []byte("|"))

			i1, _ := strconv.Atoi(string(s[0]))
			i2, _ := strconv.Atoi(string(s[1]))

			rule := Rule{i1, i2}
			rules = append(rules, rule)
		} else if bytes.Contains(line, []byte(",")) {
			// Found update
			s := bytes.Split(line, []byte(","))

			update := Update{}
			for _, si := range s {
				i, _ := strconv.Atoi(string(si))
				update = append(update, i)
			}
			updates = append(updates, update)
		}
	}

	return &rules, &updates
}

func createRuleMap(r *[]Rule) *map[int][]int {
	m := make(map[int][]int)

	for _, i := range *r {
		i0 := i[0]
		i1 := i[1]
		if _, ok := m[i0]; ok {
			m[i0] = append(m[i0], i1)
		} else {
			m[i0] = []int{i1}
		}
	}
	return &m
}

/*
{
  47: [53, 13, 61, 29],
  97: [13, 61, 47, 29, 53, 75],
  75: [29, 53, 47, 61, 13],
  61: [13, 53, 29],
  29: [13],
  53: [29, 13],
}

*/

func applyRuleMap(ruleMap map[int][]int, updates *[]Update) int {
	sum := 0

	for _, update := range *updates {
		isValid := true
		for i, u := range update {
			if i == 0 {
				continue
			}
			if slices.Contains(ruleMap[u], update[i-1]) {
				isValid = false
				break
			}
		}

		if isValid {

			m := update[len(update)/2]
			sum += m
		}
	}
	return sum
}

func p1() {
	r, u := readInput()
	rMap := createRuleMap(r)

	result := applyRuleMap(*rMap, u)
	fmt.Println(result)
}

func main() {
	p1()
}
