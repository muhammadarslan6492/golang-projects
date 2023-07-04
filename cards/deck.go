package main

import (
	"fmt"
	"strings"
)


// custome type decorator
type deck []string

// create a function for add new deck and returs

func newDexk() deck  {
	cards:= deck{}

	cardSuits:= []string {"Spades", "Diamonds", "Hearts","Clubs"}
	cardValue:= []string {"Ace", "Two", "Three", "Four"}

	for _, suit:=range cardSuits {
		for _, value:=range cardValue {
			cards = append(cards, value + " of " + suit)
		}		
	}
	return cards
}

func deal(d deck, handSize int)(deck, deck){
	return d[:handSize], d[handSize:]
}

func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}

}

// convert slice of deck into single string
func(d deck) toString() string {
return strings.Join([]string(d), ",")
} 