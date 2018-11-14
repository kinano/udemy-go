package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// Creates a deck type (slice of strings)
type deck []string

func newDeck() deck {
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	cards := deck{}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// d is a receiver on a func
// Any deck type gets access to the print method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Return two decks
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(fileName string) error {
	bytes := []byte(d.toString())
	return ioutil.WriteFile(fileName, bytes, 0444)
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func newDeckFromFile(fileName string) (deck, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return deck(strings.Split(string(b), ",")), err
}

func (d deck) shuffle() {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
