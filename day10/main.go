package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

/*
*

	| is a vertical pipe connecting north and south.
	- is a horizontal pipe connecting east and west.
	L is a 90-degree bend connecting north and east.
	J is a 90-degree bend connecting north and west.
	7 is a 90-degree bend connecting south and west.
	F is a 90-degree bend connecting south and east.
	. is ground; there is no pipe in this tile.
	S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.

*
*/

type location struct {
	x, y int
}

type chart [][]string

func nextLocation(pipeType rune) location {
	switch pipeType {
	case '|':
		return location{
			x: 1,
			y: 0,
		}
	}

	return location{}
}

func parsePipe(location) rune {
	return '0'
}

func determineStartLocation(row []string) int {
	return slices.Index(row, "S")
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	pipeMap := make(chart, 0)
	rowNum := 0

	for scanner.Scan() {
		pipeMap = append(pipeMap, make([]string, 0))
		for _, v := range scanner.Text() {
			pipeMap[rowNum] = append(pipeMap[rowNum], string(v))
		}
		rowNum++
	}
	for k, v := range pipeMap {
		l := determineStartLocation(v)
		if l >= 0 {
			fmt.Println("The starting location is: ", k, " - ", l)
			fmt.Println(pipeMap[k][l])
		}
	}
}
