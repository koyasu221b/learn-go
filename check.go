package main

import "fmt"

func checkTheNumber()  {

	//for _, number := range(numbers) {
	for number:= 1; number <=10; number++ {
		if number % 2 == 0 {
			fmt.Printf("The number %v is even \n", number)
		}else {

			fmt.Printf("The number %v is odd \n", number)
		}
	}
}