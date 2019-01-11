
package main

func main()  {

	//var card string = "Ace of Spades"

	//cards := newDeck()
	//
	//cards.print()
	//
	//hand, remainingCards := deal(cards, 5)
	//
	//hand.print()
	//remainingCards.print()

	//greeting := "Hi there!"
	//fmt.Println([]byte (greeting))

	//cards := newDeck()
	//cards.saveToFile("my_cards")
	//cards := newDeckFromFile("my_cards1")
	//fmt.Println(cards.toString())

	cards := newDeck()
	cards.shuffle()
	cards.print()

}



