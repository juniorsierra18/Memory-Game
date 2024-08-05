package views

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func TwentyCards(w fyne.Window) {

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
	})
	carta9 := widget.NewButton ("9", func ()  {
		// Interfaz con 8 cartas
	})
	carta10 := widget.NewButton ("10", func ()  {
		// Interfaz con 8 cartas
	})
	carta11 := widget.NewButton ("11", func ()  {
		// Interfaz con 8 cartas
	})
	carta12 := widget.NewButton ("12", func ()  {
		// Interfaz con 8 cartas
	})
	carta13 := widget.NewButton ("13", func ()  {
		// Interfaz con 8 cartas
	})
	carta14 := widget.NewButton ("14", func ()  {
		// Interfaz con 8 cartas
	})
	carta15 := widget.NewButton ("15", func ()  {
		// Interfaz con 8 cartas
	})
	carta16 := widget.NewButton ("16", func ()  {
		// Interfaz con 8 cartas
	})
	carta17 := widget.NewButton ("17", func ()  {
		// Interfaz con 8 cartas
	})
	carta18 := widget.NewButton ("18", func ()  {
		// Interfaz con 8 cartas
	})
	carta19 := widget.NewButton ("19", func ()  {
		// Interfaz con 8 cartas
	})
	carta20 := widget.NewButton ("20", func ()  {
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
	carta9.Importance = widget.HighImportance
	carta10.Importance = widget.HighImportance
	carta11.Importance = widget.HighImportance
	carta12.Importance = widget.HighImportance
	carta13.Importance = widget.HighImportance
	carta14.Importance = widget.HighImportance
	carta15.Importance = widget.HighImportance
	carta16.Importance = widget.HighImportance
	carta17.Importance = widget.HighImportance
	carta18.Importance = widget.HighImportance
	carta19.Importance = widget.HighImportance
	carta20.Importance = widget.HighImportance
	botonVolver.Importance = widget.DangerImportance

	grid := container.NewGridWithColumns(4,
		container.NewVBox(carta1, carta5, carta9, carta13, carta17),
		container.NewVBox(carta2, carta6, carta10, carta14, carta18),
		container.NewVBox(carta3, carta7, carta11, carta15, carta19),
		container.NewVBox(carta4, carta8, carta12, carta16, carta20),
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