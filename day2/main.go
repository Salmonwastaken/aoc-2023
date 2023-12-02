package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
}

func part1() {
	var possibleAmount int
	var gameNumber int

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		gameNumber++
		l := strings.Split(scanner.Text(), ": ")[1]
		p := strings.Split(l, "; ")

		possible := true
		for _, g := range p {
			c := strings.Split(g, ", ")
			for _, e := range c {
				n := strings.Split(e, " ")
				num, err := strconv.Atoi(n[0])
				if err != nil {
					log.Fatal(err)
				}
				if n[1] == "blue" && num > 14 {
					possible = false
				}
				if n[1] == "green" && num > 13 {
					possible = false
				}
				if n[1] == "red" && num > 12 {
					possible = false
				}
			}
		}
		if possible {
			possibleAmount += gameNumber
		}
	}
	fmt.Println(possibleAmount)
}

func part2() {
	return
}
