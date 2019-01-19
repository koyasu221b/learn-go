
package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

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

	//cards := newDeck()
	//cards.shuffle()
	//cards.print()

	//checkTheNumber();

	//alex := person {"Alex", "Anderson"}
	//alex := person {firstName: "Alex", lastName: "Anderson"}
	//fmt.Println(alex)

	//jim := person {
	//	firstName: "jim",
	//	lastName: "Party",
	//	contactInfo: contactInfo{
	//		email: "abc@gmail.com",
	//		zipCode: 94000,
	//	},
	//}
	//
	//jimPointer := &jim
	//jimPointer.updateName("jimmy")
	//jim.print()

	//jim.updateName("jimmy")
	//jim.print()


	//c := []int{1,2,3}
	//b := c[0:2]
	//
	//b[0] = 5
	//fmt.Printf("+%v", c)

	colors := map[string]string {
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "ffffff",

	}



	////var colors map[string]string
	//
	//colors := make(map[int]string)
	//
	//
	//colors[10] = "#ffffff"
	//
	//
	//delete(colors, 10)

	printMap(colors)
}

//func (p person) updateName(newFirstName string)  {
//	p.firstName = newFirstName
//}

func (pointerToPerson *person) updateName(newFirstName string)  {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print (){
	fmt.Printf("%+v", p)
}


func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

