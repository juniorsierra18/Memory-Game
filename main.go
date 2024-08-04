package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	//Crea la app
	a := app.New()
	//Crea la ventana y le pone el nombre de esta
	w := a.NewWindow("Memory Game")

	// Etiqueta para el titulo
	titulo := canvas.NewText("Memory Game", color.White)
	titulo.Alignment = fyne.TextAlignCenter
	titulo.TextSize = 24


	//Botones de Inicio
	botonJugar := widget.NewButton("Jugar", func ()  {
		//Pasar a la siguiente ventana
	})

	botonIntrucciones := widget.NewButton("Intrucciones", func ()  {
		//Mostrar las intrucciones
	})

	botonSalir := widget.NewButton("Salir", func() {
		//Cierra la aplicacion
		a.Quit()
	})

	// Estilo de los botones
	botonJugar.Importance = widget.HighImportance
	botonIntrucciones.Importance = widget.HighImportance
	botonSalir.Importance = widget.DangerImportance

	// Contenedor
	content := container.NewVBox(
		titulo,
		container.NewVBox(botonJugar, botonIntrucciones),
		botonSalir,
	)

	// Centra el contenido
	contentCenter := container.NewCenter(content)
	// Agrega el contenido a la ventana
	w.SetContent(contentCenter)
	// Define el tamaño de la ventana
	w.Resize(fyne.NewSize(500, 500))

	//Corre la aplicaion
	w.ShowAndRun()
}