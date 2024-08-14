package main

//Querido colega programador:
//Cuando escribí este código, sólo Dios y yo
//sabíamos cómo funcionaba.
//Ahora, sólo Dios lo sabe!
//Así que si está tratando de 'optimizarlo'
//y fracasa (seguramente), por favor,
//incremente el contador a continuación
//como una advertencia para su siguiente colega:
//total_horas_perdidas_aquí = 10

import (
	"myApp/views"

	"fyne.io/fyne/v2/app"
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
