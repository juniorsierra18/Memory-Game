package main

import (
	"fyne.io/fyne/v2/app"
	"myApp/views"
)

func main() {
	// Crea la app
	a := app.New()
	// Crea la ventana y le pone el nombre de esta
	w := a.NewWindow("Memory Game")

	// Inicializa la interfaz principal
	views.MainMenu(w)

	// Corre la aplicaion
	w.ShowAndRun()
}