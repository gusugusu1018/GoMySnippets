package main

import (
	"os"
	"fmt"

	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

const (
	minWidth  = 150
	minHeight = 150
)

func main() {
	qt.Logger.Println("Start")
	app := widgets.NewQApplication(len(os.Args), os.Args)
	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(minWidth, minHeight)
	window.SetWindowTitle("ImageView Snippet")
	// Status Bar Widget
	statusBar := widgets.NewQStatusBar(window)
	window.SetStatusBar(statusBar)
	statusBar.SetMaximumHeight(15)
	statusBar.SetVisible(true)

	if len(os.Args) != 2 {
		msg := fmt.Sprintf("Usage : %s <image file>\n", os.Args[0])
		qt.Logger.Printf(msg)
		statusBar.ShowMessage(msg,1000)
		os.Exit(0)
	}
	imgFileName := os.Args[1]
	msg := fmt.Sprintf("Loading image : ", imgFileName)
	qt.Logger.Printf(msg)
	statusBar.ShowMessage(msg,1000)

	// Graphics View Widget
	scene := widgets.NewQGraphicsScene(nil)
	view := widgets.NewQGraphicsView(nil)
	pixmap := gui.NewQPixmap5(imgFileName, "", core.Qt__AutoColor)
	item := widgets.NewQGraphicsPixmapItem2(pixmap, nil)
	scene.AddItem(item)
	view.SetScene(scene)

	// Layout
	mainWidget := widgets.NewQWidget(nil, 0)
	mainWidget.SetLayout(widgets.NewQHBoxLayout())
	mainWidget.Layout().AddWidget(view)
	window.SetCentralWidget(mainWidget)
	window.Show()
	app.Exec()
}
