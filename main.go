
package main

import (
	"fmt"
	"net/http"
	"time"
)

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}


type triangle struct{}
type square struct{}

type shape interface {
	getArea() float64
}

type logWriter struct{}

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

	//colors := map[string]string {
	//	"red": "#ff0000",
	//	"green": "#4bf745",
	//	"white": "ffffff",
	//
	//}



	////var colors map[string]string
	//
	//colors := make(map[int]string)
	//
	//
	//colors[10] = "#ffffff"
	//
	//
	//delete(colors, 10)

	//printMap(colors)

	//eb := englishBot{}
	//sb := spanishBot{}
	//
	//printGreeting(eb);
	//printGreeting(sb);

	//resp, err := http.Get("http://google.com")
	//if err != nil {
	//	fmt.Println("error:", err)
	//	os.Exit(1)
	//}

	//bs := make([] byte, 99999)
	//
	//resp.Body.Read(bs)
	//
	//fmt.Println(string(bs))

	//lw := logWriter{}
	//
	////io.Copy(os.Stdout, resp.Body)
	//
	//io.Copy(lw, resp.Body)


	links := []string {
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}


	//for i := 0; i < len(links); i++ {
	//	fmt.Println(<-c)
	//}

	//for  {
	//	go checkLink(<-c, c)
	//}

	for l:= range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + "might be down!")
		c <- "Might be down I think"
		return
	}

	fmt.Println(link, "is up!")
	//c <- "Yep its up"
	c <- link

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

func printGreeting(b bot)  {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return  "Hola!"
}

func (logWriter) Write(bs []byte)  (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}

func (t triangle) getArea() float64 {
	return  10.0
}

func (s square) getArea() float64 {
	return 11.0
}