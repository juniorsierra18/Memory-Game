package views

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	players        int
	numDeCards  int
	cardButtons    []*widget.Button
	poinsPlayer1   *widget.Label
	poinsPlayer2   *widget.Label
)

func startGame(w fyne.Window, numPlayers, numCards int) {
	players = numPlayers
	numDeCards = numCards

	// Etiqueta para el titulo
	titulo := canvas.NewText("Memory Game", color.White)
	titulo.Alignment = fyne.TextAlignCenter
	titulo.TextSize = 24

	cardButtons = make([]*widget.Button, numDeCards)
	cardContainer := container.NewGridWithColumns(4)

	for i := 0; i < numDeCards; i++ {
		cardButtons[i] = widget.NewButton("", func() {
			// Voltear la carta al hacer clic
		})
		cardButtons[i].Importance = widget.HighImportance
		cardContainer.Add(cardButtons[i]) // Añadir el botón al contenedor
	}

	poinsPlayer1 = widget.NewLabel("Puntuaacion Jugador 1: 0")
	poinsPlayer2 = widget.NewLabel("Puntuaacion Jugador 2: 0")

	poinsContainer := container.NewVBox(poinsPlayer1, poinsPlayer2)

	botonVolver := widget.NewButton ("Volver", func ()  {
		// Volver a NumberCards
		NumberCards(w, players)
	})

	botonVolver.Importance = widget.DangerImportance

	w.SetContent(container.NewVBox(
		titulo,
		cardContainer,
		poinsContainer,
		botonVolver,
	))
}