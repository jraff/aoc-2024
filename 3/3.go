package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func readInput() []byte {
	data, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func extractMul(data []byte) [][]int {
	re := regexp.MustCompile(`mul\((\d\d?\d?),(\d\d?\d?)\)`)
	result := re.FindAllSubmatch(data, -1)

	n := make([][]int, 0)
	for _, r := range result {
		d1 := r[1]
		d2 := r[2]

		i1, _ := strconv.Atoi(string(d1))
		i2, _ := strconv.Atoi(string(d2))

		n = append(n, []int{i1, i2})
	}
	return n
}

func extractMulAndDo(data []byte) [][]int {
	re := regexp.MustCompile(`(mul\((\d\d?\d?),(\d\d?\d?)\))|(do\(\))|(don\'t\(\))`)
	result := re.FindAllSubmatch(data, -1)

	n := make([][]int, 0)
	enabled := true

	for _, r := range result {
		p0 := r[0]
		s0 := string(p0)

		if s0 == "do()" {
			enabled = true
		} else if s0 == "don't()" {
			enabled = false
		} else {

			if enabled {
				p2 := r[2]
				s2 := string(p2)

				p3 := r[3]
				s3 := string(p3)

				i1, _ := strconv.Atoi(string(s2))
				i2, _ := strconv.Atoi(string(s3))

				n = append(n, []int{i1, i2})
			}
		}
	}
	return n
}

func p1() {
	d := readInput()
	m := extractMul(d)

	sum := 0
	for _, i := range m {
		sum = sum + i[0]*i[1]
	}
	fmt.Println(sum)
}

func p2() {
	d := readInput()
	m := extractMulAndDo(d)

	sum := 0
	for _, i := range m {
		sum = sum + i[0]*i[1]
	}
	fmt.Println(sum)
}

func main() {
	p1()
	p2()
}
