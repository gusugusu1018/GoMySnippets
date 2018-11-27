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
	window.SetWindowTitle("Push Button Snippet")
	// Status Bar Widget
	statusBar := widgets.NewQStatusBar(window)
	window.SetStatusBar(statusBar)
	statusBar.SetMaximumHeight(15)
	statusBar.SetVisible(true)
	// Button Widget
	btn := widgets.NewQPushButton2("Push Button", nil)
	btn.ConnectClicked(func(bool) {
		statusBar.ShowMessage("Clieked", 1000)
		qt.Logger.Println("Clieked")
	})
	// Layout
	mainWidget := widgets.NewQWidget(nil, 0)
	mainWidget.SetLayout(widgets.NewQHBoxLayout())
	mainWidget.Layout().AddWidget(btn)
	window.SetCentralWidget(mainWidget)

	window.Show()
	app.Exec()
}
