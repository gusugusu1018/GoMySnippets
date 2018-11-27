package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

const (
	minWidth  = 300
	minHeight = 200
)

func main() {
	qt.Logger.Println("Start")
	app := widgets.NewQApplication(len(os.Args), os.Args)
	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(minWidth, minHeight)
	window.SetWindowTitle("GroupBox Snippet")
	// Status Bar Widget
	statusBar := widgets.NewQStatusBar(window)
	window.SetStatusBar(statusBar)
	statusBar.SetMaximumHeight(15)
	statusBar.SetVisible(true)

	// ProgressBar Widget
	initValue := 100
	progBar := widgets.NewQProgressBar(nil)
	progBar.SetOrientation(core.Qt__Horizontal)
	progBar.SetMaximum(100)
	progBar.SetMinimum(0)
	progBar.SetTextVisible(false)
	progBar.SetValue(initValue)
	// Label for print progressBar value
	lblValue := widgets.NewQLabel2(fmt.Sprintf("%d%%", initValue), nil, core.Qt__Widget)
	progBar.ConnectValueChanged(func(value int) {
		lblValue.SetText(fmt.Sprintf("%d%%", value))
	})

	// Progress Bar Layout
	progBarWiget := widgets.NewQWidget(nil, 0)
	progBarWiget.SetLayout(widgets.NewQHBoxLayout())
	progBarWiget.Layout().AddWidget(lblValue)
	progBarWiget.Layout().AddWidget(progBar)

	// Button Widget
	btnUp := widgets.NewQPushButton2("Up", nil)
	btnUp.ConnectClicked(func(bool) {
		_tmp := progBar.Value() + 10
		progBar.SetValue(_tmp)
		statusBar.ShowMessage("Button Up clicked", 1000)
	})

	btnDown := widgets.NewQPushButton2("Down", nil)
	btnDown.ConnectClicked(func(bool) {
		_tmp := progBar.Value() - 10
		progBar.SetValue(_tmp)
		statusBar.ShowMessage("Button Down clicked", 1000)
	})

	// Group Layout
	var btnLayout = widgets.NewQGridLayout2()
	btnLayout.AddWidget(btnUp, 0, 0, 0)
	btnLayout.AddWidget(btnDown, 0, 1, 0)
	btnGroup := widgets.NewQGroupBox2("Button Group", nil)
	btnGroup.SetLayout(btnLayout)

	// Layout
	mainWidget := widgets.NewQWidget(nil, 0)
	mainWidget.SetLayout(widgets.NewQVBoxLayout())
	mainWidget.Layout().AddWidget(progBarWiget)
	mainWidget.Layout().AddWidget(btnGroup)
	window.SetCentralWidget(mainWidget)

	window.Show()
	app.Exec()
}
