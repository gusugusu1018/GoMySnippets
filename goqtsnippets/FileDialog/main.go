package main

import (
	"os"

	"github.com/therecipe/qt"
	"github.com/therecipe/qt/widgets"
)

const (
	minWidth  = 100
	minHeight = 50
)

func main() {
	qt.Logger.Println("Start")
	app := widgets.NewQApplication(len(os.Args), os.Args)
	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(minWidth, minHeight)
	window.SetWindowTitle("File Dialog Snippet")
	// Status Bar Widget
	statusBar := widgets.NewQStatusBar(window)
	window.SetStatusBar(statusBar)
	statusBar.SetMaximumHeight(15)
	statusBar.SetVisible(true)
	// Button Widget
	btn := widgets.NewQPushButton2("File", nil)
	btn.ConnectClicked(func(bool) {
		statusBar.ShowMessage("Clieked", 1000)
		filename := widgets.QFileDialog_GetOpenFileName(nil, "Select File", "./", "", "", widgets.QFileDialog__HideNameFilterDetails)
		if filename != "" {
			statusBar.ShowMessage(filename, 5000)
		} else {
			statusBar.ShowMessage("Not Selected", 5000)
		}
	})
	// Layout
	mainWidget := widgets.NewQWidget(nil, 0)
	mainWidget.SetLayout(widgets.NewQHBoxLayout())
	mainWidget.Layout().AddWidget(btn)
	window.SetCentralWidget(mainWidget)

	window.Show()
	app.Exec()
}
