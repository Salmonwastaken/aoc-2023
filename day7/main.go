package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type game struct {
	hand  string
	bid   int
	score int
}

type Games []*game

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var games Games
	var gameNumber int
	// Ordered by relative value, eg 2 is worth less than 3 and A is worth more than K
	possibleCards := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}

	// Build map of games
	for scanner.Scan() {
		gameNumber++
		currentLine := scanner.Text()
		number, err := strconv.Atoi(strings.Split(currentLine, " ")[1])
		if err != nil {
			log.Fatal(err)
		}
		games = append(games, &game{
			hand:  strings.Split(currentLine, " ")[0],
			bid:   number,
			score: 0,
		})
	}

	// Determine type (Assign score amount based on type?)
	// See how many hits we have
	// Determine wins inbetween type?

	// Every rank is worth 70 points
	// Inbetween ranks, get points for each seperate card as well.
	// The max you can get for just the cards is 13+13+13+13+13=65, so this way getting the cards should never trump a rank

	// Build all scores, rank + first card
	for i := 0; i < len(games); i++ {
		var maxCount int
		for _, card := range possibleCards {
			count := strings.Count(games[i].hand, card) * 70

			if count > maxCount {
				maxCount = count
			}
		}
		firstCard := string(games[i].hand[0])

		firstCardScore := slices.Index(possibleCards, firstCard)
		maxCount += firstCardScore
		games[i].score = maxCount
	}

	// Figure out which entries have duplicate values

	// // iterate again and resolve duplicate scores
	// for i := 0; i < len(games); i++ {

	// 	fmt.Println(firstCardScore)
	// 	// games[i].score += firstCard
	// }

	// Sort by score, essentially assigning rank.
	// Still allows for duplicates
	sort.Slice(games, func(i, j int) bool {
		return games[i].score < games[j].score
	})

	for i := 0; i < len(games); i++ {
		fmt.Println(i, " - ", games[i].score)
	}

	// Assign rank
	// multiply bid by rank
	// add it alltogether

}
