package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	AllowedTime    int
	recordDistance int
}

type Races map[int]*Race

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
	race := Race{}
	var lineNumber int
	for scanner.Scan() {
		lineNumber++
		currentLine := scanner.Text()
		splitLine := strings.Split(currentLine, " ")
		var s string
		for _, v := range splitLine[1:] {
			if v == "" {
				continue
			}
			s += v
		}
		number, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		if lineNumber == 1 {
			race.AllowedTime = number
		} else {
			race.recordDistance = number
		}
	}

	var raceWins int
	for timeHeld := 0; timeHeld <= race.AllowedTime; timeHeld++ {
		distanceTraveled := (race.AllowedTime - timeHeld) * timeHeld
		if distanceTraveled > race.recordDistance {
			raceWins++
		}
	}

	fmt.Println(raceWins)
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	raceMap := make(Races)
	var lineNumber int
	for scanner.Scan() {
		lineNumber++
		currentLine := scanner.Text()
		splitLine := strings.Split(currentLine, " ")
		var raceNumber int
		for _, v := range splitLine[1:] {
			if v == "" {
				continue
			}
			raceNumber++
			number, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			if lineNumber == 1 {
				raceMap[raceNumber] = &Race{AllowedTime: number}
			} else {
				raceMap[raceNumber].recordDistance = number
			}
		}
	}

	var totalWins []int
	for _, v := range raceMap {
		var raceWins int
		for timeHeld := 0; timeHeld <= v.AllowedTime; timeHeld++ {
			distanceTraveled := (v.AllowedTime - timeHeld) * timeHeld
			if distanceTraveled > v.recordDistance {
				raceWins++
			}
		}
		totalWins = append(totalWins, raceWins)
	}

	fmt.Println(totalWins)
	var multiplyWins int
	for _, v := range totalWins {
		if multiplyWins == 0 {
			multiplyWins = v
		} else {
			multiplyWins = multiplyWins * v
		}
	}
	fmt.Println(multiplyWins)
}
