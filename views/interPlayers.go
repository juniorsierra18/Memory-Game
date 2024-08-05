package views

import (
	"image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func InterPlayers(w fyne.Window) {
	// Titulo
	titulo := canvas.NewText("Seleccione el modo de juego", color.White)
	titulo.Alignment = fyne.TextAlignCenter
	titulo.TextSize = 24

	// Botones de modos de juego
	botonUnJugador := widget.NewButton("1 Jugador", func() {
		// Acción para 1 jugador
		NumberCards(w)
	})
	botonDosJugadores := widget.NewButton("2 Jugadores", func() {
		// Acción para 2 jugadores
	})
	botonVolver := widget.NewButton("Volver", func() {
		// Volver a la ventana anterior
		MainMenu(w)
	})

	// Estilo de los botones
	botonUnJugador.Importance = widget.HighImportance
	botonDosJugadores.Importance = widget.HighImportance
	botonVolver.Importance = widget.DangerImportance

	// Contenedor
	content := container.NewVBox(
		titulo,
		botonUnJugador,
		botonDosJugadores,
		botonVolver,
	)

	// Centra el contenido
	contentCenter := container.NewCenter(content)
	// Agrega el contenido a la ventana
	w.SetContent(contentCenter)
}