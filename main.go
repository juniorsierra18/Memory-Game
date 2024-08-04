package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	//Crea la app
	a := app.New()
	//Crea la ventana y le pone el nombre de esta
	w := a.NewWindow("Hello World")
	//Tamaño de la ventana ancho y alto
	w.Resize(fyne.NewSize(600, 500))
	//Texto de la ventana
	label := widget.NewLabel("Hello World")
	//Muestra el texto en la ventan
	w.SetContent(label)
	//Corre la aplicaion
	w.ShowAndRun()
}