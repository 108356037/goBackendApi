package deck

import "strconv"

type deck []string

type cardVal int

const (
	v1 cardVal = iota + 1
	v2
	v3
	v4
	v5
)

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Clubs", "Diamonds", "Hearts"}

	cardValue := []cardVal{v1, v2, v3, v4, v5}

	for _, suits := range cardSuits {
		for _, num := range cardValue {
			cards = append(cards, suits+"_"+strconv.Itoa(int(num)))
		}
	}

	return cards
}
