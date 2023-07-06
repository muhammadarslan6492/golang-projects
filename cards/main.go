package main

func main() {


	cards := newDeck()

	cards.shuffle()

	cards.print()

	// cards := newDeck()

	// cards.saveToFile("my_card")
	// cards:= newDeckFromFile("my_card")

	// cards.print()

}
