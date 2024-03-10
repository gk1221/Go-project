package main

import "fmt"


func main() {
	
	card := []string{"ACe", newCard()}
	card = append(card, "CAAA")
	
	for i, card := range card {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Fiver of Diamonds"
}