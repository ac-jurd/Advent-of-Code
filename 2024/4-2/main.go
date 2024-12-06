package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Position struct {
	x int
	y int
}

func main() {
	// Keep track of total number of X-MASes found
	total := 0
	// Load lines from file
	lines := MustLoadData("input.txt")

	// Iterate over lines in line array
	for y, line := range lines {
		// Iterate over offsets in line
		for x := range line {
			// Check if current x,y position is the center of an X-MAS
			if CheckX(lines, Position{x, y}) {
				// Increment total count of M-MASes
				total += 1
			}
		}
	}

	fmt.Println(total)
}

// CheckX reads bytes surrounding a Position for X-MAS.
// Returns true if current Position is the A in X-MAX.
func CheckX(lines []string, pos Position) bool {
	// If byte at current position is not 'A', don't bother checking
	if AtPosition(lines, pos) != 'A' {
		return false
	}

	// Calculate array indices and get byte values
	// Upper left position
	ulPos := pos
	ulPos.x -= 1
	ulPos.y -= 1
	ulByte := AtPosition(lines, ulPos)

	// Upper right position
	urPos := pos
	urPos.x += 1
	urPos.y -= 1
	urByte := AtPosition(lines, urPos)

	// Lower left position
	llPos := pos
	llPos.x -= 1
	llPos.y += 1
	llByte := AtPosition(lines, llPos)

	// Lower right position
	lrPos := pos
	lrPos.x += 1
	lrPos.y += 1
	lrByte := AtPosition(lines, lrPos)

	// If upper-left-lower-right contains MAS
	if (ulByte == 'M' && lrByte == 'S') || (ulByte == 'S' && lrByte == 'M') {
		// If lower-left-upper-right contains MAS
		if (llByte == 'M' && urByte == 'S') || (llByte == 'S' && urByte == 'M') {
			// Current position is 'A' in X-MAS
			return true
		}
	}

	// Current position is not 'A' in X-MAS
	return false
}

// AtPosition gets the byte at a specific x,y point in the lines array.
// Position contains the x,y values.
// If position is out of array bounds, AtPosition returns a garbage value.
func AtPosition(lines []string, pos Position) byte {
	// If y value is out of bounds, return garbage value
	if pos.y < 0 || pos.y > len(lines)-1 {
		return ' '
	}
	// If x value is out of bounds, return garbage value
	if pos.x < 0 || pos.x > len(lines[0])-1 {
		return ' '
	}

	// Extract and return byte at position
	line := lines[pos.y]
	return line[pos.x]
}

// MustLoadData reads lines from file given a filename.
// Returns array of strings.
// Panics on error.
func MustLoadData(filename string) []string {
	// Get file reference
	file, err := os.Open(filename)

	// If there was an error, panic
	if err != nil {
		log.Fatalln(err)
	}

	// Initialize empty array of strings to hold main output
	lines := make([]string, 0)
	// Initialize new reader
	reader := bufio.NewReader(file)

	for {
		// Read line from file split by newlines
		s, err := reader.ReadString('\n')

		// Handle errors
		if err != nil {
			// If the error was simply EOF, end file read loop
			if err.Error() == "EOF" {
				break
			} else {
				// If the error was not EOF, panic
				log.Fatalln(err)
			}
		}

		// Check length of line to filter out empty lines ("\r\n")
		if len(s) > 2 {
			// Trim newlines and carriage returns
			trimmed := strings.Trim(s, "\r\n")
			// Append read line to list of lines
			lines = append(lines, trimmed)
		}
	}

	// Return array of lines as strings
	return lines
}
