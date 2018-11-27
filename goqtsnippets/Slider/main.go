package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

const (
	minWidth  = 150
	minHeight = 50
)

func main() {
	qt.Logger.Println("Start")
	app := widgets.NewQApplication(len(os.Args), os.Args)
	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(minWidth, minHeight)
	window.SetWindowTitle("Slider Snippet")
	// Status Bar Widget
	statusBar := widgets.NewQStatusBar(window)
	window.SetStatusBar(statusBar)
	statusBar.SetMaximumHeight(15)
	statusBar.SetVisible(true)
	// Slider Widget
	initValue := 50
	sldr := widgets.NewQSlider2(core.Qt__Horizontal, nil)
	sldr.SetMaximum(100)
	sldr.SetMinimum(0)
	sldr.SetValue(50)
	lblValue := widgets.NewQLabel2(fmt.Sprintf("%d%%", initValue), nil, core.Qt__Widget)
	sldr.ConnectSliderMoved(func(value int) {
		statusBar.ShowMessage("Moved", 1000)
		lblValue.SetText(fmt.Sprintf("%d%%", value))
	})
	// Layout
	mainWidget := widgets.NewQWidget(nil, 0)
	mainWidget.SetLayout(widgets.NewQHBoxLayout())
	mainWidget.Layout().AddWidget(lblValue)
	mainWidget.Layout().AddWidget(sldr)
	window.SetCentralWidget(mainWidget)

	window.Show()
	app.Exec()
}
