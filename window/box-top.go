package window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func _top() *fyne.Container {
	startRecording := widget.NewButtonWithIcon("开始录制", theme.MediaPlayIcon(), _main_startRecording)
	stopRecording := widget.NewButtonWithIcon("停止录制", theme.MediaStopIcon(), _main_stopRecording)
	startRepeat := widget.NewButtonWithIcon("开始点击", theme.MediaPlayIcon(), _main_startRepeat)
	stopRepeat := widget.NewButtonWithIcon("停止点击", theme.MediaStopIcon(), _main_stopRepeat)

	return container.NewHBox(
		startRecording, stopRecording,
		startRepeat, stopRepeat,
	)
}
