package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Card struct {
	Suit  rune
	Value int
}

var cards = [13]string{"A", "K", "Q", "J", "10", "9", "8", "7", "6", "5", "4", "3", "2"}
var court = [4]string{"A", "K", "Q", "J"}
var numbered = [9]string{"10", "9", "8", "7", "6", "5", "4", "3", "2"}
var suits = [4]string{"C", "D", "H", "S"}

type Hand []Card

func (h Hand) evaluate() string {
	for i, card := range h {

	}

	return ""
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your poker hand: ")
	scanner.Scan()
	input := scanner.Text()

	fmt.Println("You entered: ", input)

}

func stringToHand(input string) {
	cards := [5]Card{}
	for i, char := range input {
		for i%2 == 0 {
			cards[i].Suit = char
		}
		for i%2 == 1 {
			num, err := strconv.Atoi(char)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			cards[i].Value = num
		}
	}
}

func rankToValue(r rune) int {
	switch r {
	case 'T':
		return 10
	case 'J':
		return 11
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	default:
		return int(r - '0') // rune to int conversion
	}
}

/*
func royal_flush(input string) bool {
// A-K-Q-J-T in the same suit
}
func straight_flush(input string) bool {
// five consecutive in the same suit
}
func four_of_a_kind(input string) bool {
// four cards of same rank
}
func three_of_a_kind(input string) bool {
// three cards of same rank
}
func full_house(input string) bool {
// 3 cards of one rank and 2 of another
}
func flush(input string) bool {
// five cards in same suit in any order
}
func two_pair(input string) bool {
// two cards of one rank and two of another
}
func one_pair(input string) bool {
//two card of one rank
}
func high_card(input string) bool {
// highest card if no other combination
}
*/
