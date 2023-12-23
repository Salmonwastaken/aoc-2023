package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type game struct {
	hand  string
	bid   int
	htype int
}

type Games []*game

/*
Five of a kind, where all five cards have the same label: AAAAA
Four of a kind, where four cards have the same label and one card has a different label: AA8AA
Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
High card, where all cards' labels are distinct: 23456
*/
func determineType(mainCardCount int, secondaryCardCount int) int {
	var htype int
	switch mainCardCount {
	case 5:
		// Five of a kind
		htype = 7
	case 4:
		// Four of a kind
		htype = 6
	case 3:
		if secondaryCardCount == 2 {
			// Full house
			htype = 5
		} else {
			// Three of a kind
			htype = 4
		}
	case 2:
		if secondaryCardCount == 2 {
			// Two pair
			htype = 3
		} else {
			// One pair
			htype = 2
		}
	default:
		// High card
		htype = 1
	}

	// fmt.Println()

	return htype
}

func cardToScore(card string) int {
	switch card {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		return 11
	case "T":
		return 10
	default:
		number, err := strconv.Atoi(card)
		if err != nil {
			log.Fatal(err)
		}
		return number
	}
}
func sortGames(games Games) Games {
	sort.SliceStable(games, func(i, j int) bool {
		if games[j].htype > games[i].htype {
			return true
		}
		if games[j].htype < games[i].htype {
			return false
		}
		return sortCards(games[j].hand, games[i].hand)
	})
	return games
}

func sortCards(hand1 string, hand2 string) bool {
	for k := 0; k < len(hand1); k++ {
		if cardToScore(string(hand1[k])) == cardToScore(string(hand2[k])) {
			continue
		}
		if cardToScore(string(hand1[k])) > cardToScore(string(hand2[k])) {
			return true
		}
		if cardToScore(string(hand1[k])) < cardToScore(string(hand2[k])) {
			return false
		}
	}
	return true
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	games := buildGameMap(scanner)

	sortGames(games)

	// Calculate all the points by multiplying the bid by the respective ranking
	var finalPoints int
	for i := 0; i < len(games); i++ {
		fmt.Println(games[i].hand, games[i].bid)
		finalPoints += games[i].bid * (i + 1)
	}
	fmt.Println(finalPoints)
}

func buildGameMap(scanner *bufio.Scanner) Games {
	var games Games
	var gameNumber int

	// Build map of games
	for scanner.Scan() {
		gameNumber++
		currentLine := scanner.Text()
		number, err := strconv.Atoi(strings.Split(currentLine, " ")[1])
		if err != nil {
			log.Fatal(err)
		}

		hand := strings.Split(currentLine, " ")[0]

		var mainCardCount int
		var secondaryCardCount int
		localHand := hand
		for _, card := range localHand {
			count := strings.Count(localHand, string(card))
			if count >= mainCardCount {
				secondaryCardCount = mainCardCount
				mainCardCount = count
			}
			if count > secondaryCardCount && count < mainCardCount {
				secondaryCardCount = count
			}
			// Remove any values of the current card to prevent double processing
			localHand = strings.Replace(localHand, string(card), "", -1)
		}

		htype := determineType(mainCardCount, secondaryCardCount)

		games = append(games, &game{
			hand:  hand,
			bid:   number,
			htype: htype,
		})
	}
	return games
}

func buildGameMap2(scanner *bufio.Scanner) Games {
	var games Games
	var gameNumber int

	// Build map of games
	for scanner.Scan() {
		gameNumber++
		currentLine := scanner.Text()
		number, err := strconv.Atoi(strings.Split(currentLine, " ")[1])
		if err != nil {
			log.Fatal(err)
		}

		hand := strings.Split(currentLine, " ")[0]

		var mainCardCount int
		var secondaryCardCount int
		localHand := hand
		var jokerCounter int
		for _, card := range localHand {
			if card == 'J' {
				jokerCounter++
				continue
			}

			count := strings.Count(localHand, string(card))
			if count >= mainCardCount {
				secondaryCardCount = mainCardCount
				mainCardCount = count
			}
			if count > secondaryCardCount && count < mainCardCount {
				secondaryCardCount = count
			}
			// Remove any values of the current card to prevent double processing
			localHand = strings.Replace(localHand, string(card), "", -1)
		}
		if mainCardCount != 5 {
			mainCardCount += jokerCounter
		}

		htype := determineType(mainCardCount, secondaryCardCount)

		games = append(games, &game{
			hand:  hand,
			bid:   number,
			htype: htype,
		})
	}
	return games
}

func cardToScore2(card string) int {
	switch card {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		return 1
	case "T":
		return 10
	default:
		number, err := strconv.Atoi(card)
		if err != nil {
			log.Fatal(err)
		}
		return number
	}
}

func sortCards2(hand1 string, hand2 string) bool {
	for k := 0; k < len(hand1); k++ {
		if cardToScore2(string(hand1[k])) == cardToScore2(string(hand2[k])) {
			continue
		}
		if cardToScore2(string(hand1[k])) > cardToScore2(string(hand2[k])) {
			return true
		}
		if cardToScore2(string(hand1[k])) < cardToScore2(string(hand2[k])) {
			return false
		}
	}
	return true
}

func sortGames2(games Games) Games {
	sort.SliceStable(games, func(i, j int) bool {
		if games[j].htype > games[i].htype {
			return true
		}
		if games[j].htype < games[i].htype {
			return false
		}
		return sortCards2(games[j].hand, games[i].hand)
	})
	return games
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	games := buildGameMap2(scanner)

	sortGames2(games)

	// Calculate all the points by multiplying the bid by the respective ranking
	var finalPoints int
	for i := 0; i < len(games); i++ {
		fmt.Println(games[i].hand, games[i].bid)
		finalPoints += games[i].bid * (i + 1)
	}
	fmt.Println(finalPoints)
}

func main() {
	// part1()
	part2()
}
