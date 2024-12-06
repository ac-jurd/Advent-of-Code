package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	UP int = iota
	RIGHT
	DOWN
	LEFT
)

// Map to convert direction constants to their string names
var directionNames = map[int]string{
	LEFT:  "LEFT",
	RIGHT: "RIGHT",
	UP:    "UP",
	DOWN:  "DOWN",
}

func main() {
	grid := Grid{}
	grid.MustLoad("input.txt")

	startX, startY := grid.StartPos()
	guard := NewGuard(startX, startY)

	for {
		if grid.Detect(&guard) == ' ' {
			break
		}

		grid.SetAtPosition(guard.x, guard.y, 'X')
		if grid.Detect(&guard) == '#' {
			guard.TurnRight()
		}
		guard.Step()
	}
	grid.SetAtPosition(guard.x, guard.y, 'X')

	grid.Print()
	fmt.Println("Count:", grid.Count('X'))
}

type Guard struct {
	x   int
	y   int
	dir int
}

func NewGuard(x int, y int) Guard {
	return Guard{x, y, UP}
}

func (g *Guard) Step() {
	if g.dir == LEFT {
		g.x -= 1
	} else if g.dir == RIGHT {
		g.x += 1
	} else if g.dir == UP {
		g.y -= 1
	} else if g.dir == DOWN {
		g.y += 1
	}
}

func (g *Guard) TurnRight() {
	g.dir += 1
	g.dir = g.dir % 4
}

func (g *Grid) Detect(guard *Guard) byte {
	if guard.dir == LEFT && guard.x <= 0 {
		return ' '
	}
	if guard.dir == RIGHT && guard.x >= len(g.lines[0]) {
		return ' '
	}
	if guard.dir == UP && guard.y <= 0 {
		return ' '
	}
	if guard.dir == DOWN && guard.y >= len(g.lines) {
		return ' '
	}

	if guard.dir == LEFT {
		return g.AtPosition(guard.x-1, guard.y)
	} else if guard.dir == RIGHT {
		return g.AtPosition(guard.x+1, guard.y)
	} else if guard.dir == UP {
		return g.AtPosition(guard.x, guard.y-1)
	} else if guard.dir == DOWN {
		return g.AtPosition(guard.x, guard.y+1)
	}
	return ' '
}

func (g *Guard) ToString() string {
	// Use the directionNames map to get the name of the direction
	dirName, exists := directionNames[g.dir]
	if !exists {
		dirName = "UNKNOWN" // Fallback for invalid direction values
	}
	s := fmt.Sprintf("%d,%d %s\n", g.x, g.y, dirName)
	return s
}

type Grid struct {
	lines []string
}

func (g *Grid) MustLoad(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	lines := make([]string, 0)
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				log.Fatalln(err)
			}
		}

		if len(line) > 2 {
			line = strings.Trim(line, "\r\n")
			lines = append(lines, line)
		}
	}
	g.lines = lines
}

func (g *Grid) StartPos() (x int, y int) {
	for y, line := range g.lines {
		for x := range line {
			if g.AtPosition(x, y) == '^' {
				return x, y
			}
		}
	}

	log.Fatalln("Guard not found in grid")
	return 0, 0
}

func (g *Grid) AtPosition(x int, y int) byte {
	if y < 0 || y > len(g.lines)-1 {
		return ' '
	}
	if x < 0 || x > len(g.lines[0])-1 {
		return ' '
	}

	line := g.lines[y]
	return line[x]
}

func (g *Grid) SetAtPosition(x int, y int, b byte) {
	if y < 0 || y > len(g.lines)-1 {
		return
	}
	if x < 0 || x > len(g.lines[0])-1 {
		return
	}

	line := g.lines[y]
	updated := []byte(line)
	updated[x] = b
	g.lines[y] = string(updated)
}

func (g *Grid) Print() {
	for _, line := range g.lines {
		fmt.Println(line)
	}
}

func (g *Grid) Count(b byte) int {
	total := 0
	for _, line := range g.lines {
		for x := range line {
			if line[x] == b {
				total += 1
			}
		}
	}
	return total
}
