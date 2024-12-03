package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	sum := 0
	var matches []string = readValidStrings("input.txt")

	for _, match := range matches {
		a, b := getOperands(match)
		sum += a * b
	}

	fmt.Println("Sum:", sum)
}

func getOperands(line string) (int, int) {
	re := regexp.MustCompile("\\(.*\\)")
	match := re.FindString(line)
	if match == "" {
		log.Fatalln("No match in getOperands regex find")
	}

	var operandsTrim string = strings.Trim(match, "()")
	var operandsString []string = strings.Split(operandsTrim, ",")

	a, _ := strconv.Atoi(operandsString[0])
	b, _ := strconv.Atoi(operandsString[1])
	return a, b
}

func readValidStrings(filename string) []string {
	re := regexp.MustCompile("mul\\([0-9]*,[0-9]*\\)")

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	matches := re.FindAllString(string(data), -1)
	if matches == nil {
		log.Fatalln("No matches in readValidStrings regex findall")
	}

	return matches
}
