package main

import (
	"math"
	"os"

	"github.com/therecipe/qt"
	"github.com/therecipe/qt/charts"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type Chart struct {
	core.QObject
	*charts.QChart //TODO: directly subclass QChart instead

	_ func() `constructor:"init"`

	_ func() `slot:"handleTimeout,<-(this.m_timer.timeout)"`

	m_timer  *core.QTimer
	m_series *charts.QLineSeries
	m_titles []string
	m_axis   *charts.QValueAxis
	m_x      float64
	m_y      float64
	m_dx	 float64
	m_plotxmax float64 
	m_plotxmin float64
	m_plotymax float64
	m_plotymin float64
}

func (c *Chart) init() {
	c.QChart = charts.NewQChart(nil, 0)
	c.SetTheme(charts.QChart__ChartThemeDark)
	c.m_timer = core.NewQTimer(nil)
	c.m_timer.SetInterval(100)
	c.m_axis = charts.NewQValueAxis(nil)
	c.m_x = 0.0
	c.m_y = 0.0
	c.m_dx = 0.1
	c.m_plotxmax = 10.0
	c.m_plotxmin = 0.0
	c.m_plotymax = 1.0
	c.m_plotymin = -1.0

	c.m_series = charts.NewQLineSeries(c)
	green := gui.NewQPen3(gui.NewQColor2(core.Qt__blue))
	green.SetWidth(3)
	c.m_series.SetPen(green)
	c.m_series.Append(c.m_x, c.m_y)

	c.AddSeries(c.m_series)
	c.CreateDefaultAxes()
	c.SetAxisX(c.m_axis, c.m_series)
	c.m_axis.SetTickCount(10)
	c.AxisX(nil).SetRange(core.NewQVariant12(c.m_plotxmin), core.NewQVariant12(c.m_plotxmax))
	c.AxisY(nil).SetRange(core.NewQVariant12(c.m_plotymin), core.NewQVariant12(c.m_plotymax))

	c.m_timer.Start2()
}

func (c *Chart) handleTimeout() {
	c.m_x += c.m_dx
	c.m_y = math.Sin(c.m_x)
	c.m_series.Append(c.m_x, c.m_y)
	scrollx := c.m_dx * c.PlotArea().Width() / c.m_plotxmax
	if c.m_x >= 10 {
		c.Scroll(scrollx, 0)
	}
}

const (
	minWidth  = 500
	minHeight = 250
)

func main() {
	qt.Logger.Println("Start")
	app := widgets.NewQApplication(len(os.Args), os.Args)
	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(minWidth, minHeight)
	window.SetWindowTitle("Dynamic Line Chart")

	chart := NewChart(nil)
	chart.SetTitle("Sine wave")
	chart.Legend().Hide()
	chart.SetAnimationOptions(charts.QChart__SeriesAnimations)
	chart.SetAnimationDuration(1)

	chartView := charts.NewQChartView2(chart, nil)
	chartView.SetRenderHint(gui.QPainter__Antialiasing, true)

	window.SetCentralWidget(chartView)
	window.Resize2(400, 300)
	window.Show()
	app.Exec()
}
