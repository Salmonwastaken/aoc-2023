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

	var totalNumber int

	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`[0-9]`)

	for scanner.Scan() {
		match := re.FindAll(scanner.Bytes(), -1)
		if len(match) == 0 {
			continue
		}

		byte := append(match[0], match[len(match)-1]...)
		number, err := strconv.Atoi(string(byte))
		if err != nil {
			log.Fatal(err)
		}
		totalNumber = totalNumber + number
	}

	fmt.Printf("\n%d", totalNumber)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
