package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type almaMap map[string]map[int]int

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

func transformSeed(seed int, m almaMap) int {
	stages := []string{
		"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water",
		"water-to-light", "light-to-temperature", "temperature-to-humidity",
		"humidity-to-location",
	}

	location := seed
	for _, stage := range stages {
		if val, ok := m[stage][location]; ok {
			location = val
		}
	}

	return location
}

func main() {
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
			m[category] = make(map[int]int)
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

		destinationRangeStart := lineNumbers[0]
		sourceRangeStart := lineNumbers[1]

		totalValues := lineNumbers[2]
		for i := 0; i < totalValues; i++ {
			m[category][sourceRangeStart+i] = destinationRangeStart + i
		}
	}

	answer := processSeeds(seeds, m)
	fmt.Println(answer)
}
