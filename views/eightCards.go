package views

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func EightCards(w fyne.Window) {
	// Titulo
	titulo := canvas.NewText("Memory Game", color.White)
	titulo.Alignment = fyne.TextAlignCenter
	titulo.TextSize = 24

	carta1 := widget.NewButton ("1", func ()  {
		// Interfaz con 8 cartas
	})
	carta2 := widget.NewButton ("2", func ()  {
		// Interfaz con 8 cartas
	})
	carta3 := widget.NewButton ("3", func ()  {
		// Interfaz con 8 cartas
	})
	carta4 := widget.NewButton ("4", func ()  {
		// Interfaz con 8 cartas
	})
	carta5 := widget.NewButton ("5", func ()  {
		// Interfaz con 8 cartas
	})
	carta6 := widget.NewButton ("6", func ()  {
		// Interfaz con 8 cartas
	})
	carta7 := widget.NewButton ("7", func ()  {
		// Interfaz con 8 cartas
	})
	carta8 := widget.NewButton ("8", func ()  {
		// Interfaz con 8 cartas
		win(w)
		InterPlayers(w)
	})
	botonVolver := widget.NewButton ("Volver", func ()  {
		// Volver a InterPlayers
		NumberCards(w)
	})

	//Diseñp de botones
	carta1.Importance = widget.HighImportance
	carta2.Importance = widget.HighImportance
	carta3.Importance = widget.HighImportance
	carta4.Importance = widget.HighImportance
	carta5.Importance = widget.HighImportance
	carta6.Importance = widget.HighImportance
	carta7.Importance = widget.HighImportance
	carta8.Importance = widget.HighImportance
	botonVolver.Importance = widget.DangerImportance

	grid := container.NewGridWithColumns(4,
		container.NewVBox(carta1, carta5),
		container.NewVBox(carta2, carta6),
		container.NewVBox(carta3, carta7),
		container.NewVBox(carta4, carta8),
	)

	content := container.NewVBox(
		titulo,
		/*boton8Cartas,
		boton12Cartas,
		boton16Cartas,
		boton20Cartas,*/
		grid,
		botonVolver,
	)

	// Centra el contenido
	contentCenter := container.NewCenter(content)
	// Agrega el contenido a la ventana
	w.SetContent(contentCenter)
}