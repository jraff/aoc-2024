package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Direction struct {
	dx int
	dy int
}

type Guard struct {
	x          int
	y          int
	direction  Direction
	appearance byte
	positions  map[int]bool
}

func (g *Guard) NextPosition() (int, int) {
	return g.x + g.direction.dy, g.y + g.direction.dx
}

func (g *Guard) Advance() {
	x, y := g.NextPosition()
	g.x = x
	g.y = y
}

func (g *Guard) Rotate() {
	if g.direction.dx == 0 && g.direction.dy == -1 {
		g.direction.dx = 1
		g.direction.dy = 0
		g.appearance = byte('>')
	} else if g.direction.dx == 1 && g.direction.dy == 0 {
		g.direction.dx = 0
		g.direction.dy = 1
		g.appearance = byte('v')
	} else if g.direction.dx == 0 && g.direction.dy == 1 {
		g.direction.dx = -1
		g.direction.dy = 0
		g.appearance = byte('<')
	} else if g.direction.dx == -1 && g.direction.dy == 0 {
		g.direction.dx = 0
		g.direction.dy = -1
		g.appearance = byte('^')
	}
}

type Map struct {
	d [][]byte
	w int
	h int
}

func (m *Map) IsLegalMove(x int, y int) bool {
	if x < 0 || y < 0 || x >= m.w || y >= m.h {
		return false
	}

	return true
}

func (m *Map) IsBlocking(x int, y int) bool {
	return m.d[x][y] == byte('#')
}

func (m *Map) MoveGuard(g *Guard) bool {
	nx, ny := g.NextPosition()

	if nx > m.w || ny > m.h || nx < 0 || ny < 0 {
		fmt.Println("Tried to move out of bounds")
		return false
	}

	if !m.IsLegalMove(nx, ny) {
		g.Advance()
		return false
	}

	if !m.IsBlocking(nx, ny) {
		m.d[g.x][g.y] = byte('X')
		g.Advance()
		g.positions[g.x*m.w+g.y+1] = true
		m.d[g.x][g.y] = g.appearance
	} else {
		g.Rotate()
		m.d[g.x][g.y] = g.appearance
	}

	r := m.MoveGuard(g)
	return r
}

func readInput() *Map {
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

	return &Map{
		d: r,
		w: len(r),
		h: len(r[0]),
	}
}

func isGuard(t byte) bool {
	return t == byte('^') || t == byte('>') || t == byte('<') || t == byte('v')
}

func findGuard(m [][]byte) (int, int, byte, error) {
	var gx, gy int
	var g byte

	for x := range m {
		for y := range m[x] {
			if isGuard(m[x][y]) {
				gx = x
				gy = y
				g = m[x][y]
			}
		}
	}

	if g == 0 {
		return 0, 0, 0, fmt.Errorf("Guard not found")
	}

	return gx, gy, g, nil
}

func getDirection(d byte) (int, int) {
	dx := 0
	dy := 0

	// Can I move my guard in the direction?
	if d == byte('^') {
		dy = -1
	} else if d == byte('>') {
		dx = 1
	} else if d == byte('v') {
		dy = 1
	} else if d == byte('<') {
		dx = -1
	}

	return dx, dy
}

func printMap(m *Map) {
	for _, x := range m.d {
		for _, y := range x {
			fmt.Printf("%c", y)
		}
		fmt.Println()
	}

	fmt.Println("--------------------------")
}

func p1() {
	m := readInput()
	x, y, d, err := findGuard(m.d)
	dx, dy := getDirection(d)

	g := &Guard{
		x:          x,
		y:          y,
		appearance: d,
		direction: Direction{
			dx,
			dy,
		},
		positions: make(map[int]bool),
	}

	if err != nil {
		fmt.Println("Could not find guard")
	}

	m.MoveGuard(g)
	fmt.Println(len(g.positions))
}

func main() {
	p1()
}
