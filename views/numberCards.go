package views

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NumberCards(w fyne.Window) {
	// Titulo
	titulo := canvas.NewText("Cantidad De Cartas", color.White)
	titulo.Alignment = fyne.TextAlignCenter
	titulo.TextSize = 24

	boton8Cartas := widget.NewButton ("8 Cartas", func ()  {
		// Interfaz con 8 cartas
		EightCards(w)
	})

	boton12Cartas := widget.NewButton ("12 Cartas", func ()  {
		// Interfaz con 12 cartas
		TwelveCards(w)
	})
	boton16Cartas := widget.NewButton ("16 Cartas", func ()  {
		// Interfaz con 16 cartas
		SixteenCards(w)
	})
	boton20Cartas := widget.NewButton ("20 Cartas", func ()  {
		// Interfaz con 20 cartas
		TwentyCards(w)
	})
	botonVolver := widget.NewButton ("Volver", func ()  {
		// Volver a InterPlayers
		InterPlayers(w)
	})

	//Diseñp de botones
	boton8Cartas.Importance = widget.HighImportance
	boton12Cartas.Importance = widget.HighImportance
	boton16Cartas.Importance = widget.HighImportance
	boton20Cartas.Importance = widget.HighImportance
	botonVolver.Importance = widget.DangerImportance

	grid := container.NewGridWithColumns(2,
		container.NewVBox(boton8Cartas, boton16Cartas),
		container.NewVBox(boton12Cartas, boton20Cartas),
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