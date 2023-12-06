package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type alma struct {
	sourceRangeStart      int
	destinationRangeStart int
	allowedRange          int
}

type almaMap map[string][]alma

func processSeeds(seeds []int, m almaMap) int {
	answer := 0

	for _, seed := range seeds {
		location := transformSeed(seed, m)

		if answer == 0 || location < answer {
			answer = location
		}
	}

	return answer
}

func processSeeds2(seeds []int, m almaMap) int {
	answer := 0

	start := true

	for k := range seeds {
		if start {
			for i := seeds[k]; i <= seeds[k]+seeds[k+1]-1; i++ {
				location := transformSeed(i, m)

				if answer == 0 || location < answer {
					answer = location
				}
			}
			start = false
		} else {
			start = true
			continue
		}
	}

	return answer
}

func transformSeed(seed int, m almaMap) int {
	stages := []string{
		"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water",
		"water-to-light", "light-to-temperature", "temperature-to-humidity",
		"humidity-to-location",
	}

	location := seed
	for _, stage := range stages {
		for _, v := range m[stage] {
			offset := 0
			if (location >= v.sourceRangeStart) && (location <= v.sourceRangeStart+v.allowedRange-1) {
				offset = (location - v.sourceRangeStart)
				location = v.destinationRangeStart + offset
				break
			}
		}
	}

	return location
}

func main() {
	// fmt.Println(part1())
	fmt.Println(part2())
}

func part2() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var newCategory bool
	var category string
	var seeds []int
	m := make(almaMap, 7)
	// Parse every line to build a map of the almanac file
	for scanner.Scan() {
		currentLine := scanner.Text()

		// If empty newline, a new category has been reached
		if currentLine == "" {
			newCategory = true
			continue
		}

		// Special handling for seeds
		if strings.Contains(currentLine, "seeds: ") {
			for _, v := range strings.Split(currentLine, " ")[1:] {
				number, err := strconv.Atoi(v)
				if err != nil {
					log.Fatal(err)
				}
				seeds = append(seeds, number)
			}
			continue
		}

		// If previous line was empty newline, make a new category
		if newCategory {
			// This strips the word map:
			category = currentLine[:strings.Index(currentLine, " map:")]
			m[category] = []alma{}
			newCategory = false
			continue
		}

		// Start building the category maps
		var lineNumbers [3]int

		// Turn available ranges into numbers
		for k, v := range strings.Split(currentLine, " ") {
			number, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			lineNumbers[k] = number
		}
		m[category] = append(m[category], alma{
			destinationRangeStart: lineNumbers[0],
			sourceRangeStart:      lineNumbers[1],
			allowedRange:          lineNumbers[2],
		})
	}

	answer := processSeeds2(seeds, m)
	return answer
}

func part1() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var newCategory bool
	var category string
	var seeds []int
	m := make(almaMap, 7)
	// Parse every line to build a map of the almanac file
	for scanner.Scan() {
		currentLine := scanner.Text()

		// If empty newline, a new category has been reached
		if currentLine == "" {
			newCategory = true
			continue
		}

		// Special handling for seeds
		if strings.Contains(currentLine, "seeds: ") {
			for _, v := range strings.Split(currentLine, " ")[1:] {
				number, err := strconv.Atoi(v)
				if err != nil {
					log.Fatal(err)
				}
				seeds = append(seeds, number)
			}
			continue
		}

		// If previous line was empty newline, make a new category
		if newCategory {
			// This strips the word map:
			category = currentLine[:strings.Index(currentLine, " map:")]
			m[category] = []alma{}
			newCategory = false
			continue
		}

		// Start building the category maps
		var lineNumbers [3]int

		// Turn available ranges into numbers
		for k, v := range strings.Split(currentLine, " ") {
			number, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			lineNumbers[k] = number
		}
		m[category] = append(m[category], alma{
			destinationRangeStart: lineNumbers[0],
			sourceRangeStart:      lineNumbers[1],
			allowedRange:          lineNumbers[2],
		})
	}

	answer := processSeeds(seeds, m)
	return answer
}
