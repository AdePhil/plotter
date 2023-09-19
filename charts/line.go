package charts

import (
	"io"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

type Line struct {
	XValues    []interface{}
	SeriesData []opts.LineData
	Smooth     bool
}

func (line Line) RenderGraph(w http.ResponseWriter, r *http.Request) {
	line.drawLineChart(w)
}

func (lineData Line) drawLineChart(w io.Writer) {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "All countries in the world by population", Left: "center"}),
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithLegendOpts(opts.Legend{Show: false}))

	line.SetXAxis(lineData.XValues).AddSeries("Testing", lineData.SeriesData).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: lineData.Smooth}))

	line.Render(w)
}
