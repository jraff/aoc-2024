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

func main() {
	l1 := make([]int, 0)
	l2 := make([]int, 0)

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "   ")

		s1 := split[0]
		s2 := split[1]

		i1, _ := strconv.Atoi(s1)
		i2, _ := strconv.Atoi(s2)

		l1 = append(l1, i1)
		l2 = append(l2, i2)
	}

	// Sort l1 and l2
	slices.Sort(l1)
	slices.Sort(l2)

	differences := 0
	for i := range l1 {
		i1 := l1[i]
		i2 := l2[i]

		d := i1 - i2
		if d < 0 {
			d = d * -1
		}
		differences += d
		//fmt.Println(i1, i2, d)
	}

	fmt.Println(differences)
}
