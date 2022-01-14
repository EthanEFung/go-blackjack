package main

import (
	"time"

	"github.com/ethanefung/deck"
)

type BlackJack struct {
	deck   *[]deck.Card
	dealer *Hand
	table  *[]*Hand
	index *int
	active bool
}

type Hand []deck.Card
func (h *Hand) Value() (int, int) {
	var a, b int
	for _, card := range (*h) {
		if card.Type == deck.Face {
			a += 10
			b += 10
		} else if card.Type == deck.Numerical {
			a += int(card.Rank)
			b += int(card.Rank)
		} else if card.Type == deck.High {
			a += 11
			b += 1
		}
	}
	return a, b
}
func (h *Hand) HasAce() bool {
	for _, card := range (*h) {
		if card.Rank == deck.Ace {
			return true
		}
	}
	return false
}

func (b BlackJack) Sit(player *Hand) {
	*b.table = append(*b.table, player)
}

func (b BlackJack) Deal() {
	for i := 0; i < 2; i++ {
		for _, hand := range *b.table {
			b.Hit(hand)
		}
	}
}

func (b BlackJack) Hit(hand *Hand) {
	*hand = append(*hand, (*b.deck)[0])
  *b.deck = (*b.deck)[1:]
}

func (b BlackJack) Stand() {
	(*b.index) = *b.index + 1
}

func (b BlackJack) Bust() {
	(*b.index) = *b.index + 1
}

type NewOptions struct{}

func New(options NewOptions) BlackJack {
	d := deck.New(deck.NewOptions{
		Decks:   3,
		Shuffle: 1,
		Seed:    time.Now().UnixNano(),
		Jokers:  0,
	})
	var dealer Hand
	var table []*Hand
	index := 0
	return BlackJack{
		deck:  &d,
		dealer: &dealer,
		table:  &table,
		index: &index,
	}
}
