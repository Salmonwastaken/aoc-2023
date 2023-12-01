package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	validNums := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9",
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	// Part1
	fmt.Println(calc(validNums[:9], file))
	// Part2
	fmt.Println(calc(validNums, file))

}

func calc(validNums []string, file []byte) (result int) {
	firstRe := regexp.MustCompile(`(` + strings.Join(validNums, "|") + `)`)
	lastRe := regexp.MustCompile(`.*` + firstRe.String())

	for _, s := range strings.Fields(string(file)) {
		result += 10 * (slices.Index(validNums, firstRe.FindStringSubmatch(s)[1])%9 + 1)
		result += slices.Index(validNums, lastRe.FindStringSubmatch(s)[1])%9 + 1
	}

	return result
}
