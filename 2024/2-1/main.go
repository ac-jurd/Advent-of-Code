package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines, err := parseData("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	safe := 0
	for _, line := range lines {
		if isSameDirection(line) && isValidDistance(line) {
			safe++
		}
	}

	fmt.Println("Safe:", safe)
}

// Returns true if elements in array are in increasing or decreasing order
func isSameDirection(array []int) bool {
	// Get initial direction of first two elements
	direction := getDirection(array[0], array[1])

	// Iterate over remaining elements
	for i := 1; i < len(array)-1; i++ {
		// If the direction of the current pair does not match the original direction:
		if direction != getDirection(array[i], array[i+1]) {
			// The array does not go in the same direction
			return false
		}
	}

	// If we have made it this far then the array does go in the same direction
	return true
}

// Returns -1, 0, or 1 representing decreasing, stationary, or increasing order of elements
func getDirection(a int, b int) int {
	// Increase in sequence gives positive direction
	// Decrease in sequence gives negative direction
	// No difference in sequence gives neutral direction
	if a < b {
		return 1
	} else if a > b {
		return -1
	} else {
		return 0
	}
}

// Returns true if distance between each element is within given range (1-3)
func isValidDistance(array []int) bool {
	// Iterate over elements in array
	for i := 0; i < len(array)-1; i++ {
		// Compute distance between current element and next element
		distance := getDistance(float64(array[i]), float64(array[i+1]))
		// If distance is outside of specified range:
		if distance < 1 || distance > 3 {
			// The distances are not valid
			return false
		}
	}

	// The distances are valid
	return true
}

// Returns absolute difference between two numbers
func getDistance(a float64, b float64) float64 {
	return math.Abs(a - b)
}

// Returns 2D array of integers from lines of input file
func parseData(filename string) ([][]int, error) {

	// Open file for reading
	file, err := os.Open(filename)
	if err != nil {
		return [][]int{}, errors.New("Error opening file for parsing")
	}
	defer file.Close()

	// Initialize 2D output array of integers (columns)
	lines := make([][]int, 0)
	// Create new reader
	reader := bufio.NewReader(file)

	// Iterate over lines in file until EOF is reached
	for {
		// Read line
		line, err := reader.ReadString('\n')
		if err != nil {
			// If EOF was found, end loop
			if err.Error() == "EOF" {
				break
			} else {
				// If error was not EOF, something else went wrong
				return [][]int{}, err
			}
		}

		// Initialize output row (added to lines at the end)
		splLineInt := make([]int, 0)
		// Remove extra carriage returns and newlines from processed line
		line = strings.Trim(line, "\r\n")
		// Split string by spaces into seprate strings of numbers
		splLineString := strings.Split(line, " ")

		// Foreach string in split line:
		for _, s := range splLineString {
			// Convert the string into integer type
			i, err := strconv.Atoi(s)
			if err != nil {
				// String could not be converted into an integer
				log.Fatalln(err)
			}

			// Add parsed integer to list of integers for this row
			splLineInt = append(splLineInt, i)
		}

		// Add the newly created line to the list of lines (another row in the 2D array)
		lines = append(lines, splLineInt)
	}

	return lines, nil
}
