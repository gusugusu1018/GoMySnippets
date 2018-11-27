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

	btnMsg := widgets.NewQPushButton2("Gopher", nil)
	btnMsg.ConnectClicked(func(bool) {
		ret := widgets.QMessageBox_Information(nil, "Message", "Is Gopher chan Kawaii?", widgets.QMessageBox__Yes, widgets.QMessageBox__Yes)
		if ret == widgets.QMessageBox__Yes {
			statusBar.ShowMessage("Gopher chan is very Kawaii", 1000)
		}
	})

	// Group Layout
	var btnLayout = widgets.NewQGridLayout2()
	btnLayout.AddWidget(btnUp, 0, 0, 0)
	btnLayout.AddWidget(btnDown, 0, 1, 0)
	btnLayout.AddWidget(btnMsg, 0, 2, 0)
	btnGroup := widgets.NewQGroupBox2("Button Group", nil)
	btnGroup.SetLayout(btnLayout)

	// Layout
	mainWidget := widgets.NewQWidget(nil, 0)
	mainWidget.SetLayout(widgets.NewQVBoxLayout())
	mainWidget.Layout().AddWidget(progBarWiget)
	mainWidget.Layout().AddWidget(btnGroup)
	window.SetCentralWidget(mainWidget)

	window.Show()
	style := `
		QMainWindow {
			background-color: #27313D
		}
		QLabel {
			color: white;
		}
		QGroupBox {
			background-color: #7C858E;
			border-radius: 4px;
			margin-top: 1ex;
		}
		QGroupBox:title {
			color: white
		}
		QPushButton {
			color: 08233E;
			background-color: #E6E5E1;
			border-width: 0px;
			border-color: #27313D;
			border-style: outset;
			border-radius: 4px;
		}
		QPushButton:pressed {
			background-color: #CCCCCC;
		}
		QProgressBar {
			color: #08233E;
			text-align: center;
			background-color: white;
			border: 2px solid grey;
			border-radius: 5px;
		}
		QProgressBar:chunk {
			background-color: #05B8CC;
			width: 20px;
		}
		QMessageBox {
			background-color: #CCCCCC
		}
		QMessageBox QLabel {
			color: #08233E
		}
		QMessageBox QPushButton {
			color: 08233E;
			background-color: #E6E5E1;
		}
		QMessageBox QPushButton:pressed {
			background-color: #CCCCCC;
		}
	`
	app.SetStyleSheet(style)
	app.Exec()
}
