package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
	var totalScore float64

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var cardScore float64
		var winningNumbers []string
		// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
		l := strings.Split(scanner.Text(), ": ")[1]
		// 41 48 83 86 17 | 83 86  6 31 17  9 48 53
		mw := strings.Split(l, " | ")
		// w(inning): 41 48 83 86 17
		// m(ine): 83 86 6 31 17 9 48 53
		w := mw[0]
		m := mw[1]

		// [41 48 83 86 17]
		// [83 86 6 31 17 9 48 53]
		wSlice := strings.Split(w, " ")
		mSlice := strings.Split(m, " ")
		for i := range mSlice {
			if mSlice[i] == "" {
				mSlice[i] = mSlice[len(mSlice)-1]
			}
		}
		for _, v := range wSlice {
			if slices.Contains(mSlice, v) {
				winningNumbers = append(winningNumbers, v)
			}
		}

		n := len(winningNumbers)

		if n == 0 {
			continue
		}

		cardScore = math.Pow(2, float64(n-1))
		totalScore += cardScore
	}
	fmt.Println(totalScore)
}
