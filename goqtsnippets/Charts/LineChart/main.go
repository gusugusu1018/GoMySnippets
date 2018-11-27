package main

import (
	"math"
	"os"

	"github.com/therecipe/qt/charts"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

const (
	minWidth  = 500
	minHeight = 250
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(minWidth, minHeight)
	window.SetWindowTitle("Line Chart")

	// Chart
	chart := charts.NewQChart(nil, core.Qt__Widget)
	chart.SetTheme(charts.QChart__ChartThemeDark)
	chart.SetTitle("Sin Wave")
	chart.Legend().SetVisible(false)

	xaxis := charts.NewQValueAxis(nil)
	yaxis := charts.NewQValueAxis(nil)
	series := charts.NewQLineSeries(nil)
	x := 0.0
	for x < 2.0*math.Pi {
		series.Append(x, math.Sin(x))
		x += 0.01
	}
	xaxis.SetRange(0.0, 2*math.Pi)
	yaxis.SetRange(-1.0, 1.0)
	chart.AddSeries(series)
	chart.CreateDefaultAxes()
	chart.SetAxisX(xaxis, nil)
	chart.SetAxisY(yaxis, nil)
	chartView := charts.NewQChartView2(chart, nil)

	// Layout
	mainWidget := widgets.NewQWidget(nil, 0)
	mainWidget.SetLayout(widgets.NewQHBoxLayout())
	mainWidget.Layout().AddWidget(chartView)
	window.SetCentralWidget(mainWidget)

	window.Show()
	app.Exec()
}
