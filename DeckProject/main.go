package main

// import (
// 	"fmt"
// )

 
// import ("fmt")

func main() {
	
	// cards := newDeck()
	// // fmt.Println(cards.toString())
	// cards.saveToFile("myCards")
	
	// cards := newDeckFromFile("myCards")
	// cards.print()
	cards := newDeck()
	cards.shuffle()
	cards.print()
}



