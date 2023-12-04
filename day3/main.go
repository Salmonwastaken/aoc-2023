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
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`\d+`)
	symbolRe := regexp.MustCompile(`[\x21-\x2D\x2F\x3A-\x40\x5B-\x60\x7B-\x7E]`)
	m := make(map[int]string)
	answerMap := make(map[int]int)
	var linenumber int
	var answer int

	// Build map of available lines
	for scanner.Scan() {
		linenumber++
		m[linenumber] = scanner.Text()
	}

	// For every line, determine wether they contain special symbols
	for row, _ := range m {
		var allowedIndexRange [141][2]int
		index := symbolRe.FindAllStringIndex(m[row], -1)
		if index == nil {
			continue
		}

		for _, in := range index {
			for _, ind := range in {
				// build a map of allowed indexes per row based per symbol
				for _, x := range []int{-1, 0, 1} {
					allowedIndexRange[row+x][0] = ind - 1
					allowedIndexRange[row+x][1] = ind + 1

					allowStart := allowedIndexRange[row+x][0]
					allowEnd := allowedIndexRange[row+x][1]
					// Find the indexes of all available numbers per line
					allIndexes := re.FindAllStringIndex(m[row+x], -1)
					if allIndexes == nil {
						continue
					}

					// Compare every found indexes with the set of allowedIndexes
					for _, hit := range allIndexes {
						// fmt.Println(m[row+x][hit[0]:hit[1]])
						// fmt.Printf("Current Row: %d\n", row)
						// fmt.Printf("Current indexes: %d\n", hit)
						// fmt.Printf("Allowed indexes: [%d %d]\n", allowedIndexRange[row+x][0], allowedIndexRange[row+x][1])
						allowed := false
						numberStart := hit[0]
						numberEnd := hit[1]

						if numberStart >= allowStart || allowEnd >= numberStart || allowStart >= numberEnd || numberEnd >= allowEnd {
							fmt.Printf("%d >= %d || %d >= %d || %d >= %d || %d >= %d\n\n", numberStart, allowStart, numberStart, allowEnd, numberEnd, allowStart, numberEnd, allowEnd)
							allowed = true
						}
						if !allowed {
							continue
						}
						number, err := strconv.Atoi(m[row+x][hit[0]:hit[1]])
						if err != nil {
							log.Fatal(err)
						}
						answerMap[number] = number
					}
				}
			}
		}
	}

	for _, v := range answerMap {
		// fmt.Println(v)
		answer += v
	}

	fmt.Printf("\nThe answer is: %d\n", answer)
}
