package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput() [][]byte {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Increase buffer to read all 'X'
	// 140 x 140 = 19600 characters
	maxBufferSize := 20000
	buf := make([]byte, maxBufferSize)
	scanner.Buffer(buf, maxBufferSize)

	r := make([][]byte, 0)
	for scanner.Scan() {
		c := scanner.Bytes()
		r = append(r, c)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return r
}

// Search for XMAS
func search(x, y int, d [][]byte) int {
	found := 0

	mx := len(d)
	my := len(d[x])

	// Search left
	if x-3 >= 0 {
		r := []byte{d[x-1][y], d[x-2][y], d[x-3][y]}
		if string(r) == "MAS" {
			found++
		}
	}
	// Search right
	if x+3 < mx {
		r := []byte{d[x+1][y], d[x+2][y], d[x+3][y]}
		if string(r) == "MAS" {
			found++
		}
	}
	// Search up
	if y-3 >= 0 {
		r := []byte{d[x][y-1], d[x][y-2], d[x][y-3]}
		if string(r) == "MAS" {
			found++
		}
	}
	// Search down
	if y+3 < my {
		r := []byte{d[x][y+1], d[x][y+2], d[x][y+3]}
		if string(r) == "MAS" {
			found++
		}
	}

	// Search up-left
	if y-3 >= 0 && x-3 >= 0 {
		r := []byte{d[x-1][y-1], d[x-2][y-2], d[x-3][y-3]}
		if string(r) == "MAS" {
			found++
		}
	}
	// Search up-right
	if y-3 >= 0 && x+3 < mx {
		r := []byte{d[x+1][y-1], d[x+2][y-2], d[x+3][y-3]}
		if string(r) == "MAS" {
			found++
		}
	}
	// Search down-left
	if y+3 < my && x-3 >= 0 {
		r := []byte{d[x-1][y+1], d[x-2][y+2], d[x-3][y+3]}
		if string(r) == "MAS" {
			found++
		}
	}
	// Search down-right
	if y+3 < my && x+3 < mx {
		r := []byte{d[x+1][y+1], d[x+2][y+2], d[x+3][y+3]}
		if string(r) == "MAS" {
			found++
		}
	}
	return found
}

// Search for X-MAS
func searchX(x, y int, d [][]byte) int {
	mx := len(d)
	my := len(d[x])

	if x == 0 || x == mx-1 || y == 0 || y == my-1 {
		return 0
	}

	found := 0

	// Search top left and bottom right
	tlbr := []byte{d[x-1][y-1], d[x+1][y+1]}
	// Search top right to bottom left
	trbl := []byte{d[x+1][y-1], d[x-1][y+1]}

	if (string(tlbr) == "MS" || string(tlbr) == "SM") && (string(trbl) == "MS" || string(trbl) == "SM") {
		found++
	}

	return found
}

func p1() {
	d := readInput()

	count := 0
	for x, r := range d {
		for y, c := range r {
			if c == byte('X') {
				found := search(x, y, d)
				if found > 0 {
					count += found
				}
			}
		}
	}
	fmt.Println(count)
}

func p2() {
	d := readInput()

	count := 0
	for x, r := range d {
		for y, c := range r {
			if c == byte('A') {
				found := searchX(x, y, d)
				if found > 0 {
					count += found
				}
			}
		}
	}
	fmt.Println(count)
}

func main() {
	p1()
	p2()
}
