package views

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func MainMenu(w fyne.Window) {
	// Etiqueta para el titulo
	titulo := canvas.NewText("Memory Game", color.White)
	titulo.Alignment = fyne.TextAlignCenter
	titulo.TextSize = 24

	// Botones de Inicio
	botonJugar := widget.NewButton("Jugar", func() {
		// Pasar a la siguiente ventana
		InterPlayers(w)
	})

	botonIntrucciones := widget.NewButton("Intrucciones", func() {
		// Mostrar las intrucciones
		ventIntruc(w)
	})

	botonSalir := widget.NewButton("Salir", func() {
		// Cierra la aplicacion
		w.Close()
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
}

func ventIntruc(parent fyne.Window) {
	// Titulo
	tituloIntruc := canvas.NewText("Intrucciones", color.White)
	tituloIntruc.Alignment = fyne.TextAlignCenter
	tituloIntruc.TextSize = 24

	// Intrucciones
	contenido := widget.NewLabel("Aquí van las intrucciones.")
	contenido.Wrapping = fyne.TextWrapWord

	// Boton de cierre
	botonCerrar := widget.NewButton("Salir", func() {
		parent.Canvas().Overlays().Remove(parent.Canvas().Overlays().Top())
	})
	// Estilo del boton
	botonCerrar.Importance = widget.DangerImportance

	// Contenido de las intrucciones
	contIntruc := container.NewVBox(
		tituloIntruc,
		contenido,
		botonCerrar,
	)

	// Centra el contenido
	intrucCentradas := container.NewCenter(contIntruc)
	// Crea el canva o ventana popup
	modal := widget.NewModalPopUp(intrucCentradas, parent.Canvas())
	// Asigna el tamaño del canva
	modal.Resize(fyne.NewSize(400, 300))
	modal.Show()
}

func win(parent fyne.Window) {
	// Ganador
	ganador := canvas.NewText("Ganador", color.RGBA{R: 0, G: 255, B: 0, A: 255})
	ganador.Alignment = fyne.TextAlignCenter
	ganador.TextSize = 24
	// Nombre del ganador
	contenidoG := widget.NewLabel("Jugador 1")
	contenidoG.Wrapping = fyne.TextWrapWord
	
	// Perdedor
	perdedor := canvas.NewText("Perdedor", color.RGBA{R: 255, G: 0, B: 0, A: 255})
	perdedor.Alignment = fyne.TextAlignCenter
	perdedor.TextSize = 24
	// Nombre del perdedor
	contenidoP := widget.NewLabel("Jugador 2")
	contenidoP.Wrapping = fyne.TextWrapWord

	// Boton de cierre
	botonCerrar := widget.NewButton("Salir", func() {
		parent.Canvas().Overlays().Remove(parent.Canvas().Overlays().Top())
	})
	// Estilo del boton
	botonCerrar.Importance = widget.DangerImportance

	// Contenido de las intrucciones
	contIntruc := container.NewVBox(
		ganador,
		contenidoG,
		perdedor,
		contenidoP,
		botonCerrar,
	)

	// Centra el contenido
	intrucCentradas := container.NewCenter(contIntruc)
	// Crea el canva o ventana popup
	modal := widget.NewModalPopUp(intrucCentradas, parent.Canvas())
	// Asigna el tamaño del canva
	modal.Resize(fyne.NewSize(400, 300))
	modal.Show()
}