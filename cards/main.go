package main

func main() {
	d := newDeck()
	d.shuffle()
	d.print()
	// // // hand, remainder := deal(d, 5)
	// // // hand.print()
	// // // remainder.print()
	// // d.saveToFile("my_cards")
	// d, err := newDeckFromFile("my_cards")
	// if err != nil {
	// 	fmt.Println("Could not find the deck file: ", err)
	// 	os.Exit(1)
	// }
	// d.print()
}
