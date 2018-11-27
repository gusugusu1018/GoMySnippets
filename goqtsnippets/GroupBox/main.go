package main

import (
	"os"

	"github.com/therecipe/qt"
	"github.com/therecipe/qt/widgets"
)

const (
	minWidth  = 300
	minHeight = 100
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

	// Button Widget
	btn1 := widgets.NewQPushButton2("Button 1", nil)
	btn1.ConnectClicked(func(bool) {
		ret := widgets.QMessageBox_Information(nil, "Button 1", "Are you sure?", widgets.QMessageBox__Yes, widgets.QMessageBox__No)
		if ret == widgets.QMessageBox__Yes {
			statusBar.ShowMessage("Button 1 clieked", 1000)
		}
	})

	btn2 := widgets.NewQPushButton2("Button 2", nil)
	btn2.ConnectClicked(func(bool) {
		ret := widgets.QMessageBox_Information(nil, "Button 2", "Are you sure?", widgets.QMessageBox__No, widgets.QMessageBox__Yes)
		if ret == widgets.QMessageBox__Yes {
			statusBar.ShowMessage("Button 2 clieked", 1000)
		}
	})

	btn3 := widgets.NewQPushButton2("Button 3", nil)
	btn3.ConnectClicked(func(bool) {
		ret := widgets.QMessageBox_Information(nil, "Button 3", "Are you sure?", widgets.QMessageBox__Yes, widgets.QMessageBox__Yes)
		if ret == widgets.QMessageBox__Yes {
			statusBar.ShowMessage("Button 3 clieked", 1000)
		}
	})

	// Group Layout
	var btnLayout = widgets.NewQGridLayout2()
	btnLayout.AddWidget(btn1, 0, 0, 0)
	btnLayout.AddWidget(btn2, 0, 1, 0)
	btnLayout.AddWidget(btn3, 0, 2, 0)
	btnGroup := widgets.NewQGroupBox2("Button Group", nil)
	btnGroup.SetLayout(btnLayout)

	// Layout
	mainWidget := widgets.NewQWidget(nil, 0)
	mainWidget.SetLayout(widgets.NewQHBoxLayout())
	mainWidget.Layout().AddWidget(btnGroup)
	window.SetCentralWidget(mainWidget)

	window.Show()
	app.Exec()
}
