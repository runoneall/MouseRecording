package extra

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewListWithCheckbox(
	textList binding.StringList,
	checkedList binding.BoolList,
) *widget.List {
	return widget.NewList(

		func() int {
			return textList.Length()
		},

		func() fyne.CanvasObject {
			return container.NewHBox(
				widget.NewCheck("", nil),
				widget.NewLabel(""),
				layout.NewSpacer(),
			)
		},

		func(lii widget.ListItemID, co fyne.CanvasObject) {
			_container := co.(*fyne.Container)
			_check := _container.Objects[0].(*widget.Check)
			_label := _container.Objects[1].(*widget.Label)

			text, _ := textList.GetValue(lii)
			checked, _ := checkedList.GetValue(lii)

			_label.SetText(text)
			_check.SetChecked(checked)

			_check.OnChanged = func(b bool) {
				checkedList.SetValue(lii, b)
			}
		},
	)
}
