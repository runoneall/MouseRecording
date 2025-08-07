package window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func MainWindow(app fyne.App) fyne.Window {
	win := app.NewWindow("鼠标录制器")
	win.Resize(fyne.NewSize(0, 400))

	win.SetContent(container.NewBorder(
		_top(), nil, nil, nil, _main(),
	))
	return win
}
