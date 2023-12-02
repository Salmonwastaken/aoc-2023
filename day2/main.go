package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
)

type cubes struct {
	red   int
	green int
	blue  int
}

func main() {
	part1()
}

func part1() {
	availableCubes := cubes{red: 12, green: 13, blue: 14}
	var totalNumber int

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	re := regexp.MustCompile(`^Game (?P<game>\d+)|(?P<red>\d+) red|(?P<green>\d+) green|(?P<blue>\d+) blue`)
	scanner := bufio.NewScanner(file)
lineLoop:
	for scanner.Scan() {
		out := re.FindAllStringSubmatch(scanner.Text(), -1)
		gameNumber, err := strconv.Atoi(out[0][1])
		if err != nil {
			log.Fatal(err)
		}

		for _, match := range out {
			for _, c := range []string{"red", "green", "blue"} {
				if match[re.SubexpIndex(c)] == "" {
					continue
				}
				number, err := strconv.Atoi(match[re.SubexpIndex(c)])
				if err != nil {
					log.Fatal(err)
				}
				if number > getField(&availableCubes, c) {
					continue lineLoop
				}
			}
		}
		totalNumber += gameNumber
	}
	fmt.Println(totalNumber)
}

func part2() {
	return
}

func getField(cube *cubes, c string) int {
	r := reflect.ValueOf(cube)
	f := reflect.Indirect(r).FieldByName(c)
	return int(f.Int())
}
