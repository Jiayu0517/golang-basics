package main

import "fmt"

// import "fmt"

// import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card = newCard()
	// fmt.Println(card)

	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	// hand, remainingDeck := deal(cards, 5)

	// cards := newDeckFromFile("my_cards")
	// cards.print()
		
	// hand.print()
	// remainingDeck.print()

	// cards.print()

	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

	// cards := newDeck()
	// cards.shuffle()
	// cards.print()


	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, num := range nums {
		if num % 2 == 0 {
			fmt.Printf("%v Even", i)
			fmt.Println()
		} else {
			fmt.Printf("%v Odd", i)
			fmt.Println()
		}
	}

}

// func newCard() string {
// 	return "Five of Diamonds"
// }

// Go is statically typed language