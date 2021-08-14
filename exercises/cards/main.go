package main

func main() {
	cards := newDeck()

	cards.shuffle()
	cards.print()
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()
	// cards.saveToFile("my_cards")
}
