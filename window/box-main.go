package window

import (
	"MouseRecording/extra"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var textList = binding.NewStringList()
var checkedList = binding.NewBoolList()
var clickedXList = binding.NewIntList()
var clickedYList = binding.NewIntList()

var list *widget.List = extra.NewListWithCheckbox(textList, checkedList)

var isListening bool = false
var isRepeating bool = false

func _main() *widget.List {
	return list
}

func _main_startRecording() {
	if isListening {
		return
	}

	fmt.Println("开始录制")
	isListening = true
	isMouseDown := false

	go func() {
		for {
			if !isListening {
				fmt.Println("录制已退出")
				break
			}

			if ok := hook.AddEvent("mleft"); ok {
				isMouseDown = !isMouseDown

				if isMouseDown && isListening {
					x, y := robotgo.Location()
					fyne.Do(func() {
						textList.Append(fmt.Sprintf("(%d, %d)", x, y))
						checkedList.Append(false)
						clickedXList.Append(x)
						clickedYList.Append(y)
						list.Refresh()
					})
				}
			}
		}
	}()
}

func _main_stopRecording() {
	fmt.Println("停止录制")
	isListening = false
}

func _main_startRepeat() {
	if isRepeating {
		return
	}

	fmt.Println("开始重复")
	isRepeating = true
	listSelected := []map[string]int{}

	for i := range textList.Length() {
		if checked, _ := checkedList.GetValue(i); checked {
			checkedX, _ := clickedXList.GetValue(i)
			checkedY, _ := clickedYList.GetValue(i)
			listSelected = append(listSelected, map[string]int{
				"x": checkedX,
				"y": checkedY,
			})
		}
	}

	go func() {
		for {
			if !isRepeating {
				fmt.Println("重复已退出")
				break
			}

			for _, step := range listSelected {
				fmt.Printf("执行点击 (%d, %d)\n", step["x"], step["y"])

				currentX, currentY := robotgo.Location()
				robotgo.MoveClick(step["x"], step["y"])
				robotgo.Move(currentX, currentY)

				time.Sleep(2 * time.Second)
			}
		}
	}()
}

func _main_stopRepeat() {
	fmt.Println("停止重复")
	isRepeating = false
}
