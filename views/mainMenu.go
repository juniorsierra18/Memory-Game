package views

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func MainMenu(w fyne.Window) {
	bg := canvas.NewRectangle(color.Black)
	bg.Resize(fyne.NewSize(500, 500))

	// Etiqueta para el titulo
	titulo := canvas.NewText("Memory Game", color.White)
	titulo.Alignment = fyne.TextAlignCenter
	titulo.TextSize = 40

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
	contentCenter := container.New(layout.NewStackLayout(), bg, container.NewCenter(content))
	// Agrega el contenido a la ventana
	w.SetContent(contentCenter)
	// Define el tamaño de la ventana
	w.Resize(fyne.NewSize(500, 500))
}

func ventIntruc(parent fyne.Window) {
	// Titulo
	tituloIntruc := canvas.NewText("Instrucciones", color.White)
	tituloIntruc.Alignment = fyne.TextAlignCenter
	tituloIntruc.TextSize = 40

	// Instrucciones
	contenido := widget.NewLabel(`
INTRODUCCIÓN:
Este es un juego de memoria en el que los jugadores deben encontrar pares de cartas iguales. El juego ofrece diferentes configuraciones de dificultad, con varias cantidades de cartas disponibles. Se puede jugar solo o con un amigo.

1. Menú Principal:
   - Jugar: Inicia el juego seleccionando el número de jugadores.
   - Instrucciones: Muestra las instrucciones del juego.
   - Salir: Cierra la aplicación.

SELECCIÓN DEL NÚMERO DE JUGADORES:
1. Un Jugador: Selecciona esta opción para jugar solo.
2. Dos Jugadores: Selecciona esta opción para jugar contra otro jugador.

SELECCIÓN DE LA CANTIDAD DE CARTAS:
Después de seleccionar el número de jugadores, elige la cantidad de cartas con las que deseas jugar:
- 8 Cartas
- 12 Cartas
- 16 Cartas
- 20 Cartas

OBJETIVO DEL JUEGO:
El objetivo es encontrar todos los pares de cartas iguales. En cada turno, un jugador selecciona dos cartas para voltearlas. Si las cartas coinciden, permanecen volteadas, y el jugador suma puntos. Si no coinciden, se voltean de nuevo, y el turno pasa al siguiente jugador (si hay dos jugadores).

REGLAS DEL JUEGO:
1. Volteo de Cartas:
   - Selecciona una carta para voltearla. Su valor será revelado.
   - Selecciona una segunda carta. Si ambas cartas coinciden, permanecerán volteadas. Si no coinciden, ambas se voltearán de nuevo.
2. Cambio de Turno:
   - En un juego de dos jugadores, el turno cambia automáticamente después de cada intento fallido. Si un jugador encuentra un par, puede seguir jugando.
3. Puntuación:
   - Jugador 1: Su puntuación se muestra en la parte inferior de la pantalla.
   - Jugador 2: Si se juega en modo de dos jugadores, su puntuación también se mostrará.
4. Volver al Menú: Durante el juego, puedes volver al menú principal presionando el botón "Volver".

FIN DEL JUEGO:
El juego termina cuando todos los pares han sido encontrados. El jugador con más puntos gana (en el caso de dos jugadores). Si juegas solo, tu objetivo es encontrar todos los pares en el menor tiempo posible.

CONTROLES:
- Ratón: Utiliza el ratón para hacer clic en los botones y seleccionar cartas.
- Salir: Puedes salir del juego en cualquier momento utilizando el botón "Salir" en el menú principal o la ventana de instrucciones.
`)
	contenido.Wrapping = fyne.TextWrapWord

	// Crea un contenedor de desplazamiento para las instrucciones
	scrollCont := container.NewScroll(contenido)
	scrollCont.SetMinSize(fyne.NewSize(600, 400)) // Ajusta el tamaño del contenedor de desplazamiento

	// Botón de cierre
	botonCerrar := widget.NewButton("Salir", func() {
		parent.Canvas().Overlays().Remove(parent.Canvas().Overlays().Top())
	})
	botonCerrar.Importance = widget.DangerImportance

	// Contenido de las instrucciones con desplazamiento
	contIntruc := container.NewVBox(
		tituloIntruc,
		scrollCont,
		botonCerrar,
	)

	// Centra el contenido
	intrucCentradas := container.NewCenter(contIntruc)
	// Crea la ventana popup
	modal := widget.NewModalPopUp(intrucCentradas, parent.Canvas())
	modal.Resize(fyne.NewSize(700, 500)) // Ajusta el tamaño de la ventana emergente
	modal.Show()
}

/*func win(parent fyne.Window) {
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
}*/
