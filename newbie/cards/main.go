package main

func main() {
	// cards := newDeck()
	// fmt.Println(cards)

	// hand, remainingCards := deal(cards, 5)

	// fmt.Println(hand)
	// fmt.Println(remainingCards)

	// hello := "Hasan"

	// fmt.Println([]byte(hello))

	// cards := newDeck()
	// cards.saveToFile("demo")

	d := newDeckFromFile("demo")

	d.print()

	d.shuffle()

	d.print()


}


