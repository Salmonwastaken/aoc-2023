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

func main() {
	part1()
}
