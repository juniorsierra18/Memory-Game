package views

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func InterPlayers(w fyne.Window) {
	bg := canvas.NewRectangle(color.Black)
	bg.Resize(fyne.NewSize(500, 500))
	// Titulo
	titulo := canvas.NewText("Seleccione el modo de juego", color.White)
	titulo.Alignment = fyne.TextAlignCenter
	titulo.TextSize = 40

	// Botones de modos de juego
	botonUnJugador := widget.NewButton("1 Jugador", func() {
		// Acción para 1 jugador
		time.Sleep((500) * time.Millisecond)
		NumberCards(w, 1)
	})
	botonDosJugadores := widget.NewButton("2 Jugadores", func() {
		// Acción para 2 jugadores
		time.Sleep((500) * time.Millisecond)
		NumberCards(w, 2)
	})
	botonVolver := widget.NewButton("Volver", func() {
		// Volver a la ventana anterior
		time.Sleep((500) * time.Millisecond)
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
	contentCenter := container.New(layout.NewStackLayout(), bg, container.NewCenter(content))
	// Agrega el contenido a la ventana
	w.SetContent(contentCenter)
}
