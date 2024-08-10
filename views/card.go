package views

import (
	"fmt"

	"fyne.io/fyne/v2/widget"
)

type card struct {
	value   int
	flipped bool
	button  *widget.Button
}

func newCard(value int, flipFunc func(*card)) *card {
	card := &card{value: value, flipped: false}
	card.button = widget.NewButton("", func() {
		if !card.flipped {
			flipFunc(card)
		}
	})
	card.button.Importance = widget.HighImportance
	return card
}

func (c *card) flip() {
	if c.flipped {
		c.button.SetText("")
		c.flipped = false
	} else {
		c.button.SetText(fmt.Sprintf("%d", c.value))
		c.flipped = true
	}
}

func checkMatch(card1, card2 *card) bool {
	return card1.value == card2.value
}
