package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Position struct {
	x int
	y int
}

const (
	HORIZONTAL int = iota
	VERTICAL
)

func main() {
	lines := MustLoadData("input.txt")

	total := 0
	total += CountAllDiagonal(lines)
	total += CountAllHorizontal(lines)
	total += CountAllVertical(lines)

	fmt.Println(total)
}

func CountAllDiagonal(lines []string) int {
	total := 0

	total += Diagonal(lines)
	lines = Flip(lines, HORIZONTAL)
	total += Diagonal(lines)
	lines = Flip(lines, VERTICAL)
	total += Diagonal(lines)
	lines = Flip(lines, HORIZONTAL)
	total += Diagonal(lines)

	return total
}

func CountAllHorizontal(lines []string) int {
	total := 0

	total += Horizontal(lines)
	lines = Flip(lines, HORIZONTAL)
	total += Horizontal(lines)

	return total
}

func CountAllVertical(lines []string) int {
	total := 0

	total += Vertical(lines)
	lines = Flip(lines, VERTICAL)
	total += Vertical(lines)

	return total
}

func Horizontal(lines []string) int {
	total := 0
	re := regexp.MustCompile("XMAS")

	for _, line := range lines {
		total += len(re.FindAllString(line, -1))
	}
	return total
}

func Vertical(lines []string) int {
	total := 0

	for y, line := range lines {
		for x := range line {
			tmp := make([]byte, 4)
			tmp[0] += AtPosition(lines, Position{x, y})
			tmp[1] += AtPosition(lines, Position{x, y + 1})
			tmp[2] += AtPosition(lines, Position{x, y + 2})
			tmp[3] += AtPosition(lines, Position{x, y + 3})

			if string(tmp) == "XMAS" {
				total += 1
			}
		}
	}
	return total
}

func Diagonal(lines []string) int {
	total := 0

	for y, line := range lines {
		for x := range line {
			tmp := make([]byte, 4)
			tmp[0] += AtPosition(lines, Position{x, y})
			tmp[1] += AtPosition(lines, Position{x + 1, y + 1})
			tmp[2] += AtPosition(lines, Position{x + 2, y + 2})
			tmp[3] += AtPosition(lines, Position{x + 3, y + 3})

			if string(tmp) == "XMAS" {
				total += 1
			}
		}
	}
	return total
}

func AtPosition(lines []string, pos Position) byte {
	if pos.y < 0 || pos.y > len(lines)-1 {
		return ' '
	}
	if pos.x < 0 || pos.x > len(lines[0])-1 {
		return ' '
	}

	line := lines[pos.y]
	return line[pos.x]
}

func Flip(lines []string, direction int) []string {
	newLines := make([]string, 0)

	if direction == HORIZONTAL {
		for _, line := range lines {
			newLines = append(newLines, ReverseString(line))
		}
	} else if direction == VERTICAL {
		for i := len(lines) - 1; i >= 0; i-- {
			newLines = append(newLines, lines[i])
		}
	}

	return newLines
}

func ReverseString(line string) string {
	newLine := make([]byte, len(line))
	for i := len(line) - 1; i >= 0; i-- {
		newLine[len(line)-i-1] = line[i]
	}
	return string(newLine)
}

func MustLoadData(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	lines := make([]string, 0)
	reader := bufio.NewReader(file)
	err = nil

	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				log.Fatalln(err)
			}
		}

		if len(s) > 2 {
			trimmed := strings.Trim(s, "\r\n")
			lines = append(lines, trimmed)
		}
	}

	return lines
}
