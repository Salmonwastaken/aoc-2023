package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// part1()
	part2()
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`\d+`)
	symbolRe := regexp.MustCompile(`\*`)
	m := make(map[int]string)
	answerMap := make(map[string]map[string]int)
	var gearMap [1000]int
	var linenumber int
	var answer int
	var allowed bool
	var gearcount int
	var idSymbol string

	// Build map of available lines
	for scanner.Scan() {
		linenumber++
		m[linenumber] = scanner.Text()
	}

	// For every line, determine wether they contain special symbols
	for row := range m {

		var allowedIndexRange [141][2]int
		index := symbolRe.FindAllStringIndex(m[row], -1)
		if index == nil {
			continue
		}

		for _, ind := range index {
			// build a map of allowed indexes per row based per symbol
			for _, x := range []int{-1, 0, 1} {
				allowedIndexRange[row+x][0] = ind[0] - 1
				allowedIndexRange[row+x][1] = ind[0] + 1

				allowStart := allowedIndexRange[row+x][0]
				allowEnd := allowedIndexRange[row+x][1]
				idSymbol = fmt.Sprintf("%d-%d:%d", row, allowStart, allowEnd)
				// Find the indexes of all available numbers per line
				allIndexes := re.FindAllStringSubmatchIndex(m[row+x], -1)
				if allIndexes == nil {
					continue
				}

				// Compare every found indexes with the set of allowedIndexes
				for _, hit := range allIndexes {
					allowed = false
					numberStart := hit[0]
					numberEnd := hit[1] - 1
					idNumber := fmt.Sprintf("%d-%d:%d", row+x, numberStart, numberEnd)
					if numberStart >= allowStart && numberStart <= allowEnd {
						// fmt.Printf("%d is greater than %d and smaller than %d\n", numberStart, allowStart, allowEnd)
						allowed = true
					} else if numberEnd >= allowStart && numberEnd <= allowEnd {
						// fmt.Printf("%d is greater than %d and smaller than %d\n", numberEnd, allowStart, allowEnd)
						allowed = true
					}
					if !allowed {
						continue
					}
					number, err := strconv.Atoi(m[row+x][hit[0]:hit[1]])
					if err != nil {
						log.Fatal(err)
					}
					if answerMap[idSymbol] == nil {
						answerMap[idSymbol] = make(map[string]int)
					}
					// if answerMap[id][row+x] == nil {
					// 	answerMap[id][row+x] = make(map[string]int)
					// }
					// Create a unique map that lets you store all numbers exactly once (based on row + location)
					// answerMap[id][row+x][fmt.Sprintf("%d:%d", numberStart, numberEnd)] = number
					answerMap[idSymbol][idNumber] = number
				}
			}
			if len(answerMap[idSymbol]) == 2 {
				fmt.Println(answerMap[idSymbol])
				gearcount++
				for _, v := range answerMap[idSymbol] {
					if gearMap[gearcount] == 0 {
						gearMap[gearcount] = v
					} else {
						gearMap[gearcount] = gearMap[gearcount] * v
					}
					// }
				}
				// fmt.Println(gearMap[gearcount])
			}
		}
	}
	// Add together all the numbers we have found
	for _, v := range gearMap {
		answer += v
	}
	fmt.Printf("\nThe answer is: %d\n", answer)
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`\d+`)
	symbolRe := regexp.MustCompile(`[\x21-\x2D\x2F\x3A-\x40\x5B-\x60\x7B-\x7E]`)
	m := make(map[int]string)
	answerMap := make(map[int]map[string]int)
	var linenumber int
	var answer int
	var allowed bool

	// Build map of available lines
	for scanner.Scan() {
		linenumber++
		m[linenumber] = scanner.Text()
	}

	// For every line, determine wether they contain special symbols
	for row := range m {
		var allowedIndexRange [141][2]int
		index := symbolRe.FindAllStringIndex(m[row], -1)
		if index == nil {
			continue
		}

		for _, ind := range index {
			// build a map of allowed indexes per row based per symbol
			for _, x := range []int{-1, 0, 1} {
				if answerMap[row+x] == nil {
					answerMap[row+x] = make(map[string]int)
				}
				allowedIndexRange[row+x][0] = ind[0] - 1
				allowedIndexRange[row+x][1] = ind[0] + 1

				allowStart := allowedIndexRange[row+x][0]
				allowEnd := allowedIndexRange[row+x][1]
				// Find the indexes of all available numbers per line
				allIndexes := re.FindAllStringSubmatchIndex(m[row+x], -1)
				if allIndexes == nil {
					continue
				}

				// Compare every found indexes with the set of allowedIndexes
				for _, hit := range allIndexes {
					allowed = false
					numberStart := hit[0]
					numberEnd := hit[1] - 1
					if numberStart >= allowStart && numberStart <= allowEnd {
						fmt.Printf("%d is greater than %d and smaller than %d\n", numberStart, allowStart, allowEnd)
						allowed = true
					} else if numberEnd >= allowStart && numberEnd <= allowEnd {
						fmt.Printf("%d is greater than %d and smaller than %d\n", numberEnd, allowStart, allowEnd)
						allowed = true
					}
					if !allowed {
						fmt.Println()
						continue
					}
					number, err := strconv.Atoi(m[row+x][hit[0]:hit[1]])
					if err != nil {
						log.Fatal(err)
					}
					// Create a unique map that lets you store all numbers exactly once (based on row + location)
					answerMap[row+x][fmt.Sprintf("%d:%d", numberStart, numberEnd)] = number
				}
				// }
			}
		}
	}

	// Add together all the numbers we have found
	for _, vmap := range answerMap {
		for _, v := range vmap {
			fmt.Println(v)
			answer += v
		}
	}

	fmt.Printf("\nThe answer is: %d\n", answer)
}
