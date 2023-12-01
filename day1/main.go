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

	re := regexp.MustCompile(`\d*`)

	for scanner.Scan() {
		match := re.FindAll(scanner.Bytes(), 2)

		for _, list := range match {
			if len(list) > 0 {
				number, err := strconv.Atoi(fmt.Sprintf("%d", list[0], list[len(list)-1]))
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%d\n", number)
				totalNumber = totalNumber + number
			}
		}
	}

	fmt.Printf("\n%d", totalNumber)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
