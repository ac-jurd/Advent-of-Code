package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Starting...")

	l, r, err := ParseFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	score := similarity_score(l, r)
	fmt.Println("Score", score)
}

func similarity_score(l []int, r []int) int {
	sort.Ints(l)
	sort.Ints(r)
	sum := 0

	for i := 0; i < len(l); i++ {
		// Find index of occurrence of l[i]
		index, err := binary_search(r, len(r), l[i])

		// An error was returned
		if err != nil {
			// The error was that the target was not found
			if strings.Contains(err.Error(), "not found") {
				continue
			} else {
				// The error was something else
				log.Fatalln(err)
			}
		}
		// At this point the index of the target value was returned

		// Find adjacent occurrences
		occ := 1 // Keep track of the number of adjacent occurrences in target array

		for j := index - 1; j >= 0; j-- {
			if r[j] == l[i] {
				occ++
			} else {
				break
			}
		}

		for j := index + 1; j < len(r); j++ {
			if r[j] == l[i] {
				occ++
			} else {
				break
			}
		}

		// fmt.Println("Adding", l[i], occ, "times")
		sum += l[i] * occ
	}

	return sum
}

func binary_search(r []int, len int, target int) (int, error) {
	li := 0
	ri := len - 1

	for li <= ri {
		mid := int(math.Floor((float64(li + ri)) / 2))
		if r[mid] < target {
			li = mid + 1
		} else if r[mid] > target {
			ri = mid - 1
		} else {
			return mid, nil
		}
	}
	return 0, errors.New("Target not found")
}

func sum_of_diffs(l []int, r []int) int {
	sort.Ints(l)
	sort.Ints(r)
	sum := 0

	for i := 0; i < len(l); i++ {
		diff := r[i] - l[i]
		if diff < 0 {
			diff *= -1
		}
		sum += diff
	}

	return sum
}

func ParseFile(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	l := make([]int, 0)
	r := make([]int, 0)

	for {
		// Read line from file
		line, _, err := reader.ReadLine()
		// If there is an error
		if err != nil {
			// If the error is EOF:
			if err.Error() == "EOF" {
				// End loop and return l & r
				break
			}
			// The error is not EOF, return error
			return nil, nil, err
		}

		// Split line string by 3 spaces
		spl := strings.Split(string(line), "   ")
		// If number of elements in split line is not exactly 2, return error
		if len(spl) != 2 {
			return nil, nil, errors.New("Error splitting file: expected 2 elements")
		}

		// Convert strings to integers
		l_int, _ := strconv.Atoi(spl[0])
		r_int, _ := strconv.Atoi(spl[1])

		// Append integers to left and right arrays
		l = append(l, l_int)
		r = append(r, r_int)
	}

	return l, r, nil
}
