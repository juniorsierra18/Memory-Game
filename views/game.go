package views

import (
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	players      int
	numDeCards   int
	cardButtons  []*widget.Button
	poinsPlayer1 *widget.Label
	poinsPlayer2 *widget.Label
	cards        []*card
)

func startGame(w fyne.Window, numPlayers, numCards int) {
	players = numPlayers
	numDeCards = numCards

	// Reiniciar el estado global del juego
	cards = nil

	// Sembrar la aleatoriedad
	rand.Seed(time.Now().UnixNano())

	// Generar las cartas
	for i := 0; i < numDeCards/2; i++ {
		cards = append(cards, newCard(i, nil))
		cards = append(cards, newCard(i, nil))
	}

	// Mezclar las cartas
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })

	bg := canvas.NewRectangle(color.Black)
	bg.Resize(fyne.NewSize(500, 500))

	// Etiqueta para el titulo
	titulo := canvas.NewText("Memory Game", color.White)
	titulo.Alignment = fyne.TextAlignCenter
	titulo.TextSize = 50

	//cardButtons = make([]*widget.Button, numDeCards)
	cardContainer := container.NewGridWithColumns(4)

	var flippedCards []*card

	flipFunc := func(c *card) {
		if len(flippedCards) == 0 {
			c.flip()
			flippedCards = append(flippedCards, c)
		} else if len(flippedCards) == 1 {
			// Verificar que la carta no sea la misma que la primera
			if flippedCards[0] == c {
				// Si es la misma carta, no hacer nada
				return
			}

			c.flip()
			flippedCards = append(flippedCards, c)

			// Verificar si las dos cartas coinciden
			if checkMatch(flippedCards[0], flippedCards[1]) {
				// Las cartas coinciden, deshabilitarlas
				time.Sleep(300 * time.Millisecond)
				flippedCards[0].button.Disable()
				time.Sleep(400 * time.Millisecond)
				flippedCards[1].button.Disable()
				flippedCards = nil
			} else {
				// Las cartas no coinciden, esperar antes de voltearlas de nuevo
				go func(cardsToFlip []*card) {
					time.Sleep(500 * time.Millisecond) // Espera 0.5 segundos antes de voltear las cartas
					for _, card := range cardsToFlip {
						card.flip()
					}
					flippedCards = nil // Restablecer la lista de cartas volteadas
				}(flippedCards)
			}
		}
	}

	// Asignar la función flipFunc a cada carta y añadirlas al contenedor
	for _, card := range cards {
		card.button.OnTapped = func() {
			flipFunc(card)
		}
		cardContainer.Add(card.button)
	}

	poinsPlayer1 = widget.NewLabel("Puntuacion Jugador 1: 0")
	poinsPlayer2 = widget.NewLabel("Puntuacion Jugador 2: 0")

	poinsContainer := container.NewVBox(poinsPlayer1, poinsPlayer2)

	botonVolver := widget.NewButton("Volver", func() {
		time.Sleep((500) * time.Millisecond)
		NumberCards(w, players) // Reiniciar el juego
	})

	botonVolver.Importance = widget.DangerImportance

	content := container.NewVBox(
		titulo,
		cardContainer,
		poinsContainer,
		botonVolver,
	)

	// Centra el contenido
	contentCenter := container.New(layout.NewStackLayout(), bg, container.NewCenter(content))
	// Agrega el contenido a la ventana
	w.SetContent(contentCenter)
}
