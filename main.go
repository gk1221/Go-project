package main

//go run main.go deck.go
func main() {
	
	card := deck{"ACe", newCard()}
	card = append(card, "CAAA")
	
	card.print()
}

func newCard() string {
	return "Fiver of Diamonds"
}