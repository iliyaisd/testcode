// Linked cards
// Image a set of 2.5" index cards that are a each a unique color. On the back of each card is the name of the color of the next card.
// In this way the cards have one particular order.
// Task: Create a data structure representing these cards.
// Task: Write a function that accepts a list of colors and produces these data structures in the proper order, such that if a player picked up
// only the first card (or a function took only the first card as input) they would be able to find the proper order by only looking at each card once
// (or a function would be able to traverse the cards using only the first card).
package main

import "fmt"

type CardColor string

const (
	Yellow CardColor = "yellow"
	Green  CardColor = "green"
	Blue   CardColor = "blue"
	Red    CardColor = "red"
	Empty  CardColor = ""
)

var cards = map[CardColor]CardColor{
	Yellow: Green,
	Green:  Blue,
	Blue:   Red,
	Red:    Empty,
}

func main() {
	printCardsRecursive(Yellow)
}

func printCards(firstCard CardColor) {
	fmt.Println("Cards direct print:")
	output := traceCards(firstCard)
	for _, entry := range output {
		fmt.Println(entry)
	}
}

func printCardsReverse(firstCard CardColor) {
	fmt.Println("Cards reverse print:")
	output := traceCards(firstCard)
	for i := len(output) - 1; i >= 0; i-- {
		fmt.Println(output[i])
	}
}

func traceCards(firstCard CardColor) []string {
	var output []string

	currentCard := firstCard
	for {
		nextCard, ok := cards[currentCard]
		if !ok || nextCard == Empty {
			output = append(output, fmt.Sprintf("Card of color %s with no color written\n", currentCard))
			break
		}
		output = append(output, fmt.Sprintf("Card of color %s with the color written %s\n", currentCard, nextCard))

		currentCard = nextCard
	}

	return output
}

func printCardsRecursive(currentCard CardColor) {
	nextCard, ok := cards[currentCard]
	if !ok || nextCard == Empty {
		fmt.Printf("Card of color %s with no color written\n", currentCard)
		return
	}
	printCardsRecursive(nextCard)
	fmt.Printf("Card of color %s with the color written %s\n", currentCard, nextCard)
}
