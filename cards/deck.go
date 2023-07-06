package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
	"time"
	
)

// custome type decorator
type deck []string

// create a function for add new deck and returs

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValue := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}

}

// convert slice of deck into single string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// save file in local machine
func (d deck) saveToFile(fileName string) error  {
return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

// read file form disk

func newDeckFromFile(fileName string)deck {
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
	fmt.Println("Eror:", err)
	os.Exit(1)
	}

	s := strings.Split(string(bs), ",")

	return deck(s)
}

func (d deck) shuffle() {
	source:=rand.NewSource(time.Now().Unix())
	r:=rand.New(source)
	for i :=range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]

	}
}