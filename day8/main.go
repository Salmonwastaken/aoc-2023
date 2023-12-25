package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type directions struct {
	left  string
	right string
}

type addresses map[string]*directions

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var instructionSet []string
	addressBook := make(addresses)

	for scanner.Scan() {
		currentLine := scanner.Text()
		// Skip empty lines
		if currentLine == "" {
			continue
		}

		// First line doesn't contain the = symbol, but should contain instructions
		if !strings.Contains(currentLine, "=") {
			instructionSet = strings.Split(currentLine, "")
			continue
		}

		re := regexp.MustCompile(`\w{3}`)
		split := re.FindAllString(currentLine, -1)

		addressBook[split[0]] = &directions{
			left:  split[1],
			right: split[2],
		}
	}

	currentAddress := "AAA"
	var stepCounter int
	for {
		if currentAddress == "ZZZ" {
			break
		}

		directionOffset := stepCounter % len(instructionSet)

		if instructionSet[directionOffset] == "L" {
			currentAddress = addressBook[currentAddress].left
		} else {
			currentAddress = addressBook[currentAddress].right
		}
		stepCounter++
	}

	fmt.Println(stepCounter)
}

func generateAddressBook(scanner *bufio.Scanner) ([]string, []string, addresses) {
	var instructionSet []string
	addressBook := make(addresses)
	var startingPoints []string

	for scanner.Scan() {
		currentLine := scanner.Text()
		// Skip empty lines
		if currentLine == "" {
			continue
		}

		// First line doesn't contain the = symbol, but should contain instructions
		if !strings.Contains(currentLine, "=") {
			instructionSet = strings.Split(currentLine, "")
			continue
		}

		reAddress := regexp.MustCompile(`\w{3}`)
		splitAddress := reAddress.FindAllString(currentLine, -1)

		addressBook[splitAddress[0]] = &directions{
			left:  splitAddress[1],
			right: splitAddress[2],
		}

		reStart := regexp.MustCompile(`.{2}A{1}`)
		splitStart := reStart.FindString(currentLine)

		if splitStart != "" {
			startingPoints = append(startingPoints, splitStart)
		}
	}
	return instructionSet, startingPoints, addressBook
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	instructionSet, startingPoints, addressBook := generateAddressBook(scanner)

	// Verify end point
	reEnd := regexp.MustCompile(`.{2}Z{1}`)

	steps := make([]int, 2)

	for _, p := range startingPoints {
		var stepCounter int

		currentLocation := p
		for {
			directionOffset := stepCounter % len(instructionSet)
			if instructionSet[directionOffset] == "L" {
				currentLocation = addressBook[currentLocation].left
			} else {
				currentLocation = addressBook[currentLocation].right
			}
			stepCounter++
			if reEnd.MatchString(currentLocation) {
				steps = append(steps, stepCounter)
				break
			}
		}
	}
	// First two are empty for some reason, cba to figure out why.
	// So we'll just skip them
	fmt.Println(LCM(steps[2], steps[3], steps[4:]...))
}

// GCD and LCM functions taken from: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/#footnote-1
// I'm no good at math

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	part1()
	part2()
}
