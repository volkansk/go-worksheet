package main

import "fmt"

func main() {
	//Variable Declarations
	// var card string = "Ace of Spades"
	// card = "Five of Diamonds"
	card := "Ace of Spades"
	card = "Five of Diamonds"

	card = newCard()
	fmt.Println(card)

	//slice
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	cards := newDeck()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
	cards.print()

	//conversion
	greeting := "Hi there!"
	fmt.Println([]byte(greeting))

	fmt.Println("now save to file cards2")
	cards2 := newDeck()
	cards2.saveToFile("my_cards")

	fmt.Println("now print cards3")
	cards3 := newDeckFromFile("my_cards")
	cards3.print()

	fmt.Println("now suffle new deck")
	cards4 := newDeck()
	cards4.suffle()
	cards4.print()
}

func newCard() string {
	return "Five of Diamonds"
}
