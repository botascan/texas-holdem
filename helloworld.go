package main
import("fmt")

var cards char =[12]char{'A','K','Q','J','10','9','8','7','6','5','4','3','2'}
var suits char = [3]char{'C','S','A','D'}
func royal_flush() bool {
// A-K-Q-J-T in the same suit
}
func straight_flush() bool {
// five consecutive in the same suit
}
func four_of_a_kind() bool {
// four cards of same rank
}
func three_of_a_kind() bool {
// three cards of same rank
}
func full_house() bool {
// 3 cards of one rank and 2 of another
}
func flush() bool {
// five cards in same suit in any order
}
func two_pair() bool {
// two cards of one rank and two of another
}
func one_pair() bool {
//two card of one rank
}
func high_card() bool {
// highest card if no other combination
}

func main() {
	fmt.Println("Hello World!")
}
