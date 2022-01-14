package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to blackjack")
	reader := bufio.NewReader(os.Stdin)
	game := New(NewOptions{})
	player := &Hand{}

	game.Sit(player)
	game.Sit(game.dealer)

	fmt.Println("Dealing")
	game.Deal()
	fmt.Println("Player: ", *(*game.table)[0])
	var dealerHand []interface{}
	dealerHand = append(dealerHand, (*game.dealer)[0], "===HIDDEN===")
	fmt.Println("Dealer: ", dealerHand)

	fmt.Print("Your turn: (h)it or (s)tand: ")
	for *game.index < len(*game.table) - 1 {
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		text = strings.Replace(text, "\n", "", -1)
		switch text {
		case "h":
			game.Hit(player)
			fmt.Println(*player)
			if a, b := player.Value(); a > 21 && b > 21 {
				game.Bust()
				fmt.Println("Player Busts!")
				break;
			} 
			fmt.Print("Your turn: (h)it or (s)tand: ")
		case "s":
			game.Stand()
			fmt.Println("Player Stands")
		default:
			fmt.Println("Invalid input")
		}
	}
	fmt.Println("Dealers turn: ")
	fmt.Println("Dealer has: ", (*game.dealer))
	for *game.index < len(*game.table) {
		a, b := game.dealer.Value(); 
		if (a < 16 && b < 16) || (a == 17 && game.dealer.HasAce()) {
			game.Hit(game.dealer)
			fmt.Println("Dealer hits:", *game.dealer)
		} else if a > 21 && b > 21 {
			fmt.Println("Dealer Busts!")
			game.Bust()
		} else {
			fmt.Println("Dealer Stands")
			game.Stand()
		}
	}

	a, b := player.Value()
	fmt.Println("Player has: ", (*player), a, b)
	a, b = game.dealer.Value()
	fmt.Println("Dealer has: ", *game.dealer, a, b)
}
