package main

import (
	"MouseRecording/window"

	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()

	window.MainWindow(app).Show()
	app.Run()
}
